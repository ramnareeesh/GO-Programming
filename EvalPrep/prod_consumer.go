package main

import "fmt"

func producer(ch chan<- int) {
	fmt.Println("Inside Producer...")
	for i := 0; i < 10; i++ {
		fmt.Println("Produced:", i)
		ch <- i
	}
	close(ch)
}

func consumer(ch <-chan int, done chan bool) {
	for value := range ch {
		fmt.Println("Consumed:", value)
	}
	done <- true
}

func main() {
	ch := make(chan int)
	done := make(chan bool)

	go producer(ch)
	go consumer(ch, done)

	<-done
}
