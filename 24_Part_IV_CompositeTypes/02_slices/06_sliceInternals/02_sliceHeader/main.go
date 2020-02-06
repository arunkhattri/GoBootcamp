package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

func main() {
	ages := []int{48, 42, 11, 7, 3}
	l, c, p := sliceHeader(ages)
	fmt.Printf("pntr: %s | len: %d | cap: %d\n", p, l, c)

}

func sliceHeader(slice ...interface{}) (int, int, string) {
	l := len(slice)
	c := cap(slice)
	// get pointer, convert into string
	temp := strconv.FormatInt(int64(uintptr(unsafe.Pointer(&slice))), 10)
	// get last 4 digit
	p := temp[len(temp)-4:]

	return l, c, p
}
