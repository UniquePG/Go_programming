package main;

import ("fmt"
		"math" 
	)

//globle scope (Pascal form -> First latter of variable is Capital)
var GlobleVariable = "this is string"  

//package level (camel case-> firstlatter small and next word letter capital)
var myVar bool = false  

// fmt.Println(myVar);

func main (){
	
	//local variable
	var var1 int = 20
	var var2 string = "string"


	fmt.Println(var1)
	fmt.Println(var2)

	fmt.Println(GlobleVariable)

	fmt.Println(myVar)


	fmt.Println("You number %g answere\n:", math.Sqrt(float64(var1)));

	fmt.Println("You number answere\n:", math.Sqrt(16));


}