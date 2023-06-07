package utils

import "os"

// FileExists 判断一个文件是否存在
//
// 参考 https://blog.csdn.net/leo_jk/article/details/118255913
func FileExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
