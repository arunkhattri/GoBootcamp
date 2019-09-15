package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	i, _ := strconv.Atoi(os.Args[1])

	// switch condition is by default true
	// switch true {
	switch {
	case i > 0:
		// "positive"
		fmt.Println("Positive")
	case i < 0:
		// "negative"
		fmt.Println("Negative")
	default:
		// "zero"
		fmt.Println("Zero")
	}
}
