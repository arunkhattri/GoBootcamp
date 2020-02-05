package main

// ---------------------------------------------------------
// EXERCISE: Path Searcher
//
//  Your program should search inside the path environment
//  variable.
//
//  Remove the corpus constant then get the corpus from the
//  environment variable "Path" or "PATH".
//
// EXAMPLE
//  For example, on my Mac, my PATH environment variable
//  looks like this:
//
//    "/usr/local/bin:/sbin:/Users/inanc/go/bin"
//
//  So, if the user runs the program like this:
//
//    go run main.go /sbin
//
//  It should print this:
//
//    #2 : "/sbin"
// ---------------------------------------------------------
import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// get and split the path environmental variable
	words := filepath.SplitList(os.Getenv("PATH"))

	// query
	query := os.Args[1:]

	// labelled (queries) loop
	for _, q := range query {
		// labelled (search) loop
		for i, w := range words {
			q, w = strings.ToLower(q), strings.ToLower(w)
			if strings.Contains(w, q) {
				fmt.Printf("#%-2d: %q\n", i+1, w)
			}

		}
	}

}
