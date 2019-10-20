package main

// ---------------------------------------------------------
// EXERCISE: Only Evens
//
//  1. Extend the "Sum up to N" exercise
//  2. Sum only the even numbers
//
// RESTRICTIONS
//  Skip odd numbers using the `continue` statement
//
// EXPECTED OUTPUT
//  Let's suppose that the user runs it like this:
//
//    go run main.go 1 10
//
//  Then it should print:
//
//    2 + 4 + 6 + 8 + 10 = 30
// ---------------------------------------------------------

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var (
		sum int
	)

	if len(os.Args) != 3 {
		fmt.Println("give me min and max numbers to add")
		return
	}

	min, err1 := strconv.Atoi(os.Args[1])
	max, err2 := strconv.Atoi(os.Args[2])

	if err1 != nil || err2 != nil || min >= max {
		fmt.Println("Something is wrong...have you gave min max numbers???")
		return
	}

	for i := min; i <= max; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Print(i)
		if (max%2 == 0 && i < max) || (max%2 != 0 && i < max-1) {
			fmt.Print(" + ")
		}
		sum += i
	}
	fmt.Printf(" = %d\n", sum)
}
