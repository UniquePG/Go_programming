package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {

//* range is used in the for loop for iterates over the slice or map
//* range -> it returns two values first is the index, and second is the value at that index

	for i, v := range pow {					// i -> index,  v-> value
		fmt.Printf("2^%d = %d\n", i, v)
	}


//! We can skip index or value in the range if we want
	// for i, _ := range pow  --> it skips the value 
	// for _, v := range pow  --> it skips the index
	// for i := range pow  --> it skips the value

	 aa := make([]int, 5);

	 // add elements in the slice using range(for loop)
	 for i := range aa {

		aa[i] = 2*i 	
	 }
	 
	 fmt.Println("aa after adding values: ", aa)

	 // now omit the index in range
	 for _, v := range aa {
		fmt.Printf("%d \n", v)				// print all the element of the array
	 }

}