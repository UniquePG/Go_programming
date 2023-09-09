package main

import ("fmt"
 	"math/rand" 
)

//! 2
	// func pow( x, n, num float64) float64 {

	// //* we can also give the initial value in the if statement (like for loop)
	// 	if v := math.Pow(x, n); v < num {		//* v is only accesible inside if block

	// 		return v
	// 	}

	// 	return num

	// }

//! 3 
	func demo(num int) int{

		if x := rand.Intn(10); x > num {

			fmt.Println("ramdom is greter than (inside if)", x)

		} else {
			fmt.Println("inside else block", x)		//* x is also available inside else block
		}

		return num
	}


func main(){


//! 1
	// x := 15;

	// if (x < 10) {			// we can omit the () in condition but {} are required
		
	// 	fmt.Println("less than")
	// }


//! 2
	// fmt.Println("result is: ", pow(4, 2, 20))
	// fmt.Println("result is: ", pow(5, 2, 27))


//! 3
	// fmt.Println("result is: ", demo(5))


	

}