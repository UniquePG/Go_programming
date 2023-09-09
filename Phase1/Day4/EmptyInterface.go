package main

import "fmt"

func main() {

	// empty interface -> it can take any value
	var x interface{}

	x = 43 //here interface is of int
	describe(x)

	x = "this is string"
	describe(x)

	x = false
	describe(x)

}

func describe(i interface{}) {

	fmt.Printf("(%v, %T) \n", i, i) // print value and type of the interface

}
