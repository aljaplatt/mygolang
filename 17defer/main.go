package main

import "fmt"

func main()  {
	// defer - do this last 
	// defer fmt.Println("world2")
	// defer fmt.Println("World")	
	// fmt.Println("Hello")
	/*
	Hello
	World
	world2
	*/ 
	defer fmt.Println("world")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello")
	myDefer()
	/* //* lifo 
	Hello
	4
	3
	2
	1
	0
	Two
	One
	world
	*/
}

func myDefer()  {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
	
}