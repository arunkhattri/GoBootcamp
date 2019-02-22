package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("hello, " + "how " + "are you " + "today?")
	fmt.Println(`hello, ` + "how " + `are you ` + `today?`)
	fmt.Println(`hello, ` + "how " + `are you ` + "today?")

	eq := "1 + 2 = "
	sum := 1 + 2
	fmt.Println(eq + strconv.Itoa(sum))

	bEq := strconv.FormatBool(true) + " " + strconv.FormatBool(false)
	fmt.Println(bEq)
}
