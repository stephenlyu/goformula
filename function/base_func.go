package function

import "math"

type function[T float32 | float64] interface {
	Value[T]
	Data() Value[T]
	ListOfData() []Value[T]
	BuildValueAt(index int) T
}

type funcbase[T float32 | float64] struct {
	vector[T]
	data Value[T]
}

func (this funcbase[T]) Data() Value[T] {
	return this.data
}

func (this funcbase[T]) ListOfData() []Value[T] {
	return []Value[T]{this.data}
}

type intfuncbase[T float32 | float64, I uint32 | uint16] struct {
	uintVector[T, I]
	data Value[T]
}

func (this intfuncbase[T, I]) Data() Value[T] {
	return this.data
}

func (this intfuncbase[T, I]) ListOfData() []Value[T] {
	return []Value[T]{this.data}
}

func updateLastValue[T float32 | float64](this function[T]) {
	var length int
	for _, v := range this.ListOfData() {
		if !v.IsScalar() {
			length = v.Len()
			break
		}
	}

	if length < this.Len() || length == 0 {
		return
	}

	if this.Len() == length {
		v := this.BuildValueAt(length - 1)
		this.Set(this.Len()-1, v)
	} else {
		for i := this.Len(); i < length; i++ {
			v := this.BuildValueAt(i)
			this.Append(v)
		}
	}
}

func initValues[T float32 | float64](this function[T], values []T) {
	for i := 0; i < len(values); i++ {
		v := this.BuildValueAt(i)
		values[i] = v
	}
}

func initIntValues[T float32 | float64, I uint32 | uint16](this function[T], values []I) {
	for i := 0; i < len(values); i++ {
		v := this.BuildValueAt(i)
		if math.IsNaN(float64(v)) {
			values[i] = I(uintNaN)
		} else {
			values[i] = I(v)
		}
	}
}

type simplefuncbase[T float32 | float64] struct {
	data Value[T]
}

func (this *simplefuncbase[T]) IsScalar() bool {
	return false
}

func (this *simplefuncbase[T]) Len() int {
	return this.data.Len()
}

func (this *simplefuncbase[T]) Get(index int) float64 {
	panic("Not implemented")
}

func (this *simplefuncbase[T]) Set(index int, v T) {
	panic("Not implemented")
}

func (this *simplefuncbase[T]) UpdateLastValue() {
}

func (this *simplefuncbase[T]) Append(v T) {
	panic("Not implemented")
}

type binaryfuncbase[T float32 | float64] struct {
	simplefuncbase[T]
	data1 Value[T]
}

func (this *binaryfuncbase[T]) Len() int {
	if !this.data.IsScalar() {
		return this.data.Len()
	}
	return this.data1.Len()
}
