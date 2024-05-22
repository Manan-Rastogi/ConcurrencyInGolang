package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var msg string

// go run -race .
func updateMessage(s string, mux *sync.Mutex) {
	defer wg.Done()
	// mux.Lock()
	msg = s
	// mux.Unlock()
}

func main() {
	msg = "Hello Pluto!"

	var mux sync.Mutex

	wg.Add(2)
	go updateMessage("Hello Saturn!", &mux)
	go updateMessage("Hello Neptune!", &mux)
	wg.Wait()

	fmt.Printf("msg: %v\n", msg)

	// money()
}
