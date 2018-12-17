package main

import (
	"fmt"
)

func celciusToFahrenheit(x float64) float64 {
	return (x*(9/5) + 32)
}

func main() {
	// var safe = true  // type inference
	// short declaration
	safe := true
	name := "Arun"
	age := 47
	temperature := 19.5
	tempF := .0
	tempF = celciusToFahrenheit(temperature)
	firstName, lastName := "Aarush", "Khattri"

	fmt.Println(name, age, safe)
	fmt.Printf("Temperature in Celcius is %.2f\n", temperature)
	fmt.Printf("Temperature in Fahrenheit is %.2f\n", tempF)
	fmt.Println(firstName, lastName)
}
