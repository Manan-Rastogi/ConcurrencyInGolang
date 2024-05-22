package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func Test_printIt(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()

	os.Stdout = w

	var wg sync.WaitGroup
	
	wg.Add(1)
	go printIt1(1, "One", &wg)
	wg.Wait()

	w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "One"){
		t.Errorf("expected to find One, but it is not present")
	}

}