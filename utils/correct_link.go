package utils

import (
	v1 "github.com/damiensedgwick/op/api/v1"
	"log"
)

func ReturnCorrectLink(q string) string {
	links := v1.Links()
	exists, link := CheckLinkExists(q, links)

	if exists == false {
		log.Fatal("The documentation you seek cannot be found.")
	}

	return link
}

func CheckLinkExists(q string, links map[string]string) (bool, string) {
	l := links

	for k, v := range l {
		if k == q {
			return true, v
		}
	}

	return false, ""
}
