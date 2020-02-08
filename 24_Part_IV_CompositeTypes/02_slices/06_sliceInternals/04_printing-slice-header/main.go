package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

// Interface a type, typically a collection
type Interface interface {
	// number of elements in the collection
	Len() int
	// capacity of the collections
	Capacity() int
	// memory address in int of the collections
	Pntr() string
}

// convenient types for common cases

// IntSlice attaches the method of Interface to []int
type IntSlice []int

// Len method for IntSlice
func (p IntSlice) Len() int { return len(p) }

// Capacity method for IntSlice
func (p IntSlice) Capacity() int { return cap(p) }

// Pntr method for IntSlice
func (p IntSlice) Pntr() string {
	temp := strconv.FormatInt(int64(uintptr(unsafe.Pointer(&p))), 10)
	pnt := temp[len(temp)-4:]
	return pnt
}

// StringSlice attaches the method of Interface to []string
type StringSlice []string

// Len method for StringSlice
func (p StringSlice) Len() int { return len(p) }

// Capacity method for StringSlice
func (p StringSlice) Capacity() int { return cap(p) }

// Pntr method for StringSlice
func (p StringSlice) Pntr() string {
	temp := strconv.FormatInt(int64(uintptr(unsafe.Pointer(&p))), 10)
	pnt := temp[len(temp)-4:]
	return pnt
}

func sliceHeader(data Interface) string {

	length := data.Len()
	capacity := data.Capacity()
	pntr := data.Pntr()
	res := fmt.Sprintf("pntr: %s | len: %d | cap: %d\n", pntr, length, capacity)
	return res
}

func main() {
	ages := []int{48, 42, 11, 7, 3}
	names := []string{"arun", "minoo", "booboo", "inoo", "mia"}
	shA := sliceHeader(IntSlice(ages))
	shN := sliceHeader(StringSlice(names))
	fmt.Println(shA)
	fmt.Println(shN)

}
