package main

import "fmt"

func main() {

	// var colors map[string]string  // Does not allocate memory colors["red"] = "#ff0000" // ❌ panic: assignment to entry in nil map

	// colors := make(map[string]string) // allocates memory
	colors := make(map[int]string)

	colors[10] = "#fffff"

	// Built in function

	delete(colors, 10)
	fmt.Println(colors)
	// Iterating over a map
	coloring := map[string]string{
		"red":   "#ff000",
		"green": "#4bf745",
	}
	printMap(coloring)

}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex Code for", color, "is", hex)

	}

}
