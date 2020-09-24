package main

import (
	"fmt"
)

func main() {
	fmt.Print("Hello")
	fmt.Print("world!\n")

	fmt.Println("Ola")
	fmt.Println("mundo!")

	x := 3.1415
	// error ==> fmt.Println("Value of X: " + x)

	xs := fmt.Sprint(x)
	fmt.Println("Value of X: " + xs)
	fmt.Println("Value of X: ", x)
	fmt.Printf("Value of X: %.2f", x)

	a := 1
	b := 1.99999
	c := false
	d := "YO!"

	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)
	fmt.Printf("\n%v %v %v %v", a, b, c, d)

}
