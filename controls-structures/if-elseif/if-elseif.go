package main

import "fmt"

func getNote(note float64) string {
	if note >= 9 && note <= 10 {
		return "A"
	} else if note >= 8 && note < 9 {
		return "B"
	} else if note >= 5 && note < 8 {
		return "C"
	} else if note >= 3 && note < 5 {
		return "D"
	} else {
		return "E"
	}
}

func print(data string) {
	fmt.Println(data)
}

func main() {
	print(getNote(17))
	print(getNote(7))
}
