package main

func sum(numbers []int, c chan int) {
	sum := 0
	for _, i := range numbers {
		sum += i
	}

	c <- sum
}
