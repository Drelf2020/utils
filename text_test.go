package utils_test

import (
	"testing"

	"github.com/Drelf2020/utils"
)

func TestCut(t *testing.T) {
	if utils.Cut("abc", "abc", "abc", 0) != "" {
		t.Fail()
	}
	if utils.Cut("abc", "abc", "abc", 1) != "" {
		t.Fail()
	}
	if utils.Cut("abc", "abc", "abc", 2) != "" {
		t.Fail()
	}
	if utils.Cut("abc", "abc", "abc", 3) != "abc" {
		t.Fail()
	}
	if utils.Cut("abc", "x", "y", 0) != "abc" {
		t.Fail()
	}
	if utils.Cut("abc", "x", "y", 1) != "abc" {
		t.Fail()
	}
	if utils.Cut("abc", "x", "y", 2) != "abc" {
		t.Fail()
	}
	if utils.Cut("abc", "x", "y", 3) != "abc" {
		t.Fail()
	}
	if utils.Cut("/", "/", "/", 0) != "" {
		t.Fail()
	}
	if utils.Cut("/", "/", "/", 1) != "" {
		t.Fail()
	}
	if utils.Cut("/", "/", "/", 2) != "" {
		t.Fail()
	}
	if utils.Cut("/", "/", "/", 3) != "/" {
		t.Fail()
	}
}
