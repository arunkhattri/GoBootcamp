package main

import (
	"fmt"
	"strings"
)

func main() {

	books := [...]string{
		"Anna Karenina",
		"Madam Bovary",
		"War and Peace",
		"The Great Gatsby",
		"Lolita",
		"Middlemarch",
		"The adventures of Huckleberry Finn",
	}

	seperator := "\n" + strings.Repeat("=", 35) + "\n"

	fmt.Print("books", seperator)
	fmt.Printf("books: %#v\n", books)

	// create two more copies
	upper := books
	lower := books

	// change the book titles to uppercase in upper array only
	for i, v := range books {
		upper[i] = strings.ToUpper(v)
	}

	fmt.Print("upper", seperator)
	fmt.Printf("upper: %#v\n", upper)
	// change the book titles to lowercase in lower array only
	for i, v := range books {
		lower[i] = strings.ToLower(v)
	}
	fmt.Print("lower", seperator)
	fmt.Printf("lower: %#v\n", lower)
}
