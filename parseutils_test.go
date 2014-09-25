package main

import (
	"testing"
	"time"
)

func Test_RunTimeout(t *testing.T) {
	tTimeOut := 1 * time.Second
	tWaitToFail := tTimeOut + 200*time.Millisecond

	endlessLoop := "while true { var_a }"
	timer := time.NewTimer(tWaitToFail)
	hasResult := make(chan bool)
	go func() {
		_, err := ExecuteString(endlessLoop, tTimeOut)
		hasResult <- err.Error() == "timeout"
	}()

	select {
	case <-timer.C:
		t.Error("Timeout did not happen!")
	case res := <-hasResult:
		if res {
			t.Log("Timeout worked!")
		} else {
			t.Error("something went wrong in timeout test!")
		}
	}
}
