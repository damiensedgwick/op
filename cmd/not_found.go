package cmd

import (
	"github.com/cli/browser"
	"github.com/damiensedgwick/op/utils"
	"log"
)

func NotFound(arg string) {
	err := browser.OpenURL(utils.ReturnCorrectLink(arg))
	if err != nil {
		log.Fatal("Could not open documents")
	}
}
