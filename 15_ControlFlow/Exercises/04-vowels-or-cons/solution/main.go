package main

// ----------------------------------------------
// Detect whether a letter is vowel or consonant.
// Note: y or w is called a semi-vowel
// ----------------------------------------------

import (
	"fmt"
	"os"
)

const msg = `Give me a letter.`

func main() {
	args := os.Args

	if len(args) != 2 || len(args[1]) != 1 {
		fmt.Println(msg)
		return
	}

	// I didn't use a short-if here because, it's already
	// hard to read. Do not make it harder.

	s := args[1]
	if s == "a" || s == "e" || s == "i" || s == "o" || s == "u" {
		fmt.Printf("%q is a vowel.\n", s)
	} else if s == "w" || s == "y" {
		fmt.Printf("%q is sometimes a vowel, sometimes not.\n", s)
	} else {
		fmt.Printf("%q is a consonant.\n", s)
	}

}
