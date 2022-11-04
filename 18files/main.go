package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main()  {
	fmt.Println("Welcome to files")
	content := "This needs to go in a file - https://go.dev/tour/list"

	// create file in this directory ./ - os package
	file, err := os.Create("./mygofile.txt")

	if err != nil {
		// panic will shut down program and show error 
		panic(err)
	}

	//write to the file - io package 
	// writestring 2 args - what file, what string 
	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("length is:", length)

	//close file
	defer file.Close()

	readFile("./mygofile.txt")
}

func readFile(filename string){
	// the  file we read is not a string we get back, its in bytes
	databyte, err := ioutil.ReadFile(filename)
	// if err != nil {
	// 	panic(err)
	// }
	checkNilErr(err)
	fmt.Println("Text data in file is: \n", string(databyte))
	// databyte is raw data, result below - we can wrap in string - 
	// string(databyte) - This needs to go in a file - https://go.dev/tour/list
	/*
	//! RAW - Text data in file is: 
 [84 104 105 115 32 110 101 101 100 115 32 116 111 32 103 111 32 105 110 32 97 32 102 105 108 101 32 45 32 104 116 116 112 115 58 47 47 103 111 46 100 101 118 47 116 111 117 114 47 108 105 115 116]
	*/
}

// we're being quite repetitive with the err checking - function instead
func checkNilErr(err error){
	if err != nil {
		panic(err)
	}
}