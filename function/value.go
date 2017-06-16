package function

import (
	"errors"
	"fmt"
)

type Value interface {
	IsScalar() bool
	Len() int
	Get(index int) float64
	Set(index int, v float64)
	UpdateLastValue()
	Append(v float64)
}


type scalar struct {
	value float64
}

func (this scalar) IsScalar() bool {
	return true
}

func (this scalar) Len() int {
	return -1
}

func (this scalar) Get(index int) float64 {
	return this.value
}

func (this scalar) Set(index int, v float64) {
	this.value = v
}

func (this scalar) Append(v float64) {
	this.value = v
}

func (this scalar) UpdateLastValue() {
}

func Scalar(v float64) *scalar {
	return &scalar{value: v}
}

type vector struct {
	Values 	[]float64
}


func (this vector) IsScalar() bool {
	return false
}

func (this vector) Len() int {
	return len(this.Values)
}

func (this vector) Get(index int) float64 {
	if index < 0 || index >= this.Len() {
		panic(errors.New("index out of range"))
	}
	return this.Values[index]
}

func (this vector) Set(index int, v float64) {
	if index < 0 || index >= this.Len() {
		panic(errors.New("index out of range"))
	}
	this.Values[index] = v
}

func (this vector) UpdateLastValue() {
}

func (this *vector) Append(v float64) {
	this.Values = append(this.Values, v)
}

func (this vector) String() string {
	return fmt.Sprintf("%v", this.Values)
}

func Vector(values []float64) *vector {
	return &vector{Values: values}
}
