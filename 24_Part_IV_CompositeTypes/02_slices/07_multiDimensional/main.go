package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// spendings
	// 1st day: ₹200, ₹100
	// 2nd day: ₹500
	// 3rd day: ₹50, ₹25, ₹75
	// spendings := [][]int{
	// 	{200, 100},
	// 	{500},
	// 	{50, 25, 75},
	// }

	// let's say we want to fetch our spendings for last 5 days from a file
	// spendings := make([][]int, 0, 5)
	// spendings = append(spendings, []int{200, 100})         // 1st day
	// spendings = append(spendings, []int{500})              // 2nd day
	// spendings = append(spendings, []int{50, 25, 75})       // 3rd day
	// spendings = append(spendings, []int{150, 125, 35, 65}) // 4th day
	// spendings = append(spendings, []int{50, 125, 85})      // 5th day

	// usign a function
	spendings := fetch()

	for i, daily := range spendings {
		var total int
		for _, spent := range daily {
			total += spent
		}
		fmt.Printf("Day %d: %d\n", i+1, total)
	}

}

func fetch() [][]int {
	content := `200 100
500
50 25 75
150 125 35 65
50 125 85`

	lines := strings.Split(content, "\n")
	spendings := make([][]int, len(lines))

	for i, line := range lines {
		// fmt.Printf("%d: %#v\n", i+1, line)
		fields := strings.Fields(line)
		spendings[i] = make([]int, len(fields))
		// let's see what is there in fields
		for j, field := range fields {
			// fmt.Printf("\t%d: %#v\n", j+1, field)
			spending, _ := strconv.Atoi(field)
			spendings[i][j] = spending
		}
	}
	return spendings

}
