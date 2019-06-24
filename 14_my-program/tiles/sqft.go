// Package tiles contains utility function for working with tiles and related calculation.
package tiles

// Sqft returns coverage in a box for given sizes in mm and number of pcs in box.
func Sqft(w, l, c float64) float64 {
	const sqft float64 = 10.764

	res := ((w * l) / 10000 * c) / sqft
	return res

	// fmt.Printf("%gx%g coverage in SqFt.: %.2f\n", w, l, res)
}
