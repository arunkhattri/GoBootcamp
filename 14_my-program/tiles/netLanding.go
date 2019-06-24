// Package tiles contains utility function for working with tiles and related calculation.
package tiles

// Netlanding returns net landing price in sqft
func NetLanding(b, f float64) float64 {
	const gst float64 = 0.18
	res := b + (b * gst) + f
	return res
}
