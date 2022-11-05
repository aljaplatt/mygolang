package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const URL = "https://lco.dev"

func main()  {
	fmt.Println("Welcome to web requests")

	// HTTP GET REQUEST from net/http package - need to handle err 
	response, err := http.Get(URL)
	
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of type: %T\n", response)
	//? Response is of type: *http.Response - pointer

	//! response has a public object - Body, we must close this connection.
	defer response.Body.Close()

	// read the response body using ioutils - 
	data, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}
	// convert data into content 
	content := string(data)
	fmt.Println(content)

}