package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 10; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	go f("goroutine")
	f("direct")
	go func(msg string) { // nameless function
		fmt.Println(msg)
	}("going") // implicit call for nameless functions

	time.Sleep(time.Second)
	fmt.Println("done")
}
