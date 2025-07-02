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

func IsTrue[T float32 | float64](v T) bool {
	return v != 0 && !math.IsNaN(float64(v))
}

func sum[T float32 | float64](values Value[T], start, end int) T {
	sum := T(0.0)
	for i := start; i < end; i++ {
		sum += values.Get(i)
	}
	return sum
}

func ma_[T float32 | float64](values Value[T], start, end int) T {
	if start >= end {
		return 0
	}

	return sum[T](values, start, end) / T(end-start)
}

func minv[T float32 | float64](v1 T, v2 T) T {
	if v1 <= v2 {
		return v1
	}
	return v2
}

func maxv[T float32 | float64](v1 T, v2 T) T {
	if v1 <= v2 {
		return v2
	}
	return v1
}

func min[T float32 | float64](values Value[T], start, end int) T {
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

func max[T float32 | float64](values Value[T], start, end int) T {
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

func iif[T float32 | float64](condition bool, yesValue, noValue T) T {
	if condition {
		return yesValue
	}
	return noValue
}

func round[T float32 | float64](val T, places int) T {
	var t T
	f := T(math.Pow10(places))
	x := val * f
	if math.IsInf(float64(x), 0) || math.IsNaN(float64(x)) {
		return val
	}
	if x >= 0.0 {
		t = T(math.Ceil(float64(x)))
		if (t - x) > 0.50000000001 {
			t -= 1.0
		}
	} else {
		t = T(math.Ceil(-float64(x)))
		if (t + x) > 0.50000000001 {
			t -= 1.0
		}
		t = -t
	}
	x = t / f

	if !math.IsInf(float64(x), 0) {
		return x
	}

	return t
}

func LinearRegression[T float32 | float64](x, y Value[T]) (err error, slope T, intercept T) {
	if x.Len() != y.Len() {
		return errors.New("Length mismatch"), 0, 0
	}

	// Placeholder for the math to be done
	var sum [5]T

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
	f := T(i)
	slope = (f*sum[3] - sum[0]*sum[1]) / (f*sum[2] - sum[0]*sum[0])
	intercept = (sum[1] / f) - (slope * sum[0] / f)

	err = nil
	return
}

func Interpolate[T float32 | float64](values []T, from int, to int, fromValue T, toValue T) {
	Assert(from >= 0, "from >= 0 required")
	Assert(to < len(values), "to < len(values) required")

	if math.IsNaN(float64(fromValue)) || math.IsNaN(float64(toValue)) {
		for i := from + 1; i < to; i++ {
			values[i] = T(math.NaN())
		}
		return
	}

	if from == to {
		values[from] = fromValue
		return
	}

	slope := (toValue - fromValue) / T(to-from)
	values[from] = fromValue
	values[to] = toValue
	for i := from + 1; i < to; i++ {
		values[i] = fromValue + T(i-from)*slope
	}
}
