package main

import (
	"fmt"
	"log"
	"net/http"

	// this works as we have already said 'go get gorilla/mux'
	"github.com/gorilla/mux"
)

func main()  {
	fmt.Println("Welcome to mod in Go")
	greeter()
	// router from gorilla mux 
	r := mux.NewRouter()
    r.HandleFunc("/", serveHome).Methods("GET")
	// run as a server - 2 params, 1. port, 2. router 
	// handle errors here using log package and fatal method 
	log.Fatal(http.ListenAndServe(":4000", r))
}

func greeter()  {
	fmt.Println("Hey modder")
}

// 2 params - w from http package for sending response, r - us receiving a request , pointer 
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>This is a golang response</h1>"))
}