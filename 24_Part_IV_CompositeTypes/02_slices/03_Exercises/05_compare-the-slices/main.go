package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	namesA := "Da Vinci, Wozniak, Carmack"
	namesB := []string{"Wozniak", "Da Vinci", "Carmack"}

	namesC := strings.Split(namesA, ", ")
	sort.Strings(namesB)
	sort.Strings(namesC)
	fmt.Printf("namesB : %v\n", namesB)
	fmt.Printf("namesB : %v\n", namesC)

	if len(namesB) == len(namesC) {
		for i := range namesC {
			if namesB[i] != namesC[i] {
				return
			}
		}
		fmt.Println("They are equal.")
	}

}
