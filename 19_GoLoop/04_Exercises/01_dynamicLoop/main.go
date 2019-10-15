package main

//  Get the size of the table from the command-line
//    Passing 5 should create a 5x5 table
//    Passing 10 for a 10x10 tableCreate an infinite loop
//    Passing 10 for a 10x10 table
//

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Give me the size of the table you want")
		return
	}

	size, err := strconv.Atoi(os.Args[1])
	if err != nil || size <= 0 {
		fmt.Println("Give me a number from 1 to 10")
		return
	}

	fmt.Printf("%5s", "X")
	for i := 1; i <= size; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println()

	for i := 1; i <= size; i++ {
		fmt.Printf("%5d", i)

		for j := 1; j <= size; j++ {
			fmt.Printf("%5d", i*j)
		}
		fmt.Println()
	}
}
