package main

import "fmt"

// generate function generates numbers starting from 2 and sends them to the channel.
func generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i
	}
}

// filter function filters out numbers that are divisible by the prime number.
func filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in
		if i%prime != 0 {
			out <- i
		}
	}
}

func main() {
	ch := make(chan int)
	go generate(ch)

	for i := 0; i < 100; i++ {
		prime := <-ch
		fmt.Println(prime)
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
	}
}
