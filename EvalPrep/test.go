package main

import "fmt"

type TokenM struct {
	Token bool
	LN    []int
	Queue []int
	Owner int
}

func main() {
	token := TokenM{
		Token: false,
		LN:    nil,
		Queue: nil,
		Owner: -1,
	}

	fmt.Println(token)
}
