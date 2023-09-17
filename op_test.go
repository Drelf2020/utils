package utils_test

import (
	"testing"

	"github.com/Drelf2020/utils"
)

func TestTernary(t *testing.T) {
	if utils.Ternary(1 > 2, 1, 2) != 2 {
		t.Fail()
	}
	if utils.Ternary(2 > 1, 1, 2) != 1 {
		t.Fail()
	}
	var a, b = "abc", "xyz"
	utils.Update(1 > 2, &a, b)
	if a != "abc" {
		t.Fail()
	}
	utils.Update(2 > 1, &a, b)
	if a != "xyz" {
		t.Fail()
	}
}
