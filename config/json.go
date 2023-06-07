package flag2config

import (
	"encoding/json"
	"os"

	"github.com/Drelf2020/utils"
)

type Json struct {
	Config
}

// 读取
func (js Json) Read() error {
	file, err := os.ReadFile(js.Path)
	if utils.CheckErr(err) {
		return err
	}
	return json.Unmarshal(file, js.Data)
}

// 写入
func (js Json) Write() error {
	data, err := json.Marshal(js.Data)
	if utils.CheckErr(err) {
		return err
	}
	return os.WriteFile(js.Path, data, 0666)
}
