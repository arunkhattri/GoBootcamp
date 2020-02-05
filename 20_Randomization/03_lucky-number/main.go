package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// Produce a different random number in each run
// seed with a different value
// time.Now().UnixNano() yields a constantly changing number
const (
	maxTurns = 5
	usage    = `Welcome to the Lucky Number Game!
The program will pick %d random numbers.
Your mission is to guess two those numbers.

The greater your numbers is, the harder it gets.

Wanna play?`
)

func main() {
	rand.Seed(time.Now().UnixNano())
	args := os.Args[1:]

	// check if number is given
	if len(args) != 2 {
		fmt.Printf(usage, maxTurns)
		return
	}

	guess1, err1 := strconv.Atoi(args[0])
	guess2, err2 := strconv.Atoi(args[1])

	// if not a number
	if err1 != nil && err2 != nil {
		fmt.Println("Not a number.")
		return
	}

	// guess < 0
	if guess1 < 0 && guess2 < 0 {
		fmt.Println("Please pick a positive number.")
		return
	}

	for turn := 0; turn < maxTurns; turn++ {
		maxGuess := math.Max(float64(guess1), float64(guess2))
		n := rand.Intn(int(maxGuess) + 1)
		if turn == 0 && (n == guess1 || n == guess2) {
			fmt.Println("[^-^] Bravo! you did it in first turn...")
			return
		}

		// Case: random number is equal to guess
		if n == guess1 || n == guess2 {
			// emoji.Println(":tada: YOU WIN!")
			fmt.Println("[!!!] YOU WIN!!!")
			return
		}
	}
	// Case: random number != guess
	// emoji.Println(":skull: YOU LOST... Try Again?")
	fmt.Println("[;w;] YOU LOST... Try Again?")
}
