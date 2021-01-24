package main

import (
	"fmt"
)

func main() {
	tittle := "Golang the Best Language"
	for index, letter := range tittle {
		letterString := string(letter)

		switch string(letter) {
		case "a", "i", "u", "e", "o":
			fmt.Println("index", index, "letter :", letterString)
		}
	}
}
