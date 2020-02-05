package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	args := os.Args[1:]

	// if no argument given; return
	if len(args) != 1 {
		fmt.Println("[-] Your Name")
		return
	}

	name := args[0]

	moods := [...]string{
		"happy 😀", "good 👍", "awesome 😎",
		"sad 😞", "bad 👎", "terrible 😟",
	}

	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(len(moods))

	fmt.Printf("%s feels %s\n", name, moods[n])

}
