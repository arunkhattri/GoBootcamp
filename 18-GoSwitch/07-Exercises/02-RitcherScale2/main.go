package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Tell me the magnitude of the earhquake in human terms")
		return
	}

	r, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println("alien's richter scale is unknown")
		return
	}

	var msg string

	switch {
	case r < 2.0:
		msg = "micro's richter scale is less than 2.0"
	case r < 3.0:
		msg = "very minor's richter scale is 2 - 2.9"
	case r < 4.0:
		msg = "minor's richter scale is 3 - 3.9"
	case r < 5.0:
		msg = "Light's richter scale is 4 - 4.9"
	case r < 6.0:
		msg = "moderate's richter scale is 5 - 5.9"
	case r < 7.0:
		msg = "strong's richter scale is 6 - 6.9"
	case r < 8.0:
		msg = "major's richter scale is 7 - 7.9"
	case r < 10.0:
		msg = "great's richter scale is 8 - 9.9"
	default:
		msg = "massive's richter scale is 10+"
	}
	fmt.Println(msg)
}
