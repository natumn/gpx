package main

import (
	"io"
	"os"
)

type Cli struct {
	version string
	errors  io.Writer
}

type Config struct {
}

func (c *Cli) Run(args []string, config Config) int {
	i, err := NewInstaller(config)
	if err != nil {
		return err
	}

	path, err := i.Install(args[1:])
	if err != nil {
		return err
	}

	if err := Exec(path); err != nil {
		return err
	}

	return 0
}
