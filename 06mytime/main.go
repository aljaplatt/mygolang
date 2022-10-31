package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to Time in Golang")

	presentTime := time.Now()
	fmt.Println(presentTime) // 2022-10-31 09:03:30.229348 +0000 GMT m=+0.000106167
	// .Format("01-02-2006") - we use this as standard 
	fmt.Println(presentTime.Format("01-02-2006")) // 10-31-2022
	// have to use monday - to get what ever the current day is - weird but the way it is
	fmt.Println(presentTime.Format("01-02-2006 Monday")) // 10-31-2022 Monday
	// to get time - have to use 15:04:05
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday")) // 10-31-2022 09:13:00 Monday

	createdDate := time.Date(2020, time.September, 10, 23, 12, 0, 0, time.UTC )
	fmt.Println("CD: ",createdDate) // 2020-09-10 23:12:00 +0000 UTC
	fmt.Println(createdDate.Format("01-02-2006 Monday")) // 09-10-2020 Thursday
}