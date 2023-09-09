package main

import "fmt"

type Student struct {
	x, y int
}

// method is a function with a special 'receiver' arguement
func (s Student) multiply() int {
	return s.x * s.y // return the value of x times two
}


// we can only define methods with a receiver whoes type is defines in the same package as method
// we can not directly define int, string , float(built-in types ) in the receiver

	type Name string	// first we have to define built-in datatypes (as a variable)

func (n Name) demo () int {
	return len(n)
}

func main(){
	
	Std := Student{5, 7}
	fmt.Println("multiply", Std.multiply())


	//* built-in types
	str := Name("pavan gupta");			// now define value for built-in type before call method
	fmt.Println("built-int types: ", str.demo())
}