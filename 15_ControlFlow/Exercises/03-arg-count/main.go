package main

import (
	"fmt"
	"os"
	"strings"
)

const usage = `
Give me args`

// ---------------------------------------------------------
// EXERCISE: Arg Count
//
//  1. Get arguments from command-line.
//  2. Print the expected outputs below depending on the number
//     of arguments.
//
// EXPECTED OUTPUT
//  go run main.go
//    Give me args
//
//  go run main.go hello
//    There is one: "hello"
//
//  go run main.go hi there
//    There are two: "hi there"
//
//  go run main.go i wanna be a gopher
//    There are 5 arguments
// ---------------------------------------------------------

func main() {
	args := os.Args[1:]
	// if len(args) == 5 {
	// 	fmt.Println("There are 5 arguments")
	// } else if len(args) == 2 {
	// 	fmt.Printf("There are %d arguments: %s\n", len(args), args)
	// } else {
	// 	fmt.Println(strings.TrimSpace(usage))
	// }
	if len(args) >= 2 {
		fmt.Printf("There are %d arguments: %s\n", len(args), strings.Join(args, " "))
	} else {
		fmt.Println(usage)
	}
}
