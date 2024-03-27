package main

import (
	"fmt"
)

func conditional() (z int) {
	if x, y := 1, 2; x < y {
		fmt.Println("x is less than y")
	}
	z = 10
	return
}

type messageToSend struct {
	phoneNumber int
	message     string
}

type sender struct { // example of an embedded struct
	user
	rateLimit int
}

type user struct {
	name   string
	number int
}

type authenticationInfo struct {
	username string
	password string
}

func getBasicAuth(a authenticationInfo) string { // this is just another fn with struct as input
	return a.username + ":" + a.password
}

func (a authenticationInfo) getBasicAuth() string { // this is a method of the struct
	return a.username + ":" + a.password
}

func getMonthlyPrice(tier string) int {
	// ?
	if tier == "basic" {
		return 100
	} else if tier == "premium" {
		return 150
	} else if tier == "enterprise" {
		return 500
	}
	return 0
}

// main function

func main() {
	fmt.Printf("%d\n", conditional())
	fmt.Printf("Dollars : %d\n", getMonthlyPrice("pre"))

	// Create a new sender
	s := sender{
		user: user{
			name:   "John",
			number: 1234567890,
		},
		rateLimit: 100,
	}

	fmt.Printf("Sender name: %v\n", s.user)

	auth := authenticationInfo{
		username: "admin",
		password: "password",
	}

	fmt.Printf("Authorization basic: %s\n", getBasicAuth(auth))
	fmt.Printf("Authorization basic: %s\n", auth.getBasicAuth())
}
