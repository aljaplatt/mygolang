package main

import "fmt"

func main()  {
	fmt.Println("Maps in go")
	lang := make(map[string]string)

	lang["js"] = "javascript"
	lang["go"] = "go"
	lang["jv"] = "java"
	lang["py"] = "python"
	
	fmt.Println("List of all langs: ", lang)
	// List of all langs:  map[go:go js:javascript jv:java py:python]
	fmt.Println("js is the key for:", lang["js"]) // js is the key for: javascript
	fmt.Println("py is the key for:", lang["py"]) // js is the key for: python
	
	// delete
	delete(lang, "jv")
	fmt.Println("List of all langs: ", lang)
	// List of all langs:  map[go:go js:javascript py:python]

	//* looping maps
	for key, value := range lang {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
	// For key js, value is javascript
	// For key go, value is go
	// For key py, value is python
}