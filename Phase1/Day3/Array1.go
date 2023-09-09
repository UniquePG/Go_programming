package main

import "fmt"

func main() {

//! declair and initialize array values
	var arr [3]int // declair array of 5 elements of int

	arr[0] = 12
	arr[1] = 13
	arr[2] = 14

	fmt.Println("array1: ", arr)
	fmt.Println("values are array1: ", arr[0], arr[2], arr[1])

//! declair and initilize at the same line
	// arr1 := [5]int{5,7,9,11,13}		
	arr1 := [...]int{5,7,9,11,13,12,11,3} 	

	arr1[3] = 15	// change the value of index 3 

	fmt.Println("arr1 is: ", arr1)
	fmt.Println("value is: ", arr1[3])

}
