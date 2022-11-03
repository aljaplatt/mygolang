package main

import "fmt"

func main()  {
	fmt.Println("If else in Go")

	loginCount := 23

	var result string 

	if loginCount < 10 {
		result = "peasant user"
	} else if loginCount  > 10 {
		result = "Baller user"
	} else {
		result = "meh user"
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Num is even")
	} else {
		fmt.Println("Num is odd")
	}

	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is not less than 10")
	}

	// if err != nil {

	// }
}