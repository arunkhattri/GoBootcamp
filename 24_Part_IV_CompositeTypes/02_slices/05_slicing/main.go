package main

import (
	"fmt"
)

func main() {
	items := []string{
		"pacman", "mario", "tetris", "doom",
		"galaga", "frogger", "asteroids", "simcity",
		"metroid", "defender", "rayman", "tempest",
		"ultima"}

	fmt.Printf("len: %d\nitems: %v\n", len(items), items)

	// slice
	top3 := items[:3]
	fmt.Printf("len: %d\nitems: %v\n", len(top3), top3)

	// pagination

	const pageSize = 4

	l := len(items)

	for from := 0; from < l; from += pageSize {
		to := from + pageSize
		if to > l {
			to = l
		}
		currentPage := items[from:to]

		head := fmt.Sprintf("Page #%d | %d:%d", (from/pageSize)+1, from, to)
		fmt.Printf("%s\n%v\n", head, currentPage)
	}

}
