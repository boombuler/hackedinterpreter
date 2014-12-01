package main

import (
	"./runtime"
	"testing"
	"time"
)

func Test_VarIsZeroByDefault(t *testing.T) {
	tests := map[string]runtime.Value{
		"var_a": 0,
		"input": 0,
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

func Test_RunTimeout(t *testing.T) {
	tTimeOut := 1 * time.Second
	tWaitToFail := tTimeOut + 200*time.Millisecond

	endlessLoop := "while true { var_a }"
	timer := time.NewTimer(tWaitToFail)
	hasResult := make(chan error)
	go func() {
		_, err := execString(endlessLoop, tTimeOut)
		hasResult <- err
	}()

	select {
	case <-timer.C:
		t.Error("Timeout did not happen!")
	case <-hasResult:
		t.Log("Timeout worked!")
	}
}

func Test_Comments(t *testing.T) {
	tests := map[string]runtime.Value{
		"1 // var_a = -3\r\n + 2": 3,
		"1 // var_a = -3":         1,
		"3 * /*3*/ 2":             6,
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
