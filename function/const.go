package function

import "math"

var NaN = math.NaN()

func IsNaN(v float64) bool {
	return math.IsNaN(v)
}
