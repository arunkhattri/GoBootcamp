package main

import "fmt"

func main() {
	myAge := 42
	yourAge := 35
	var avg float64
	var ratio float64 = 3 / 2 // Note print statement
	var ratio1 float64 = 3.0 / 2

	avg = float64(myAge+yourAge) / 2
	fmt.Printf("Average: %.2f\n", avg)
	fmt.Println("sum: ", 2+3)
	fmt.Println("sum: ", 2+3.5)
	fmt.Println("diff: ", 3-2)
	fmt.Println("diff: ", 3-2.5)
	fmt.Println("quot: ", 8/2)
	fmt.Println("quot: ", 8/1.5)
	fmt.Println("rem: ", 8%2)
	// fmt.Println("rem: ", 8%2.5)
	fmt.Printf("Ratio : %v, type: %[1]T\n", ratio)
	fmt.Printf("Ratio : %v, type: %[1]T\n", ratio1)

}
