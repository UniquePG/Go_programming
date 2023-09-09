package main

import ( "fmt"
		// "math/rand"
)

func main () {

	// value := "applee"

	//* switch case does neet break statment (only one of them case is execute if it is match)
	// switch value {
	// case "mango": 
	// 	fmt.Println("it is ", value)

	// case "apple":
	// 	fmt.Println("it is ", value)
	
	// default:
	// 	fmt.Println("any case doesnt match")
	// }



	//* switch without a condition (means switch true)
		// value := rand.Intn(20)
		// switch {
		// 	case value < 5 :
		// 		fmt.Println("less than 5")
		// 	case value > 10:
		// 		fmt.Println("grater than 10")

		// 	default: 
		// 		fmt.Println("default case")
		// }

		// fmt.Println(value)



//! Defer -> it executes the statement at the very end of the function
		// defer fmt.Println("hellooo 1")

		// fmt.Println("hellooo 2")

		// defer fmt.Println("hellooo 3")
		// defer fmt.Println("hellooo 4")

//! stacking defer-> it executes the statement last in first out (LIFO) order
		fmt.Println("start")

		for i := 1; i<=10; i++ {
			defer fmt.Println(i)
		}

		fmt.Println("end")
}