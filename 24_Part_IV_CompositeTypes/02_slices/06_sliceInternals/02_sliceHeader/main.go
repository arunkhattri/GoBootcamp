package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

func main() {
	ages := []int{48, 42, 11, 7, 3}
	l, c, p := sliceHeader(ages)
	fmt.Println(p, l, c)
	header := fmt.Sprintf("pntr: %s | len: %d | cap: %d\n", p, l, c)
	fmt.Println(header)
	fmt.Printf("len: %d\n", len(ages))
	fmt.Printf("cap: %d\n", cap(ages))

	// part := ages[:]

	// for cap(part) != 0 {
	// 	part = part[1:cap(part)]
	// 	fmt.Println(sliceHeader(part))
	// 	fmt.Println(part)
	// }

}

func sliceHeader(slice []int) (int, int, string) {
	l := len(slice)
	c := cap(slice)
	// get pointer, convert into string
	temp := strconv.FormatInt(int64(uintptr(unsafe.Pointer(&slice))), 10)
	// get last 4 digit
	p := temp[len(temp)-4:]

	// header := fmt.Sprintf("pntr: %s | len: %d | cap: %d\n", p, l, c)
	return l, c, p
}
