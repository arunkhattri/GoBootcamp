package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("usage: [command] [string]")
		return
	}

	c, s := os.Args[1], os.Args[2]

	switch {
	case c == "lower":
		fmt.Println(strings.ToLower(s))
	case c == "upper":
		fmt.Println(strings.ToUpper(s))
	case c == "title":
		fmt.Println(strings.ToTitle(s))
	default:
		fmt.Printf("Unknown command: %q\n", c)
	}

}
