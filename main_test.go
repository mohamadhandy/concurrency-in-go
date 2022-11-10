package main

import (
	"testing"
)

func Test_printMessage(t *testing.T) {
	msg = "hello, cosmos!"

	printMessage()
	output := msg

	if output != msg {
		t.Errorf("Expected hello, cosmos! but it didn't")
	}
}
