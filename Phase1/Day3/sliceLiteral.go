package main

import "fmt"

func main() {

	arr := []int{3, 4, 5, 6, 7}
	fmt.Println("array is: ", arr)

	s := []struct {
		i int
		j string
	}{
		{2, "string1"},
		{5, "string2"},
		{7, "string3"},
		{9, "string4"},
	}


	fmt.Println("slice of structure (literal)", s)


	// slice bound
	sb := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(sb)

	a := s[1:4]
	fmt.Println(a)

	b := s[:2]
	fmt.Println(b)

	c := s[1:]
	fmt.Println(c)

}
