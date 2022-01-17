package utils

import (
	"strings"
)

func QueryBuilder(arr []string) string {
	str := ""

	for _, word := range arr {
		str += "+" + word
	}

	query := strings.Replace(str, "+", "", 1)

	return query
}
