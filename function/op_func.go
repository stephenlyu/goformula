package function

import (
	"math"
	"fmt"
	"github.com/stephenlyu/tds/util"
)

type not struct {
	funcbase
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

func (this not) BuildValueAt(index int) float64 {
	return BuildNotValueAt(this.data, index)
}

func (this *not) UpdateLastValue() {
	updateLastValue(this)
}

func NOT(data Value) Value {
	if data.IsScalar() {
		return Scalar(BuildNotValueAt(data, 0))
	}

	ret := &not{
		funcbase {
			data: data,
		},
	}
	ret.Values = make([]float64, data.Len())
	initValues(ret, ret.Values)

	return ret
}

type minus struct {
	funcbase
}

func (this minus) BuildValueAt(index int) float64 {
	return -this.data.Get(index)
}

func (this *minus) UpdateLastValue() {
	updateLastValue(this)
}

func MINUS(data Value) Value {
	if data.IsScalar() {
		return Scalar(-data.Get(0))
	}

	ret := &minus{
		funcbase {
			data: data,
		},
	}
	ret.Values = make([]float64, data.Len())
	initValues(ret, ret.Values)

	return ret
}

type add struct {
	funcbase
	data1 Value
}

func (this add) BuildValueAt(index int) float64 {
	return this.data.Get(index) + this.data1.Get(index)
}

func (this *add) UpdateLastValue() {
	updateLastValue(this)
}

func ADD(data Value, data1 Value) Value {
	if data.IsScalar() && data1.IsScalar() {
		return Scalar(data.Get(0) + data1.Get(0))
	}
	var length int
	if !data.IsScalar() {
		length = data.Len()
	} else {
		length = data1.Len()
	}

	ret := &add{
		funcbase: funcbase {
			data: data,
		},
		data1: data1,
	}
	ret.Values = make([]float64, length)
	initValues(ret, ret.Values)

	return ret
}

func (this *add) ListOfData() []Value {
	return []Value{this.data, this.data1}
}

type sub struct {
	funcbase
	data1 Value
}

func (this sub) BuildValueAt(index int) float64 {
	return this.data.Get(index) - this.data1.Get(index)
}

func (this *sub) UpdateLastValue() {
	updateLastValue(this)
}

func SUB(data Value, data1 Value) Value {
	if data.IsScalar() && data1.IsScalar() {
		return Scalar(data.Get(0) - data1.Get(0))
	}
	var length int
	if !data.IsScalar() {
		length = data.Len()
	} else {
		length = data1.Len()
	}

	ret := &sub{
		funcbase: funcbase {
			data: data,
		},
		data1: data1,
	}
	ret.Values = make([]float64, length)
	initValues(ret, ret.Values)

	return ret
}

func (this *sub) ListOfData() []Value {
	return []Value{this.data, this.data1}
}

type mul struct {
	funcbase
	data1 Value
}

func (this mul) BuildValueAt(index int) float64 {
	return this.data.Get(index) * this.data1.Get(index)
}

func (this *mul) UpdateLastValue() {
	updateLastValue(this)
}

func MUL(data Value, data1 Value) Value {
	if data.IsScalar() && data1.IsScalar() {
		return Scalar(data.Get(0) * data1.Get(0))
	}
	var length int
	if !data.IsScalar() {
		length = data.Len()
	} else {
		length = data1.Len()
	}

	ret := &mul{
		funcbase: funcbase {
			data: data,
		},
		data1: data1,
	}
	ret.Values = make([]float64, length)
	initValues(ret, ret.Values)

	return ret
}

func (this *mul) ListOfData() []Value {
	return []Value{this.data, this.data1}
}

type div struct {
	funcbase
	data1 Value
}

func (this div) BuildValueAt(index int) float64 {
	if this.data1.Get(index) == 0 {
		return NaN
	}
	return this.data.Get(index) / this.data1.Get(index)
}

func (this *div) UpdateLastValue() {
	updateLastValue(this)
}

func DIV(data Value, data1 Value) Value {
	if data.IsScalar() && data1.IsScalar() {
		if data1.Get(0) == 0 {
			return Scalar(NaN)
		}
		return Scalar(data.Get(0) / data1.Get(0))
	}
	var length int
	if !data.IsScalar() {
		length = data.Len()
	} else {
		length = data1.Len()
	}

	ret := &div{
		funcbase: funcbase {
			data: data,
		},
		data1: data1,
	}
	ret.Values = make([]float64, length)
	initValues(ret, ret.Values)

	return ret
}

func (this *div) ListOfData() []Value {
	return []Value{this.data, this.data1}
}

type and struct {
	funcbase
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

func (this and) BuildValueAt(index int) float64 {
	return BuildAndValueAt(this.data, this.data1, index)
}

func (this *and) UpdateLastValue() {
	updateLastValue(this)
}

func AND(data Value, data1 Value) Value {
	if data.IsScalar() && data1.IsScalar() {
		return Scalar(BuildAndValueAt(data, data1, 0))
	}
	var length int
	if !data.IsScalar() {
		length = data.Len()
	} else {
		length = data1.Len()
	}

	ret := &and{
		funcbase: funcbase {
			data: data,
		},
		data1: data1,
	}
	ret.Values = make([]float64, length)
	initValues(ret, ret.Values)

	return ret
}

func (this *and) ListOfData() []Value {
	return []Value{this.data, this.data1}
}

type or struct {
	funcbase
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

func (this or) BuildValueAt(index int) float64 {
	return BuildOrValueAt(this.data, this.data1, index)
}

func (this *or) UpdateLastValue() {
	updateLastValue(this)
}

func OR(data Value, data1 Value) Value {
	if data.IsScalar() && data1.IsScalar() {
		return Scalar(BuildOrValueAt(data, data1, 0))
	}
	var length int
	if !data.IsScalar() {
		length = data.Len()
	} else {
		length = data1.Len()
	}

	ret := &or{
		funcbase: funcbase {
			data: data,
		},
		data1: data1,
	}
	ret.Values = make([]float64, length)
	initValues(ret, ret.Values)

	return ret
}

func (this *or) ListOfData() []Value {
	return []Value{this.data, this.data1}
}

type lt struct {
	funcbase
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

func (this lt) BuildValueAt(index int) float64 {
	return BuildLtValueAt(this.data, this.data1, index)
}

func (this *lt) UpdateLastValue() {
	updateLastValue(this)
}

func LT(data Value, data1 Value) Value {
	if data.IsScalar() && data1.IsScalar() {
		return Scalar(BuildLtValueAt(data, data1, 0))
	}
	var length int
	if !data.IsScalar() {
		length = data.Len()
	} else {
		length = data1.Len()
	}

	ret := &lt{
		funcbase: funcbase {
			data: data,
		},
		data1: data1,
	}
	ret.Values = make([]float64, length)
	initValues(ret, ret.Values)

	return ret
}

func (this *lt) ListOfData() []Value {
	return []Value{this.data, this.data1}
}

type le struct {
	funcbase
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

func (this le) BuildValueAt(index int) float64 {
	return BuildLEValueAt(this.data, this.data1, index)
}

func (this *le) UpdateLastValue() {
	updateLastValue(this)
}

func LE(data Value, data1 Value) Value {
	if data.IsScalar() && data1.IsScalar() {
		return Scalar(BuildLEValueAt(data, data1, 0))
	}
	var length int
	if !data.IsScalar() {
		length = data.Len()
	} else {
		length = data1.Len()
	}
	ret := &le{
		funcbase: funcbase {
			data: data,
		},
		data1: data1,
	}
	ret.Values = make([]float64, length)
	initValues(ret, ret.Values)

	return ret
}

func (this *le) ListOfData() []Value {
	return []Value{this.data, this.data1}
}

type gt struct {
	funcbase
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

func (this gt) BuildValueAt(index int) float64 {
	return BuildGtValueAt(this.data, this.data1, index)
}

func (this *gt) UpdateLastValue() {
	updateLastValue(this)
}

func GT(data Value, data1 Value) Value {
	if data.IsScalar() && data1.IsScalar() {
		return Scalar(BuildGtValueAt(data, data1, 0))
	}
	var length int
	if !data.IsScalar() {
		length = data.Len()
	} else {
		length = data1.Len()
	}
	ret := &gt{
		funcbase: funcbase {
			data: data,
		},
		data1: data1,
	}
	ret.Values = make([]float64, length)
	initValues(ret, ret.Values)

	return ret
}

func (this *gt) ListOfData() []Value {
	return []Value{this.data, this.data1}
}

type ge struct {
	funcbase
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

func (this ge) BuildValueAt(index int) float64 {
	return BuildGeValueAt(this.data, this.data1, index)
}

func (this *ge) UpdateLastValue() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(string(util.PanicTrace(24)))
		}
	}()
	updateLastValue(this)
}

func GE(data Value, data1 Value) Value {
	if data.IsScalar() && data1.IsScalar() {
		return Scalar(BuildGeValueAt(data, data1, 0))
	}
	var length int
	if !data.IsScalar() {
		length = data.Len()
	} else {
		length = data1.Len()
	}
	ret := &ge{
		funcbase: funcbase {
			data: data,
		},
		data1: data1,
	}
	ret.Values = make([]float64, length)
	initValues(ret, ret.Values)

	return ret
}

func (this *ge) ListOfData() []Value {
	return []Value{this.data, this.data1}
}

type eq struct {
	funcbase
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

func (this eq) BuildValueAt(index int) float64 {
	return BuildEqValueAt(this.data, this.data1, index)
}

func (this *eq) UpdateLastValue() {
	updateLastValue(this)
}

func EQ(data Value, data1 Value) Value {
	if data.IsScalar() && data1.IsScalar() {
		return Scalar(BuildEqValueAt(data, data1, 0))
	}
	var length int
	if !data.IsScalar() {
		length = data.Len()
	} else {
		length = data1.Len()
	}
	ret := &eq{
		funcbase: funcbase {
			data: data,
		},
		data1: data1,
	}
	ret.Values = make([]float64, length)
	initValues(ret, ret.Values)

	return ret
}

func (this *eq) ListOfData() []Value {
	return []Value{this.data, this.data1}
}

type neq struct {
	funcbase
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

func (this neq) BuildValueAt(index int) float64 {
	return BuildNeqValueAt(this.data, this.data1, index)
}

func (this *neq) UpdateLastValue() {
	updateLastValue(this)
}

func NEQ(data Value, data1 Value) Value {
	if data.IsScalar() && data1.IsScalar() {
		return Scalar(BuildNeqValueAt(data, data1, 0))
	}
	var length int
	if !data.IsScalar() {
		length = data.Len()
	} else {
		length = data1.Len()
	}
	ret := &neq{
		funcbase: funcbase {
			data: data,
		},
		data1: data1,
	}
	ret.Values = make([]float64, length)
	initValues(ret, ret.Values)

	return ret
}

func (this *neq) ListOfData() []Value {
	return []Value{this.data, this.data1}
}
