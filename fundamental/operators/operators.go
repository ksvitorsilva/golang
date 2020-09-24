package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("banana" == "banana")
	fmt.Println("banana" != "banana")
	fmt.Println(2 != 3)
	fmt.Println(2 > 3)
	fmt.Println(2 >= 3)
	fmt.Println(2 <= 3)

	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)
	fmt.Println(d1 == d2)
	fmt.Println(d1.Equal(d2))

	type User struct {
		Name string
	}
	p1 := User{"John Doe"}
	p2 := User{"John Doe"}
	fmt.Println(p1 == p2)
}
