package utils

const mdnPrefix = "https://developer.mozilla.org/en-US/search?q="

func MDNSearch(q []string) string {
	url := mdnPrefix + QueryBuilder(q)

	return url
}
