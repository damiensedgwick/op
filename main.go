package main

import (
	"github.com/damiensedgwick/op/cmd"
	"github.com/damiensedgwick/op/utils"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		cmd.Menu()
		os.Exit(0)
	}

	arg := args[0]

	if len(args) == 1 {
		cmd.OpenDocs(arg)
	}

	if len(args) > 1 && arg != "mdn" {
		q := utils.ReplaceSpaces(args)
		cmd.OpenDocs(q)
	}

	if len(args) > 1 && arg == "mdn" {
		cmd.MDN(args)
	}
}
