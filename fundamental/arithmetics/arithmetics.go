package main

import (
	"fmt"
	"math"
)

func main() {
	a := 2
	b := 3

	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a / b)
	fmt.Println(a * b)
	fmt.Println(a % b)
	// bitwise
	fmt.Println(a & b)
	fmt.Println(a | b)
	fmt.Println(a ^ b)

	c := 3.0
	d := 4.0

	fmt.Println(math.Max(float64(a), float64(b)))
	fmt.Println(math.Min(c, d))
	fmt.Println(math.Pow(c, d))

}
