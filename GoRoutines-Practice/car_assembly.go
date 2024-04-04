package main

import (
	"fmt"
	"sync"
	"time"
)

func engineWork(wg *sync.WaitGroup) {
	fmt.Println("Working on Engine")
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Done: Engine job")
	wg.Done()
}

func acc_brake_clutchWork(wg *sync.WaitGroup) {
	fmt.Println("Working on ABC")
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Done: ABC job")
	wg.Done()
}

func bodyWork(wg *sync.WaitGroup) {
	fmt.Println("Working on body")
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Done: Body job")
	wg.Done()
}
