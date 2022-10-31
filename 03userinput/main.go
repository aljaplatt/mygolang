package main

import (
	"bufio"
	"fmt"
	"os"
)

// bufio - reader, scanner

func main() {
	welcome := "Welcome"
	fmt.Println(welcome)
	
	// use bufio to read
	// where do you want to read from? os.Stdin
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Score Pizza: ")

	// whatever the reader reads, store into variable
	//! comma ok, err ok syntax - try/catch replacement 
	// we will expect input or err - can handle err, or _ if not needed
	// reader.ReadString("\n") - read until new line 
	// input, err := reader.ReadString("\n") - cannot use double quotes
	input, _ := reader.ReadString('\n') // err not needed _
	fmt.Println("Thanks for rating", input)
	fmt.Printf("Type of this rating is %T", input)
}