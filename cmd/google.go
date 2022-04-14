package cmd

import (
	"github.com/cli/browser"
	"github.com/damiensedgwick/op/utils"
	"log"
)

func Google(args []string) {
	err := browser.OpenURL(utils.GoogleSearch(args[1:]))
	if err != nil {
		log.Fatal("Could not perform google search")
	}
}
