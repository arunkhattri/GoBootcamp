package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
prime:
	for _, arg := range os.Args[1:] {
		n, err := strconv.Atoi(arg)
		if err != nil {
			// skip non-numerics
			continue
		}
		switch {
		// prime
		case n == 2 || n == 3:
			fmt.Print(n, " ")
			continue
		// not prime
		case n <= 1 || n%2 == 0 || n%3 == 0:
			continue
		}

		for i, w := 5, 2; i*i <= n; {
			// not prime
			if n%i == 0 {
				continue prime
			}
			i += w
			w = 6 - w
		}
		// all checks ok: it's a prime
		fmt.Print(n, " ")
	}
	fmt.Println()
}
