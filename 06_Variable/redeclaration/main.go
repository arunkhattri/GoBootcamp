package main

import "fmt"

func main() {
	safe := false // existing variable
	// new variable 'speed' declares and initialized
	// redeclaraed 'safe' assigns true
	safe, speed := true, 50 // at least one of the variable should be a new variable

	fmt.Println(safe, speed)

	name := "Nikola"
	name, age := "Marie", 66 // redeclaration of name

	fmt.Println(name, age)

	// name = "Albert" // assignment
	// birth := 1879   // declares a new variable

	// ^-- same as above
	name, birth := "Albert", 1879

	fmt.Println(name, birth)
}
