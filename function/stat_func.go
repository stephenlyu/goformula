package function

import (
	"math"

	_math "github.com/chanxuehong/util/math"
)

// LLV

type llv[T float32 | float64] struct {
	funcbase[T]
	N Value[T]
}

func (this llv[T]) BuildValueAt(index int) T {
	N := this.N.Get(index)
	if math.IsNaN(float64(N)) || N <= 0 {
		return T(math.NaN())
	}
	end := index + 1
	start := int(_math.Max(0, int64(end-int(N))))

	return min(this.data, start, end)
}

func (this *llv[T]) UpdateLastValue() {
	updateLastValue(this)
}

func (this *llv[T]) GetCacheKey() string {
	return "llv_" + ValueCacheBuildKey([]Value[T]{this.data, this.N})
}

func LLV[T float32 | float64](data Value[T], N Value[T]) Value[T] {
	if data.IsScalar() {
		return Scalar(data.Get(0))
	}

	key := "llv_" + ValueCacheBuildKey([]Value[T]{data, N})
	if v := GlobalValueCache.Get(key); v != nil {
		return v.(Value[T])
	}

	ret := &llv[T]{
		funcbase: funcbase[T]{
			data: data,
		},
		N: N,
	}
	ret.Values = make([]T, data.Len())
	initValues(ret, ret.Values)
	ret.IncRef()
	GlobalValueCache.Set(key, ret)
	return ret
}

// LLVBARS

type llvbars[T float32 | float64] struct {
	intfuncbase[T, UintVecValueType]
	N Value[T]
}

func (this llvbars[T]) BuildValueAt(index int) T {
	N := this.N.Get(index)
	if math.IsNaN(float64(N)) || N <= 0 {
		return T(index)
	}

	end := index + 1
	start := int(_math.Max(0, int64(end-int(N))))

	low := T(math.MaxInt64)
	low_pos := -1
	for j := start; j < end; j++ {
		v := this.data.Get(j)
		if v < low {
			low = v
			low_pos = j
		}
	}

	return T(end - 1 - low_pos)
}

func (this *llvbars[T]) UpdateLastValue() {
	updateLastValue(this)
}

func (this *llvbars[T]) GetCacheKey() string {
	return "llvbars_" + ValueCacheBuildKey([]Value[T]{this.data, this.N})
}

func LLVBARS[T float32 | float64](data Value[T], N Value[T]) Value[T] {
	if data.IsScalar() {
		return Scalar(T(0))
	}

	key := "llvbars_" + ValueCacheBuildKey([]Value[T]{data, N})
	if v := GlobalValueCache.Get(key); v != nil {
		return v.(Value[T])
	}

	ret := &llvbars[T]{
		intfuncbase: intfuncbase[T, UintVecValueType]{
			data: data,
		},
		N: N,
	}
	ret.Values = make([]UintVecValueType, data.Len())
	initIntValues(ret, ret.Values)
	ret.IncRef()
	GlobalValueCache.Set(key, ret)
	return ret
}

// HHV

type hhv[T float32 | float64] struct {
	funcbase[T]
	N Value[T]
}

func (this hhv[T]) BuildValueAt(index int) T {
	N := this.N.Get(index)
	if math.IsNaN(float64(N)) || N <= 0 {
		return T(math.NaN())
	}
	end := index + 1
	start := int(_math.Max(0, int64(end-int(N))))

	return max(this.data, start, end)
}

func (this *hhv[T]) UpdateLastValue() {
	updateLastValue(this)
}

func (this *hhv[T]) GetCacheKey() string {
	return "hhv_" + ValueCacheBuildKey([]Value[T]{this.data, this.N})
}

func HHV[T float32 | float64](data Value[T], N Value[T]) Value[T] {
	if data.IsScalar() {
		return Scalar(data.Get(0))
	}
	key := "hhv_" + ValueCacheBuildKey([]Value[T]{data, N})
	if v := GlobalValueCache.Get(key); v != nil {
		return v.(Value[T])
	}
	ret := &hhv[T]{
		funcbase: funcbase[T]{
			data: data,
		},
		N: N,
	}
	ret.Values = make([]T, data.Len())
	initValues(ret, ret.Values)
	ret.IncRef()
	GlobalValueCache.Set(key, ret)
	return ret
}

// HHVBARS

type hhvbars[T float32 | float64] struct {
	intfuncbase[T, UintVecValueType]
	N Value[T]
}

func (this hhvbars[T]) BuildValueAt(index int) T {
	N := this.N.Get(index)
	if math.IsNaN(float64(N)) || N <= 0 {
		return T(index)
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

	return T(end - 1 - high_pos)
}

func (this *hhvbars[T]) UpdateLastValue() {
	updateLastValue(this)
}

func (this *hhvbars[T]) GetCacheKey() string {
	return "hhvbars_" + ValueCacheBuildKey([]Value[T]{this.data, this.N})
}

func HHVBARS[T float32 | float64](data Value[T], N Value[T]) Value[T] {
	if data.IsScalar() {
		return Scalar(T(0))
	}
	key := "hhvbars_" + ValueCacheBuildKey([]Value[T]{data, N})
	if v := GlobalValueCache.Get(key); v != nil {
		return v.(Value[T])
	}
	ret := &hhvbars[T]{
		intfuncbase: intfuncbase[T, UintVecValueType]{
			data: data,
		},
		N: N,
	}
	ret.Values = make([]UintVecValueType, data.Len())
	initIntValues(ret, ret.Values)
	ret.IncRef()
	GlobalValueCache.Set(key, ret)
	return ret
}

// HHVLLV

type hhvllv[R float32 | float64] struct {
	funcbase[R]
	T  Value[R]
	N1 Value[R]
	N2 Value[R]
}

func (this hhvllv[R]) BuildValueAt(index int) R {
	T := this.T.Get(index)
	N1 := this.N1.Get(index)
	N2 := this.N2.Get(index)
	if math.IsNaN(float64(N1)) || N1 < 0 {
		return R(math.NaN())
	}
	if math.IsNaN(float64(N2)) || N2 < 0 {
		return R(math.NaN())
	}

	var start int = 0
	var end int = index + 1
	if N1 > 0 {
		start = int(_math.Max(0, int64(index+1-int(N1))))
	}
	if N2 > 0 {
		end = int(_math.Max(0, int64(index+1-int(N2))))
	}

	if start > end {
		start, end = end, start
	}

	if T == 0 {
		return max(this.data, start, end)
	}
	return min(this.data, start, end)
}

func (this *hhvllv[R]) UpdateLastValue() {
	updateLastValue(this)
}

func (this *hhvllv[R]) GetCacheKey() string {
	return "hhvllv_" + ValueCacheBuildKey([]Value[R]{this.data, this.T, this.N1, this.N2})
}

func HHVLLV[R float32 | float64](data Value[R], T Value[R], N1 Value[R], N2 Value[R]) Value[R] {
	if data.IsScalar() {
		return Scalar(data.Get(0))
	}
	key := "hhvllv_" + ValueCacheBuildKey([]Value[R]{data, T, N1, N2})
	if v := GlobalValueCache.Get(key); v != nil {
		return v.(Value[R])
	}
	ret := &hhvllv[R]{
		funcbase: funcbase[R]{
			data: data,
		},
		T:  T,
		N1: N1,
		N2: N2,
	}
	ret.Values = make([]R, data.Len())
	initValues(ret, ret.Values)
	ret.IncRef()
	GlobalValueCache.Set(key, ret)
	return ret
}

// STD

type std[T float32 | float64] struct {
	funcbase[T]
	N Value[T]
}

func (this std[T]) BuildValueAt(index int) T {
	if math.IsNaN(float64(this.N.Get(index))) {
		return T(math.NaN())
	}
	N := int(this.N.Get(index))
	if index < N-1 {
		return 0
	}

	end := index + 1
	start := int(_math.Max(0, int64(end-N)))

	average := ma_(this.data, start, end)

	sum_ := T(0.0)
	for i := start; i < end; i++ {
		v := this.data.Get(i)
		diff := v - average
		sum_ += diff * diff
	}

	return T(math.Sqrt(float64(sum_) / float64(end-start)))
}

func (this *std[T]) UpdateLastValue() {
	updateLastValue(this)
}

func (this *std[T]) GetCacheKey() string {
	return "std_" + ValueCacheBuildKey([]Value[T]{this.data, this.N})
}

func STD[T float32 | float64](data Value[T], N Value[T]) Value[T] {
	if data.IsScalar() {
		return Scalar(T(0))
	}
	key := "std_" + ValueCacheBuildKey([]Value[T]{data, N})
	if v := GlobalValueCache.Get(key); v != nil {
		return v.(Value[T])
	}
	ret := &std[T]{
		funcbase: funcbase[T]{
			data: data,
		},
		N: N,
	}
	ret.Values = make([]T, data.Len())
	initValues(ret, ret.Values)
	ret.IncRef()
	GlobalValueCache.Set(key, ret)
	return ret
}

// AVEDEV

type avedev[T float32 | float64] struct {
	funcbase[T]
	N Value[T]
}

func (this avedev[T]) BuildValueAt(index int) T {
	if math.IsNaN(float64(this.N.Get(index))) {
		return T(math.NaN())
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
		diff := math.Abs(float64(v - average))
		vsum += diff
	}

	return T(vsum / float64(count))
}

func (this *avedev[T]) UpdateLastValue() {
	updateLastValue(this)
}

func (this *avedev[T]) GetCacheKey() string {
	return "avedev_" + ValueCacheBuildKey([]Value[T]{this.data, this.N})
}

func AVEDEV[T float32 | float64](data Value[T], N Value[T]) Value[T] {
	if data.IsScalar() {
		return Scalar(T(0))
	}
	key := "avedev_" + ValueCacheBuildKey([]Value[T]{data, N})
	if v := GlobalValueCache.Get(key); v != nil {
		return v.(Value[T])
	}
	ret := &avedev[T]{
		funcbase: funcbase[T]{
			data: data,
		},
		N: N,
	}
	ret.Values = make([]T, data.Len())
	initValues(ret, ret.Values)
	ret.IncRef()
	GlobalValueCache.Set(key, ret)
	return ret
}

// SUM

type sumf[T float32 | float64] struct {
	funcbase[T]
	N Value[T]
}

func (this sumf[T]) BuildValueAt(index int) T {
	if math.IsNaN(float64(this.N.Get(index))) {
		return T(math.NaN())
	}
	end := index + 1
	start := int(_math.Max(0, int64(end-int(this.N.Get(index)))))
	return sum(this.data, start, end)
}

func (this *sumf[T]) UpdateLastValue() {
	updateLastValue(this)
}

func (this *sumf[T]) GetCacheKey() string {
	return "sumf_" + ValueCacheBuildKey([]Value[T]{this.data, this.N})
}

func SUM[T float32 | float64](data Value[T], N Value[T]) Value[T] {
	if data.IsScalar() {
		return Scalar(data.Get(0) * N.Get(0))
	}

	key := "sumf_" + ValueCacheBuildKey([]Value[T]{data, N})
	if v := GlobalValueCache.Get(key); v != nil {
		return v.(Value[T])
	}
	ret := &sumf[T]{
		funcbase: funcbase[T]{
			data: data,
		},
		N: N,
	}
	ret.Values = make([]T, data.Len())
	initValues(ret, ret.Values)
	ret.IncRef()
	GlobalValueCache.Set(key, ret)
	return ret
}

// CROSS

type cross[T float32 | float64] struct {
	binaryfuncbase[T]
}

func BuildCrossValueAt[T float32 | float64](data, data1 Value[T], index int) T {
	if index == 0 {
		return 0
	}

	return iif[T](data.Get(index-1) < data1.Get(index-1) && data.Get(index) >= data1.Get(index), 1, 0)
}

func (this cross[T]) Get(index int) T {
	return BuildCrossValueAt(this.data, this.data1, index)
}

func CROSS[T float32 | float64](data, data1 Value[T]) Value[T] {
	if data.IsScalar() && data1.IsScalar() {
		return Scalar(T(0))
	}
	ret := &cross[T]{
		binaryfuncbase: binaryfuncbase[T]{
			simplefuncbase: simplefuncbase[T]{
				data: data,
			},
			data1: data1,
		},
	}
	return ret
}

// COUNT

type count[T float32 | float64] struct {
	funcbase[T]
	N Value[T]
}

func (this count[T]) BuildValueAt(index int) T {
	if math.IsNaN(float64(this.N.Get(index))) {
		return T(math.NaN())
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

	return T(c)
}

func (this *count[T]) UpdateLastValue() {
	updateLastValue(this)
}

func (this *count[T]) GetCacheKey() string {
	return "count_" + ValueCacheBuildKey([]Value[T]{this.data, this.N})
}

func COUNT[T float32 | float64](data Value[T], N Value[T]) Value[T] {
	if data.IsScalar() {
		if data.Get(0) == 0 {
			return Scalar(T(0))
		}
		return Scalar(N.Get(0))
	}

	key := "count_" + ValueCacheBuildKey([]Value[T]{data, N})
	if v := GlobalValueCache.Get(key); v != nil {
		return v.(Value[T])
	}

	ret := &count[T]{
		funcbase: funcbase[T]{
			data: data,
		},
		N: N,
	}
	ret.Values = make([]T, data.Len())
	initValues(ret, ret.Values)
	ret.IncRef()
	GlobalValueCache.Set(key, ret)
	return ret
}

// IF

type iff[T float32 | float64] struct {
	simplefuncbase[T]
	yesData Value[T]
	noData  Value[T]
}

func (this iff[T]) Get(index int) T {
	return iif(IsTrue(this.data.Get(index)), this.yesData.Get(index), this.noData.Get(index))
}

func IF[T float32 | float64](data, yesData, noData Value[T]) Value[T] {
	if data.IsScalar() && yesData.IsScalar() && noData.IsScalar() {
		return Scalar(iif(IsTrue(data.Get(0)), yesData.Get(0), noData.Get(0)))
	}

	ret := &iff[T]{
		simplefuncbase: simplefuncbase[T]{
			data: data,
		},
		yesData: yesData,
		noData:  noData,
	}
	return ret
}

// EVERY

type every[T float32 | float64] struct {
	funcbase[T]
	N Value[T]
}

func (this every[T]) BuildValueAt(index int) T {
	if math.IsNaN(float64(this.N.Get(index))) {
		return T(math.NaN())
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

func (this *every[T]) UpdateLastValue() {
	updateLastValue(this)
}

func (this *every[T]) GetCacheKey() string {
	return "every_" + ValueCacheBuildKey([]Value[T]{this.data, this.N})
}

func EVERY[T float32 | float64](data Value[T], N Value[T]) Value[T] {
	if data.IsScalar() {
		if data.Get(0) == 0 {
			return Scalar(T(0))
		}
		return Scalar(T(1))
	}

	key := "every_" + ValueCacheBuildKey([]Value[T]{data, N})
	if v := GlobalValueCache.Get(key); v != nil {
		return v.(Value[T])
	}

	ret := &every[T]{
		funcbase: funcbase[T]{
			data: data,
		},
		N: N,
	}
	ret.Values = make([]T, data.Len())
	initValues(ret, ret.Values)
	ret.IncRef()
	GlobalValueCache.Set(key, ret)
	return ret
}

// BARSLAST

type barslast[T float32 | float64] struct {
	intfuncbase[T, UintVecValueType]
}

func (this barslast[T]) BuildValueAt(index int) T {
	for j := index; j >= 0; j-- {
		if IsTrue(this.data.Get(j)) {
			return T(index - j)
		}
	}
	return T(math.NaN())
}

func (this *barslast[T]) UpdateLastValue() {
	updateLastValue(this)
}

func (this *barslast[T]) GetCacheKey() string {
	return "barslast_" + ValueCacheBuildKey([]Value[T]{this.data})
}

func BARSLAST[T float32 | float64](data Value[T]) Value[T] {
	if data.IsScalar() {
		if data.Get(0) == 0 {
			return Scalar(T(math.NaN()))
		}
		return Scalar(T(0))
	}

	key := "barslast_" + ValueCacheBuildKey([]Value[T]{data})
	if v := GlobalValueCache.Get(key); v != nil {
		return v.(Value[T])
	}

	ret := &barslast[T]{
		intfuncbase: intfuncbase[T, UintVecValueType]{
			data: data,
		},
	}
	ret.Values = make([]UintVecValueType, data.Len())
	lastTruePos := -1
	for i := 0; i < data.Len(); i++ {
		if IsTrue(data.Get(i)) {
			lastTruePos = i
			ret.Values[i] = 0
		} else if lastTruePos == -1 {
			ret.Values[i] = UintVecValueType(uintNaN)
		} else {
			ret.Values[i] = UintVecValueType(i - lastTruePos)
		}
	}
	ret.IncRef()
	GlobalValueCache.Set(key, ret)
	return ret
}

func BARSLASTOLD[T float32 | float64](data Value[T]) Value[T] {
	if data.IsScalar() {
		if data.Get(0) == 0 {
			return Scalar(T(math.NaN()))
		}
		return Scalar(T(0))
	}

	key := "barslast_" + ValueCacheBuildKey([]Value[T]{data})
	if v := GlobalValueCache.Get(key); v != nil {
		return v.(Value[T])
	}

	ret := &barslast[T]{
		intfuncbase: intfuncbase[T, UintVecValueType]{
			data: data,
		},
	}
	ret.Values = make([]UintVecValueType, data.Len())
	initIntValues(ret, ret.Values)
	ret.IncRef()
	GlobalValueCache.Set(key, ret)
	return ret
}

// BARSLASTS

type barslasts[T float32 | float64] struct {
	intfuncbase[T, UintVecValueType]
	N int
}

func (this barslasts[T]) BuildValueAt(index int) T {
	count := 0
	for j := index; j >= 0; j-- {
		if IsTrue(this.data.Get(j)) {
			count++
			if count == this.N {
				return T(index - j)
			}
		}
	}
	return T(math.NaN())
}

func (this *barslasts[T]) UpdateLastValue() {
	updateLastValue(this)
}

func (this *barslasts[T]) GetCacheKey() string {
	return "barslasts_" + ValueCacheBuildKey([]Value[T]{this.data, Scalar(T(this.N))})
}

func BARSLASTS[T float32 | float64](data Value[T], N Value[T]) Value[T] {
	if data.IsScalar() {
		if data.Get(0) == 0 {
			return Scalar(T(math.NaN()))
		}
		return Scalar(T(0))
	}

	key := "barslasts_" + ValueCacheBuildKey([]Value[T]{data, N})
	if v := GlobalValueCache.Get(key); v != nil {
		return v.(Value[T])
	}

	ret := &barslasts[T]{
		intfuncbase: intfuncbase[T, UintVecValueType]{
			data: data,
		},
		N: int(N.Get(0)),
	}
	ret.Values = make([]UintVecValueType, data.Len())
	for i := 0; i < data.Len(); i++ {
		v := ret.BuildValueAt(i)
		if math.IsNaN(float64(v)) {
			ret.Values[i] = UintVecValueType(uintNaN)
		} else {
			ret.Values[i] = UintVecValueType(v)
		}
	}
	ret.IncRef()
	GlobalValueCache.Set(key, ret)
	return ret
}

// BARSCOUNT

type barscount[T float32 | float64] struct {
	scalar[T]
	data Value[T]
}

func (this barscount[T]) Get(index int) T {
	return T(this.data.Len())
}

func BARSCOUNT[T float32 | float64](data Value[T]) *barscount[T] {
	if data.IsScalar() {
		panic("BARSCOUNT do not support scalar data")
	}
	ret := &barscount[T]{
		data: data,
	}
	return ret
}

// ROUND2

type roundf[T float32 | float64] struct {
	simplefuncbase[T]
	N Value[T]
}

func (this roundf[T]) Get(index int) T {
	return round(this.data.Get(index), int(this.N.Get(index)))
}

func ROUND2[T float32 | float64](data Value[T], N Value[T]) Value[T] {
	if data.IsScalar() {
		return Scalar(round(data.Get(0), int(N.Get(0))))
	}

	if N != nil {
		N = Scalar(T(2))
	}

	ret := &roundf[T]{
		simplefuncbase: simplefuncbase[T]{
			data: data,
		},
		N: N,
	}
	return ret
}

// REF

type ref[T float32 | float64] struct {
	simplefuncbase[T]
	N Value[T]
}

func (this ref[T]) Get(index int) T {
	if math.IsNaN(float64(this.N.Get(index))) {
		return T(math.NaN())
	}
	N := int(this.N.Get(index))
	if index < N {
		return 0
	}

	return this.data.Get(index - N)
}

func REF[T float32 | float64](data Value[T], N Value[T]) Value[T] {
	if data.IsScalar() {
		return data
	}

	ret := &ref[T]{
		simplefuncbase: simplefuncbase[T]{
			data: data,
		},
		N: N,
	}
	return ret
}

// MIN

type minf[T float32 | float64] struct {
	simplefuncbase[T]
	data1 Value[T]
}

func (this minf[T]) Get(index int) T {
	return minv(this.data.Get(index), this.data1.Get(index))
}

func MIN[T float32 | float64](data Value[T], data1 Value[T]) Value[T] {
	if data.IsScalar() && data1.IsScalar() {
		return Scalar(minv(data.Get(0), data1.Get(0)))
	}

	ret := &minf[T]{
		simplefuncbase: simplefuncbase[T]{
			data: data,
		},
		data1: data1,
	}
	return ret
}

// MAX

type maxf[T float32 | float64] struct {
	simplefuncbase[T]
	data1 Value[T]
}

func (this maxf[T]) Get(index int) T {
	return maxv(this.data.Get(index), this.data1.Get(index))
}

func MAX[T float32 | float64](data Value[T], data1 Value[T]) Value[T] {
	if data.IsScalar() && data1.IsScalar() {
		return Scalar(maxv(data.Get(0), data1.Get(0)))
	}

	ret := &maxf[T]{
		simplefuncbase: simplefuncbase[T]{
			data: data,
		},
		data1: data1,
	}
	return ret
}

// ABS

type absf[T float32 | float64] struct {
	simplefuncbase[T]
}

func (this absf[T]) Get(index int) T {
	return T(math.Abs(float64(this.data.Get(index))))
}

func ABS[T float32 | float64](data Value[T]) Value[T] {
	if data.IsScalar() {
		return Scalar(T(math.Abs(float64(data.Get(0)))))
	}
	ret := &absf[T]{
		simplefuncbase: simplefuncbase[T]{
			data: data,
		},
	}
	return ret
}

// SLOPE

type slopef[T float32 | float64] struct {
	funcbase[T]
	N Value[T]
}

func (this slopef[T]) BuildValueAt(index int) T {
	if math.IsNaN(float64(this.N.Get(index))) {
		return T(math.NaN())
	}
	N := int(this.N.Get(index))
	if index < N-1 {
		return 0
	}

	x := make([]T, N)
	y := make([]T, N)

	for i := 0; i < N; i++ {
		x[i] = T(i + 1)
		y[i] = this.data.Get(index + 1 - (N - i))
	}

	_, slope, _ := LinearRegression(Vector(x), Vector(y))

	return slope
}

func (this *slopef[T]) UpdateLastValue() {
	updateLastValue(this)
}

func (this *slopef[T]) GetCacheKey() string {
	return "slope_" + ValueCacheBuildKey([]Value[T]{this.data, this.N})
}

func SLOPE[T float32 | float64](data Value[T], N Value[T]) Value[T] {
	if data.IsScalar() {
		return Scalar(T(0))
	}
	key := "slope_" + ValueCacheBuildKey([]Value[T]{data, N})
	if v := GlobalValueCache.Get(key); v != nil {
		return v.(Value[T])
	}
	ret := &slopef[T]{
		funcbase: funcbase[T]{
			data: data,
		},
		N: N,
	}
	ret.Values = make([]T, data.Len())
	initValues(ret, ret.Values)
	ret.IncRef()
	GlobalValueCache.Set(key, ret)
	return ret
}
