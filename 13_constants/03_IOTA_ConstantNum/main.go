package main

import "fmt"

func main() {
	// iota is a built-on constant generator which generates
	// ever increasing number
	// const (
	// 	monday    = 0
	// 	tuesday   = 1
	// 	wednesday = 2
	// 	thursday  = 3
	// 	friday    = 4
	// 	saturday  = 5
	// 	sunday    = 6
	// )

	const (
		monday = iota
		tuesday
		wednesday
		thursday
		friday
		saturday
		sunday
	)
	fmt.Println(monday, tuesday, wednesday, thursday, friday, saturday, sunday)
}
