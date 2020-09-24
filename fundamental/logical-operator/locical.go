package main

import "fmt"

func buys(job1, job2 bool) (bool, bool, bool) {
	tv50 := job1 && job2
	tv32 := job1 != job2
	radio := job1 || job2

	return tv50, tv32, radio
}

func main() {
	a, b, c := buys(true, true)
	fmt.Println(a, b, c)
	a, b, c = buys(false, true)
	fmt.Println(a, b, c)
	a, b, c = buys(true, false)
	fmt.Println(a, b, c)
	a, b, c = buys(false, false)
	fmt.Println(a, b, c)
}
