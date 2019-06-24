package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/arunkhattri/GoBootcamp/14_my-program/tiles"
)

// func netLanding(b, f float64) float64 {
// 	const gst float64 = 0.18
// 	res := b + (b * gst) + f
// 	return res
// }

// func sqft(w, l, c float64) float64 {
// 	const sqft float64 = 10.764

// 	res := ((w * l) / 10000 * c) / sqft
// 	return res
// }

func main() {
	w, _ := strconv.ParseFloat(os.Args[1], 64)
	l, _ := strconv.ParseFloat(os.Args[2], 64)
	c, _ := strconv.ParseFloat(os.Args[3], 64)

	tc := tiles.Sqft(w, l, c)
	fmt.Printf("Coverage %gx%g: %.2f\n", w, l, tc)

	var (
		basic   float64
		freight float64
	)

	scanner := bufio.NewReader(os.Stdin)
	r, _ := scanner.ReadString('\n')
	// for scanner.Scan() {
	// 	r := scanner.Text()
	// }
	// if err := scanner.Err(); err != nil {
	// 	fmt.Fprintln(os.Stderr, "reading standart input:", err)
	// }
	_, err := fmt.Fscanf(r, "%.2f %.2f", &basic, &freight)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Fscanf: %v\n", err)
	}

	nl := tiles.NetLanding(basic, freight)
	fmt.Printf("Net Landing (Sq.ft.): ₹ %.2f\n", nl)
}
