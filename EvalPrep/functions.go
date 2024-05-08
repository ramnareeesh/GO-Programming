package main

import (
	"errors"
	"fmt"
)

func add(x int, y int) (int, int, int) {
	return x + y, x, y
}

func divide(x int, y int) (int, int, int, error) {
	if y == 0 {
		return x, y, 0, errors.New("cannot divide by zero")
	}
	return x, y, x / y, nil
}

func main() {
	fmt.Println("main is active...")
	var a = int(3)
	var b = int(0)

	arr1 := [4]string{"A", "B", "C", "D"}

	for i := range arr1 {
		fmt.Println(arr1[i])
	}

	ages := make(map[string]int)
	ages["Ram"] = 20
	ages["Aravi"] = 19

	_, ok := ages["Ram"]

	if !ok {
		fmt.Println("Not found!")
	} else {
		fmt.Println("Found!")
	}

	sum, val1, val2 := add(a, b)
	val3, val4, quotient, error_ := divide(a, b)

	fmt.Println("result:", val1, "+", val2, "=", sum)

	fmt.Println("result:", val3, "/", val4, "=", quotient, "\n", error_)
}
