package function

import "errors"

type fieldbase[T float32 | float64] struct {
	data RVectorReader[T]
}

func (this fieldbase[T]) IsScalar() bool {
	return false
}

func (this fieldbase[T]) Len() int {
	return this.data.Len()
}

func (this fieldbase[T]) Get(index int) T {
	panic(errors.New("Not Implemented"))
}

func (this fieldbase[T]) Set(index int, v T) {
	panic(errors.New("Not Implemented"))
}

func (this fieldbase[T]) UpdateLastValue() {
}

func (this fieldbase[T]) Append(v T) {
	panic(errors.New("Not Implemented"))
}

// CLOSE

type close[T float32 | float64] struct {
	fieldbase[T]
}

func (this close[T]) Get(index int) T {
	return this.data.Get(index).GetClose()
}

func CLOSE[T float32 | float64](data RVectorReader[T]) *close[T] {
	ret := &close[T]{
		fieldbase[T]{
			data: data,
		},
	}
	return ret
}

// OPEN

type open[T float32 | float64] struct {
	fieldbase[T]
}

func (this open[T]) Get(index int) T {
	return this.data.Get(index).GetOpen()
}

func OPEN[T float32 | float64](data RVectorReader[T]) *open[T] {
	ret := &open[T]{
		fieldbase[T]{
			data: data,
		},
	}
	return ret
}

// LOW

type low[T float32 | float64] struct {
	fieldbase[T]
}

func (this low[T]) Get(index int) T {
	return this.data.Get(index).GetLow()
}

func LOW[T float32 | float64](data RVectorReader[T]) *low[T] {
	ret := &low[T]{
		fieldbase[T]{
			data: data,
		},
	}
	return ret
}

// HIGH

type high[T float32 | float64] struct {
	fieldbase[T]
}

func (this high[T]) Get(index int) T {
	return this.data.Get(index).GetHigh()
}

func HIGH[T float32 | float64](data RVectorReader[T]) *high[T] {
	ret := &high[T]{
		fieldbase[T]{
			data: data,
		},
	}
	return ret
}

// AMOUNT

type amount[T float32 | float64] struct {
	fieldbase[T]
}

func (this amount[T]) Get(index int) T {
	return this.data.Get(index).GetAmount()
}

func AMOUNT[T float32 | float64](data RVectorReader[T]) *amount[T] {
	ret := &amount[T]{
		fieldbase[T]{
			data: data,
		},
	}
	return ret
}

// VOLUME

type volume[T float32 | float64] struct {
	fieldbase[T]
}

func (this volume[T]) Get(index int) T {
	return this.data.Get(index).GetVolume()
}

func VOLUME[T float32 | float64](data RVectorReader[T]) *volume[T] {
	ret := &volume[T]{
		fieldbase[T]{
			data: data,
		},
	}
	return ret
}

// Period

type fPeriod[T float32 | float64] struct {
	fieldbase[T]

	value T
}

func (this fPeriod[T]) Get(index int) T {
	return this.value
}

func PERIOD[T float32 | float64](data RVectorReader[T]) *fPeriod[T] {
	ret := &fPeriod[T]{
		fieldbase: fieldbase[T]{
			data: data,
		},
		value: T(GetPeriodIndex(data.Period())),
	}
	return ret
}

// ISLASTBAR

type islastbar[T float32 | float64] struct {
	fieldbase[T]
}

func (this islastbar[T]) Get(index int) T {
	if this.data.Len()-1 == index {
		return 1
	}
	return 0
}

func ISLASTBAR[T float32 | float64](data RVectorReader[T]) *islastbar[T] {
	ret := &islastbar[T]{
		fieldbase[T]{
			data: data,
		},
	}
	return ret
}
