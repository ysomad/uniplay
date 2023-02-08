package floatrounder

import "math"

// Round rounds float64 to 2 decimal places.
func Round(n float64) float64 {
	return math.Round(n*100) / 100
}
