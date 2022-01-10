package utils

var links = map[string]string{
	"javascript": "https://developer.mozilla.org/en-US/docs/Web/JavaScript",
	"typescript": "https://www.typescriptlang.org/docs/handbook/typescript-in-5-minutes.html",
	"react":      "https://reactjs.org/docs/getting-started.html",
}

func ReturnCorrectLink(q string) string {
	res := ""

	for k, v := range links {
		if k == q {
			res = v
		}
	}

	return res
}
