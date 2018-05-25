package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {

	f, err := os.Open("hello")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	cmd := exec.Command("ls")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.ExtraFiles = append(cmd.ExtraFiles, f)
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", cmd)
}
