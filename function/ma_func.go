package function

import "math"

// MA - moving average

type ma[T float32 | float64] struct {
	funcbase[T]
	N Value[T]
}

// 假定按照index递增的顺序依次调用，计算index时，index-1已经计算，可以从this.data[index-1]访问到
func (this ma[T]) BuildValueAt(index int) T {
	n := int(this.N.Get(index))
	if index < n-1 {
		return T(NaN)
	}
	prev := this.Values[index-1]
	if index == n-1 || math.IsNaN(float64(prev)) {
		end := index + 1
		start := int(end - n)
		return ma_(this.data, start, end)
	}

	N := this.N.Get(index)
	sum := prev * N
	sum += this.data.Get(index) - this.data.Get(index-n)
	return sum / N
}

func (this *ma[T]) UpdateLastValue() {
	updateLastValue(this)
}

func (this *ma[T]) GetCacheKey() string {
	return "ma_" + ValueCacheBuildKey([]Value[T]{this.data, this.N})
}

func MA[T float32 | float64](data Value[T], N Value[T]) Value[T] {
	if data.IsScalar() {
		return Scalar(data.Get(0))
	}

	if N == nil {
		N = Scalar(T(5))
	}

	key := "ma_" + ValueCacheBuildKey([]Value[T]{data, N})
	if v := GlobalValueCache.Get(key); v != nil {
		return v.(Value[T])
	}

	ret := &ma[T]{
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

// SMA

type sma[T float32 | float64] struct {
	funcbase[T]
	N Value[T]
	M Value[T]
}

func (this sma[T]) BuildValueAt(index int) T {
	if index == 0 {
		return this.data.Get(index)
	}

	p := this.data.Get(index)
	lv := this.Get(index - 1)

	return (p*this.M.Get(index) + lv*(this.N.Get(index)-this.M.Get(index))) / this.N.Get(index)
}

func (this *sma[T]) UpdateLastValue() {
	updateLastValue(this)
}

func (this *sma[T]) GetCacheKey() string {
	return "sma_" + ValueCacheBuildKey([]Value[T]{this.data, this.N, this.M})
}

func SMA[T float32 | float64](data Value[T], N, M Value[T]) Value[T] {
	if data.IsScalar() {
		return Scalar(data.Get(0))
	}
	if N == nil {
		N = Scalar(T(12))
	}

	if M == nil {
		M = Scalar(T(1))
	}
	key := "sma_" + ValueCacheBuildKey([]Value[T]{data, N, M})
	if v := GlobalValueCache.Get(key); v != nil {
		return v.(Value[T])
	}
	ret := &sma[T]{
		funcbase: funcbase[T]{
			data: data,
		},
		N: N,
		M: M,
	}
	ret.Values = make([]T, data.Len())
	initValues(ret, ret.Values)
	ret.IncRef()
	GlobalValueCache.Set(key, ret)
	return ret
}

// SMA

type dma[T float32 | float64] struct {
	funcbase[T]
	A Value[T]
}

func (this dma[T]) BuildValueAt(index int) T {
	if index == 0 {
		return this.data.Get(0)
	}

	p := this.data.Get(index)
	lv := this.Get(index - 1)

	return p*this.A.Get(index) + lv*(1-this.A.Get(index))
}

func (this *dma[T]) UpdateLastValue() {
	updateLastValue(this)
}

func (this *dma[T]) GetCacheKey() string {
	return "dma_" + ValueCacheBuildKey([]Value[T]{this.data, this.A})
}

func DMA[T float32 | float64](data Value[T], A Value[T]) Value[T] {
	if data.IsScalar() {
		return Scalar(data.Get(0))
	}
	if A == nil {
		A = Scalar(T(0.5))
	}

	key := "dma_" + ValueCacheBuildKey([]Value[T]{data, A})
	if v := GlobalValueCache.Get(key); v != nil {
		return v.(Value[T])
	}

	ret := &dma[T]{
		funcbase: funcbase[T]{
			data: data,
		},
		A: A,
	}
	ret.Values = make([]T, data.Len())
	initValues(ret, ret.Values)
	ret.IncRef()
	GlobalValueCache.Set(key, ret)
	return ret
}

// EMA

type ema[T float32 | float64] struct {
	funcbase[T]
	N Value[T]
}

func (this ema[T]) BuildValueAt(index int) T {
	if index == 0 {
		return this.data.Get(0)
	}

	p := this.data.Get(index)
	lv := this.Get(index - 1)
	alpha := 2.0 / (this.N.Get(index) + 1)
	return p*alpha + lv*(1-alpha)
}

func (this *ema[T]) UpdateLastValue() {
	updateLastValue(this)
}

func (this *ema[T]) GetCacheKey() string {
	return "ema_" + ValueCacheBuildKey([]Value[T]{this.data, this.N})
}

func EMA[T float32 | float64](data Value[T], N Value[T]) Value[T] {
	if data.IsScalar() {
		return Scalar(data.Get(0))
	}
	if N == nil {
		N = Scalar(T(5))
	}

	key := "ema_" + ValueCacheBuildKey([]Value[T]{data, N})
	if v := GlobalValueCache.Get(key); v != nil {
		return v.(Value[T])
	}

	ret := &ema[T]{
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

// EXPMEMA

type expmema[T float32 | float64] struct {
	funcbase[T]
	N Value[T]
}

func (this expmema[T]) BuildValueAt(index int) T {
	N := int(this.N.Get(index))
	if index < N-1 {
		return 0
	}

	if index == N-1 {
		return ma_[T](this.data, 0, N)
	}

	p := this.data.Get(index)
	lv := this.Get(index - 1)

	alpha := 2.0 / (this.N.Get(index) + 1)
	return p*alpha + lv*(1-alpha)
}

func (this *expmema[T]) UpdateLastValue() {
	updateLastValue(this)
}

func (this *expmema[T]) GetCacheKey() string {
	return "expmema_" + ValueCacheBuildKey([]Value[T]{this.data, this.N})
}

func EXPMEMA[T float32 | float64](data Value[T], N Value[T]) Value[T] {
	if data.IsScalar() {
		return Scalar(data.Get(0))
	}
	if N == nil {
		N = Scalar(T(0.5))
	}

	key := "expmema_" + ValueCacheBuildKey([]Value[T]{data, N})
	if v := GlobalValueCache.Get(key); v != nil {
		return v.(Value[T])
	}

	ret := &expmema[T]{
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
