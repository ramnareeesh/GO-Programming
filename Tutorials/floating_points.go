package main

import "fmt"
import "math"

func main() {
	var num float32 = 3.0999
	fmt.Println(num)

	fmt.Println("Maximum value for float 64", math.MaxFloat64)
	// this is imported from the math library to show the max value of a fp number

	var scientific float64 = 6.02e3
	fmt.Println("Scientific representation using floating point:", scientific)
}
