package main

import "fmt"

func main()  {
	fmt.Println("Welcome to Go arrays")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	// [2] is missing
	fruitList[3] = "Mango"

	fmt.Println("Fruit list is: ", fruitList) //  [Apple Banana  Mango]
	fmt.Println("Fruit list length is: ", len(fruitList)) //?  4, even with only 3 values

	var vegList = [3]string{"broccoli", "spinach", "carrot"}
	fmt.Println("VegList is: ", vegList)
	fmt.Println("VegList length is: ", len(vegList))
	
}