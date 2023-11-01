package utils

import (
	"io"
	"os"

	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
)

var (
	// 日志文件
	logfile *os.File
	// 基础输出格式
	log = &logrus.Logger{
		Out: os.Stderr,
		Formatter: &nested.Formatter{
			HideKeys:        true,
			NoColors:        true,
			TimestampFormat: "15:04:05",
		},
		Level: logrus.DebugLevel,
	}
)

// 获取 log
func Logger() *logrus.Logger {
	return log
}

// 关闭日志文件
func CloseLogFile() error {
	if logfile != nil {
		return logfile.Close()
	}
	return nil
}

// 导出文件
func SetOutputFile(path string) *logrus.Logger {
	CloseLogFile()
	logfile = Panic(os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm))
	log.SetOutput(io.MultiWriter(logfile, os.Stdout))
	return log
}

// 设置时间格式
func SetTimestampFormat(format string) *logrus.Logger {
	log.SetFormatter(&nested.Formatter{
		HideKeys:        true,
		NoColors:        true,
		TimestampFormat: format,
	})
	return log
}
