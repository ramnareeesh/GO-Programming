package main

import "fmt"

func main() {
	arr1 := [...]int{100, 200, 300, 400, 500}
	arr2 := arr1
	arr1[0] = 500
	fmt.Println(arr1 == arr2)
	fmt.Println(arr1)
	fmt.Println(arr2)
}
