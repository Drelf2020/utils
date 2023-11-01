package utils

import (
	"fmt"
	"reflect"
	"runtime"
)

// 存在错误则 panic
func PanicErr(err error) {
	if err != nil {
		panic(err)
	}
}

func Panic[T any](data T, err error) T {
	PanicErr(err)
	return data
}

// 存在错误则打印错误并返回 true
func LogErr(err error) bool {
	if err != nil {
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

// The Try function allow you run a code in f function without catching the error.
//
// When an error raised, it will be automatic capture
// and send into catch functions as a parameter.
//
// If there is no error, nothing will happen.
func Try(f func(), catch ...func(error)) (e error) {
	defer func() {
		if r := recover(); r != nil {
			e = fmt.Errorf("%v", r)
			for _, c := range catch {
				if c != nil {
					c(e)
				}
			}
		}
	}()
	if f != nil {
		f()
	}
	return nil
}

// The CopyAny function copy value from the src into des.
//
// An error will be return when Their types are different
// or dst can't be set.
func CopyAny(dst, src any) error {
	return Try(func() { reflect.ValueOf(dst).Elem().Set(reflect.ValueOf(src)) })
}

// The Copy function with type parameter.
func Copy[T any](dst *T, src T) error {
	return CopyAny(dst, src)
}
