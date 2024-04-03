package main

func fib_buff(c chan int) {
	x, y := 0, 1
	for i := 0; i < cap(c); i++ {
		c <- x
		x, y = y, x+y
	}

	close(c)
}
