package main

import (
	"fmt"
)

func main() {
	// same type and static length
	var notes [5]int
	notes[0], notes[1], notes[2], notes[3], notes[4] = 3, 6, 2, 8, 4
	fmt.Println(notes)

	total := 0

	for i := 0; i < len(notes); i++ {
		total += notes[i]
	}

	average := total / len(notes)

	fmt.Println(average)
}
