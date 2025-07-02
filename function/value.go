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

func (this vector[T]) String() string {
	return fmt.Sprintf("%v", this.Values)
}

func Vector[T float32 | float64](values []T) *vector[T] {
	return &vector[T]{Values: values}
}

var intNaN int = 0xdeadbeaf

type intVector[T float32 | float64] struct {
	Values []int
}

func (this intVector[T]) IsScalar() bool {
	return false
}

func (this intVector[T]) Len() int {
	return len(this.Values)
}

func (this intVector[T]) Get(index int) T {
	if index < 0 || index >= this.Len() {
		panic(errors.New("index out of range"))
	}
	if this.Values[index] == intNaN {
		return T(math.NaN())
	}
	return T(this.Values[index])
}

func (this intVector[T]) Set(index int, v T) {
	if index < 0 || index >= this.Len() {
		panic(errors.New("index out of range"))
	}
	if math.IsNaN(float64(v)) {
		this.Values[index] = intNaN
	} else {
		this.Values[index] = int(v)
	}
}

func (this intVector[T]) UpdateLastValue() {
}

func (this *intVector[T]) Append(v T) {
	this.Values = append(this.Values, int(v))
}

func (this intVector[T]) String() string {
	return fmt.Sprintf("%v", this.Values)
}

func IntVector[T float32 | float64](values []int) *intVector[T] {
	return &intVector[T]{Values: values}
}
