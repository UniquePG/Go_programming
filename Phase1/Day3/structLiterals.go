package main

import "fmt"

type structure struct {
	i, j, k int
}

var (
	v1 = structure{5, 10, 15}    // give all values
	v2 = structure{i: 25}        // here j: 0 & k: 0
	v3 = structure{j: 32, k: 35} // here i: 0
	v4 = structure{}             // i=0, j: 0, k: 0

	p = &structure{51, 52, 53} // pointer of struct
)

func main() {

	fmt.Println("value when give all value: ", v1)
	fmt.Println("value when give 1 value: ", v2)
	fmt.Println("value when give 2 value: ", v3)
	fmt.Println("value when give NO value: ", v4)
	fmt.Println("value of point struct: ", *p)

}
