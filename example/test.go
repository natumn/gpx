package main

import (
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("printenv", "GOPATH")
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Run()
}
