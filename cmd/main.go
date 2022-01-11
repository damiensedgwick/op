package main

import (
	"fmt"
	"github.com/cli/browser"
	"github.com/damiensedgwick/op/utils"
	"log"
	"os"
)

const mdnPrefix = "https://developer.mozilla.org/en-US/search?q="

func MDNSearch(q string) string {
	url := mdnPrefix + q

	return url
}

func main() {
	args := os.Args[1:]
	arg := args[0]

	fmt.Println(args)
	fmt.Println(len(args))

	fmt.Println(MDNSearch("div"))

	if len(args) == 1 {
		fmt.Println("Searching for " + arg + " documents...")
		fmt.Println(arg + " documents found, opening browser...")
		err := browser.OpenURL(utils.ReturnCorrectLink(arg))
		if err != nil {
			log.Fatal("Could not open documents")
		}
	}

	if len(args) > 1 {
		if arg == "mdn" {
			err := browser.OpenURL(MDNSearch(args[1]))
			if err != nil {
				log.Fatal("Could not open mdn")
			}
		}
	}
}
