package main

import (
	"fmt"
	"net/http"
	"sync"
	// "time"
)

var signals = []string{"test"}

var wg sync.WaitGroup //pointer
var mut sync.Mutex    // pointer

func main()  {
	fmt.Println("Welcome to goroutines")
	//*1
	// greeter("Hello")
	// greeter("world")
	//* 2 - create goroutine using 'go'
	//? fire up a thread to be responsible for first greeter but only second greeter call prints! why? 
	// go greeter("Hello")
	// greeter("world")

	websitelist := []string{
		"https://go.dev",
		"https://google.com",
		"https://bbc.co.uk",
		"https://github.com",
	}

	for _, web := range websitelist {
		// loop thru and grab status code for each site
		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait()
	fmt.Println(signals)
}

func greeter(s string) {
	// log.Println("sup")
	for i := 0; i < 6; i++ {
		// this will work but is not good practice
		// time.Sleep(3 * time.Millisecond)
		fmt.Printf("Str%v: %v\n", i, s)
		
	}
}

/**
//* 1
Str0: Hello
Str1: Hello
Str2: Hello
Str3: Hello
Str4: Hello
Str5: Hello
Str0: world
Str1: world
Str2: world
Str3: world
Str4: world
Str5: world
*/

// makes request to website, gets status code
func getStatusCode(ep string)  {
	// ep = endpoint 
	res, err := http.Get(ep)
	if err != nil {
		fmt.Println("Problem in endpoint")
	} else {
		mut.Lock()
		signals = append(signals, ep)
		mut.Unlock()
		
		fmt.Printf("%d status code for %s", res.StatusCode, ep)
	}
}