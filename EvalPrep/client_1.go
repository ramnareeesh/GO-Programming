package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type User struct {
	Username string
	Password string
	Token    string
	Plan     string
}

func main() {
	var reply string

	client, err := rpc.DialHTTP("tcp", "localhost:3000")

	if err != nil {
		log.Fatalf("Error in dialing!\n")
	}

	client.Call("API.Greet", "", &reply)
	fmt.Println(reply)

	user := User{
		Username: "ramnaresh",
		Password: "Naresh",
		Token:    "",
		Plan:     "",
	}
	err = client.Call("API.Register", user, &reply)
	if err != nil {
		fmt.Println("Error1:", err)
	}
	fmt.Println(reply)

	err = client.Call("API.Register", user, &reply)
	if err != nil {
		fmt.Println("Error2:", err)
	}
	fmt.Println(reply)

	err = client.Call("API.Login", user, &reply)
	if err != nil {
		fmt.Println("Error3:", err)
	}
	fmt.Println("Your token:", reply)

}
