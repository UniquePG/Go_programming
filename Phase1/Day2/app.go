package main

import "fmt"

func main() {

//! 1
	// for i := 1; i<= 10; i++ {
	// 	fmt.Println(i * 2)
	// }

//! 2
	sum := 1
	for ; sum < 10; {		// here we omit the initial state and increment of the for syntax
		sum += sum
	}

	// for sum < 20 {		// we can also omit the semicolon(;) (for works as a while loop)
	// 	sum *=3
	// }

	fmt.Println(sum)


}