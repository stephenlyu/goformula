package function

import "math"

func sum(values Value, start, end int) float64 {
	sum := 0.0
	for i := start; i < end; i++ {
		sum += values.Get(i)
	}
	return sum
}

func ma_(values Value, start, end int) float64 {
	if start >= end {
		return 0
	}

	return sum(values, start, end) / float64(end - start)
}

func min(values Value, start, end int) float64 {
	if start >= end {
		return 0
	}

	ret := values.Get(start)
	for i := start + 1; i < end; i++ {
		if values.Get(i) < ret {
			ret = values.Get(i)
		}
	}
	return ret
}

func max(values Value, start, end int) float64 {
	if start >= end {
		return 0
	}

	ret := values.Get(start)
	for i := start + 1; i < end; i++ {
		if values.Get(i) > ret {
			ret = values.Get(i)
		}
	}
	return ret
}

func iif(condition bool, yesValue, noValue float64) float64 {
	if (condition) {
		return yesValue
	}
	return noValue
}

func round(val float64, places int) float64 {
	var t float64
	f := math.Pow10(places)
	x := val * f
	if math.IsInf(x, 0) || math.IsNaN(x) {
		return val
	}
	if x >= 0.0 {
		t = math.Ceil(x)
		if (t - x) > 0.50000000001 {
			t -= 1.0
		}
	} else {
		t = math.Ceil(-x)
		if (t + x) > 0.50000000001 {
			t -= 1.0
		}
		t = -t
	}
	x = t / f

	if !math.IsInf(x, 0) {
		return x
	}

	return t
}
