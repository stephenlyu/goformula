package function

import (
	"errors"
	"math"
)

func Assert(cond bool, msg string) {
	if !cond {
		panic(msg)
	}
}

func IsTrue(v float64) bool {
	return v != 0 && !math.IsNaN(v)
}

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

	return sum(values, start, end) / float64(end-start)
}

func min(values Value, start, end int) float64 {
	if start >= end {
		return 0
	}

	ret := values.Get(start)
	for i := start; i < end; i++ {
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
	for i := start; i < end; i++ {
		if values.Get(i) > ret {
			ret = values.Get(i)
		}
	}
	return ret
}

func iif(condition bool, yesValue, noValue float64) float64 {
	if condition {
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

func LinearRegression(x, y Value) (err error, slope float64, intercept float64) {
	if x.Len() != y.Len() {
		return errors.New("Length mismatch"), 0, 0
	}

	// Placeholder for the math to be done
	var sum [5]float64

	// Loop over data keeping index in place
	i := 0
	for ; i < x.Len(); i++ {
		sum[0] += x.Get(i)
		sum[1] += y.Get(i)
		sum[2] += x.Get(i) * x.Get(i)
		sum[3] += x.Get(i) * y.Get(i)
		sum[4] += y.Get(i) * y.Get(i)
	}

	// Find gradient and intercept
	f := float64(i)
	slope = (f*sum[3] - sum[0]*sum[1]) / (f*sum[2] - sum[0]*sum[0])
	intercept = (sum[1] / f) - (slope * sum[0] / f)

	err = nil
	return
}

func Interpolate(values []float64, from int, to int, fromValue float64, toValue float64) {
	Assert(from >= 0, "from >= 0 required")
	Assert(to < len(values), "to < len(values) required")

	if math.IsNaN(fromValue) || math.IsNaN(toValue) {
		for i := from + 1; i < to; i++ {
			values[i] = math.NaN()
		}
		return
	}

	if from == to {
		values[from] = fromValue
		return
	}

	slope := (toValue - fromValue) / float64(to-from)
	values[from] = fromValue
	values[to] = toValue
	for i := from + 1; i < to; i++ {
		values[i] = fromValue + float64(i-from)*slope
	}
}
