package main

import (
	"testing"
)

func Test_GetIndex(t *testing.T) {
	tests := map[string]Value{
		"[0, 1, 2].length":                  3,
		"[0, 1, 2][1]":                      1,
		"var_a = [1]; var_a[0] = 2; var[0]": 2,
	}
	for code, val := range tests {
		value, err := ExecuteString(code, DefaultTimeout)
		if err != nil {
			t.Error(err)
		}
		if !compare(value, val) {
			t.Errorf("Failed got %v expected %v\nTest:%v", value, val, code)
		}
	}
}
