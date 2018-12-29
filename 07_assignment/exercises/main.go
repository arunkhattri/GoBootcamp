package main

import (
	"fmt"
)

func main() {
	// make it blue
	color := "green"
	fmt.Println(color)
	color = "blue"
	fmt.Println(color)

	// Variables to Variables
	color = "dark " + color
	fmt.Println(color)

	// Assign with Expressions
	n := 0.
	n = 3.14 * 2
	fmt.Println(n)

	// find the rectangle's peimeter
	var (
		perimeter     int
		width, height = 5, 6
	)
	perimeter = 2 * (width + height)
	fmt.Println(perimeter)

	// Multi Assign
	var (
		lang    string
		version int
	)
	lang, version = "go", 1
	fmt.Println(lang, "version:", version)

	// Multi Assign #2
	var (
		planet string
		isTrue bool
		temp   float64
	)
	planet, isTrue, temp = "Mars", true, 80.15
	fmt.Println(planet, isTrue, temp)

	// Multi short func
	a, b := 8, 9
	fmt.Println(a, b)
	_, b = multi()
	fmt.Println("b", b)

	// Swapper
	c1, c2 := "green", "orange"
	fmt.Println("c1:", c1, "c2:", c2)
	c1, c2 = c2, c1
	fmt.Println("c1:", c1, "c2:", c2)

	// Swapper #2
	red, blue := "red", "blue"
	red, blue = blue, red
	fmt.Println(red, blue)

}

// Multi is a function that returns multiple int values
func multi() (int, int) {
	return 5, 4
}
