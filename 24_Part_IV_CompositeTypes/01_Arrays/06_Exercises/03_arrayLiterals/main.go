package main

import (
	"fmt"
	"strings"
)

func main() {

	names := [3]string{
		"Adolf Hitler",
		"Bennito Musolini",
		"Saddam Hussain",
	}
	distances := [5]int{262, 250, 245, 3152, 1524}
	data := [5]byte{'H', 'E', 'L', 'L', 'O'}
	ratios := [1]float64{3.14}
	alives := [4]bool{true, false, true, true}
	var zero [0]byte

	seperator := "\n" + strings.Repeat("=", 20) + "\n"

	// ordinary loop array
	fmt.Print("names", seperator)
	for i := 0; i < len(names); i++ {
		fmt.Printf("names[%d] : %s\n", i, names[i])
	}

	// range loops
	fmt.Print("distances", seperator)
	for i, v := range distances {
		fmt.Printf("distances[%d]: %d\n", i, v)
	}

	// data
	fmt.Print("data", seperator)
	for i, v := range data {
		fmt.Printf("data[%d]: %d\n", i, v)
	}

	// ratios
	fmt.Print("ratios", seperator)
	for i, v := range ratios {
		fmt.Printf("ratios[%d] = %.2f\n", i, v)
	}

	// alives
	fmt.Print("alives", seperator)
	for i, v := range alives {
		fmt.Printf("alives[%d] = %t\n", i, v)
	}

	// zero
	fmt.Print("zero", seperator)
	for i, v := range zero {
		fmt.Printf("zero[%d] = %d\n", i, v)
	}
}
