package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	// "strings"
)

func main()  {
	fmt.Println("Creating a get requests in Go")
	PerformGetRequest()
}

func PerformGetRequest()  {
	const myUrl = "http://localhost:8000/get"

	response, err := http.Get(myUrl)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("Status code:", response.StatusCode)
	fmt.Println("Content length is:", response.ContentLength)

	//? 2 ways to handle converting content from bytes.
	//* 1
	// var responseString strings.Builder

	// better to handle err 
	// content, _ := ioutil.ReadAll(response.Body)
	// byteCount, _ := responseString.Write(content)
	// fmt.Println("ByteCount is:", byteCount)
	// ByteCount is: 43
	// fmt.Println(responseString.String())
	// {"message":"Hello from learnCodeonline.in"}
	// responseString. - this way gives you access to some more methods

	//* 2
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
	/*
	Status code: 200
	Content length is: 43
	{"message":"Hello from learnCodeonline.in"}
	*/
	

}