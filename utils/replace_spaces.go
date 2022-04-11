package utils

import "strings"

func ReplaceSpaces(q []string) string {
	return strings.Join(q, "-")
}
