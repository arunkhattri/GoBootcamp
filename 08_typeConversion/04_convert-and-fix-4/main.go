package main

// ---------------------------------------------------------
// EXERCISE: Convert and Fix #4
//
//  Fix the code.
//
// EXPECTED OUTPUT
//  9.5
// ---------------------------------------------------------
import "fmt"

func main() {
	age := 2
	// fmt.Println(int(7.5) + int(age))
	fmt.Println(7.5 + float64(age))
}
