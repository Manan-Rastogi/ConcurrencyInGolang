package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func Test_printMessage(t *testing.T){
	stdOut := os.Stdout

	r, w, _ := os.Pipe()

	os.Stdout = w

	
	msg = `updated`
	printMessage()
	
	w.Close()

	result, _ := io.ReadAll(r)

	resp := string(result)

	os.Stdout = stdOut

	if !strings.Contains(resp, "updated"){
		t.Errorf("Updated not found!!")
	}

}


func Test_updateMessage(t *testing.T){
	var wg sync.WaitGroup

	wg.Add(1)
	updateMessage("updated", &wg)
	wg.Wait()

	if msg != "updated"{
		t.Error("updated not found!!")
	}

}

func TestMain(t *testing.T){
	stdOut := os.Stdout

	r, w, _ := os.Pipe()

	os.Stdout = w

	main()

	w.Close()

	os.Stdout = stdOut
	result, _ := io.ReadAll(r)
	output := string(result)

	if !strings.Contains(output, "Hello, universe!"){
		t.Errorf("universe not present")
	}

	if !strings.Contains(output, "Hello, cosmos!"){
		t.Errorf("cosmos not present")
	}

	if !strings.Contains(output, "Hello, world!"){
		t.Errorf("world not present")
	}

}