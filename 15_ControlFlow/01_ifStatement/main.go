package main

import "fmt"

func main() {
	score := 14
	valid := true

	if score > 20 && valid {
		fmt.Println("good")
	} else if score == 20 {
		fmt.Println("on the edge")
	} else if score < 20 && score >= 15 {
		fmt.Println("average")
	} else {
		fmt.Println("low")
	}
}
