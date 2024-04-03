package main

func buff_chan(c chan int) {
	size := cap(c)
	for i := 0; i < size; i++ {
		c <- i
	}
	close(c)
}
