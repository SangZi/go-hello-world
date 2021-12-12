package main

import (
	"fmt"
	"html"
	"net/http"
)

//There’s no key-word class in Go
func main() {
	/*******************
	Example of printing hello world in console
	********************/
	var x string = "hello" //two ways of declaration and initialize variable
	y := "world"

	/*******************
	Example of package http
	Visiting in browser http://localhost:3030/foo
	You will see following message in browser:
	Hello, "/foo"
	********************/
	http.HandleFunc("/foo", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "Hello, %q", html.EscapeString(r.URL.Path))
	})
	http.ListenAndServe(":3030", nil) //handler is nil

	/*******************
	Example of calling function
	********************/
	fmt.Println("My greeting: " + x + " " + y + "!")
	fmt.Println("call function sum(), result is", sum(10, 20))

	/*******************
	Example of slice and for-loop
	Slice is a flexible and extensible data structure to implement and manage collections of data.
	Slices are made up of multiple elements, all of same type.
	********************/
	mySlice := []string{"one", "two", "three"}
	mySlice = append(mySlice, "four")
	for i, v := range mySlice { //range is a key-word
		//index starts with 0. In Println you can use comma "," to separate multiple args with space
		fmt.Println("index is", i, "and value is", v)
	}

	/*******************
	Example of map
	********************/
	myMap := make(map[int]string) //use make(map[<KEY-TYPE>]<VALUE-TYPE>) to declare a map
	myMap[100] = "100th"
	myMap[200] = "200th"
	for k, v := range myMap {
		fmt.Println("Key is", k, "and value is", v)
	}
}

//method signature, return type comes after method name; it's also allowed to have multiple return type, separate them with comma
func sum(a int, b int) int {
	return a + b
}
