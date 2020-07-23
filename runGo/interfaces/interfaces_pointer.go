package main

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rect struct {
	width  float64
	height float64
}

// Area method has pointer receiver
func (r *Rect) Area() float64 {
	return r.width * r.height
}

func (r Rect) Perimeter() float64 {

	return 2 * (r.width + r.height)
}

func main() {
	r := Rect{5.0, 4.0}
	// var s Shape = r
	var s Shape = &r // assingned pointer
	area := s.Area()

	perimeter := s.Perimeter()
	// s.Perimeter() call did not fail even though Perimeter is not implemented by
	// Area. Seems like Go is happy to pass a *copy of the pointer value* as the
	// receiver to the Perimeter method perhaps because it is not a very dangerous
	// idea, it is just a copy and nobody can mutate it accidentally.

	fmt.Println("area of rectangle is", area)
	fmt.Println("perimeter of rectangle is", perimeter)
}
