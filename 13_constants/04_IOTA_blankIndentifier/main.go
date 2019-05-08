package main

import "fmt"

func main() {
	// const (
	// 	EST = -(5 + iota)
	// 	_   // blank identifier
	// 	MST
	// 	PST
	// )
	const (
		yes = (iota * 5) + 2
		no
		both
	)
	// fmt.Println(EST, MST, PST)
	fmt.Println(yes, no, both)
}
