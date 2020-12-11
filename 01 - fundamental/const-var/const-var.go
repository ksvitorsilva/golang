package main

import (
	"fmt"
	// rename import
	m "math"
)

func main() {
	const PI float64 = 3.1415
	var diameter = 3.2 // auto type float64

	// or
	// = assign
	// := create and assign
	area := PI * m.Pow(diameter, 2)

	fmt.Println("value:", area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	var e, f bool = true, false

	fmt.Println(e, f)

	g, h, i := 2, false, "Supp"

	fmt.Println(g, h, i)

}
