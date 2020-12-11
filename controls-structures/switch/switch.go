package main

import (
	"fmt"
	"time"
)

func getNote(note float64) string {
	var intNote = int(note)
	switch intNote {
	case 10:
		fallthrough
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"
	default:
		return "invalid"
	}
}

func morning() string {
	t := time.Now()
	switch {
	case t.Hour() < 13:
		return "Bom dia"
	case t.Hour() < 18:
		return "Boa tarde"
	default:
		return "Boa noite"
	}
}

func getNoteV2(note float64) string {
	switch {
	case note >= 9 && note <= 10:
		return "A"
	case note >= 8 && note < 9:
		return "B"
	case note >= 5 && note < 8:
		return "C"
	case note >= 3 && note < 5:
		return "D"
	default:
		return "E"
	}
}

func print(data string) {
	fmt.Println(data)
}
func geType(value interface{}) string {
	switch value.(type) {
	case int:
		return "INT"
	case float32, float64:
		return "FLOAT"
	case string:
		return "STRING"
	case func():
		return "FUNC"
	default:
		return "UNKNOWN"
	}
}

func main() {
	print(getNote(17))
	print(getNote(7))
	print(getNote(10))
	print(morning())
	print(getNoteV2(7))
	print(getNoteV2(10))
	print(geType(10))
	print(geType(10.3))
	print(geType("10.3"))
	print(geType(func() {}))
}
