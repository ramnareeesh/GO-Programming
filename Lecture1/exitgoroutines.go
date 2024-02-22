package main

import (
	"fmt"
	"time"
)

func main() {
	go fmt.Println("Go Routine") // the keyword go mentions that it is a Go routine
	// a separate thread will be created to execute the Println("Go Routine") call
	time.Sleep(time.Second)
	// when a sleep of some time is given, both the thread has the chance of getting executed.
	fmt.Println("Main Routine")
}
