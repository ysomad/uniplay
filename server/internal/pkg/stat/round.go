package stat

import "math"

// round rounds float64 to 2 decimal places.
func round(n float64) float64 { return math.Round(n*100) / 100 }
