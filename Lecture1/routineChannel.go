package main

import "fmt"

func main() {

	messages := make(chan string)

	go func(msg string) {
		messages <- "ping"
		fmt.Println(msg)
	}("In func")

	//fmt.Println("Pinged message")
	msg := <-messages
	fmt.Println(msg)
}
