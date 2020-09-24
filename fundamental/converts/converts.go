package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y))

	a := 6.9
	b := int(a)
	fmt.Println(b)

	// warning...
	fmt.Println("Test " + string(97))

	// int to string
	fmt.Println("Test " + strconv.Itoa(123))

	// string para int
	// 2 params, num and error but _ is to ignore
	num, _ := strconv.Atoi("123")
	fmt.Println(num - 122)

	h, _ := strconv.ParseBool("true")

	if h {
		fmt.Println("TRUE")
	}
}
