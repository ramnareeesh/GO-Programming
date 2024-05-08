package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type API int

func (a *API) Greet(empty string, reply *string) error {
	*reply = "Hello World!"
	return nil
}

func main() {
	api := new(API)
	err := rpc.Register(api)

	if err != nil {
		log.Fatalf("Error in registering!\n")
	}

	rpc.HandleHTTP()

	port := ":3000"

	listening, err := net.Listen("tcp", ":3000")

	if err != nil {
		log.Fatalf("Error in listening!\n")
	}

	log.Printf("RPC is listening in port%s.\n\n", port)

	err = http.Serve(listening, nil)
	if err != nil {
		log.Fatalf("Error in serving!\n")
	}

}
