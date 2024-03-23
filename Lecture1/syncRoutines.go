package main

import (
	"fmt"
	"sync"
)

func foo(wg *sync.WaitGroup) {
	fmt.Println("New Routine 1")
	wg.Done()
}
func foo1(wg *sync.WaitGroup) {
	fmt.Println("New Routine 2")
	wg.Done()
}
func foo2(wg *sync.WaitGroup) {
	fmt.Println("New Routine 3")
	wg.Done()
}
func foo3(wg *sync.WaitGroup) {
	fmt.Println("New Routine 4")
	wg.Done()
}
func foo4(wg *sync.WaitGroup) {
	fmt.Println("New Routine 5")
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(5)
	go foo(&wg)
	go foo1(&wg)
	go foo2(&wg)
	go foo3(&wg)
	go foo4(&wg)
	wg.Wait()

	fmt.Println("Main Routine")

}
