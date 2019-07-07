package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// s := strconv.Itoa(42)
	// fmt.Println(s)
	n, err := strconv.Atoi(os.Args[1])
	fmt.Println("Converted number    :", n)
	fmt.Println("Returned error value:", err)
}
