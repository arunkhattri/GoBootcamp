package main

import (
	"fmt"
	"os"
)

func main() {
	city := os.Args[1]

	switch city {
	case "Paris":
		fmt.Println("France")
	case "Tokyo":
		fmt.Println("Japan")
	case "Delhi":
		fmt.Println("India")
		vip := true
		fmt.Println("VIP tour:", vip)
		// multiple conditions
	case "Kathmandu", "Pokhara":
		fmt.Println("Nepal")
	// default case, one per switch statement
	default:
		fmt.Println("Where?")
	}
}
