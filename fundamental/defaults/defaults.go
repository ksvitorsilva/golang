package main

import (
	"fmt"
)

func main() {
	var a int
	var b float64
	var c bool
	var d string
	var e *int

	fmt.Println(a, b, c, d, e)
	fmt.Printf("%v %v %v %v %v", a, b, c, d, e)
	fmt.Println()
	fmt.Printf("%v %v %v %q %v", a, b, c, d, e)
}
