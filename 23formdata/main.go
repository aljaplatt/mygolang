package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	// "strings"
)

func main()  {
	fmt.Println("sending form data in Go")
	PerformPostFormReq()
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

func PerformPostFormReq()  {
	const myurl = "http://localhost:8000/postform"

	// formdata
	data := url.Values{}
	data.Add("firstname", "alex")
	data.Add("lastname", "platt")
	data.Add("email", "alex@gomail.com")

	// postform - url encoded post request
	response, err := http.PostForm(myurl, data)

	// handle error if exists
	if err != nil {
		panic(err)
	}
	// when everything is done close the connection
	defer response.Body.Close()

	// read response if no error 
	content, _ := ioutil.ReadAll(response.Body)
	// convert content to string and print out
	fmt.Println(string(content))
	// {"email":"alex@gomail.com","firstname":"alex","lastname":"platt"}
}