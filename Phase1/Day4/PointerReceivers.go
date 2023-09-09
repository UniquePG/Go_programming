package main

import (
	"fmt"
	"math"
)

type Numbers struct {
	x, y float64
}

// make a method for operation 
func (n Numbers) SquareRoot() float64{ 

	return math.Sqrt(n.x*n.x + n.y*n.y)
}

//! Make method to change the values of Numbers

// 1. method with a 'value receiver'-> //* It does not change the original values it simply make copy of the values and operates on them
func (n Numbers) Change1 (f float64) {		

	n.x = n.x * f;	
	n.y = n.y * f;

}

// 2. method with a 'pointer receiver'-> //* It change the original values it simply take original value and change the original values
func (n *Numbers) Change2 (f float64) {		

	n.x = n.x * f;	
	n.y = n.y * f;

}

func main() {
//! using value receiver
	n1 := Numbers{3, 4}		// give our numbers(original values)

	n1.Change1(10)	// change values using value receiver -> //* original value will not change

	fmt.Println("Result after using value receiver: ", n1.SquareRoot())	

//! using pointer receiver

	n2 := Numbers{3, 4}		// give our numbers(original values)

	n2.Change2(10)	// change values using pointer receiver	//* original values will be changed

	fmt.Println("Result after using pointer receiver: ", n2.SquareRoot())

}

