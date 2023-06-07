package utils

import "strings"

func Startswith(s, prefix string) bool {
	return strings.HasPrefix(s, prefix)
}

func Endswith(s, suffix string) bool {
	return strings.HasSuffix(s, suffix)
}
