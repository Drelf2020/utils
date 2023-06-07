package utils

import "strings"

// 起始匹配
func Startswith(s, prefix string) bool {
	return strings.HasPrefix(s, prefix)
}

// 结束匹配
func Endswith(s, suffix string) bool {
	return strings.HasSuffix(s, suffix)
}
