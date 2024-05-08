package main

import (
	"fmt"
)

func first(ch chan<- int) {
	fmt.Println("First")
	ch <- 1
}

func second(ch1 <-chan int, ch2 chan<- int) {
	<-ch1
	fmt.Println("Second")
	ch2 <- 1
}

func third(ch <-chan int, chDone chan<- bool) {
	<-ch
	fmt.Println("Third")
	chDone <- true
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	chDone := make(chan bool)

	go third(ch2, chDone)
	go first(ch1)
	go second(ch1, ch2)

	<-chDone
}
