package utils

import (
	"io"
	"os"

	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
)

// 基础输出格式
var log = &logrus.Logger{
	Out: os.Stderr,
	Formatter: &nested.Formatter{
		HideKeys:        true,
		NoColors:        true,
		TimestampFormat: "15:04:05",
	},
	Level: logrus.DebugLevel,
}

// 获取 log
func GetLog() *logrus.Logger {
	return log
}

// 导出文件
func SetOutputFile(path string) {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0666)
	PanicErr(err)
	writers := []io.Writer{
		file,
		os.Stdout,
	}
	fileAndStdoutWriter := io.MultiWriter(writers...)
	log.SetOutput(fileAndStdoutWriter)
}
