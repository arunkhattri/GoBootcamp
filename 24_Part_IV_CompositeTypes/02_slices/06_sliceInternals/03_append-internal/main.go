package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

func main() {
	words := []string{"lucy", "in", "the", "sky", "with", "diamonds"}
	fmt.Println(sliceHeader(words))
	fmt.Println(words)

	words = append(words[:3], "crystals")
	fmt.Println(sliceHeader(words))
	fmt.Println(words)

	words = append(words[:1], "is", "everywhere")
	fmt.Println(sliceHeader(words))
	fmt.Println(words)

	words = append(words, words[len(words)+1:cap(words)]...)
	fmt.Println(sliceHeader(words))
	fmt.Println(words)
}

func sliceHeader(slice []string) string {
	temp := strconv.FormatInt(int64(uintptr(unsafe.Pointer(&slice))), 10)
	p := temp[len(temp)-4:]
	l := len(slice)
	c := cap(slice)
	res := fmt.Sprintf("pntr: %s | len: %d | cap: %d\n", p, l, c)
	return res
}
