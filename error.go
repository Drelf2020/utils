package utils

import (
	"runtime"
)

// 检查错误 存在错误返回 true
func CheckErr(err error) bool {
	return err != nil
}

// 存在错误则  panic
func PanicErr(err error) {
	if CheckErr(err) {
		panic(err)
	}
}

// 存在错误则打印错误并返回 true
func LogErr(err error) bool {
	if CheckErr(err) {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			log.Errorf("%v occured in %v:%v", err, file, line)
		} else {
			log.Error(err)
		}
		return true
	}
	return false
}
