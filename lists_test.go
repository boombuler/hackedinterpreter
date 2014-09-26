package main

import (
	"testing"
)

func Test_Lists(t *testing.T) {
	tests := map[string]Value{
		"new_list(1).map(function var_a -> var_a + 2)[0]": 2,
		"[0, 1, 2].length":                                3,
		"[0, 1, 2][1]":                                    1,
		"var_a = [1]; var_a[0] = 2; var[0]":               2,
		"[[1, 2, 3]].push(1).length":                      2,
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
