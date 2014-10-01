package main

import (
	"./runtime"
	"testing"
	"time"
)

func Test_VarIsZeroByDefault(t *testing.T) {
	value, err := execString("var_a", runtime.DefaultTimeout)
	if err != nil {
		t.Error(err)
	}
	if value != 0 {
		t.Errorf("Failed got %v expected %v", value, 0)
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
