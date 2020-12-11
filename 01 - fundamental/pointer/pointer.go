package main

import "fmt"

func main() {
	a := 1

	var p *int = nil
	p = &a

	// NOP p++
	*p++
	a++

	fmt.Println(p, *p, a, &a)
}
