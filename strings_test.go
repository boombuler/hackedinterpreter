package main

import (
	"./runtime"
	"testing"
)

func Test_String(t *testing.T) {
	tests := map[string]runtime.Value{
		`"\\hello\twor\"ld\r\n"`: "\\hello\twor\"ld\r\n",
		`"Hello" != "World"`:     true,
		`"Hello" == "Hello"`:     true,
		`"Hello"[0]`:             "H",
		`"Foo" + "bar"`:          "Foobar",
		`"a" < "b"`:              true,
		`"ab" > "bc"`:            false,
		`"a" + "b" > "aa"`:       true,
		`is_list("foo")`:         false,
	}
	for code, val := range tests {
		value, err := execString(code, runtime.DefaultTimeout)
		if err != nil {
			t.Errorf("%v\nCode: %v", err.Error(), code)
		} else if !runtime.Equals(value, val) {
			t.Errorf("Failed got %v expected %v\nTest:%v", value, val, code)
		}
	}
}

func Test_StringFail(t *testing.T) {
	tests := []string{
		`"Hallo"[1] = "A"`, `-"a"`, `"a"-"b"`,
	}
	for _, code := range tests {
		if val, err := execString(code, runtime.DefaultTimeout); err == nil {
			t.Errorf("Failed to fail! \nResult: %v\nCode:%v", val, code)
		}
	}
}
