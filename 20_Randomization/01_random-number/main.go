package main

import (
	"fmt"
	"math/rand"
)

// Produce a random number

func main() {
	guess := 10

	for n := 0; n != guess; {
		n = rand.Intn(guess + 1)
		fmt.Printf("%d ", n)
	}
	fmt.Println()
}
