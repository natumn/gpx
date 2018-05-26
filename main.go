package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/janeczku/go-spinner"
	sw "github.com/mattn/go-shellwords"
	"github.com/pkg/errors"
	"github.com/urfave/cli"
)

// Command has execute command infomation.
type Command struct {
	name string
	args []string
	repo string
	path string
}

func main() {
	app := cli.NewApp()
	app.Name = "gpx"
	app.Usage = "Execute Go package binaries."

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "businesscard, b",
			Usage: "Load cli business card from github.com/USERNAME/`USERNAME`",
		},
	}

	app.Action = func(c *cli.Context) error {

		cmd := NewCommand(c)

		path, _ := exec.LookPath(cmd.name)

		if path == "" {
			if err := run(cmd); err != nil {
				return err
			}
		} else {
			if err := execCmd(cmd); err != nil {
				return err
			}
		}

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}

func run(cmd Command) error {
	s := spinner.StartNew("Start installing...")
	st := time.Now()

	gobinPath, err := getCommandPath("$GOPATH")
	cmd.path = gobinPath + cmd.name
	if err != nil {
		return err
	}

	err = install(cmd.repo)
	if err != nil {
		return err
	}

	et := time.Now()
	s.Stop()
	fmt.Printf("finish! Install time: %v.\n", et.Sub(st))

	err = execCmd(cmd)
	if err != nil {
		return err
	}

	err = uninstall(cmd.path)
	if err != nil {
		return err
	}

	return nil
}

func install(repo string) error {

	installCmd := exec.Command("go", "get", repo)

	if err := installCmd.Run(); err != nil {
		return errors.Wrap(err, "Install pharse: ")
	}

	return nil
}

func execCmd(cmd Command) error {

	execCmd := exec.Command(cmd.name, cmd.args...)

	execCmd.Stdin = os.Stdin
	execCmd.Stdout = os.Stdout
	execCmd.Stderr = os.Stderr

	if err := execCmd.Run(); err != nil {
		uninstall(cmd.path)
		return errors.Wrap(err, "Exec phase: ")
	}

	return nil
}

func getCommandPath(env string) (string, error) {
	p := sw.NewParser()
	p.ParseEnv = true

	args, err := p.Parse(env)

	if err != nil {
		return "", errors.Wrap(err, "Cannot get command path:")
	} else if len(args) == 0 {
		message := fmt.Sprintf("Not set up %v.", env)
		return "", errors.Wrap(errors.New(message), "Cannot get command path:")
	}

	return args[0] + "/bin/", nil
}

func uninstall(path string) error {

	uninstallCmd := exec.Command("rm", path)

	if err := uninstallCmd.Run(); err != nil {
		return errors.Wrap(err, "Uninstall pharse:")
	}

	return nil
}

// NewCommand is make command struct
func NewCommand(c *cli.Context) Command {
	var cmd Command
	b := c.GlobalString("businesscard")

	if b == "" {
		args := c.Args()
		cmd.repo = args[0]
		cmd.name = getName(cmd.repo)
		cmd.args = args[1:]
	} else {
		cmd.name = b
		cmd.repo = "github.com/" + b + "/" + b
	}

	return cmd
}

func getName(repo string) string {
	slice := strings.Split(repo, "/")
	return slice[2]
}
