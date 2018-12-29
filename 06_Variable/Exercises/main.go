package main

import "fmt"

func main() {

	// Short Declare
	i := 314
	f := 3.14
	s := "Hello"
	b := true
	fmt.Println(i, f, s, b)

	// Multiple Short Declare
	width, truth := 14, true
	fmt.Println(width, truth)

	// Multiple Short Declare # 2
	a, c := 42, "good"
	fmt.Println(a, c)

	// Short with expression
	sum := 27 + 3.5
	fmt.Println(sum)

	// Redeclare
	age, yourAge := 40, 47
	fmt.Printf("age = %d\nyourAge = %d\n", age, yourAge)

	ratio, age := 3.14, 42
	fmt.Println(age, yourAge, ratio)
}
