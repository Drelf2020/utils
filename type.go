package utils

import (
	"strconv"
	"strings"
)

// 判断字符是否为整数
func IsDigit(v string) bool {
	_, err := strconv.Atoi(v)
	return err == nil
}

// 判断字符是否为实数
func IsNumber(v string) bool {
	// 允许一个小数点
	return IsDigit(strings.Replace(v, ".", "", 1))
}
