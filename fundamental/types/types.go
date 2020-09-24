package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// ints
	fmt.Println(1, 2, 333)
	fmt.Println("tupeof", reflect.TypeOf(12))

	/*
		positives
			uint8 = 1byte
			uint16 = 2bytes / sort
			uint32 = 4bytes / int
			uint64 = 81bytes / log
	*/

	var b byte = 255
	fmt.Println("tupeof b", reflect.TypeOf(b))

	i1 := math.MaxInt64
	fmt.Println("Max of i1", i1)
	fmt.Println("tupeof i1", reflect.TypeOf(i1))

	var i2 rune = 'a' // map to unicode ref
	fmt.Println("tupeof i2", reflect.TypeOf(i2))
	fmt.Println(i2)

	// floats
	var x float32 = 49.99
	fmt.Println("tupeof i2 x", reflect.TypeOf(x))
	fmt.Println("tupeof literal", reflect.TypeOf(49.99)) // float64

	// boolean
	bo := true
	fmt.Println("tupeof bo", reflect.TypeOf(bo))
	fmt.Println(!bo)

	// string
	s1 := "Hello world"
	fmt.Println(s1 + "!")
	fmt.Println("length is", len(s1))

	// string com multiplas linhas
	s2 := `Hello
		world
		!
		`
	fmt.Println("length is", len(s2))

	// char???
	// var x char = 'b'
	char := 'a'
	fmt.Println("typeof char", reflect.TypeOf(char))
	fmt.Println(char) // unicode value
}
