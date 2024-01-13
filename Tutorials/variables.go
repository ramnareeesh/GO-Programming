package main

import "fmt"

func main() {
	x := "Johnny"
	var num int
	fmt.Println(x)
	fmt.Println(num)

	// this is called a varaible block -> to declare variables

	// also here we can see 2 string variables are declared in one go
	var (
		firstname, lastname string = "Ramnaresh", "Ulaganathan"
		age                 int    = 20
	)

	fmt.Println("Name:", firstname, lastname)
	fmt.Println("Age:", age)
}
