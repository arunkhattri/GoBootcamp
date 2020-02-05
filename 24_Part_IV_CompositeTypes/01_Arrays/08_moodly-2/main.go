package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	args := os.Args[1:]

	// if no argument given; return
	if len(args) != 2 {
		fmt.Println("[-] Your Name and Mood")
		return
	}

	name, mood := args[0], strings.ToLower(args[1])

	moods := [...][3]string{
		{"happy 😀", "good 👍", "awesome 😎"},
		{"sad 😞", "bad 👎", "terrible 😟"},
	}

	rand.Seed(time.Now().UnixNano())

	n := rand.Intn(len(moods[0]))

	// if mood == "positive" {
	//	fmt.Printf("%s feels %s\n", name, moods[0][n])

	// } else {
	//	fmt.Printf("%s feels %s\n", name, moods[1][n])
	// }

	// another way to do it
	var mi int // by default initiated to 0
	if mood != "positive" {
		mi = 1
	}

	fmt.Printf("%s feels %s\n", name, moods[mi][n])

}
