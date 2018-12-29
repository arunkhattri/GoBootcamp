package main

import (
	"fmt"
	"path"
)

func main() {
	var dir, file string

	dir, file = path.Split("goBootCamp/main.go")
	fmt.Printf("dir = %s\nfile = %s\n", dir, file)

	// unwanted variables
	var filename string
	_, filename = path.Split("gopher/main.go") // blank identifier
	// also could be done via short assignment
	fmt.Printf("filename = %s\n", filename)
}
