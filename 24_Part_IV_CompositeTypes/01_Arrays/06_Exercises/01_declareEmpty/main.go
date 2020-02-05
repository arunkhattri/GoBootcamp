package main

// ----------------------------------------------------------------------
// EXERCISE: Declare empty arrays
//
//  1. Declare and print the following arrays with their types:
//
//    1. The names of your best three friends (names array)
//
//    2. The distances to five different locations (distances array)
//
//    3. A data buffer with five bytes of capacity (data array)
//
//    4. Currency exchange ratios only for a single currency (ratios array)
//
//    5. Up/Down status of four different web servers (alives array)
//
//    6. A byte array that doesn't occupy memory space (zero array)
//
//  2. Print only the types of the same arrays.
//
//  3. Print only the elements of the same arrays.
//
// ----------------------------------------------------------------------

import (
	"fmt"
)

func main() {

	var (
		names     [3]string
		distances [5]int
		data      [5]byte
		ratios    [1]float64
		alives    [4]bool
		zero      [0]byte
	)

	// print elements of the array along with type
	fmt.Printf("names: %#v\n", names)
	fmt.Printf("distances: %#v\n", distances)
	fmt.Printf("data: %#v\n", data)
	fmt.Printf("ratios: %#v\n", ratios)
	fmt.Printf("alives: %#v\n", alives)
	fmt.Printf("zero: %#v\n", zero)

	// print type of the array
	fmt.Printf("names: %T\n", names)
	fmt.Printf("distances: %T\n", distances)
	fmt.Printf("data: %T\n", data)
	fmt.Printf("ratios: %T\n", ratios)
	fmt.Printf("alives: %T\n", alives)
	fmt.Printf("zero: %T\n", zero)

	// print element of the array
	fmt.Printf("names: %q\n", names)
	fmt.Printf("distances: %d\n", distances)
	fmt.Printf("data: %d\n", data)
	fmt.Printf("ratios: %.2f\n", ratios)
	fmt.Printf("alives: %t\n", alives)
	fmt.Printf("zero: %d\n", zero)
}
