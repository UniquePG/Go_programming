package main

import "fmt"

func main() {

	a, b := 2, 4

	p1 := &a	// p1 is the pointer of a

	fmt.Println("value of pointer p1 is: ", *p1)	// we can access pointer value as *p1
	
	*p1 = 90	// set value to pointer p1

	p2 := &b	// p2 is the pointer of b

	fmt.Println("value of pointer p2 is: ", *p2)
	fmt.Println("value of pointer p1 is: ", *p1)

	// change the variable a value using pointer
	*p2 = *p2 * 5

	fmt.Println("vlaue of variable b: ", b)

}
