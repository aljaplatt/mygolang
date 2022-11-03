package main

import "fmt"

func main()  {
	fmt.Println("Structs in Go")
	// no inheritance in golang - no super, or parent
	// idea is to simplify understanding - no inheriting from other files etc

	alex := User{"Alex", "ap@gmail.com", true, 25}
	fmt.Println(alex) 
	// {Alex ap@gmail.com true 25}
	fmt.Printf("Alex details are: %+v\n", alex)
	// {Name:Alex Email:ap@gmail.com Status:true Age:25}
	fmt.Printf("Name is %v and email is %v\n", alex.Name, alex.Email)
}

type User struct {
	Name 	string
	Email 	string
	Status 	bool
	Age		int
}