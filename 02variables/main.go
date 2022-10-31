package main

import "fmt"

// first letter caps means public variable 
const LoginToken string = "ghabbhhjd" // Public


func main()  {
	// if you create a variable but do not use it, go will throw an error
	var username string = "alex"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)
	
	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)
	
	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)
	
	var smallFloat float32 = 255.45553453
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	// default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	// implicit type

	var website = "learncodeonline.in"
	fmt.Println(website)

	// no var style
	// walrus operator only inside function 
	numberOfUser := 300000.0
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)
}