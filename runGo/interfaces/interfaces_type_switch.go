package main

// Type switch. The syntax for type switch is similar to type assertion and it
// is i.(type) where i is interface and type is a fixed keyword. Using this we
// can get the dynamic type of the interface instead of the dynamic value.

import (
	"fmt"
	"strings"
)

func explain(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Println("i stored string ", strings.ToUpper(i.(string)))
	case int:
		fmt.Println("i stored int", i)
	default:
		fmt.Println("i stored something else", i)
	}
}

func main() {
	explain("Hello World")
	explain(52)
	explain(true)
}
