package main

import "fmt"

func main()  {
	fmt.Println("Methods in Go")
	// no inheritance in golang - no super, or parent
	// idea is to simplify understanding - no inheriting from other files etc

	alex := User{"Alex", "ap@gmail.com", true, 25}
	fmt.Println(alex) 
	// {Alex ap@gmail.com true 25}
	fmt.Printf("Alex details are: %+v\n", alex)
	// {Name:Alex Email:ap@gmail.com Status:true Age:25}
	fmt.Printf("15: Name is %v and email is %v\n", alex.Name, alex.Email)
	alex.GetStatus() // Is user active:  tru
	alex.NewMail() // Email of this user is:  test@go.dev
	//! original value not changed - only copy
	fmt.Printf("18: Name is %v and email is %v\n", alex.Name, alex.Email)
}

type User struct {
	Name 	string
	Email 	string
	Status 	bool
	Age		int

}

//?  method is missing receiver - u User
func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is: ", u.Email)
	//? Does this change the actual property or not??
	//! No its not - its updating a copy, the original is not passed in pointers etc - not changes og value
}