package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	// "strings"
)

func main()  {
	fmt.Println("Creating a post requests in Go")
	PerformPostJsonReq()
}

func PerformPostJsonReq()  {
	const myurl = "http://localhost:8000/post"
	
	// fake json payload
	reqBody := strings.NewReader(`
		{
			"coursename" : "Golang 101",
			"price" : 0,
			"platform" : "udemy"
		}
	`)
	// 3 params - 1. url, 2. content type, 3. data to send
	response, err := http.Post(myurl, "application/json", reqBody)

	if err  != nil {
		panic(err)
	} 
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
	// {"coursename":"Golang 101","price":0,"platform":"udemy"}
}