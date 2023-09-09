package main

import "fmt"

type generic interface {
	area() float64
}

type square struct {
	side float64
}

type rectangle struct {
	height float64
	width  float64

}

// make method of area

func (s square) area() float64 {
	return s.side * s.side // return the value of side squared
}

func (s rectangle) area() float64 {
	return s.height * s.width
}


func calculate(g generic) {
	fmt.Println("values are: ", g )
	fmt.Println("Area is: ", g.area())
}

func main() {

	sq := square{5}

	rec := rectangle{10, 5}

	calculate(sq)
	calculate(rec)
}