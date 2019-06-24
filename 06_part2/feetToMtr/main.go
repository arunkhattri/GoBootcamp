package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	const (
		feetInMeter float64 = 0.3048
		feetInYards         = feetInMeter / 0.9144
	)

	arg := os.Args[1]
	feet, _ := strconv.ParseFloat(arg, 64)

	meter := feet * feetInMeter
	yards := feet * feetInYards

	fmt.Printf("%g feet is %g meters\n", feet, meter)
	fmt.Printf("%g feet is %.2g yards\n", feet, yards)
	fmt.Printf("%g feet is %.3g yards\n", feet, yards)
	fmt.Printf("%g feet is %g yards\n", feet, yards)
}
