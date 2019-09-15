package main

import "fmt"

func main() {
	switch i := 10; {
	case i > 0:
		// positive
		fmt.Println("positive")
	case i < 0:
		// negative
		fmt.Println("negative")
	default:
		// zero
		fmt.Println("zero")
	}

}
