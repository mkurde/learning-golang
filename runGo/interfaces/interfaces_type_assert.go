package main

// We can find out the underlying dynamic value of an interface using the syntax
// i.(Type) where i is a variable of type interface and Type is a type that
// implements the interface. Go will check if dynamic type of i is identical to
// the Type and return the dynamic value is possible.

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
	var s Shape = Cube{3}
	c := s.(Cube)
	fmt.Println("area of c of type Cube is", c.Area())
	fmt.Println("volume of c of type Cube is", c.Volume())
}
