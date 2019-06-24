package main

import "fmt"

func main() {
	// logical "And" and "OR" operator can only be used with bool values
	speed := 100
	upperLimit := 55
	lowerLimit := 40
	fmt.Println("Within Limits?", speed <= upperLimit && speed >= lowerLimit)
	// short-circuit
	// false && true => false
	// ^       ^
	// |       |
	// +------ only runs this
	//         |
	//         +---- it doesn't run this one

	fmt.Println("Within Limits?", speed >= lowerLimit || speed <= upperLimit)
	// short-circuit
	// true || false => true
	// ^       ^
	// |       |
	// +------ only runs this
	//         |
	//         +---- it doesn't run this one

	speed = 20
	fmt.Println("Within Limits?", speed >= lowerLimit && speed <= upperLimit)
	fmt.Println("Within Limits?", speed >= lowerLimit || speed <= upperLimit)

	speed = 50
	fmt.Println("Within Limits?", speed >= lowerLimit && speed <= upperLimit)
	fmt.Println("Within Limits?", speed >= lowerLimit || speed <= upperLimit)

	// "NOT" operator
	var on bool
	on = !on
	fmt.Println(on)
	fmt.Println("a" > "b")
}
