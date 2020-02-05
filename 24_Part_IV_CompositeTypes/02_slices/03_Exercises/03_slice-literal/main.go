package main

import (
	"fmt"
)

func main() {
	var (
		names     []string
		distances []int
		data      []byte
		ratios    []float64
		alives    []bool
	)

	names = []string{"booboo", "inoo", "mia"}
	distances = []int{3, 1, 1}
	data = []byte{7, 3, 1}
	ratios = []float64{2008.5, 2013.1, 2016.10}
	alives = []bool{true, true, true}

	fmt.Printf("names: %T | %d | %t\n", names, len(names), names == nil)
	fmt.Printf("distances %T | %d | %t\n", distances, len(distances), distances == nil)
	fmt.Printf("data %T | %d | %t\n", data, len(data), data == nil)
	fmt.Printf("ratios: %T | %d | %t\n", ratios, len(ratios), ratios == nil)
	fmt.Printf("alives: %T | %d | %t\n", alives, len(alives), alives == nil)

	if len(distances) == len(data) {
		println("length of distances and data are same")
	}

}
