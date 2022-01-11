package main

import (
	"github.com/cli/browser"
	"github.com/damiensedgwick/op/utils"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]
	arg := args[0]

	if len(args) == 1 {
		err := browser.OpenURL(utils.ReturnCorrectLink(arg))
		if err != nil {
			log.Fatal("Could not open documents")
		}
	}

	if len(args) > 1 && arg == "mdn" {
		err := browser.OpenURL(utils.MDNSearch(args[1]))
		if err != nil {
			log.Fatal("Could not open mdn")
		}
	}
}
