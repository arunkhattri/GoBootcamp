package main

import (
	"fmt"
	"strings"
)

func main() {
	// for i := 1; i < len(os.Args); i++ {
	// fmt.Printf("%q\n", os.Args[i])
	// }

	words := strings.Fields(
		"lazy cat jumps again and again and again",
	)

	for j := 0; j < len(words); j++ {
		fmt.Printf("#%-2d: %q\n", j+1, words[j])
	}
}
