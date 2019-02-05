package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Print the Type
//
//  Print the type and value of 3 using fmt.Printf
//
// EXPECTED OUTPUT
//  Type of 3 is int
// ---------------------------------------------------------

func main() {
	n := 3
	fmt.Printf("Type of %d is %[1]T\n", n)
	//  Print the type and value of 3.14 using fmt.Printf
	nF := 3.14
	fmt.Printf("Type of %.2f is %[1]T\n", nF)

	//  Print the type and value of "hello" using fmt.Printf
	msg := "hello"
	fmt.Printf("Type of %s is %[1]T\n", msg)

	//  Print the type and value of true using fmt.Printf
	bt := true
	fmt.Printf("Type of %t is %[1]T\n", bt)
}
