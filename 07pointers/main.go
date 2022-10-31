package main

import "fmt"

func main()  {
	fmt.Println("Welcome to pointers")
	
	//? Default value is nil
	// var ptr *int
	// fmt.Println("Value of pointer is ", ptr) // Value of pointer is  <nil>
	
	myNumber := 23
	// & creating a pointer to reference another variable
	var ptr = &myNumber
	fmt.Println(&myNumber)
	fmt.Println("myNumber memory address: ", ptr) //0x14000016080
	fmt.Println("Value the pointer is referencing:", *ptr) // 23

	//* we can manipulate the value of myNumber using the pointer...
	//? action will be performed on the actual values, not a copy
	*ptr = *ptr + 10
	fmt.Println("New value: ", myNumber) //? 33 
}