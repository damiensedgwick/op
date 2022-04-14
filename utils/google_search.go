package utils

const googlePrefix = "https://www.google.com/search?q="

func GoogleSearch(q []string) string {
	return googlePrefix + QueryBuilder(q)
}
