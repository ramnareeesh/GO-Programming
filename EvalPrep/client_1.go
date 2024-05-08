package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	var reply string
	client, err := rpc.DialHTTP("tcp", "localhost:3000")

	if err != nil {
		log.Fatalf("Error in dialing!\n")
	}

	client.Call("API.Greet", "", &reply)
	fmt.Printf(reply)

}
