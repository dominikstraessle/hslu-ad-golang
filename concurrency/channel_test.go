package main

import "testing"

func Test_channel(t *testing.T) {
	got := channel()
	if got != 499500 {
		t.Logf("Failed to compute summarize: %d", got)
		t.Fail()
	}
}
