package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name string `json:"coursename"`
	Price int	
	Platform string `json:"website"`
	Password string `json:"-"`
	Tags []string `json:"tags,omitempty"`
}

func main()  {
	fmt.Println("consume or decode json data")
		DecodeJson()
}

func DecodeJson()  {
	// consume data from web as json, comes in as bytes
	jsonDataFromWeb := []byte(`
	{
		"coursename": "ReactJS",
		"Price": 15,
		"website": "Udemy",
		"tags": ["web-dev","js"]
	}
	`)

	var currentCourse course

	// check if json is valid
	checkValidJson := json.Valid(jsonDataFromWeb)

	if checkValidJson {
		fmt.Println("JSON is valid")
		json.Unmarshal(jsonDataFromWeb, &currentCourse)
		fmt.Printf("%#v\n", currentCourse)
	} else {
		fmt.Println("JSON NOT VALID")
	}
		/*
		JSON is valid
		main.course{Name:"ReactJS", Price:15, Platform:"Udemy", Password:"", Tags:[]string{"web-dev", "js"}}
		*/
		
		//* case: just want to add data to key:value pair
		// not sure what return data is - use interface 
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	// map[string]interface {}{"Price":15, "coursename":"ReactJS", "tags":[]interface {}{"web-dev", "js"}, "website":"Udemy"}

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and type is %T\n", k, v, v)
	}
	/*
	Key is coursename and value is ReactJS and type is string
	Key is Price and value is 15 and type is float64
	Key is website and value is Udemy and type is string
	Key is tags and value is [web-dev js] and type is []interface {}
	*/
}