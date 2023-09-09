package main

import "fmt"

func main() {

	a := [5]int{2, 3, 4, 5, 6}

	var s []int = a[1:3] // make a slice of range index 1 to 3

	fmt.Println(s)
	
	// var slice []int
	
	// slice[0] = 11
	// slice[1] = 15

	// fmt.Println("slice is: ", slice)


	// we can also make slices using make() function

	dd := make([]string, 5)		// it creates the slice of length 5
	printSlice("dd", dd);

	ff := make([]int, 0, 4)		// it crate a slice of length 0 and capacity 4
	printSlice1("ff", ff)


// append elements in the new slice
	var as []int;
	printSlice1("as", as)

	// append an new element in the slice
	as = append(as ,78 )
	printSlice1("as after append in empty slice", as)
	
	// we can append more that one element at a time
	as = append(as, 34, 44, 54, 12)
	printSlice1("as after appending multiple values", as)

	
}

// make a func that print our slice
func printSlice(s string, x []string) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
func printSlice1(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
