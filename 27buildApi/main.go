package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Model for course - file
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	// *Author - not a direct reference - signifies that the type is an Author pointer
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

//fake DB - slice of course structs 
var courses []Course

// middleware, helper - check if passed in course has a coursename
// own file
func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}
func main()  {
	fmt.Println("Building an API with Go ")
	
}

//controllers - own file
// serve homeroute - reader - get value from request,- writer - response
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API by LearnCodeOnline</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	// set headers
	w.Header().Set("Content-Type", "application/json")
	// encode courses to json
	// new encoder works with the writer - w
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "applicatioan/json")

	// grab id/params values from request - vars = variables extracted from request into 'object'
	params := mux.Vars(r)
	fmt.Println("PARAMS: ", params)

	// loop through courses, find matching id and return the response
	for _, course := range courses {
		// find course with matching id to request id in params
		if course.CourseId == params["id"] {
			// if found, encode it to json 
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course found with given id")
	return // optional ?
}