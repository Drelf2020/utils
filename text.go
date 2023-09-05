package utils

import "strings"

// 首字母大写
func Capitalize(s string) string {
	return strings.ToUpper(s[:1]) + s[1:]
}

// 字符串切片
//
// cut 表示区间开闭
//
// 0(00) 左右都不保留
//
// 1(01) 左不保留右保留
//
// 2(10) 左保留右不保留
//
// 3(11) 左右都保留
func Cut(s, start, end string, cut int) string {
	st := strings.Index(s, start)
	sp := strings.LastIndex(s, end)
	if st == -1 || sp == -1 {
		return ""
	}
	st += (cut>>1 ^ 1) * len(start)
	sp += (cut & 1) * len(end)
	return s[st:sp]
}
