package flag2config

import (
	"flag"
	"fmt"

	"github.com/Drelf2020/utils"
)

// 配置
type Config struct {
	// 配置文件路径
	Path string
	// 导出的数据结构体
	Data any
}

var configs = make([]*Config, 0)

// 新建配置
func NewConfig(data any) *Config {
	c := &Config{"", data}
	configs = append(configs, c)
	return c
}

// 读取配置文件
func GetConfig() {
	if !flag.Parsed() {
		flag.Parse()
	}
	for _, config := range configs {
		ac := auto(*config)
		if ac == nil {
			utils.GetLog().Errorf("不支持的文件 %v", config.Path)
		} else if !utils.FileExists(config.Path) {
			ac.Write()
		} else {
			ac.Read()
		}
	}
}

// 设置 flag
func SetFlag(name, value string, data any) {
	flag.StringVar(&NewConfig(data).Path, name, value, fmt.Sprintf("配置文件 %v 路径", name))
}
