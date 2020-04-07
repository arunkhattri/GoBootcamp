package main

import (
	s "github.com/inancgumus/prettyslice"
)

func main() {
	// 1. Create a string slice, names wiht a length and capacity of 5 and print it
	names := make([]string, 5, 5)
	s.Show("names", names)

	// 2. Append the following names to the names slice
	// "einstein", "tesla", "aristo"
	names = append(names, []string{"einstein", "tesla", "aristo"}...)
	s.Show("2nd step", names)

	// 3. Append the following naems to the beginning of the names slice:
	// "einstein", "tesla", "aristo"
	names = make([]string, 0, 5)
	names = append(names, []string{"einstein", "tesla", "aristo"}...)
	s.Show("3rd step", names)

	// 4. copy only the first two elements of the following
	// array to the last two elements of the names slice.
	moreNames := [...]string{"plato", "khayyam", "ptolemy"}
	copy(names[3:5], moreNames[:2])
	names = names[:cap(names)]
	s.Show("4th step", names)

	// 5. Only copy the last 3 elements of the names slice to a new slice: clone
	// Append the first two elements of the names to the clone
	// clone := append([]string(nil), names[2:]...)
	clone := make([]string, 3, 5)
	copy(clone, names[len(names)-3:])
	s.Show("clone", clone)
	clone = append(clone, names[:2]...)
	s.Show("5th step", clone)

	// 6. slice the clone slice between [2, 4)
	sliced := clone[2:4]
	// append "hypatia" to the sliced
	sliced = append(sliced, "hypatia")
	s.Show("6th Step", sliced)
	// change the 3rd element of the clone slice to "elder"
	clone[2] = "elder"
	s.Show("6th step...", clone, sliced)

}
