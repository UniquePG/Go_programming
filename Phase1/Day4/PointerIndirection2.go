package main

import "fmt"

type Number struct {
	x, y float64
}

// method with value receiver	-> //* It can take either value or Pointer as a receiver
func (n Number) ScaleMethod(f float64) {
	n.x = n.x * f
	n.y = n.y * f
}


// function with value argument -> //* it must need the value of Number struct
func ScaleFunc(n Number, f float64) {
	n.x = n.x * f
	n.y = n.y * f
}

func main() {

//! method with value argument
	n1 := Number{2, 3};		// give values to the struct as a "value"

	n1.ScaleMethod(10)	

	//* we can also write this( it is also valid-> althrough method receiver need only value but it cal also take a pointer) 
	// (&n1).ScaleMethod(10)	//! Go interprets it like this-> n1.ScaleMethod(10)	

	fmt.Println("With value receiver: ", n1);		// NOT modifies original values


//! function with value argument
	n2 := Number{4, 5}		// give values to the struct as a value
	ScaleFunc(n2, 10)		// must need value as argument
		// ScaleFunc(&n2, 10)		//! ERROR

	fmt.Println("With value argument: ", n2)		// NOT modifies original values
}