package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// const max = 5
	max, _ := strconv.Atoi(os.Args[1])

	// var uniques [max]int
	var uniques []int // as slices

loop:
	// for found := 0; found < max; {
	for len(uniques) < max {
		n := rand.Intn(max) + 1
		fmt.Print(n, " ")

		for _, u := range uniques {
			if u == n {
				continue loop
			}
		}
		// uniques[found] = n
		uniques = append(uniques, n)
		// found++
	}
	fmt.Println("\nuniques:", uniques)
	sort.Ints(uniques)
	fmt.Println("sorted uniques:", uniques)

	nums := [5]int{5, 4, 3, 2, 1}
	sort.Ints(nums[:])
	fmt.Println("\nnums:", nums)

	var tasks []string
	fmt.Println(len(tasks))

}
