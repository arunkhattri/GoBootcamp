package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	name := "arun"
	nameHindi := "अरुण"
	fmt.Println(len(name))
	fmt.Println(len(nameHindi))
	fmt.Println(utf8.RuneCountInString(nameHindi))
}
