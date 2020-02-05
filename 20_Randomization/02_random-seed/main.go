package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Produce a different random number in each run
// seed with a different value
// time.Now().UnixNano() yields a constantly changing number

func main() {
	rand.Seed(time.Now().UnixNano())
	guess := 10

	for n := 0; n != guess; {
		n = rand.Intn(guess + 1)
		fmt.Printf("%d ", n)
	}
	fmt.Println()
}
