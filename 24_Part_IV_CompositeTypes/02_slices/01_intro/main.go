package main

import (
	"fmt"
)

func main() {
	// arrays can't grow or shrink
	// length of an array is part of its type.
	// automatically sets uninitaialized array elements to theri zero values
	// slices can grow and shrink
	// slice's zero value is nil
	var books [5]string
	books[0] = "dracula"
	books[1] = "1984"
	books[2] = "island"

	games := []string{"pokemon", "roblox"}

	fmt.Printf("[-] books: %#v\n", books)
	fmt.Printf("[-] games: %#v\n", games)
	fmt.Printf("[-] games: %T\n", games)

	var sports []string

	fmt.Printf("sports : %#v\n", sports)
	fmt.Printf("sports len: %d\n", len(sports))

	sports = append(sports, "soccer")
	fmt.Printf("sports : %#v\n", sports)

	sports = append(sports, "cricket")
	fmt.Printf("sports : %#v\n", sports)

	// newBooks := [...]string{"Excel 2010", "Roger's Thesarus"}

	// mismatched type
	// fmt.Printf("Books and newBooks are same: %t\n", books == newBooks)

	// comparing slices
	var ok string
	for i, game := range games {
		if game != sports[i] {
			ok = "not"
			break
		}
	}

	fmt.Printf("games and sports are %s equal\n\n", ok)

}
