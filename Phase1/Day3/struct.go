package main

import "fmt"

type structure struct {
	i     int
	value string
}

type structure2 struct {
	j int
	k int
}

func main() {

	fmt.Println(structure{5, "this is the value"})

	v := structure{10, "this is string"}

	v.i = 25
	v.value = "changed string"

	fmt.Println(v.i)
	fmt.Println(v.value)


//! struct with pointer
	ab := structure2{3, 6}

	p := &ab
	p.j = 30

	fmt.Println("structure 2 is: ", ab)

}
