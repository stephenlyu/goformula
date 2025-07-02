package function

import (
	"github.com/chanxuehong/util/math"
)

// MA - moving average

type ma[T float32 | float64] struct {
	funcbase[T]
	N Value[T]
}

func (this ma[T]) BuildValueAt(index int) T {
	n := int(this.N.Get(index))
	if index < n-1 {
		return T(NaN)
	}

	end := index + 1
	start := int(math.Max(0, int64(end-n)))

	return ma_[T](this.data, start, end)
}

func (this *ma[T]) UpdateLastValue() {
	updateLastValue(this)
}

func MA[T float32 | float64](data Value[T], N Value[T]) Value[T] {
	if data.IsScalar() {
		return Scalar(data.Get(0))
	}

	if N == nil {
		N = Scalar(T(5))
	}

	ret := &ma[T]{
		funcbase: funcbase[T]{
			data: data,
		},
		N: N,
	}
	ret.Values = make([]T, data.Len())
	initValues(ret, ret.Values)
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

	ret := &sma[T]{
		funcbase: funcbase[T]{
			data: data,
		},
		N: N,
		M: M,
	}
	ret.Values = make([]T, data.Len())
	initValues(ret, ret.Values)

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

func DMA[T float32 | float64](data Value[T], A Value[T]) Value[T] {
	if data.IsScalar() {
		return Scalar(data.Get(0))
	}
	if A == nil {
		A = Scalar(T(0.5))
	}

	ret := &dma[T]{
		funcbase: funcbase[T]{
			data: data,
		},
		A: A,
	}
	ret.Values = make([]T, data.Len())
	initValues(ret, ret.Values)

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

func EMA[T float32 | float64](data Value[T], N Value[T]) Value[T] {
	if data.IsScalar() {
		return Scalar(data.Get(0))
	}
	if N == nil {
		N = Scalar(T(5))
	}

	ret := &ema[T]{
		funcbase: funcbase[T]{
			data: data,
		},
		N: N,
	}
	ret.Values = make([]T, data.Len())
	initValues(ret, ret.Values)
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

func EXPMEMA[T float32 | float64](data Value[T], N Value[T]) Value[T] {
	if data.IsScalar() {
		return Scalar(data.Get(0))
	}
	if N == nil {
		N = Scalar(T(0.5))
	}

	ret := &expmema[T]{
		funcbase: funcbase[T]{
			data: data,
		},
		N: N,
	}
	ret.Values = make([]T, data.Len())
	initValues(ret, ret.Values)

	return ret
}
