package utils

var flags = map[string]string{
	"-g": "Google",
	"-m": "MDN",
}

func IsValidFlag(flag string) bool {
	if _, ok := flags[flag]; ok {
		return true
	}

	return false
}
