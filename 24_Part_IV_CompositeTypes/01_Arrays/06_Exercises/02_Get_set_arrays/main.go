package main

import "fmt"

func main() {

	// array declaration
	var (
		names     [3]string
		distances [5]int
		data      [5]byte
		ratios    [1]float64
		alives    [4]bool
		// zero      [0]byte
	)

	// assign names to names array
	names[0] = "Adolf Hitler"
	names[1] = "Bennito Musolini"
	names[2] = "Saddam Hussain"

	// assign distances
	distances[0] = 262
	distances[1] = 250
	distances[2] = 245
	distances[3] = 3152
	distances[4] = 1524

	// assign data
	data[0] = 72
	data[1] = 69
	data[2] = 76
	data[3] = 79
	data[4] = 72

	// assign ratio
	ratios[0] = 3.14

	// assign alives
	alives[0] = true
	alives[1] = false
	alives[2] = true
	alives[3] = true

	// ordinary loop array
	for i := 0; i < len(names); i++ {
		fmt.Printf("names[%d] : %s\n", i, names[i])
	}

	// range loops
	for i, v := range distances {
		fmt.Printf("distances[%d]: %d\n", i, v)
	}

}
