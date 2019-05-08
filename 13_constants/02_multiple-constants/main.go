package main

import "fmt"

func main() {
	const (
		gst     float64 = 0.18
		freight         = 4.5 // Constants get their types and expression from the prev const
	)

	basic := 45.0
	netLanding := basic + (basic * gst) + freight

	fmt.Printf("Net Landing: %0.2f\n", netLanding)

}
