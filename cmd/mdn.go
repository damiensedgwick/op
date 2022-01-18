package cmd

import (
	"github.com/cli/browser"
	"github.com/damiensedgwick/op/utils"
	"log"
)

func MDN(args []string) {
	err := browser.OpenURL(utils.MDNSearch(args[1:]))
	if err != nil {
		log.Fatal("Could not open mdn")
	}
}
