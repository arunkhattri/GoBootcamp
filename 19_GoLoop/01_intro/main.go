package main

import (
	"fmt"
)

func main() {
	var sum int

	// for loop

	// for i := 1; i <= 1000; i++ {
	// 	sum += i
	// }
	// sum++
	// sum += 2
	// sum += 3
	// sum += 4
	// sum += 5

	// using break statement
	i := 1
	for {
		if i > 5 {
			break
		}
		sum += i
		i++
	}

	fmt.Println(sum)
}
