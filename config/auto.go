package flag2config

import "github.com/Drelf2020/utils"

type Parser interface {
	Read() error
	Write() error
}

func auto(config Config) Parser {
	if utils.Endswith(config.Path, ".yml") {
		return Yml{config}
	} else if utils.Endswith(config.Path, ".json") {
		return Json{config}
	}
	return nil
}
