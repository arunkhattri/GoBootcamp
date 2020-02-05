package main

import (
	"fmt"
)

func main() {
	// blue bookcase
	blue := [...]int{6, 9, 3}
	// red bookcase
	red := [...]int{6, 9, 3}

	fmt.Printf("blue bookcase: %v\n", blue)
	fmt.Printf("red bookcase: %v\n", red)

	fmt.Printf("red is equal to blue? %t", blue == red)

}
