package function

import "math"

type not struct {
	funcbase
}

func (this not) BuildValueAt(index int) float64 {
	v := this.data.Get(index)
	if math.IsNaN(v) {
		return math.NaN()
	}
	if v > 0 || v < 0 {
		return 0
	}

	return 1
}

func (this *not) UpdateLastValue() {
	updateLastValue(this)
}

func NOT(data Value) *not {
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

func MINUS(data Value) *minus {
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

func ADD(data Value, data1 Value) *add {
	ret := &add{
		funcbase: funcbase {
			data: data,
		},
		data1: data1,
	}
	ret.Values = make([]float64, data.Len())
	initValues(ret, ret.Values)

	return ret
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

func SUB(data Value, data1 Value) *sub {
	ret := &sub{
		funcbase: funcbase {
			data: data,
		},
		data1: data1,
	}
	ret.Values = make([]float64, data.Len())
	initValues(ret, ret.Values)

	return ret
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

func MUL(data Value, data1 Value) *mul {
	ret := &mul{
		funcbase: funcbase {
			data: data,
		},
		data1: data1,
	}
	ret.Values = make([]float64, data.Len())
	initValues(ret, ret.Values)

	return ret
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

func DIV(data Value, data1 Value) *div {
	ret := &div{
		funcbase: funcbase {
			data: data,
		},
		data1: data1,
	}
	ret.Values = make([]float64, data.Len())
	initValues(ret, ret.Values)

	return ret
}

type and struct {
	funcbase
	data1 Value
}

func (this and) BuildValueAt(index int) float64 {
	a := this.data.Get(index)
	b := this.data1.Get(index)
	if IsTrue(a) && IsTrue(b) {
		return 1
	}
	return 0
}

func (this *and) UpdateLastValue() {
	updateLastValue(this)
}

func AND(data Value, data1 Value) *and {
	ret := &and{
		funcbase: funcbase {
			data: data,
		},
		data1: data1,
	}
	ret.Values = make([]float64, data.Len())
	initValues(ret, ret.Values)

	return ret
}

type or struct {
	funcbase
	data1 Value
}

func (this or) BuildValueAt(index int) float64 {
	a := this.data.Get(index)
	b := this.data1.Get(index)
	if IsTrue(a) || IsTrue(b) {
		return 1
	}
	return 0
}

func (this *or) UpdateLastValue() {
	updateLastValue(this)
}

func OR(data Value, data1 Value) *or {
	ret := &or{
		funcbase: funcbase {
			data: data,
		},
		data1: data1,
	}
	ret.Values = make([]float64, data.Len())
	initValues(ret, ret.Values)

	return ret
}

type lt struct {
	funcbase
	data1 Value
}

func (this lt) BuildValueAt(index int) float64 {
	a := this.data.Get(index)
	b := this.data1.Get(index)
	if a < b {
		return 1
	}
	return 0
}

func (this *lt) UpdateLastValue() {
	updateLastValue(this)
}

func LT(data Value, data1 Value) *lt {
	ret := &lt{
		funcbase: funcbase {
			data: data,
		},
		data1: data1,
	}
	ret.Values = make([]float64, data.Len())
	initValues(ret, ret.Values)

	return ret
}

type le struct {
	funcbase
	data1 Value
}

func (this le) BuildValueAt(index int) float64 {
	a := this.data.Get(index)
	b := this.data1.Get(index)
	if a <= b {
		return 1
	}
	return 0
}

func (this *le) UpdateLastValue() {
	updateLastValue(this)
}

func LE(data Value, data1 Value) *le {
	ret := &le{
		funcbase: funcbase {
			data: data,
		},
		data1: data1,
	}
	ret.Values = make([]float64, data.Len())
	initValues(ret, ret.Values)

	return ret
}

type gt struct {
	funcbase
	data1 Value
}

func (this gt) BuildValueAt(index int) float64 {
	a := this.data.Get(index)
	b := this.data1.Get(index)
	if a > b {
		return 1
	}
	return 0
}

func (this *gt) UpdateLastValue() {
	updateLastValue(this)
}

func GT(data Value, data1 Value) *gt {
	ret := &gt{
		funcbase: funcbase {
			data: data,
		},
		data1: data1,
	}
	ret.Values = make([]float64, data.Len())
	initValues(ret, ret.Values)

	return ret
}

type ge struct {
	funcbase
	data1 Value
}

func (this ge) BuildValueAt(index int) float64 {
	a := this.data.Get(index)
	b := this.data1.Get(index)
	if a >= b {
		return 1
	}
	return 0
}

func (this *ge) UpdateLastValue() {
	updateLastValue(this)
}

func GE(data Value, data1 Value) *ge {
	ret := &ge{
		funcbase: funcbase {
			data: data,
		},
		data1: data1,
	}
	ret.Values = make([]float64, data.Len())
	initValues(ret, ret.Values)

	return ret
}

type eq struct {
	funcbase
	data1 Value
}

func (this eq) BuildValueAt(index int) float64 {
	a := this.data.Get(index)
	b := this.data1.Get(index)
	if a == b {
		return 1
	}
	return 0
}

func (this *eq) UpdateLastValue() {
	updateLastValue(this)
}

func EQ(data Value, data1 Value) *eq {
	ret := &eq{
		funcbase: funcbase {
			data: data,
		},
		data1: data1,
	}
	ret.Values = make([]float64, data.Len())
	initValues(ret, ret.Values)

	return ret
}

type neq struct {
	funcbase
	data1 Value
}

func (this neq) BuildValueAt(index int) float64 {
	a := this.data.Get(index)
	b := this.data1.Get(index)
	if a != b {
		return 1
	}
	return 0
}

func (this *neq) UpdateLastValue() {
	updateLastValue(this)
}

func NEQ(data Value, data1 Value) *neq {
	ret := &neq{
		funcbase: funcbase {
			data: data,
		},
		data1: data1,
	}
	ret.Values = make([]float64, data.Len())
	initValues(ret, ret.Values)

	return ret
}