package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// for i := 1; i < len(os.Args); i++ {
	// 	fmt.Printf("%q\n", os.Args[i])
	// }
	// convert above in range

	for _, v := range os.Args[1:] {
		fmt.Printf("%q\n", v)
	}

	words := strings.Fields(
		"lazy cat jumps again and again and again",
	)

	// for j := 0; j < len(words); j++ {
	// 	fmt.Printf("#%-2d: %q\n", j+1, words[j])
	// }
	// convert above in range

	for i, v := range words {
		fmt.Printf("#%-2d: %q\n", i+1, v)
	}
}
