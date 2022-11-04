package main

import "fmt"

func main()  {
	fmt.Println("Loops in Go")

	days := []string{"sunday", "tuesday", "wednesday", "friday", "saturday"}
	fmt.Println(days) // [sunday tuesday wednesday friday saturday]
	
	//* 1. 
	// for d:=0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }
	/*
		sunday
		tuesday
		wednesday
		friday
		saturday
		*/

	//* 2. - same result as 1
	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	//* 3. for each type loop - using Printf - USEFUL
	// for index, value := range days {
	// 	fmt.Printf("index is %v and value is %v\n", index, value)
	// 	// index is 0 and value is sunday
	// }
	// if you dont need the index
	// for _, value := range days {
	// 	fmt.Printf("index is not needed and value is %v\n", value)
	// 	// index is not needed and value is sunday
	// }

	//* 4. like a while loop 
	// someValue := 1

	// for someValue < 10 {
	// 	fmt.Println("Value is: ", someValue)
	// 	someValue++
	// }
	/*
	Value is:  1
	Value is:  2
	Value is:  3
	Value is:  4
	Value is:  5
	Value is:  6
	Value is:  7
	Value is:  8
	Value is:  9
	*/

	//* we can use break continue as well 
	// someValue := 1

	// for someValue < 10 {
	// 	if someValue == 3 {
	// 		break
	// 	}
	// 	fmt.Println("Value is: ", someValue)
	// 	someValue++
	// }
	/*
	Value is:  1
	Value is:  2
	*/
	//? skip item using continue
	someValue := 1

	for someValue < 5 {
		if someValue == 3 {
			someValue++
			continue
		}
		fmt.Println("Value is: ", someValue)
		someValue++
	}
	/*
	Value is:  1
	Value is:  2
	Value is:  4
	*/
}