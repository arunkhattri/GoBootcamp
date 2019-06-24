package main

// Comparision Operators

import (
	"fmt"
	"strings"
)

func main() {
	// Equality Operators
	// can be used wiht any comparable type and they yield a bool value
	speed := 100
	fast := speed >= 80 // yield bool
	slow := speed < 20

	fmt.Printf("fast's type is %T\n", fast)
	fmt.Printf("going fast? %t\n", fast)
	fmt.Printf("going slow? %t\n", slow)

	fmt.Printf("Is it 100 mph? %t\n", speed == 100)
	fmt.Printf("Is it 100 mph? %t\n", speed != 100)

	fmt.Println(strings.Repeat("-", 25))

	speedB := 150.5
	faster := speedB > 100

	fmt.Println("is the other car going faster?", faster)
	fmt.Println(false >= true)

	// if something is not assignable it's uncomparable too

}
