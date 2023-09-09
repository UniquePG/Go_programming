package main

import "fmt"

// make structure
type Student struct {
	name  string
	phone int
}

func main() {

	// make a map using struct
	studMap := map[int]Student{
		1: {"John", 54545353},			//* key: {struct-Name, struct-phone}
		2: {"Poul", 42343434},
		3: {"Gong", 98765432},
	}

	// now iterate map using for range loop
	for i, v := range studMap {

		fmt.Printf("key: %d, Name: %v, Number: %d \n", i, v.name, v.phone)
	}

}
