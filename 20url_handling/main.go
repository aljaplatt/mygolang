package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=gsdfgsdfg09f"

func main()  {
	fmt.Println("Welcome to handling urls")
	fmt.Println(myUrl)
	
	// parsing url with url package/library
	result, _ := url.Parse(myUrl)

	fmt.Println(result.Scheme) // https
	fmt.Println(result.Host) // lco.dev:3000
	fmt.Println(result.Path) // /learn
	fmt.Println(result.Port()) // 3000
	fmt.Println(result.RawQuery) // coursename=reactjs&paymentid=gsdfgsdfg09f

	queryParams := result.Query()
	fmt.Printf("Type of queryParams: %T\n", queryParams)
	// Type of queryParams: url.Values - url.Values = key value pairs

	// we know this is an object/map/dictionary, we can access its values with a key eg: queryParams["coursename"])
	fmt.Println(queryParams["coursename"]) // [reactjs]
	fmt.Println(queryParams["paymentid"]) // [gsdfgsdfg09f]

	// we can iterate over this with the for loop and range keyword
	// do not need inx, _
	for _, value := range queryParams {
		fmt.Println("Param is:", value)
	}
	/*
	Param is: [reactjs]
	Param is: [gsdfgsdfg09f]
	*/
	//? Construct URL ??  
	//! IMPORTANT - & = you have to pass on reference, not copy !
	partsOfUrl := &url.URL{ 
	// partsOfUrl := url.URL{ - will work but problem ?
		Scheme: "https",
		Host: "lco.dev",
		Path: "/tutcss",
		RawPath: "user=hitesh",
	}

	anotherURL := partsOfUrl.String()
	fmt.Println("anotherURL:", anotherURL)
	// anotherURL: https://lco.dev/tutcss
}