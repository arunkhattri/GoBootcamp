package main

// Contains tells whether slice a contains x.

import (
	"fmt"
	"os"
)

func contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func main() {

	vowels := []string{"a", "e", "i", "o", "u"}
	args := os.Args

	if len(args) != 2 || len(args[1]) != 1 {
		fmt.Println("Give me a letter.")
		return
	}

	s := args[1]

	if contains(vowels, s) {
		fmt.Printf("%q is vowel.\n", s)
	} else {
		fmt.Printf("%q is consonant.\n", s)
	}
}
