package main

import "fmt"

func main() {
	slice := make([]int, 10)
	slice[9] = 99
	fmt.Println(slice)

	slice = make([]int, 10, 20)
	fmt.Println(slice, len(slice), cap(slice))
	slice = append(slice, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(slice, len(slice), cap(slice))
	slice = append(slice, 2)
	fmt.Println(slice, len(slice), cap(slice))
}
