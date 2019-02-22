package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	msg := os.Args[1]
	l := len(msg) // works with only non-unicode chars

	sr := strings.Repeat("!", l)
	s := sr + msg + sr
	s = strings.ToUpper(s)
	fmt.Println(s)
}
