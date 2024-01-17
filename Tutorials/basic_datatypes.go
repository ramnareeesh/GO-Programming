package main

import "fmt"

func main() {
	var num1 int8 = 36
	var num2 int16 = 1455

	fmt.Println(num1, num2)
	fmt.Printf("%T\n", num1)
	fmt.Printf("%T\n", num2)

	//rune := 'G' // rune is a alias for int32 which returns the unicode of the character
	//fmt.Println(rune)

	var unsigned uint = -10
	// this throws a compilation error
	// variable overflows because uint declaration
	fmt.Println(unsigned)
}
