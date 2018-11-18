package main

import "fmt"

const ok = true // package scoped

func main() {
	var hello = "Hello!" // block scoped
	fmt.Println(hello, ok)

}
