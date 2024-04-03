package main

import "fmt"

func main() {
	c := make(chan int, 5)

	go fib_buff(c)

	for i := range c {
		fmt.Print(i, ",")
	}

}
