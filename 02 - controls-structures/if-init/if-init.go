package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomNumber() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	return r.Intn(10)
}

func print(data string) {
	fmt.Println(data)
}

func main() {
	if i := randomNumber(); i > 5 {
		fmt.Println(i)
		print("WIN")
	} else {
		fmt.Println(i)
		print("LOSE")
	}
}
