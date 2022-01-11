package main

import (
	"fmt"
	"github.com/cli/browser"
	"github.com/damiensedgwick/op/utils"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		log.Fatal("oops, it looks like you didn't add a search term.")
	}

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

	if len(args) > 1 && arg != "mdn" {
		fmt.Println("unfortunately you can only search mdn docs at the moment.")
		fmt.Println("you can use a single search term 'op react' or using mdn 'op mdn div'.")
	}
}
