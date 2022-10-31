package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our Pizza app")
	fmt.Println("Please rate our Pizza between 1 and 5:")

	reader := bufio.NewReader(os.Stdin)
	// this returns 2 values - so you have to account for that 
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)
 
	//! numRating = input + 1 // input is string
	//! invalid operation: input + 1 (mismatched types string and untyped int)
	//? can convert using strconv - req 2 values, value to be converted and the bit size to be converted to
	// numRating = strconv.ParseFloat(input, 64) //! cannot assign 2 values to 1 variables
	//* Convert string, parse input  
	// numRating, err = strconv.ParseFloat(input, 64) // must use walrus or undeclared err 
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64) 

	//? if err is not empty, do this
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to your rating: ", numRating+1)
	}
	//! strconv.ParseFloat: parsing "2\n": invalid syntax 
	//? when we parse the input, their is a trailing \n that happens when the user hits enter. To fix this we use the strings package, and the TrimSpace feature
}