package main

import "fmt"

func main() {

	c_engine := make(chan int)
	c_body := make(chan int)
	c_paint := make(chan int)

	go engine_Work(c_engine)
	<-c_engine

	go body_Work(c_body)
	<-c_body

	go paint_Work(c_paint)
	<-c_paint

	fmt.Println("All Jobs Done!")
}
