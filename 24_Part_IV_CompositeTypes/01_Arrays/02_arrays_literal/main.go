package main

import "fmt"

func main() {
	// array literals
	books := [4]string{
		"Anna Kareina",
		"Madam Bovary",
		"War and Peace",
		"The Great Gatsby",
	}

	fmt.Printf("books: %#v\n", books)

	// elipses (...) set the length of the array by looking at the length
	// of the given elements
	authors := [...]string{
		"Leo Tolstoy",
		"Gustave Flaubert",
		"Leo Tolstoy",
		"F.Scott Fitzgerald",
	}

	fmt.Printf("authors: %#v\n", authors)

	for i := 0; i < 4; i++ {
		fmt.Printf("Book: %s, Author: %s\n", books[i], authors[i])
	}

	// another way
	for i, _ := range books {
		fmt.Printf("Books: %s, Author: %s\n", books[i], authors[i])
	}

}
