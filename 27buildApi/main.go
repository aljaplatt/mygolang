package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

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
	// want to manually create coure id, not rely on user to provide
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

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	// what if: body is empty
	if r.Body == nil {
		// json sends a json response 
		json.NewEncoder(w).Encode("Please send some data")
	}
	// what about - {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}

	//TODO: check only if title is duplicate
	// loop, title matches with course.coursename, JSON
	// generate unique id, convert to string
	// append course into courses slice 

	rand.Seed(time.Now().UnixNano())
	// rand.Intn(100) - type int - need to convert to string
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")

	// first - grab id from req
	params := mux.Vars(r)

	// loop, id, remove, add with my ID

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			// create new course and update with incoming info 
			var course Course
			// grab values from request body
			_ = json.NewDecoder(r.Body).Decode(&course)
			// set the course id to be the incoming id 
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	//TODO: send a response when id is not found
}
// delete course
func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	//loop, id, remove (index, index+1)
	// for a course in the courses slice, if the course id matches the id given in the request params, 
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			// TODO: send a confirm or deny response
			break
		}
	}
}