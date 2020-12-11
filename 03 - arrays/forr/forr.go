package main

import "fmt"

func main() {
	// same type and static length
	notes := [...]int{1, 2, 3, 4, 5, 6}

	for i, v := range notes {
		fmt.Println(i, "---", v)
	}

	for _, v := range notes {
		fmt.Println(v)
	}
}
