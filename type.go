package utils

import (
	"strconv"
	"strings"
)

// 判断字符是否为数字
func IsNumber(v string) bool {
	// 允许一个小数点
	v = strings.Replace(v, ".", "", 1)
	_, err := strconv.Atoi(v)
	return err == nil
}

// 自动类型
func AutoType(k, v string) (typ string, val any) {
	if k != "" && k != "auto" {
		typ = k
	} else if v == "true" || v == "false" {
		typ = "bool"
	} else if v == "{" {
		typ = "dict"
	} else if IsNumber(v) {
		typ = "num"
	} else {
		typ = "str"
	}
	if v == "" {
		return
	}
	switch typ {
	case "bool":
		val = v == "true"
	case "num":
		val, _ = strconv.ParseFloat(v, 64)
	case "str":
		val = v
	}
	return
}
