package main

import (
	"fmt"
	"math/big"

	// "math/rand"
	// "time"
	"crypto/rand"
)

func main() {

	fmt.Println("Welcome to maths in golang")

	// var numOne int = 2
	// var numTwo float64 = 4.2  //! the .2 will be dropped by the auto conversion 

	// will auto convert float to int - not good in this case
	// fmt.Println("The sum is: ", numOne + int(numTwo))

	//* Random number using rand 
	// random numbers - math/rand & crypto/rand
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5) + 1) //? without seed you only get 1 - need seed
	//* (5) alone gives 0 - 4?

	//* Random num using crypto

	randNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(randNum)
}