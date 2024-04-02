package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		c <- a
		a, b = b, a+b
	}
	close(c)
}

func main() {
	c := make(chan int)
	go fibonacci(10, c)
	for i := range c {
		fmt.Println(i)
	}

	c_b := make(chan int, 10)
	go fibonacci(cap(c_b), c_b) // cap() fn is used just to input no. of terms that the channel can hold
	for i := range c_b {
		fmt.Println(i)
	}
}
