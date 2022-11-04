package main

import "fmt"

// main acts as an entry point for your program - do not need to call it to run program - main() - just - go run main.go
func main()  {
	fmt.Println("Welcome to functions")
	greeter()

	//! define other functions outside of main
	// func greeterTwo() { // expected expression
	// 	fmt.Println("another method")
	// }
	greeterTwo()
	
	result := adder(3, 5)
	fmt.Println("Result is: ", result) // Result is:  8

	// proRes, _ := proAdder(2,5,9,12,183)
	proRes, myMess := proAdder(2,5,9,12,183)
	// fmt.Println("Pro adder total: ", proRes) // Pro adder total:  211
	fmt.Printf(" %v Pro adder total: %v\n", myMess, proRes) 
	//  Hi pro result function Pro adder total: 211 
	// 
}


func greeter() {
	fmt.Println("Namaste from golang")
}


func greeterTwo() {
	fmt.Println("another method")
}

// if we're returning a value, we must specify the type we want to return - int in this case ...............\/
func adder(val1 int, val2 int) int {
	return val1 + val2
}

//? What if we dont know how many arguments we're going to get? 
func proAdder(values ...int) (int, string) {
	total := 0
	
	for _, val := range values {
		total += val
	}
	return total, "Hi pro result function"
}