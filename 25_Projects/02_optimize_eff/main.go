package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Provide a directory")
		return
	}

	files, err := ioutil.ReadDir(args[0])
	if err != nil {
		fmt.Println(err)
	}

	// optimize the allocation of backing array of the name slices only once
	// program appends the names slice in a loop. Doing so allocates a new
	// backing array when its full.

	// heuristic, simple but wastes unnecessary memory

	// total := len(files) * 256       // average filename size + 1 newline char
	// names := make([]byte, 0, total) // length 0 to add elements at the beginning
	// fmt.Printf("Total required Space: %d bytes\n", total)

	// actually calculate the filename length, memory-wise efficient however wastes CPU time
	var total int

	for _, file := range files {
		// check if file has 0 bytes, its empty
		if file.Size() == 0 {
			total += len(file.Name()) + 1
		}
	}
	fmt.Printf("Total required Space: %d bytes\n", total)

	names := make([]byte, 0, total) // length 0 to add elements at the beginning

	for _, file := range files {
		// check if file has 0 bytes, its empty
		if file.Size() == 0 {
			name := file.Name()
			names = append(names, name...)
			names = append(names, '\n')
		}
	}
	// write file
	err = ioutil.WriteFile("output.txt", names, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s", names)
}
