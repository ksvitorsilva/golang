package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	for k := 0; k < 11; k++ {
		fmt.Println(k)
	}

	for k := 0; k <= 10; k++ {
		if k%2 == 0 {
			fmt.Println("PAR")
		} else {
			fmt.Println("IMPAR")
		}
	}

	for {
		fmt.Println("PARA O INFINITO...")
		time.Sleep(time.Second * 2)
	}
}
