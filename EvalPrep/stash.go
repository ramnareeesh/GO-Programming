package main

import (
	"fmt"
	"sync"
	"time"
)

func say(s string, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go say("world", &wg)
	go say("hello", &wg)

	wg.Wait()
}
