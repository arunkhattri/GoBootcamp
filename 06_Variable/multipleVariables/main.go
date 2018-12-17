package main

import "fmt"

func main() {
	var (
		speed int
		heat  float64

		off   bool
		brand string
	)

	// can be done like parallel declaration
	// var speed, velocity int

	fmt.Println(speed, heat)
	fmt.Printf("%q\n", brand)
	fmt.Println(off)
}
