package main

import ( "fmt" )

// function declaration
func sub(x int, y int) int{

	return x - y
}

func dividee(x , y int) int{

	return x / y
}

func main(){

	// function calling
	var sub int = sub(15, 5);

	fmt.Println(sub)

	divi := dividee(25, 3);

	fmt.Println("result is ", divi)

	a := 19
	b := 3

	ans := a / b 
	fmt.Println(ans)
}