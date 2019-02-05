package main

import (
	"fmt"
)

func main() {

	var (
		planet   = "venus"
		distance = 261
		orbital  = 224.701
		hasLife  = false
	)

	// type of variables
	// fmt.Printf("Planet: %v\n", planet)
	// fmt.Printf("Distance: %v millions Kms\n", distance)
	// fmt.Printf("Orbital Period: %v days\n", orbital)
	// fmt.Printf("Does %v has life? %v\n", planet, hasLife)
	// fmt.Printf("%v is %v away. Think! %[2]v Kms! %[1]v OMG!\n", planet, distance)
	// fmt.Printf("Planet: %s\n", planet)

	// type safe printing
	fmt.Printf("Distance: %d millions Kms\n", distance)
	fmt.Printf("Orbital Period: %.2f days\n", orbital)
	fmt.Printf("Does %v has life? %t\n", planet, hasLife)
	fmt.Printf("%s is %d away. Think! %[2]d Kms! %[1]s OMG!\n", planet, distance)
}
