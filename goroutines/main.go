package main

import (
	"fmt"
	"net/http"
	"sync"
	// "time"
)

//mutex video 
// mutual exclusion lock
// lock memory until this one thing is working
// lock unlock - can also read/write 
var signals = []string{"test"}

// variable for waitgroup 
var wg sync.WaitGroup // creates a pointer
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
		// need to wait in main, use sync package
		go getStatusCode(web)
		// when the goroutine is created, add it to the waitgroup
		// 1 - 1 goroutine 
		wg.Add(1)
	}

	// wait - pause main method until waigroup tasks have finished 
	// this will know when the done signal comes from getStatusCode()
	wg.Wait()
	fmt.Println(signals)
}

// func greeter(s string) {
// 	// log.Println("sup")
// 	for i := 0; i < 6; i++ {
// 		// this will work but is not good practice
// 		// time.Sleep(3 * time.Millisecond)
// 		fmt.Printf("Str%v: %v\n", i, s)
		
// 	}
// }

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
func getStatusCode(endpoint string)  {
	// when everything is finished pass on the signal that it is done.
	defer wg.Done()
	// endpoint = endpoint 
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("Problem in endpoint")
	} else {
		// mutex lock 
		mut.Lock()
		// mutex - add endpoint to signals slice
		signals = append(signals, endpoint)
		// when done, unlock 
		mut.Unlock()
		
		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}
}