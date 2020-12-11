package main

import "fmt"

func printResult(note float64) {
	if note >= 10 {
		fmt.Println("OK", note)
	} else {

		fmt.Println("NOT OK", note)
	}
}

func main() {
	printResult(17)
	printResult(7)
}
