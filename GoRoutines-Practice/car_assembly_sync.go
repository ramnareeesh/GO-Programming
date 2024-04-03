package main

import (
	"fmt"
	"time"
)

func engine_Work(engineChannel chan int) {
	fmt.Println("Working: Engine Job")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Done: Engine Job")
	engineChannel <- 1
}

func body_Work(bodyChannel chan int) {
	fmt.Println("Working: Body Job")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Done: Engine Job")
	bodyChannel <- 1
}

func paint_Work(paintChannel chan int) {
	fmt.Println("WorkingL: Paint Job")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Done: Paint Job")
	paintChannel <- 1
}
