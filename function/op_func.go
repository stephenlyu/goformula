package function

import (
	"math"
)

type not struct {
	simplefuncbase
}

func BuildNotValueAt(data Value, index int) float64 {
	v := data.Get(index)
	if math.IsNaN(v) {
		return math.NaN()
	}
	if v > 0 || v < 0 {
		return 0
	}

	return 1
}

func (this *not) Get(index int) float64 {
	return BuildNotValueAt(this.data, index)
}

func NOT(data Value) Value {
	if data.IsScalar() {
		return Scalar(BuildNotValueAt(data, 0))
	}

	ret := &not{
		simplefuncbase {
			data: data,
		},
	}

	return ret
}

type minus struct {
	simplefuncbase
}

func (this minus) Get(index int) float64 {
	return -this.data.Get(index)
}

func MINUS(data Value) Value {
	if data.IsScalar() {
		return Scalar(-data.Get(0))
	}

	ret := &minus{
		simplefuncbase {
			data: data,
		},
	}

	return ret
}

type add struct {
	simplefuncbase
	data1 Value
}

func (this add) Get(index int) float64 {
	return this.data.Get(index) + this.data1.Get(index)
}

func ADD(data Value, data1 Value) Value {
	if data.IsScalar() && data1.IsScalar() {
		return Scalar(data.Get(0) + data1.Get(0))
	}

	ret := &add{
		simplefuncbase: simplefuncbase {
			data: data,
		},
		data1: data1,
	}

	return ret
}

type sub struct {
	simplefuncbase
	data1 Value
}

func (this sub) Get(index int) float64 {
	return this.data.Get(index) - this.data1.Get(index)
}

func SUB(data Value, data1 Value) Value {
	if data.IsScalar() && data1.IsScalar() {
		return Scalar(data.Get(0) - data1.Get(0))
	}

	ret := &sub{
		simplefuncbase: simplefuncbase {
			data: data,
		},
		data1: data1,
	}

	return ret
}

type mul struct {
	simplefuncbase
	data1 Value
}

func (this mul) Get(index int) float64 {
	return this.data.Get(index) * this.data1.Get(index)
}

func MUL(data Value, data1 Value) Value {
	if data.IsScalar() && data1.IsScalar() {
		return Scalar(data.Get(0) * data1.Get(0))
	}

	ret := &mul{
		simplefuncbase: simplefuncbase {
			data: data,
		},
		data1: data1,
	}

	return ret
}

type div struct {
	simplefuncbase
	data1 Value
}

func (this div) Get(index int) float64 {
	if this.data1.Get(index) == 0 {
		return NaN
	}
	return this.data.Get(index) / this.data1.Get(index)
}

func DIV(data Value, data1 Value) Value {
	if data.IsScalar() && data1.IsScalar() {
		if data1.Get(0) == 0 {
			return Scalar(NaN)
		}
		return Scalar(data.Get(0) / data1.Get(0))
	}

	ret := &div{
		simplefuncbase: simplefuncbase {
			data: data,
		},
		data1: data1,
	}

	return ret
}

type and struct {
	simplefuncbase
	data1 Value
}

func BuildAndValueAt(data, data1 Value, index int) float64 {
	a := data.Get(index)
	b := data1.Get(index)
	if IsTrue(a) && IsTrue(b) {
		return 1
	}
	return 0
}

func (this and) Get(index int) float64 {
	return BuildAndValueAt(this.data, this.data1, index)
}

func AND(data Value, data1 Value) Value {
	if data.IsScalar() && data1.IsScalar() {
		return Scalar(BuildAndValueAt(data, data1, 0))
	}

	ret := &and{
		simplefuncbase: simplefuncbase {
			data: data,
		},
		data1: data1,
	}

	return ret
}

type or struct {
	simplefuncbase
	data1 Value
}

func BuildOrValueAt(data, data1 Value, index int) float64 {
	a := data.Get(index)
	b := data1.Get(index)
	if IsTrue(a) || IsTrue(b) {
		return 1
	}
	return 0
}

func (this *or) Get(index int) float64 {
	return BuildOrValueAt(this.data, this.data1, index)
}

func OR(data Value, data1 Value) Value {
	if data.IsScalar() && data1.IsScalar() {
		return Scalar(BuildOrValueAt(data, data1, 0))
	}

	ret := &or{
		simplefuncbase: simplefuncbase {
			data: data,
		},
		data1: data1,
	}

	return ret
}

type lt struct {
	simplefuncbase
	data1 Value
}

func BuildLtValueAt(data, data1 Value, index int) float64 {
	a := data.Get(index)
	b := data1.Get(index)
	if a < b {
		return 1
	}
	return 0
}

func (this lt) Get(index int) float64 {
	return BuildLtValueAt(this.data, this.data1, index)
}

func LT(data Value, data1 Value) Value {
	if data.IsScalar() && data1.IsScalar() {
		return Scalar(BuildLtValueAt(data, data1, 0))
	}

	ret := &lt{
		simplefuncbase: simplefuncbase {
			data: data,
		},
		data1: data1,
	}

	return ret
}

type le struct {
	simplefuncbase
	data1 Value
}

func BuildLEValueAt(data, data1 Value, index int) float64 {
	a := data.Get(index)
	b := data1.Get(index)
	if a <= b {
		return 1
	}
	return 0
}

func (this le) Get(index int) float64 {
	return BuildLEValueAt(this.data, this.data1, index)
}

func LE(data Value, data1 Value) Value {
	if data.IsScalar() && data1.IsScalar() {
		return Scalar(BuildLEValueAt(data, data1, 0))
	}
	ret := &le{
		simplefuncbase: simplefuncbase {
			data: data,
		},
		data1: data1,
	}

	return ret
}

type gt struct {
	simplefuncbase
	data1 Value
}

func BuildGtValueAt(data, data1 Value, index int) float64 {
	a := data.Get(index)
	b := data1.Get(index)
	if a > b {
		return 1
	}
	return 0
}

func (this gt) Get(index int) float64 {
	return BuildGtValueAt(this.data, this.data1, index)
}

func GT(data Value, data1 Value) Value {
	if data.IsScalar() && data1.IsScalar() {
		return Scalar(BuildGtValueAt(data, data1, 0))
	}
	ret := &gt{
		simplefuncbase: simplefuncbase {
			data: data,
		},
		data1: data1,
	}

	return ret
}

type ge struct {
	simplefuncbase
	data1 Value
}

func BuildGeValueAt(data, data1 Value, index int) float64 {
	a := data.Get(index)
	b := data1.Get(index)
	if a >= b {
		return 1
	}
	return 0
}

func (this ge) Get(index int) float64 {
	return BuildGeValueAt(this.data, this.data1, index)
}

func GE(data Value, data1 Value) Value {
	if data.IsScalar() && data1.IsScalar() {
		return Scalar(BuildGeValueAt(data, data1, 0))
	}
	ret := &ge{
		simplefuncbase: simplefuncbase {
			data: data,
		},
		data1: data1,
	}

	return ret
}

type eq struct {
	simplefuncbase
	data1 Value
}

func BuildEqValueAt(data, data1 Value, index int) float64 {
	a := data.Get(index)
	b := data1.Get(index)
	if a == b {
		return 1
	}
	return 0
}

func (this eq) Get(index int) float64 {
	return BuildEqValueAt(this.data, this.data1, index)
}

func EQ(data Value, data1 Value) Value {
	if data.IsScalar() && data1.IsScalar() {
		return Scalar(BuildEqValueAt(data, data1, 0))
	}
	ret := &eq{
		simplefuncbase: simplefuncbase {
			data: data,
		},
		data1: data1,
	}

	return ret
}

type neq struct {
	simplefuncbase
	data1 Value
}

func BuildNeqValueAt(data, data1 Value, index int) float64 {
	a := data.Get(index)
	b := data1.Get(index)
	if a != b {
		return 1
	}
	return 0
}

func (this neq) Get(index int) float64 {
	return BuildNeqValueAt(this.data, this.data1, index)
}

func NEQ(data Value, data1 Value) Value {
	if data.IsScalar() && data1.IsScalar() {
		return Scalar(BuildNeqValueAt(data, data1, 0))
	}
	ret := &neq{
		simplefuncbase: simplefuncbase {
			data: data,
		},
		data1: data1,
	}

	return ret
}
