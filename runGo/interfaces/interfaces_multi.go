package main

// Shape interface with Area method and Object interface with Volume method.
// Since struct type Cube implements both these methods, it implements both
// these interfaces. Hence we can assign a value of struct type Cube to the
// variable of type Shape or Object.

import "fmt"

type Shape interface {
	Area() float64
}

type Object interface {
	Volume() float64
}

type Cube struct {
	side float64
}

func (c Cube) Area() float64 {
	return 6 * (c.side * c.side)
}

func (c Cube) Volume() float64 {
	return c.side * c.side * c.side
}

func main() {
	c := Cube{3}
	var s Shape = c
	var o Object = c
	fmt.Println("volume of s of interface type Shape is", s.Area())
	fmt.Println("area of o of interface type Object is", o.Volume())

	//fmt.Println("area of s of interface type Shape is", s.Volume())
	//fmt.Println("volume of o of interface type Object is", o.Area())
}
