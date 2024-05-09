package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type User struct {
	Username string
	Password string
	Token    string
	Plan     string
}

var database []User

type API int

func (a *API) Greet(empty string, reply *string) error {
	*reply = "Hello World!"
	return nil
}

func (a *API) Register(user_data User, reply *string) error {
	for _, user := range database {
		if user.Username == user_data.Username {
			*reply = "Register failure! User already exists, try logging in."
			return nil
		}
	}
	database = append(database, user_data)
	*reply = "Registered successfully!"
	return nil
}

func (a *API) Login(user_data User, reply *string) error {
	for _, user := range database {
		if user.Username == user_data.Username {
			user_data.Token = "Bond007"
			*reply = user_data.Token
			return nil
		}
	}
	*reply = "Login failed! User does not exist."
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
