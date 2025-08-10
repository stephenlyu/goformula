package function

import (
	"errors"
	"fmt"
	"math"
)

type Value[T float32 | float64] interface {
	IsScalar() bool
	Len() int
	Get(index int) T
	Set(index int, v T)
	UpdateLastValue()
	Append(v T)
	LTrim(count int) // 移除左侧的count个元素
}

type scalar[T float32 | float64] struct {
	value T
}

func (this scalar[T]) IsScalar() bool {
	return true
}

func (this scalar[T]) Len() int {
	return -1
}

func (this scalar[T]) Get(index int) T {
	return this.value
}

func (this scalar[T]) Set(index int, v T) {
	this.value = v
}

func (this scalar[T]) Append(v T) {
	this.value = v
}

func (this scalar[T]) UpdateLastValue() {
}

func (this scalar[T]) LTrim(count int) {
	// Do nothing
}

func Scalar[T float32 | float64](v T) *scalar[T] {
	return &scalar[T]{value: v}
}

type vector[T float32 | float64] struct {
	Values []T
}

func (this vector[T]) IsScalar() bool {
	return false
}

func (this vector[T]) Len() int {
	return len(this.Values)
}

func (this vector[T]) Get(index int) T {
	if index < 0 || index >= this.Len() {
		panic(errors.New("index out of range"))
	}
	return this.Values[index]
}

func (this vector[T]) Set(index int, v T) {
	if index < 0 || index >= this.Len() {
		panic(errors.New("index out of range"))
	}
	this.Values[index] = v
}

func (this vector[T]) UpdateLastValue() {
}

func (this *vector[T]) Append(v T) {
	this.Values = append(this.Values, v)
}

func (this *vector[T]) LTrim(count int) {
	if len(this.Values) < count {
		panic("count is greater than vector length")
	}

	this.Values = append([]T{}, this.Values[count:]...)
}

func (this vector[T]) String() string {
	return fmt.Sprintf("%v", this.Values)
}

func Vector[T float32 | float64](values []T) *vector[T] {
	return &vector[T]{Values: values}
}

var uintNaN uint32 = 0xFFFFFFFF

type uintVector[T float32 | float64, I uint32 | uint16] struct {
	Values []I
}

func (this uintVector[T, I]) IsScalar() bool {
	return false
}

func (this uintVector[T, I]) Len() int {
	return len(this.Values)
}

func (this uintVector[T, I]) Get(index int) T {
	if index < 0 || index >= this.Len() {
		panic(errors.New("index out of range"))
	}
	if this.Values[index] == I(uintNaN) {
		return T(math.NaN())
	}
	return T(this.Values[index])
}

func (this uintVector[T, I]) Set(index int, v T) {
	if index < 0 || index >= this.Len() {
		panic(errors.New("index out of range"))
	}
	if math.IsNaN(float64(v)) {
		this.Values[index] = I(uintNaN)
	} else {
		this.Values[index] = I(v)
	}
}

func (this uintVector[T, I]) UpdateLastValue() {
}

func (this *uintVector[T, I]) Append(v T) {
	this.Values = append(this.Values, I(v))
}

func (this *uintVector[T, I]) LTrim(count int) {
	if len(this.Values) < count {
		panic("count is greater than vector length")
	}

	this.Values = append([]I{}, this.Values[count:]...)
}

func (this uintVector[T, I]) String() string {
	return fmt.Sprintf("%v", this.Values)
}

func IntVector[T float32 | float64, I uint32 | uint16](values []I) *uintVector[T, I] {
	return &uintVector[T, I]{Values: values}
}
