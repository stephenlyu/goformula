package function

import (
	"math"

	_math "github.com/chanxuehong/util/math"
)

// LLV

type llv struct {
	funcbase
	N Value
}

func (this llv) BuildValueAt(index int) float64 {
	N := this.N.Get(index)
	if math.IsNaN(N) || N <= 0 {
		return math.NaN()
	}
	end := index + 1
	start := int(_math.Max(0, int64(end-int(N))))

	return min(this.data, start, end)
}

func (this *llv) UpdateLastValue() {
	updateLastValue(this)
}

func LLV(data Value, N Value) Value {
	if data.IsScalar() {
		return Scalar(data.Get(0))
	}

	ret := &llv{
		funcbase: funcbase{
			data: data,
		},
		N: N,
	}
	ret.Values = make([]float64, data.Len())
	initValues(ret, ret.Values)
	return ret
}

// LLVBARS

type llvbars struct {
	funcbase
	N Value
}

func (this llvbars) BuildValueAt(index int) float64 {
	N := this.N.Get(index)
	if math.IsNaN(N) || N <= 0 {
		return float64(index)
	}

	end := index + 1
	start := int(_math.Max(0, int64(end-int(N))))

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

func LLVBARS(data Value, N Value) Value {
	if data.IsScalar() {
		return Scalar(0)
	}

	ret := &llvbars{
		funcbase: funcbase{
			data: data,
		},
		N: N,
	}
	ret.Values = make([]float64, data.Len())
	initValues(ret, ret.Values)
	return ret
}

// HHV

type hhv struct {
	funcbase
	N Value
}

func (this hhv) BuildValueAt(index int) float64 {
	N := this.N.Get(index)
	if math.IsNaN(N) || N <= 0 {
		return math.NaN()
	}
	end := index + 1
	start := int(_math.Max(0, int64(end-int(N))))

	return max(this.data, start, end)
}

func (this *hhv) UpdateLastValue() {
	updateLastValue(this)
}

func HHV(data Value, N Value) Value {
	if data.IsScalar() {
		return Scalar(data.Get(0))
	}
	ret := &hhv{
		funcbase: funcbase{
			data: data,
		},
		N: N,
	}
	ret.Values = make([]float64, data.Len())
	initValues(ret, ret.Values)
	return ret
}

// HHVBARS

type hhvbars struct {
	funcbase
	N Value
}

func (this hhvbars) BuildValueAt(index int) float64 {
	N := this.N.Get(index)
	if math.IsNaN(N) || N <= 0 {
		return float64(index)
	}

	end := index + 1
	start := int(_math.Max(0, int64(end-int(N))))

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

func HHVBARS(data Value, N Value) Value {
	if data.IsScalar() {
		return Scalar(0)
	}
	ret := &hhvbars{
		funcbase: funcbase{
			data: data,
		},
		N: N,
	}
	ret.Values = make([]float64, data.Len())
	initValues(ret, ret.Values)
	return ret
}

// HHVLLV

type hhvllv struct {
	funcbase
	T  Value
	N1 Value
	N2 Value
}

func (this hhvllv) BuildValueAt(index int) float64 {
	T := this.T.Get(index)
	N1 := this.N1.Get(index)
	N2 := this.N2.Get(index)
	if math.IsNaN(N1) || N1 <= 0 {
		return math.NaN()
	}
	if math.IsNaN(N2) || N2 <= 0 {
		return math.NaN()
	}

	var start int = 0
	var end int = index + 1
	if N1 > 0 {
		start = int(_math.Max(0, int64(index+1-int(N1))))
	}
	if N2 > 0 {
		end = int(_math.Max(0, int64(index+1-int(N2))))
	}

	if T == 0 {
		return max(this.data, start, end)
	}
	return min(this.data, start, end)
}

func (this *hhvllv) UpdateLastValue() {
	updateLastValue(this)
}

func HHVLLV(data Value, T Value, N1 Value, N2 Value) Value {
	if data.IsScalar() {
		return Scalar(data.Get(0))
	}
	ret := &hhvllv{
		funcbase: funcbase{
			data: data,
		},
		T:  T,
		N1: N1,
		N2: N2,
	}
	ret.Values = make([]float64, data.Len())
	initValues(ret, ret.Values)
	return ret
}

// STD

type std struct {
	funcbase
	N Value
}

func (this std) BuildValueAt(index int) float64 {
	if math.IsNaN(this.N.Get(index)) {
		return math.NaN()
	}
	N := int(this.N.Get(index))
	if index < N-1 {
		return 0
	}

	end := index + 1
	start := int(_math.Max(0, int64(end-N)))

	average := ma_(this.data, start, end)

	sum_ := 0.0
	for i := start; i < end; i++ {
		v := this.data.Get(i)
		diff := v - average
		sum_ += diff * diff
	}

	return math.Sqrt(sum_ / float64(end-start))
}

func (this *std) UpdateLastValue() {
	updateLastValue(this)
}

func STD(data Value, N Value) Value {
	if data.IsScalar() {
		return Scalar(0)
	}

	ret := &std{
		funcbase: funcbase{
			data: data,
		},
		N: N,
	}
	ret.Values = make([]float64, data.Len())
	initValues(ret, ret.Values)
	return ret
}

// AVEDEV

type avedev struct {
	funcbase
	N Value
}

func (this avedev) BuildValueAt(index int) float64 {
	if math.IsNaN(this.N.Get(index)) {
		return math.NaN()
	}
	N := int(this.N.Get(index))
	if index < N-1 {
		return 0
	}

	end := index + 1
	start := int(_math.Max(0, int64(end-N)))

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

func AVEDEV(data Value, N Value) Value {
	if data.IsScalar() {
		return Scalar(0)
	}

	ret := &avedev{
		funcbase: funcbase{
			data: data,
		},
		N: N,
	}
	ret.Values = make([]float64, data.Len())
	initValues(ret, ret.Values)
	return ret
}

// SUM

type sumf struct {
	funcbase
	N Value
}

func (this sumf) BuildValueAt(index int) float64 {
	if math.IsNaN(this.N.Get(index)) {
		return math.NaN()
	}
	end := index + 1
	start := int(_math.Max(0, int64(end-int(this.N.Get(index)))))
	return sum(this.data, start, end)
}

func (this *sumf) UpdateLastValue() {
	updateLastValue(this)
}

func SUM(data Value, N Value) Value {
	if data.IsScalar() {
		return Scalar(data.Get(0) * N.Get(0))
	}

	ret := &sumf{
		funcbase: funcbase{
			data: data,
		},
		N: N,
	}
	ret.Values = make([]float64, data.Len())
	initValues(ret, ret.Values)
	return ret
}

// CROSS

type cross struct {
	binaryfuncbase
}

func BuildCrossValueAt(data, data1 Value, index int) float64 {
	if index == 0 {
		return 0
	}

	return iif(data.Get(index-1) < data1.Get(index-1) && data.Get(index) >= data1.Get(index), 1, 0)
}

func (this cross) Get(index int) float64 {
	return BuildCrossValueAt(this.data, this.data1, index)
}

func CROSS(data, data1 Value) Value {
	if data.IsScalar() && data1.IsScalar() {
		return Scalar(0)
	}
	ret := &cross{
		binaryfuncbase: binaryfuncbase{
			simplefuncbase: simplefuncbase{
				data: data,
			},
			data1: data1,
		},
	}
	return ret
}

// COUNT

type count struct {
	funcbase
	N Value
}

func (this count) BuildValueAt(index int) float64 {
	if math.IsNaN(this.N.Get(index)) {
		return math.NaN()
	}
	N := int(this.N.Get(index))

	end := index + 1
	start := int(_math.Max(0, int64(end-N)))

	c := 0
	for i := start; i < end; i++ {
		v := this.data.Get(i)
		if IsTrue(v) {
			c++
		}
	}

	return float64((c))
}

func (this *count) UpdateLastValue() {
	updateLastValue(this)
}

func COUNT(data Value, N Value) Value {
	if data.IsScalar() {
		if data.Get(0) == 0 {
			return Scalar(0)
		}
		return Scalar(N.Get(0))
	}

	ret := &count{
		funcbase: funcbase{
			data: data,
		},
		N: N,
	}
	ret.Values = make([]float64, data.Len())
	initValues(ret, ret.Values)
	return ret
}

// IF

type iff struct {
	simplefuncbase
	yesData Value
	noData  Value
}

func (this iff) Get(index int) float64 {
	return iif(IsTrue(this.data.Get(index)), this.yesData.Get(index), this.noData.Get(index))
}

func IF(data, yesData, noData Value) Value {
	if data.IsScalar() && yesData.IsScalar() && noData.IsScalar() {
		return Scalar(iif(IsTrue(data.Get(0)), yesData.Get(0), noData.Get(0)))
	}

	ret := &iff{
		simplefuncbase: simplefuncbase{
			data: data,
		},
		yesData: yesData,
		noData:  noData,
	}
	return ret
}

// EVERY

type every struct {
	funcbase
	N Value
}

func (this every) BuildValueAt(index int) float64 {
	if math.IsNaN(this.N.Get(index)) {
		return math.NaN()
	}
	N := int(this.N.Get(index))

	end := index + 1
	start := int(_math.Max(0, int64(end-N)))

	for i := start; i < end; i++ {
		if !IsTrue(this.data.Get(i)) {
			return 0
		}
	}

	return 1
}

func (this *every) UpdateLastValue() {
	updateLastValue(this)
}

func EVERY(data Value, N Value) Value {
	if data.IsScalar() {
		if data.Get(0) == 0 {
			return Scalar(0)
		}
		return Scalar(1)
	}

	ret := &every{
		funcbase: funcbase{
			data: data,
		},
		N: N,
	}
	ret.Values = make([]float64, data.Len())
	initValues(ret, ret.Values)
	return ret
}

// BARSLAST

type barslast struct {
	funcbase
}

func (this barslast) BuildValueAt(index int) float64 {
	for j := index; j >= 0; j-- {
		if IsTrue(this.data.Get(j)) {
			return float64(index - j)
		}
	}
	return math.NaN()
}

func (this *barslast) UpdateLastValue() {
	updateLastValue(this)
}

func BARSLAST(data Value) Value {
	if data.IsScalar() {
		if data.Get(0) == 0 {
			return Scalar(math.NaN())
		}
		return Scalar(0)
	}

	ret := &barslast{
		funcbase: funcbase{
			data: data,
		},
	}
	ret.Values = make([]float64, data.Len())
	lastTruePos := -1
	for i := 0; i < data.Len(); i++ {
		if IsTrue(data.Get(i)) {
			lastTruePos = i
			ret.Values[i] = 0
		} else if lastTruePos == -1 {
			ret.Values[i] = math.NaN()
		} else {
			ret.Values[i] = float64(i - lastTruePos)
		}
	}
	return ret
}

func BARSLASTOLD(data Value) Value {
	if data.IsScalar() {
		if data.Get(0) == 0 {
			return Scalar(math.NaN())
		}
		return Scalar(0)
	}

	ret := &barslast{
		funcbase: funcbase{
			data: data,
		},
	}
	ret.Values = make([]float64, data.Len())
	initValues(ret, ret.Values)
	return ret
}

// BARSLASTS

type barslasts struct {
	funcbase
	N int
}

func (this barslasts) BuildValueAt(index int) float64 {
	count := 0
	for j := index; j >= 0; j-- {
		if IsTrue(this.data.Get(j)) {
			count++
			if count == this.N {
				return float64(index - j)
			}
		}
	}
	return math.NaN()
}

func (this *barslasts) UpdateLastValue() {
	updateLastValue(this)
}

func BARSLASTS(data Value, N Value) Value {
	if data.IsScalar() {
		if data.Get(0) == 0 {
			return Scalar(math.NaN())
		}
		return Scalar(0)
	}

	ret := &barslasts{
		funcbase: funcbase{
			data: data,
		},
		N: int(N.Get(0)),
	}
	ret.Values = make([]float64, data.Len())
	for i := 0; i < data.Len(); i++ {
		ret.Values[i] = ret.BuildValueAt(i)
	}
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
	if data.IsScalar() {
		panic("BARSCOUNT do not support scalar data")
	}
	ret := &barscount{
		data: data,
	}
	return ret
}

// ROUND2

type roundf struct {
	simplefuncbase
	N Value
}

func (this roundf) Get(index int) float64 {
	return round(this.data.Get(index), int(this.N.Get(index)))
}

func ROUND2(data Value, N Value) Value {
	if data.IsScalar() {
		return Scalar(round(data.Get(0), int(N.Get(0))))
	}

	if N != nil {
		N = Scalar(2)
	}

	ret := &roundf{
		simplefuncbase: simplefuncbase{
			data: data,
		},
		N: N,
	}
	return ret
}

// REF

type ref struct {
	simplefuncbase
	N Value
}

func (this ref) Get(index int) float64 {
	if math.IsNaN(this.N.Get(index)) {
		return math.NaN()
	}
	N := int(this.N.Get(index))
	if index < N {
		return 0
	}

	return this.data.Get(index - N)
}

func REF(data Value, N Value) Value {
	if data.IsScalar() {
		return data
	}

	ret := &ref{
		simplefuncbase: simplefuncbase{
			data: data,
		},
		N: N,
	}
	return ret
}

// MIN

type minf struct {
	simplefuncbase
	data1 Value
}

func (this minf) Get(index int) float64 {
	return math.Min(this.data.Get(index), this.data1.Get(index))
}

func MIN(data Value, data1 Value) Value {
	if data.IsScalar() && data1.IsScalar() {
		return Scalar(math.Min(data.Get(0), data1.Get(0)))
	}

	ret := &minf{
		simplefuncbase: simplefuncbase{
			data: data,
		},
		data1: data1,
	}
	return ret
}

// MAX

type maxf struct {
	simplefuncbase
	data1 Value
}

func (this maxf) Get(index int) float64 {
	return math.Max(this.data.Get(index), this.data1.Get(index))
}

func MAX(data Value, data1 Value) Value {
	if data.IsScalar() && data1.IsScalar() {
		return Scalar(math.Max(data.Get(0), data1.Get(0)))
	}

	ret := &maxf{
		simplefuncbase: simplefuncbase{
			data: data,
		},
		data1: data1,
	}
	return ret
}

// ABS

type absf struct {
	simplefuncbase
}

func (this absf) Get(index int) float64 {
	return math.Abs(this.data.Get(index))
}

func ABS(data Value) Value {
	if data.IsScalar() {
		return Scalar(math.Abs(data.Get(0)))
	}
	ret := &absf{
		simplefuncbase: simplefuncbase{
			data: data,
		},
	}
	return ret
}

// SLOPE

type slopef struct {
	funcbase
	N Value
}

func (this slopef) BuildValueAt(index int) float64 {
	if math.IsNaN(this.N.Get(index)) {
		return math.NaN()
	}
	N := int(this.N.Get(index))
	if index < N-1 {
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

func SLOPE(data Value, N Value) Value {
	if data.IsScalar() {
		return Scalar(0)
	}

	ret := &slopef{
		funcbase: funcbase{
			data: data,
		},
		N: N,
	}
	ret.Values = make([]float64, data.Len())
	initValues(ret, ret.Values)
	return ret
}
