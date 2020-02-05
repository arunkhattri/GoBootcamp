package main

import (
	"fmt"
)

func main() {
	const (
		ETH = iota
		WAN
	)

	rates := [...]float64{
		ETH: 25.5,
		WAN: 120.5,
	}

	fmt.Printf("1 BTC is %g ETH\n", rates[ETH])
	fmt.Printf("1 BTC is %g WAN\n", rates[WAN])

}
