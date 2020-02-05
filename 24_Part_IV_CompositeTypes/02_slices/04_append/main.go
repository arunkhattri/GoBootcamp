package main

import "fmt"

func main() {

	var todo []string

	todo = append(todo, "sing")

	fmt.Printf("todo: %v\n", todo)

	tomorrow := []string{"see mom", "learn go"}
	todo = append(todo, tomorrow...)

	fmt.Printf("todo: %v\n", todo)

}
