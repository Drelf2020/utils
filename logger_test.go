package utils_test

import (
	"testing"

	"github.com/Drelf2020/utils"
)

func TestLog(t *testing.T) {
	log := utils.SetOutputFile(".log")
	log.Info("test")
}
