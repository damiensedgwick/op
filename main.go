package main

import (
	"github.com/damiensedgwick/op/cmd"
	"os"
)

func main() {
	arg := os.Args[0]
	args := os.Args[1:]

	if len(args) == 0 {
		cmd.Menu()
		os.Exit(0)
	}

	if len(args) == 1 {
		cmd.NotFound(arg)
	}

	if len(args) > 1 && arg == "mdn" {
		cmd.MDN(args)
	}

	if len(args) > 1 && arg != "mdn" {
		cmd.Menu()
		os.Exit(0)
	}
}
