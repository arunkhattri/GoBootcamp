package main

import "fmt"

func main() {
	width, height := 5., 12.
	area := width * height
	fmt.Printf("area = %g\n", area)

	area = area - 10 // decrease the area by 10
	area = area + 10 // increase the area by 10
	fmt.Printf("area = %g\n", area)

	area = area * 2 // double the area
	area = area / 2 // half the area
	fmt.Printf("area = %g\n", area)

	// assignment operator
	area -= 10 // decrease the area by 10
	fmt.Printf("area = %g\n", area)
	area += 10 // increase the area by 10
	area *= 2  //  double the area
	fmt.Printf("area=%g\n", area)

	fmt.Printf("area = %g\n", area)
	area /= 2 //  half the area	fmt.Printf("area=%g\n", area)

	fmt.Printf("area = %g\n", area)

	// remainder of area

	//area %= 7
	area = float64(int(area) % 7)
	fmt.Printf("area = %g\n", area)

}
