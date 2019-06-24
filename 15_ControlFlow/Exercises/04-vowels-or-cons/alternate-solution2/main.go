package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args

	if len(args) < 2 || len(args[1]) != 1 {
		fmt.Println("Give me letter.")
		return
	}

	l := args[1]

	if strings.IndexAny("aeiou", l) != -1 {
		fmt.Printf("%q is a vowel.\n", l)
	} else if l == "y" || l == "w" {
		fmt.Printf("%q is a semi-vowel.\n", l)
	} else {
		fmt.Printf("%q is a consonant.\n", l)
	}
}
