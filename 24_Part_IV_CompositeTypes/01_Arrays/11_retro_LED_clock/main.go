package main

// Project Retro LED Clock

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

func main() {

	type placeholder [5]string

	zero := placeholder{
		"███",
		"█ █",
		"█ █",
		"█ █",
		"███",
	}
	one := placeholder{
		"██ ",
		" █ ",
		" █ ",
		" █ ",
		"███",
	}
	two := placeholder{
		"███",
		"  █",
		"███",
		"█  ",
		"███",
	}
	three := placeholder{
		"███",
		"  █",
		"███",
		"  █",
		"███",
	}
	four := placeholder{
		"█ █",
		"█ █",
		"███",
		"  █",
		"  █",
	}
	five := placeholder{
		"███",
		"█  ",
		"███",
		"  █",
		"███",
	}
	six := placeholder{
		"███",
		"█  ",
		"███",
		"█ █",
		"███",
	}
	seven := placeholder{
		"███",
		"  █",
		"  █",
		"  █",
		"  █",
	}
	eight := placeholder{
		"███",
		"█ █",
		"███",
		"█ █",
		"███",
	}
	nine := placeholder{
		"███",
		"█ █",
		"███",
		"  █",
		"███",
	}
	colon := placeholder{
		"   ",
		" ░ ",
		"   ",
		" ░ ",
		"   ",
	}

	digits := [...]placeholder{
		zero, one, two, three, four, five, six, seven, eight, nine,
	}

	// for line := range digits[0] {
	//	// print a line for each placeholder in digits
	//	for digit := range digits {
	//		fmt.Print(digits[digit][line], " ")
	//	}
	//	fmt.Println()
	// }

	//  Get the current time
	screen.Clear()

	for {
		// move the cursor to top left corner
		screen.MoveTopLeft()

		currTime := time.Now()
		// fmt.Println("Current Time", currTime)

		// current hour, minute & second
		currHour := currTime.Hour()
		currMin := currTime.Minute()
		currSec := currTime.Second()
		// fmt.Printf("Current Hour: %d\n", currHour)
		// fmt.Printf("Current Minute: %d\n", currMin)
		// fmt.Printf("Current Second: %d\n", currSec)

		clock := [...]placeholder{
			digits[currHour/10], digits[currHour%10],
			colon,
			digits[currMin/10], digits[currMin%10],
			colon,
			digits[currSec/10], digits[currSec%10],
		}

		for line := range clock[0] {
			for index, digit := range clock {
				next := clock[index][line]
				// blink the colon
				if digit == colon && currSec%2 == 0 {
					next = "   "
				}
				fmt.Print(next, " ")
			}
			fmt.Println()
		}
		fmt.Println()
		time.Sleep(time.Second)

	}
}
