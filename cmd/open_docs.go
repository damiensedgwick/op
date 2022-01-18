package cmd

import (
	"github.com/cli/browser"
	"github.com/damiensedgwick/op/utils"
	"log"
)

func OpenDocs(arg string) {
	err := browser.OpenURL(utils.ReturnCorrectLink(arg))
	if err != nil {
		log.Fatal("Could not open documents")
	}
}
