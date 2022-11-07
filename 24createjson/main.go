package main

import (
	"encoding/json"
	"fmt"
)

// lowercase course, not making it public - not exporting outside this file
type course struct {
	// Name string
	// Price int
	// Platform string
	// Password string
	// Tags []string
	//? can use alias for keys 
	Name string `json:"coursename"`
	Price int	
	Platform string `json:"website"`
	// - dash will protect field, not consumable, cannot be viewed
	Password string `json:"-"`
	// omitempty - if no value = nil / null 
	//! Tags []string `json:"tags, omitempty"`
	Tags []string `json:"tags,omitempty"`
	/*
	! struct field tag `json:"tags, omitempty"` not compatible with reflect.StructTag.Get: suspicious space in struct tag value - WATCH FOR SPACE BEFORE OMITEMPTY
	*/
}

func main()  {
	fmt.Println("Welcome to json")
	EncodeJson()
}

func EncodeJson()  {
	courses := []course{
		{"ReactJS", 15, "Udemy", "password", []string{ "web-dev", "js"}},
		{"Python", 0, "youtube", "grassword", []string{ "backend", "python"}},
		{"java", 150, "coursera", "sassword", nil},
	}

	// package data as json data - using json library
	// finalJson, err := json.Marshal(courses)
	// second argument is prefix - not really needed
	//? \t tab i think?
	finalJson, err := json.MarshalIndent(courses, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
	/*
	marshal
	[{"Name":"ReactJS","Price":15,"Platform":"Udemy","Password":"password","Tags":["web-dev","js"]},{"Name":"Python","Price":0,"Platform":"youtube","Password":"grassword","Tags":["backend","python"]},{"Name":"java","Price":150,"Platform":"coursera","Password":"sassword","Tags":null}]

	marshal indent
	[
        {
                "Name": "ReactJS",
                "Price": 15,
                "Platform": "Udemy",
                "Password": "password",
                "Tags": [
                        "web-dev",
                        "js"
                ]
        },
        {
                "Name": "Python",
                "Price": 0,
                "Platform": "youtube",
                "Password": "grassword",
                "Tags": [
                        "backend",
                        "python"
                ]
        },
        {
                "Name": "java",
                "Price": 150,
                "Platform": "coursera",
                "Password": "sassword",
                "Tags": null
        }
]
	*/
}