package main

import (
	"fmt"
	"github.com/cli/browser"
	"github.com/damiensedgwick/op/utils"
	"log"
	"os"
)

func main() {
	arg := os.Args[1]
	//flags := os.Args[2:]

	fmt.Println("Searching for " + arg + " documents...")
	fmt.Println(arg + " documents found, opening browser...")

	err := browser.OpenURL(utils.ReturnCorrectLink(arg))
	if err != nil {
		log.Fatal("Could not open documents")
	}

	// when completed, the goal is to open the link in the browser and print it for reference.
	//fmt.Println(utils.ReturnCorrectLink(arg))
}
