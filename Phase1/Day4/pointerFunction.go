package main

import (
	"fmt"
	"math"
)

type Numbers struct {
	x, y float64
}

// make a method for operation 
func  SquareRoot(n Numbers) float64{ 

	return math.Sqrt(n.x*n.x + n.y*n.y)
}

//! Make method to change the values of Numbers

// 1. method with a 'value argument '-> //* It does not change the original values it simply make copy of the values and operates on them
func Change1 (n Numbers ,f float64) {		

	n.x = n.x * f;	
	n.y = n.y * f;

}

// 2. method with a 'pointer argument '-> //* It change the original values it simply take original value and change the original values
func Change2 (n *Numbers, f float64) {		

	n.x = n.x * f;	
	n.y = n.y * f;

}

func main() {
//! using value argument 
	n1 := Numbers{3, 4}		// give our numbers(original values)

	Change1( n1, 10 )	// change values using value argument  -> //* original value will not change

	fmt.Println("Result after using value argument : ", SquareRoot(n1))	

//! using pointe argument 

	n2 := Numbers{3, 4}		// give our numbers(original values)

	Change2( &n2, 10 )	// change values using pointer argument 	//* original values will be changed

	fmt.Println("Result after using pointer argument : ", SquareRoot(n2))

}

