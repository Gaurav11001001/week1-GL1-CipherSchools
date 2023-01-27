package main

import "fmt"

//main function will be automatically called when you start the application
func main() {
	// var Varaiable_name datatype
	var data int
	data = 10

	data1 := 100
	Data := 1000
	//data := 1000  //redeclare the data variable
	data = 1000
	// data1 = false // wrong assignment value type
	fmt.Println(data)
	fmt.Println(data1)
	fmt.Println(Data)
	// data type
	//1. Primitive
	// byte, uint, int, float64, string, bool, complex
	//2. Non-Primitive
	// struct, map, chan, interface, array, slice
	//functional programming

	name := "rahul"
	fmt.Println(name)
}
