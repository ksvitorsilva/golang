package main

import "fmt"

func main() {
	array := [6]int{1, 2, 3, 4, 5, 6}
	slice := []int{1, 2, 3}

	fmt.Println(array, "---", slice)

	// slice is not an array

	s2 := array[3:5]
	fmt.Println(s2)
	s3 := array[:4]
	fmt.Println(s3)
	s4 := s3[:2]
	fmt.Println(s4)
}
