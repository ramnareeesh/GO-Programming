package main

import (
	"fmt"
	"sync"
	"time"
)

func isEven(n int) bool {
	return n%2 == 0
}
func main() {
	n := 2
	var m sync.Mutex

	// now, both goroutines call m.Lock() before accessing `n`
	// and call m.Unlock once they are done
	go func() { // anonymous function
		m.Lock()
		defer m.Unlock() // defer is used for mutual exclusion
		nIsEven := isEven(n)
		time.Sleep(5 * time.Millisecond)
		if nIsEven {
			fmt.Println(n, " is even")
			return
		}
		fmt.Println(n, "is odd")
	}()

	go func() { // anonymous function
		m.Lock()
		n++
		m.Unlock()
	}()

	time.Sleep(time.Second)
}
