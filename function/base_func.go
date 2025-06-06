package function

type function interface {
	Value
	Data() Value
	ListOfData() []Value
	BuildValueAt(index int) float64
}

type funcbase struct {
	vector
	data Value
}

func (this funcbase) Data() Value {
	return this.data
}

func (this funcbase) ListOfData() []Value {
	return []Value{this.data}
}

func updateLastValue(this function) {
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

func initValues(this function, values []float64) {
	for i := 0; i < len(values); i++ {
		v := this.BuildValueAt(i)
		values[i] = v
	}
}

type simplefuncbase struct {
	data Value
}

func (this *simplefuncbase) IsScalar() bool {
	return false
}

func (this *simplefuncbase) Len() int {
	return this.data.Len()
}

func (this *simplefuncbase) Get(index int) float64 {
	panic("Not implemented")
}

func (this *simplefuncbase) Set(index int, v float64) {
	panic("Not implemented")
}

func (this *simplefuncbase) UpdateLastValue() {
}

func (this *simplefuncbase) Append(v float64) {
	panic("Not implemented")
}

type binaryfuncbase struct {
	simplefuncbase
	data1 Value
}

func (this *binaryfuncbase) Len() int {
	if !this.data.IsScalar() {
		return this.data.Len()
	}
	return this.data1.Len()
}
