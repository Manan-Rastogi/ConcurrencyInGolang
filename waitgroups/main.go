package main

import (
	"fmt"
	"sync"
	"time"
)

var words = []string{
	"Zero",
	"One",
	"Two",
	"Three",
	"Four",
	"Five",
	"Six",
	"Seven",
	"Eight",
}

// Normal printIt
func printIt(i int, s string){
	fmt.Printf("%v: %v\n", i, s)
}

// printIt with waitgroup
func printIt1(i int, s string, wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Printf("%v: %v\n", i, s)
}


func normalPrint(){
	for i := 0; i < len(words); i++ {
		printIt(i, words[i])
	}
}

// Using sleep is a very bad way to handle routine.
func sleepRoutine() {
	for i := 0; i < len(words); i++ {
		go printIt(i, words[i])
	}

	time.Sleep(1 * time.Second)
}

// Using waitgroups is a simple way to deal with goroutines
func waitRoutine(){
	var wg sync.WaitGroup
	for i := 0; i < len(words); i++ {
		wg.Add(1)
		go printIt1(i, words[i], &wg)
	}
	wg.Wait()
}


func main() {
	
	normalPrint()

	fmt.Println("///////////////// With Time Sleep /////////////////")

	sleepRoutine()

	fmt.Println(`\\\\\\\\\\\\\\\\ Wait Groups \\\\\\\\\\\\\\\\\`)

	waitRoutine()
}