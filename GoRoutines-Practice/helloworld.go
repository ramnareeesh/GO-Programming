package main

import (
	"fmt"
	"time"
)

func say(msg string) {
	for i := 0; i < 5; i++ {

		time.Sleep(100 * time.Millisecond)
		fmt.Println(msg)

	}
}

func main() {
	// go say("hello")
	say("hello")
	say("world")
}
