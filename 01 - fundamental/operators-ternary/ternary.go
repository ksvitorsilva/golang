package main

import "fmt"

// there are no ternaries :tada:

func handleNote(note float64) bool {
	// no no no return note >= 10 ? true : false
	if note >= 10 {
		return true
	}

	return false
}

func main() {
	fmt.Println(handleNote(8))
	fmt.Println(handleNote(13))
}
