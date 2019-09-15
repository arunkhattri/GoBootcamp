package main

import (
	"fmt"
)

func main() {
	i := 142
	// switch {
	// case i > 100:
	// 	fmt.Print("big positive number")
	// case i > 0:
	// 	fmt.Print("positive number")
	// default:
	// 	fmt.Print("number")
	// }
	switch {
	case i > 100:
		fmt.Print("big ")
		// fallthrough executes the next clause without checking
		// for its condition
		fallthrough
	case i > 0:
		fmt.Print("positive ")
		fallthrough
	default:
		fmt.Print("number")
	}
	fmt.Println()
}
