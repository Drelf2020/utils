package utils

import "os"

// FileExist 判断一个文件是否存在
//
// 参考  https://blog.csdn.net/leo_jk/article/details/118255913
func FileExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil
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
	return os.WriteFile(path, []byte(s), os.ModePerm)
}
