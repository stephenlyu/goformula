package function

import (
	_math "github.com/chanxuehong/util/math"
	"math"
)


// LLV

type llv struct {
	funcbase
	N Value
}

func (this llv) BuildValueAt(index int) float64 {
	end := index + 1
	start := int(_math.Max(0, int64(end - int(this.N.Get(index)))))

	return min(this.data, start, end)
}

func (this *llv) UpdateLastValue() {
	updateLastValue(this)
}

func LLV(data Value, N Value) *llv {
	ret := &llv{
		funcbase: funcbase {
			data: data,
		},
		N: N,
	}
	ret.Values = initValues(ret)
	return ret
}

// LLVBARS

type llvbars struct {
	funcbase
	N Value
}

func (this llvbars) BuildValueAt(index int) float64 {
	end := index + 1
	start := int(_math.Max(0, int64(end - int(this.N.Get(index)))))

	low := math.MaxFloat64
	low_pos := -1
	for j := start; j < end; j++ {
		v := this.data.Get(j)
		if v < low {
			low = v
			low_pos = j
		}
	}

	return float64(end - 1 - low_pos)
}

func (this *llvbars) UpdateLastValue() {
	updateLastValue(this)
}

func LLVBARS(data Value, N Value) *llvbars {
	ret := &llvbars{
		funcbase: funcbase {
			data: data,
		},
		N: N,
	}
	ret.Values = initValues(ret)
	return ret
}

// HHV

type hhv struct {
	funcbase
	N Value
}

func (this hhv) BuildValueAt(index int) float64 {
	end := index + 1
	start := int(_math.Max(0, int64(end - int(this.N.Get(index)))))

	return max(this.data, start, end)
}

func (this *hhv) UpdateLastValue() {
	updateLastValue(this)
}

func HHV(data Value, N Value) *hhv {
	ret := &hhv{
		funcbase: funcbase {
			data: data,
		},
		N: N,
	}
	ret.Values = initValues(ret)
	return ret
}

// HHVBARS

type hhvbars struct {
	funcbase
	N Value
}

func (this hhvbars) BuildValueAt(index int) float64 {
	end := index + 1
	start := int(_math.Max(0, int64(end - int(this.N.Get(index)))))

	high := this.data.Get(start)
	high_pos := -1
	for j := start + 1; j < end; j++ {
		v := this.data.Get(j)
		if v > high {
			high = v
			high_pos = j
		}
	}

	return float64(end - 1 - high_pos)
}

func (this *hhvbars) UpdateLastValue() {
	updateLastValue(this)
}

func HHVBARS(data Value, N Value) *hhvbars {
	ret := &hhvbars{
		funcbase: funcbase {
			data: data,
		},
		N: N,
	}
	ret.Values = initValues(ret)
	return ret
}

// STD

type std struct {
	funcbase
	N Value
}

func (this std) BuildValueAt(index int) float64 {
	N := int(this.N.Get(index))
	if index < N - 1 {
		return 0
	}

	end := index + 1
	start := int(_math.Max(0, int64(end - N)))

	average := ma_(this.data, start, end)

	sum_ := 0.0
	for i := start; i < end; i++ {
		v := this.data.Get(i)
		diff := v - average
		sum_ += diff * diff
	}

	return math.Sqrt(sum_ / float64(end - start))
}

func (this *std) UpdateLastValue() {
	updateLastValue(this)
}

func STD(data Value, N Value) *std {
	ret := &std{
		funcbase: funcbase {
			data: data,
		},
		N: N,
	}
	ret.Values = initValues(ret)
	return ret
}

// AVEDEV

type avedev struct {
	funcbase
	N Value
}

func (this avedev) BuildValueAt(index int) float64 {
	N := int(this.N.Get(index))
	if index < N - 1 {
		return 0
	}

	end := index + 1
	start := int(_math.Max(0, int64(end - N)))

	count := end - start
	average := ma_(this.data, start, end)

	vsum := 0.0
	for i := start; i < end; i++ {
		v := this.data.Get(i)
		diff := math.Abs(v - average)
		vsum += diff
	}

	return vsum / float64(count)
}

func (this *avedev) UpdateLastValue() {
	updateLastValue(this)
}

func AVEDEV(data Value, N Value) *avedev {
	ret := &avedev{
		funcbase: funcbase {
			data: data,
		},
		N: N,
	}
	ret.Values = initValues(ret)
	return ret
}

// SUM

type sumf struct {
	funcbase
	N Value
}

func (this sumf) BuildValueAt(index int) float64 {
	end := index + 1
	start := int(_math.Max(0, int64(end - int(this.N.Get(index)))))
	return sum(this.data, start, end)
}

func (this *sumf) UpdateLastValue() {
	updateLastValue(this)
}

func SUM(data Value, N Value) *sumf {
	ret := &sumf{
		funcbase: funcbase {
			data: data,
		},
		N: N,
	}
	ret.Values = initValues(ret)
	return ret
}

// CROSS

type cross struct {
	funcbase
	data1 Value
	N Value
}

func (this cross) BuildValueAt(index int) float64 {
	if index == 0 {
		return 0
	}

	return iif(this.data.Get(index - 1) < this.data1.Get(index - 1) && this.data.Get(index) >= this.data1.Get(index), 1, 0)
}

func (this *cross) UpdateLastValue() {
	updateLastValue(this)
}

func CROSS(data, data1 Value, N Value) *cross {
	ret := &cross{
		funcbase: funcbase {
			data: data,
		},
		data1: data1,
		N: N,
	}
	ret.Values = initValues(ret)
	return ret
}

// COUNT

type count struct {
	funcbase
	N Value
}

func (this count) BuildValueAt(index int) float64 {
	N := int(this.N.Get(index))

	end := index + 1
	start := int(_math.Max(0, int64(end - N)))

	c := 0
	for i := start; i < end; i++ {
		v := this.data.Get(i)
		if v != 0 {
			c++
		}
	}

	return float64((c))
}

func (this *count) UpdateLastValue() {
	updateLastValue(this)
}

func COUNT(data Value, N Value) *count {
	ret := &count{
		funcbase: funcbase {
			data: data,
		},
		N: N,
	}
	ret.Values = initValues(ret)
	return ret
}

// IF

type iff struct {
	funcbase
	yesData Value
	noData Value
}

func (this iff) BuildValueAt(index int) float64 {
	return iif(this.data.Get(index) != 0, this.yesData.Get(index), this.noData.Get(index))
}

func (this *iff) UpdateLastValue() {
	updateLastValue(this)
}

func IF(data, yesData, noData Value) *iff {
	ret := &iff{
		funcbase: funcbase {
			data: data,
		},
		yesData: yesData,
		noData: noData,
	}
	ret.Values = initValues(ret)
	return ret
}

// EVERY

type every struct {
	funcbase
	N Value
}

func (this every) BuildValueAt(index int) float64 {
	N := int(this.N.Get(index))

	end := index + 1
	start := int(_math.Max(0, int64(end - N)))

	for i := start; i < end; i++ {
		if this.data.Get(i) == 0 {
			return 0
		}
	}

	return 1
}

func (this *every) UpdateLastValue() {
	updateLastValue(this)
}

func EVERY(data Value, N Value) *every {
	ret := &every{
		funcbase: funcbase {
			data: data,
		},
		N: N,
	}
	ret.Values = initValues(ret)
	return ret
}

// BARSLAST

type barslast struct {
	funcbase
	N Value
}

func (this barslast) BuildValueAt(index int) float64 {
	for j := index; j >= 0; j-- {
		if this.data.Get(j) != 0 {
			return float64(index - j)
		}
	}
	// TODO: BUGGY?
	return float64(index + 1)
}

func (this *barslast) UpdateLastValue() {
	updateLastValue(this)
}

func BARSLAST(data Value, N Value) *barslast {
	ret := &barslast{
		funcbase: funcbase {
			data: data,
		},
		N: N,
	}
	ret.Values = initValues(ret)
	return ret
}

// BARSCOUNT

type barscount struct {
	scalar
	data Value
}

func (this barscount) Get(index int) float64 {
	return float64(this.data.Len())
}

func BARSCOUNT(data Value) *barscount {
	ret := &barscount{
		data: data,
	}
	return ret
}

// ISLASTBAR

type islastbar struct {
	scalar
	data Value
}

func (this islastbar) Get(index int) float64 {
	return iif(this.data.Len() - 1 == index, 1, 0)
}

func ISLASTBAR(data Value) *islastbar {
	ret := &islastbar{
		data: data,
	}
	return ret
}

// ROUND2

type roundf struct {
	funcbase
	N Value
}

func (this roundf) BuildValueAt(index int) float64 {
	return round(this.data.Get(index), int(this.N.Get(index)))
}

func (this *roundf) UpdateLastValue() {
	updateLastValue(this)
}

func ROUND2(data Value, N Value) *roundf {
	if N != nil {
		N = Scalar(2)
	}

	ret := &roundf{
		funcbase: funcbase {
			data: data,
		},
		N: N,
	}
	ret.Values = initValues(ret)
	return ret
}

// REF

type ref struct {
	funcbase
	N Value
}

func (this ref) BuildValueAt(index int) float64 {
	N := int(this.N.Get(index))
	if index < N {
		return 0
	}

	return this.data.Get(index - N)
}

func (this *ref) UpdateLastValue() {
	updateLastValue(this)
}

func REF(data Value, N Value) *ref {
	ret := &ref{
		funcbase: funcbase {
			data: data,
		},
		N: N,
	}
	ret.Values = initValues(ret)
	return ret
}

// MIN

type minf struct {
	funcbase
	data1 Value
}

func (this minf) BuildValueAt(index int) float64 {
	return math.Min(this.data.Get(index), this.data1.Get(index))
}

func (this *minf) UpdateLastValue() {
	updateLastValue(this)
}

func MIN(data Value, data1 Value) *minf {
	ret := &minf{
		funcbase: funcbase {
			data: data,
		},
		data1: data1,
	}
	ret.Values = initValues(ret)
	return ret
}

// MAX

type maxf struct {
	funcbase
	data1 Value
}

func (this maxf) BuildValueAt(index int) float64 {
	return math.Max(this.data.Get(index), this.data1.Get(index))
}

func (this *maxf) UpdateLastValue() {
	updateLastValue(this)
}

func MAX(data Value, data1 Value) *maxf {
	ret := &maxf{
		funcbase: funcbase {
			data: data,
		},
		data1: data1,
	}
	ret.Values = initValues(ret)
	return ret
}

// ABS

type absf struct {
	funcbase
}

func (this absf) BuildValueAt(index int) float64 {
	return math.Abs(this.data.Get(index))
}

func (this *absf) UpdateLastValue() {
	updateLastValue(this)
}

func ABS(data Value) *absf {
	ret := &absf{
		funcbase: funcbase {
			data: data,
		},
	}
	ret.Values = initValues(ret)
	return ret
}

// SLOPE

type slopef struct {
	funcbase
	N Value
}

func (this slopef) BuildValueAt(index int) float64 {
	N := int(this.N.Get(index))
	if index < N - 1 {
		return 0
	}

	x := make([]float64, N)
	y := make([]float64, N)

	for i := 0; i < N; i++ {
		x[i] = float64((i + 1))
		y[i] = this.data.Get(index + 1 - (N - i))
	}

	_, slope, _ := LinearRegression(Vector(x), Vector(y))

	return slope
}

func (this *slopef) UpdateLastValue() {
	updateLastValue(this)
}

func SLOPE(data Value, N Value) *slopef {
	ret := &slopef{
		funcbase: funcbase {
			data: data,
		},
		N: N,
	}
	ret.Values = initValues(ret)
	return ret
}