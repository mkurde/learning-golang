package main

// Show type empty interface and explains the dynamic value & type of the
// interface

import "fmt"

type MyString string

type Rect struct {
	width  float64
	height float64
}

func explain(i interface{}) {
	fmt.Printf("value given to explain function is of type '%T' with value %v\n", i, i)
}

func main() {
	ms := MyString("Hello World!")
	r := Rect{5.5, 4.5}
	explain(ms)
	explain(r)
}
