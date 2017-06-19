package function

import (
	"github.com/chanxuehong/util/math"
)

// MA - moving average

type ma struct {
	funcbase
	N Value
}

func (this ma) BuildValueAt(index int) float64 {
	n := int(this.N.Get(index))
	if index < n -1 {
		return 0
	}

	end := index + 1
	start := int(math.Max(0, int64(end - n)))

	return ma_(this.data, start, end)
}

func (this *ma) UpdateLastValue() {
	updateLastValue(this)
}

func MA(data Value, N Value) *ma {
	if N == nil {
		N = Scalar(5)
	}

	ret := &ma{
		funcbase: funcbase {
			data: data,
		},
		N: N,
	}
	ret.Values = make([]float64, data.Len())
	initValues(ret, ret.Values)
	return ret
}

// SMA

type sma struct {
	funcbase
	N Value
	M Value
}

func (this sma) BuildValueAt(index int) float64 {
	if index == 0 {
		return this.data.Get(index)
	}

	p := this.data.Get(index)
	lv := this.Get(index - 1)

	return (p * this.M.Get(index) + lv * (this.N.Get(index) - this.M.Get(index))) / this.N.Get(index)
}

func (this *sma) UpdateLastValue() {
	updateLastValue(this)
}

func SMA(data Value, N, M Value) *sma {
	if N == nil {
		N = Scalar(12)
	}

	if M == nil {
		M = Scalar(1)
	}

	ret := &sma{
		funcbase: funcbase {
			data: data,
		},
		N: N,
		M: M,
	}
	ret.Values = make([]float64, data.Len())
	initValues(ret, ret.Values)

	return ret
}

// SMA

type dma struct {
	funcbase
	A Value
}

func (this dma) BuildValueAt(index int) float64 {
	if index == 0 {
		return this.data.Get(0)
	}

	p := this.data.Get(index)
	lv := this.Get(index - 1)

	return p * this.A.Get(index) + lv * (1 - this.A.Get(index))
}

func (this *dma) UpdateLastValue() {
	updateLastValue(this)
}

func DMA(data Value, A Value) *dma {
	if A == nil {
		A = Scalar(0.5)
	}

	ret := &dma{
		funcbase: funcbase {
			data: data,
		},
		A: A,
	}
	ret.Values = make([]float64, data.Len())
	initValues(ret, ret.Values)

	return ret
}

// EMA

type ema struct {
	funcbase
	N Value
}

func (this ema) BuildValueAt(index int) float64 {
	if index == 0 {
		return this.data.Get(0)
	}

	p := this.data.Get(index)
	lv := this.Get(index - 1)
	alpha := 2.0 / (this.N.Get(index) + 1)
	return p * alpha + lv * (1 - alpha)
}

func (this *ema) UpdateLastValue() {
	updateLastValue(this)
}

func EMA(data Value, N Value) *ema {
	if N == nil {
		N = Scalar(5)
	}

	ret := &ema{
		funcbase: funcbase {
			data: data,
		},
		N: N,
	}
	ret.Values = make([]float64, data.Len())
	initValues(ret, ret.Values)
	return ret
}

// EXPMEMA

type expmema struct {
	funcbase
	N Value
}

func (this expmema) BuildValueAt(index int) float64 {
	N := int(this.N.Get(index))
	if index < N - 1 {
		return 0
	}

	if index == N - 1 {
		return ma_(this.data, 0, N)
	}

	p := this.data.Get(index)
	lv := this.Get(index - 1)

	alpha := 2.0 / (this.N.Get(index) + 1)
	return p * alpha + lv * (1 - alpha)
}

func (this *expmema) UpdateLastValue() {
	updateLastValue(this)
}

func EXPMEMA(data Value, N Value) *expmema {
	if N == nil {
		N = Scalar(0.5)
	}

	ret := &expmema{
		funcbase: funcbase {
			data: data,
		},
		N: N,
	}
	ret.Values = make([]float64, data.Len())
	initValues(ret, ret.Values)

	return ret
}