package main

import (
	"os"
)

func main() {
	cli := &CLI{
		version: "v1.0.0",
		errors:  "error",
	}
	os.Exit(cli.Run(os.Args))
}
