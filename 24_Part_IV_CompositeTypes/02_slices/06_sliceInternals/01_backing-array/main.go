package main

import (
	"fmt"
	"unsafe"
)

func main() {
	ages := []int{35, 15, 25}
	grades := [...]float64{40, 10, 20, 50, 60, 70}
	pA := (uintptr(unsafe.Pointer(&ages)))
	pG := uintptr(unsafe.Pointer(&grades))
	// last4Digits := int64(pA) / 10000

	fmt.Printf("len: %d | cap: %d | mem: %d\n", len(grades), cap(grades), pG)
	fmt.Printf("len: %d | cap: %d | mem: %d\n", len(ages), cap(ages), pA)

	front := grades[:3]
	pF := uintptr(unsafe.Pointer(&front))
	fmt.Printf("len: %d | cap: %d | mem: %d\n", len(front), cap(front), pF)

	firstAge := ages[:1]
	pFA := uintptr(unsafe.Pointer(&firstAge))
	fmt.Printf("len: %d | cap: %d | mem: %d\n", len(firstAge), cap(firstAge), pFA)
}
