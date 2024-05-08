package main

import "fmt"

// Functions
func mul(a1, a2 int) int {

	res := a1 * a2
	fmt.Println("Result: ", res)
	return 0
}

func show() {
	fmt.Println("Hello!, GeeksforGeeks")
}

func main() {

	defer mul(23, 45)
	defer mul(23, 56) // as it is defer, it will be executed at the end
	show()
}
