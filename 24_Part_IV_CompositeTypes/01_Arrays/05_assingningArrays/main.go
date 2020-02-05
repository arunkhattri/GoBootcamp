package main

import (
	"fmt"
)

func main() {
	books := [...]string{
		"Anna Karenina",
		"Madame Bovary",
		"War and Peace",
		"The Great Gatsby",
	}

	newBooks := books

	for i := range books {
		newBooks[i] += " 2nd Ed."
	}

	fmt.Printf("New Books: %q\n", newBooks)
	fmt.Printf("Books: %#v\n", books)

}
