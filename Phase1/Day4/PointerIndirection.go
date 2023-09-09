package main

import "fmt"

type Number struct {
	x, y float64
}

// method with pointer receiver	-> //* It can take either value or pointer as a receiver
func (n *Number) ScaleMethod(f float64) {
	n.x = n.x * f
	n.y = n.y * f
}


// function with pointer argument -> //* it must need the pointer of Number struct
func ScaleFunc(n *Number, f float64) {
	n.x = n.x * f
	n.y = n.y * f
}

func main() {

//! Method with pointer receiver 
	n1 := Number{2, 3};		// give values to the struct as a "value"

//* n1 is a value not a pointer(but method need pointer as a receriver) but the method with the pointer receiver is called automatically by GoCompiler
	n1.ScaleMethod(10)	//! Go interpretes it like this-> (&n1).ScaleMethod(10)
		// (&n1).ScaleMethod(10)	//* we can also write this

	fmt.Println("With pointer receiver: ", n1);		// modifies original values


//! function with pointer argument
	// n2 := &Number{4, 5}	// give values to the struct as a" pointer" (n2 is the pointer of struct Number)
	// ScaleFunc(n2 , 10) // call the func with a pointer arguement

							//! OR

	n2 := Number{4, 5}		// give values to the struct as a value
	ScaleFunc(&n2, 10)		// must need pointer as argument
		// ScaleFunc(n2, 10)		//! ERROR

	fmt.Println("With pointer argument: ", n2)		// modifies original values
}