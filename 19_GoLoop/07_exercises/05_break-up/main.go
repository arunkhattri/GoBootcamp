package main

// ---------------------------------------------------------
// EXERCISE: Break Up
//
//  1. Extend the "Only Evens" exercise
//  2. This time, use an infinite loop.
//  3. Break the loop when it reaches to the `max`.
//
// RESTRICTIONS
//  You should use the `break` statement once.
//
// HINT
//  Do not forget incrementing the `i` before the `continue`
//  statement and at the end of the loop.
//
// EXPECTED OUTPUT
//    go run main.go 1 10
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

	i := min
	for {
		if i > max {
			break
		}
		if i%2 != 0 {
			i++
			continue
		}
		fmt.Print(i)
		if (max%2 == 0 && i < max) || (max%2 != 0 && i < max-1) {
			fmt.Print(" + ")
		}
		sum += i
		i++
	}
	fmt.Printf(" = %d\n", sum)
}
