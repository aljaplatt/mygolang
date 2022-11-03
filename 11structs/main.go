package main

import "fmt"

func main()  {
	fmt.Println("Structs in Go")
	// no inheritance in golang - no super, or parent
	// idea is to simplify understanding - no inheriting from other files etc
}

type User struct {
	Name string
	Email string
	Status bool
	Age	int
}