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
		this.Set(this.Len() - 1, v)
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
