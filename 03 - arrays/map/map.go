package main

import "fmt"

func main() {
	// same type and static length
	data := make(map[int]string)
	data[1] = "ola"
	data[2] = "mundo"

	for i, v := range data {
		fmt.Println(i, "---", v)
	}

	fmt.Println(data[2])
	delete(data, 2)
	data[2] = "world"
	fmt.Println(data)

	salaries := map[string]float64{
		"ola":   123,
		"mundo": 234,
		"code":  432,
	}
	salaries["visual"] = 233

	fmt.Println(salaries)

	delete(salaries, "unknown")

	for i, v := range salaries {
		fmt.Println(i, "---", v)
	}

	mapOfMap := map[string]map[string]float64{
		"H": {
			"Hello": 123,
			"Hi":    134,
		},
		"W": {
			"world": 322,
		},
		"M": {
			"mundo": 234,
		},
	}

	fmt.Println(mapOfMap)
	delete(mapOfMap, "M")
	for i, v := range mapOfMap {
		fmt.Println(i)
		for k, y := range v {
			fmt.Println(k, "---", y)
		}
	}
}
