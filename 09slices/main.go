package main

import (
	"fmt"
	"sort"
)

func main()  {
	fmt.Println("Welcome to Go slices")

	var fruitList = []string{"Orange", "Kiwi", "Date"}
	fmt.Printf("Type of fruitList is: %T\n", fruitList)
	// Type of fruitList is: []string

	//* Add to slice 
	fruitList = append(fruitList, "Pineapple", "Banana")
	fmt.Println(fruitList) // [Orange Kiwi Date Pineapple Banana]

	//* Slice a slice
	fruitList = append(fruitList[1:]) 
	fmt.Println(fruitList) // [Kiwi Date Pineapple Banana]
	
	fruitList = append(fruitList[1:3]) 
	fmt.Println(fruitList) // [Date Pineapple]
	
	fruitList = append(fruitList[:1]) 
	fmt.Println(fruitList) // [Date]

	//* Use make() to initialise an array
	highScores := make([]int, 4)

	highScores[0] = 141
	highScores[1] = 942
	highScores[2] = 243
	highScores[3] = 544
	// highScores[4] = 429 //! panic: runtime error: index out of range [4] with length 4
	
	//? HOWEVER, this works...
	highScores = append(highScores, 101, 303, 999) // [141 942 243 544 101 303 999]

	fmt.Println(highScores) // [141 942 243 544]

	//* Slices
	fmt.Println(sort.IntsAreSorted(highScores)) // false
	sort.Ints(highScores)
	fmt.Println(highScores) // [101 141 243 303 544 942 999]
	fmt.Println(sort.IntsAreSorted(highScores)) // true

	//* remove value from slice based on index
	var courses = []string{"go", "python", "javascript", "java"}
	fmt.Println("courses: ", courses) // [go python javascript java]
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println("updated courses: ",courses) // [go python java]
}
