package main

import (
	"os"
	"os/exec"
)

func Exec(path string, args ...string) error {
	defer Uninstall(path)

	cmd := exec.Command(path, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
