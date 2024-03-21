package main

import (
	"fmt"
	"sync"
)

var x = 0

func increment(wg *sync.WaitGroup) {
	x = x + 1
	wg.Done()
}
func main() {
	var w sync.WaitGroup       // WaitGroup is a struct type
	for i := 0; i < 500; i++ { // means that there are 500 go routines
		w.Add(1)
		go increment(&w)
	}
	w.Wait()
	fmt.Println("final value of x", x)
}
