package utils

import "os"

// FileExists 判断一个文件是否存在
//
// 参考  https://blog.csdn.net/leo_jk/article/details/118255913
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

// 读取文件
func ReadFile(path string) string {
	data, err := os.ReadFile(path)
	if LogErr(err) {
		return ""
	}
	return string(data)
}

// 写入文件
func WriteFile(path, s string) error {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if LogErr(err) {
		return err
	}

	_, err = file.WriteString(s)
	if LogErr(err) {
		return err
	}
	return nil
}
