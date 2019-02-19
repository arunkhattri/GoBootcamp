package main

import "fmt"

func main() {
	var n int //initialised with 0 value
	fmt.Println(n)

	n = n + 1
	fmt.Println(n)

	n += 1
	fmt.Println(n)

	n++ //incdec statement
	fmt.Println(n)

	// decrease counter
	n--
	fmt.Println(n)

}
