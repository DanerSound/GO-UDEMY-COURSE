package main

import (
	"fmt"
)

func printMap(m map[string]string) {
	for color, hex := range m {
		fmt.Println("color:", color, " hex: ", hex)
	}
}
func main() {
	//namemaps :=map[keyType]valueType{}
	// var colors map[string]string
	// colors := make(map[string]string)

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#verde",
	}

	colors["white"] = "bianco" // non ce bisogno di specificare l'aggiunta
	//fmt.Print(colors)
	//delete(colors,"white")
	//fmt.Print(colors)
	printMap(colors)

}
