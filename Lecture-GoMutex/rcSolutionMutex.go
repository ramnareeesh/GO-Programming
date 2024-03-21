package main

import (
	"fmt"
	"sync"
)

var x = 0

func increment(wg *sync.WaitGroup, m *sync.Mutex) { // we're passing the mutex as a pointer
	m.Lock() // here we are locking the mutex
	x = x + 1
	m.Unlock() // here we are unlocking the mutex
	wg.Done()
}
func main() {
	var w sync.WaitGroup
	var m sync.Mutex // here we are creating a mutex M
	for i := 0; i < 500; i++ {
		w.Add(1)
		go increment(&w, &m)
	}
	w.Wait()
	fmt.Println("final value of x", x)

}
