package flag2config

import (
	"os"

	"gopkg.in/yaml.v2"

	"github.com/Drelf2020/utils"
)

type Yml struct {
	Config
}

// 读取
func (yml Yml) Read() error {
	file, err := os.ReadFile(yml.Path)
	if utils.CheckErr(err) {
		return err
	}
	return yaml.Unmarshal(file, yml.Data)
}

// 写入
func (yml Yml) Write() error {
	data, err := yaml.Marshal(yml.Data)
	if utils.CheckErr(err) {
		return err
	}
	return os.WriteFile(yml.Path, data, 0666)
}
