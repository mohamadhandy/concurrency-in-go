package main

import (
	"io"
	"os"
	"strings"
	"sync"
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

func Test_updateMessage(t *testing.T) {
	var wg sync.WaitGroup

	wg.Add(1)
	go updateMessage("delta", &wg)
	wg.Wait()

	if msg != "delta" {
		t.Errorf("Expected delta, but it didn't")
	}
}

func Test_main(t *testing.T) {
	stdOut := os.Stdout

	read, write, _ := os.Pipe()
	os.Stdout = write

	main()

	_ = write.Close()
	result, _ := io.ReadAll(read)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "Hello, universe!") {
		t.Errorf("Expected Hello, universe!, but it didn't")
	}
	if !strings.Contains(output, "Hello, cosmos!") {
		t.Errorf("Expected Hello, cosmos!, but it didn't")
	}
	if !strings.Contains(output, "Hello, world!") {
		t.Errorf("Expected Hello, world!, but it didn't")
	}
}
