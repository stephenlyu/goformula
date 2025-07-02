package function

import (
	"math"
)

type not[T float32 | float64] struct {
	simplefuncbase[T]
}

func BuildNotValueAt[T float32 | float64](data Value[T], index int) T {
	v := data.Get(index)
	if math.IsNaN(float64(v)) {
		return T(math.NaN())
	}
	if v > 0 || v < 0 {
		return 0
	}

	return 1
}

func (this *not[T]) Get(index int) T {
	return BuildNotValueAt(this.data, index)
}

func NOT[T float32 | float64](data Value[T]) Value[T] {
	if data.IsScalar() {
		return Scalar(BuildNotValueAt(data, 0))
	}

	ret := &not[T]{
		simplefuncbase[T]{
			data: data,
		},
	}

	return ret
}

type minus[T float32 | float64] struct {
	simplefuncbase[T]
}

func (this minus[T]) Get(index int) T {
	return -this.data.Get(index)
}

func MINUS[T float32 | float64](data Value[T]) Value[T] {
	if data.IsScalar() {
		return Scalar(-data.Get(0))
	}

	ret := &minus[T]{
		simplefuncbase[T]{
			data: data,
		},
	}

	return ret
}

type add[T float32 | float64] struct {
	binaryfuncbase[T]
}

func (this add[T]) Get(index int) T {
	return this.data.Get(index) + this.data1.Get(index)
}

func ADD[T float32 | float64](data Value[T], data1 Value[T]) Value[T] {
	if data.IsScalar() && data1.IsScalar() {
		return Scalar(data.Get(0) + data1.Get(0))
	}

	ret := &add[T]{
		binaryfuncbase: binaryfuncbase[T]{
			simplefuncbase: simplefuncbase[T]{
				data: data,
			},
			data1: data1,
		},
	}

	return ret
}

type sub[T float32 | float64] struct {
	binaryfuncbase[T]
}

func (this sub[T]) Get(index int) T {
	return this.data.Get(index) - this.data1.Get(index)
}

func SUB[T float32 | float64](data Value[T], data1 Value[T]) Value[T] {
	if data.IsScalar() && data1.IsScalar() {
		return Scalar(data.Get(0) - data1.Get(0))
	}

	ret := &sub[T]{
		binaryfuncbase: binaryfuncbase[T]{
			simplefuncbase: simplefuncbase[T]{
				data: data,
			},
			data1: data1,
		},
	}

	return ret
}

type mul[T float32 | float64] struct {
	binaryfuncbase[T]
}

func (this mul[T]) Get(index int) T {
	return this.data.Get(index) * this.data1.Get(index)
}

func MUL[T float32 | float64](data Value[T], data1 Value[T]) Value[T] {
	if data.IsScalar() && data1.IsScalar() {
		return Scalar(data.Get(0) * data1.Get(0))
	}

	ret := &mul[T]{
		binaryfuncbase: binaryfuncbase[T]{
			simplefuncbase: simplefuncbase[T]{
				data: data,
			},
			data1: data1,
		},
	}

	return ret
}

type div[T float32 | float64] struct {
	binaryfuncbase[T]
}

func (this div[T]) Get(index int) T {
	if this.data1.Get(index) == 0 {
		return T(NaN)
	}
	return this.data.Get(index) / this.data1.Get(index)
}

func DIV[T float32 | float64](data Value[T], data1 Value[T]) Value[T] {
	if data.IsScalar() && data1.IsScalar() {
		if data1.Get(0) == 0 {
			return Scalar(T(NaN))
		}
		return Scalar(data.Get(0) / data1.Get(0))
	}

	ret := &div[T]{
		binaryfuncbase: binaryfuncbase[T]{
			simplefuncbase: simplefuncbase[T]{
				data: data,
			},
			data1: data1,
		},
	}

	return ret
}

type and[T float32 | float64] struct {
	binaryfuncbase[T]
}

func BuildAndValueAt[T float32 | float64](data, data1 Value[T], index int) T {
	a := data.Get(index)
	b := data1.Get(index)
	if IsTrue(a) && IsTrue(b) {
		return 1
	}
	return 0
}

func (this and[T]) Get(index int) T {
	return BuildAndValueAt[T](this.data, this.data1, index)
}

func AND[T float32 | float64](data Value[T], data1 Value[T]) Value[T] {
	if data.IsScalar() && data1.IsScalar() {
		return Scalar(BuildAndValueAt(data, data1, 0))
	}

	ret := &and[T]{
		binaryfuncbase: binaryfuncbase[T]{
			simplefuncbase: simplefuncbase[T]{
				data: data,
			},
			data1: data1,
		},
	}

	return ret
}

type or[T float32 | float64] struct {
	binaryfuncbase[T]
}

func BuildOrValueAt[T float32 | float64](data, data1 Value[T], index int) T {
	a := data.Get(index)
	b := data1.Get(index)
	if IsTrue(a) || IsTrue(b) {
		return 1
	}
	return 0
}

func (this *or[T]) Get(index int) T {
	return BuildOrValueAt(this.data, this.data1, index)
}

func OR[T float32 | float64](data Value[T], data1 Value[T]) Value[T] {
	if data.IsScalar() && data1.IsScalar() {
		return Scalar(BuildOrValueAt(data, data1, 0))
	}

	ret := &or[T]{
		binaryfuncbase: binaryfuncbase[T]{
			simplefuncbase: simplefuncbase[T]{
				data: data,
			},
			data1: data1,
		},
	}

	return ret
}

type lt[T float32 | float64] struct {
	binaryfuncbase[T]
}

func BuildLtValueAt[T float32 | float64](data, data1 Value[T], index int) T {
	a := data.Get(index)
	b := data1.Get(index)
	if a < b {
		return 1
	}
	return 0
}

func (this lt[T]) Get(index int) T {
	return BuildLtValueAt(this.data, this.data1, index)
}

func LT[T float32 | float64](data Value[T], data1 Value[T]) Value[T] {
	if data.IsScalar() && data1.IsScalar() {
		return Scalar(BuildLtValueAt(data, data1, 0))
	}

	ret := &lt[T]{
		binaryfuncbase: binaryfuncbase[T]{
			simplefuncbase: simplefuncbase[T]{
				data: data,
			},
			data1: data1,
		},
	}

	return ret
}

type le[T float32 | float64] struct {
	binaryfuncbase[T]
}

func BuildLEValueAt[T float32 | float64](data, data1 Value[T], index int) T {
	a := data.Get(index)
	b := data1.Get(index)
	if a <= b {
		return 1
	}
	return 0
}

func (this le[T]) Get(index int) T {
	return BuildLEValueAt(this.data, this.data1, index)
}

func LE[T float32 | float64](data Value[T], data1 Value[T]) Value[T] {
	if data.IsScalar() && data1.IsScalar() {
		return Scalar(BuildLEValueAt(data, data1, 0))
	}
	ret := &le[T]{
		binaryfuncbase: binaryfuncbase[T]{
			simplefuncbase: simplefuncbase[T]{
				data: data,
			},
			data1: data1,
		},
	}

	return ret
}

type gt[T float32 | float64] struct {
	binaryfuncbase[T]
}

func BuildGtValueAt[T float32 | float64](data, data1 Value[T], index int) T {
	a := data.Get(index)
	b := data1.Get(index)
	if a > b {
		return 1
	}
	return 0
}

func (this gt[T]) Get(index int) T {
	return BuildGtValueAt(this.data, this.data1, index)
}

func GT[T float32 | float64](data Value[T], data1 Value[T]) Value[T] {
	if data.IsScalar() && data1.IsScalar() {
		return Scalar(BuildGtValueAt(data, data1, 0))
	}
	ret := &gt[T]{
		binaryfuncbase: binaryfuncbase[T]{
			simplefuncbase: simplefuncbase[T]{
				data: data,
			},
			data1: data1,
		},
	}

	return ret
}

type ge[T float32 | float64] struct {
	binaryfuncbase[T]
}

func BuildGeValueAt[T float32 | float64](data, data1 Value[T], index int) T {
	a := data.Get(index)
	b := data1.Get(index)
	if a >= b {
		return 1
	}
	return 0
}

func (this ge[T]) Get(index int) T {
	return BuildGeValueAt(this.data, this.data1, index)
}

func GE[T float32 | float64](data Value[T], data1 Value[T]) Value[T] {
	if data.IsScalar() && data1.IsScalar() {
		return Scalar(BuildGeValueAt(data, data1, 0))
	}
	ret := &ge[T]{
		binaryfuncbase: binaryfuncbase[T]{
			simplefuncbase: simplefuncbase[T]{
				data: data,
			},
			data1: data1,
		},
	}

	return ret
}

type eq[T float32 | float64] struct {
	binaryfuncbase[T]
}

func BuildEqValueAt[T float32 | float64](data, data1 Value[T], index int) T {
	a := data.Get(index)
	b := data1.Get(index)
	if a == b {
		return 1
	}
	return 0
}

func (this eq[T]) Get(index int) T {
	return BuildEqValueAt(this.data, this.data1, index)
}

func EQ[T float32 | float64](data Value[T], data1 Value[T]) Value[T] {
	if data.IsScalar() && data1.IsScalar() {
		return Scalar(BuildEqValueAt(data, data1, 0))
	}
	ret := &eq[T]{
		binaryfuncbase: binaryfuncbase[T]{
			simplefuncbase: simplefuncbase[T]{
				data: data,
			},
			data1: data1,
		},
	}

	return ret
}

type neq[T float32 | float64] struct {
	simplefuncbase[T]
	data1 Value[T]
}

func BuildNeqValueAt[T float32 | float64](data, data1 Value[T], index int) T {
	a := data.Get(index)
	b := data1.Get(index)
	if a != b {
		return 1
	}
	return 0
}

func (this neq[T]) Get(index int) T {
	return BuildNeqValueAt(this.data, this.data1, index)
}

func NEQ[T float32 | float64](data Value[T], data1 Value[T]) Value[T] {
	if data.IsScalar() && data1.IsScalar() {
		return Scalar(BuildNeqValueAt(data, data1, 0))
	}
	ret := &neq[T]{
		simplefuncbase: simplefuncbase[T]{
			data: data,
		},
		data1: data1,
	}

	return ret
}

type between[T float32 | float64] struct {
	simplefuncbase[T]
	data1 Value[T]
	data2 Value[T]
}

func BuildBetweenValueAt[T float32 | float64](data, data1 Value[T], data2 Value[T], index int) T {
	a := data.Get(index)
	b := data1.Get(index)
	c := data2.Get(index)
	if a < b {
		return 0
	}
	if a > c {
		return 0
	}
	return 1
}

func (this *between[T]) Len() int {
	if !this.data.IsScalar() {
		return this.data.Len()
	}
	if !this.data1.IsScalar() {
		return this.data1.Len()
	}
	return this.data2.Len()
}

func (this *between[T]) Get(index int) T {
	return BuildBetweenValueAt(this.data, this.data1, this.data2, index)
}

func BETWEEN[T float32 | float64](data Value[T], data1 Value[T], data2 Value[T]) Value[T] {
	if data.IsScalar() && data1.IsScalar() && data2.IsScalar() {
		return Scalar(BuildBetweenValueAt(data, data1, data2, 0))
	}
	ret := &between[T]{
		simplefuncbase: simplefuncbase[T]{
			data: data,
		},
		data1: data1,
		data2: data2,
	}

	return ret
}
