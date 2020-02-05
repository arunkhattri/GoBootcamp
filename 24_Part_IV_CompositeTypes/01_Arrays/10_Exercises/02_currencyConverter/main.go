package main

// ---------------------------------------------------------
// EXERCISE: Currency Converter
//
//   In this exercise, you're going to display currency exchange ratios
//   against USD.
//
//   1. Declare a few constants with iota. They're going to be the keys
//      of the array.
//
//   2. Create an array that contains the conversion ratios.
//
//      You should use keyed elements and the contants you've declared before.
//
//   3. Get the USD amount to be converted from the command line.
//
//   4. Handle the error cases for missing or invalid input.
//
//   5. Print the exchange ratios.
// ---------------------------------------------------------

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	const (
		EUR = iota
		USD
		GBP
		JPY
		BDT
	)
	// currency converter INR to ---
	currencyConverter := [...]float64{
		EUR: 79.31,
		USD: 71.50,
		GBP: 93.78,
		JPY: 0.66,
		BDT: 1.19,
	}

	args := os.Args[1:]

	// sanity check: there is args
	if len(args) != 1 {
		fmt.Println("Please provide the amount to be converted.")
		return
	}

	num, err := strconv.ParseFloat(args[0], 64)

	// sanity check: args is a number
	if err != nil {
		fmt.Println("Please provide a number")
		return
	}

	fmt.Printf("%.2f INR is %.2f EUR\n", num, currencyConverter[EUR]*num)
	fmt.Printf("%.2f INR is %.2f USD\n", num, currencyConverter[USD]*num)
	fmt.Printf("%.2f INR is %.2f GBP\n", num, currencyConverter[GBP]*num)
	fmt.Printf("%.2f INR is %.2f JPY\n", num, currencyConverter[JPY]*num)
	fmt.Printf("%.2f INR is %.2f BDT\n", num, currencyConverter[BDT]*num)

}
