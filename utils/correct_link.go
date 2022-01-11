package utils

import v1 "github.com/damiensedgwick/op/api/v1"

func ReturnCorrectLink(q string) string {
	links := v1.Links()

	res := ""

	for k, v := range links {
		if k == q {
			res = v
		}
	}

	return res
}
