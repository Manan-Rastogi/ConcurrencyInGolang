package main

import (
	"fmt"
	"sync"
)

var msgs map[string]interface{}
var wg1 sync.WaitGroup
var mux sync.Mutex

func updateMsg(s string, channel1 chan string, i string) {
	// defer wg1.Done()
	mux.Lock()
	defer mux.Unlock()
	msgs["msg"] = s
	channel1 <- s + i
	channel1 <- s + i
}

func whatisMsg() {
	delete(msgs, "msg")
}

func main() {
	msgs = map[string]interface{}{
		"msg": "Hello Manan",
	}

	channel1 := make(chan string)
	channel2 := make(chan string)
	channel3 := make(chan string)
	channel4 := make(chan string)

	// wg1.Add(3)

	go updateMsg("Hello Ambuj", channel1, "0")
	go updateMsg("Hello Vanshaj", channel2, "1")
	go updateMsg("Hello Pinnu", channel3, "2")
	go updateMsg("Hello Saieesha", channel4, "2")

	select {
	case <-channel1:
		fmt.Println(<-channel1)
	case <-channel2:
		fmt.Println(<-channel2)
	case <-channel3:
		fmt.Println(<-channel3)
	case <-channel4:
		fmt.Println(<-channel4)
	}

	// wg1.Wait()
	// x := <-channel1

}
