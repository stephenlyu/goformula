package function


type function interface {
	Value
	Data() Value
	BuildValueAt(index int) float64
}

type funcbase struct {
	vector
	data Value
}

func (this funcbase) Data() Value {
	return this.data
}

func updateLastValue(this function) {
	if this.Data().Len() < this.Len() {
		return
	}
	v := this.BuildValueAt(this.Data().Len() - 1)
	if this.Len() == this.Data().Len() {
		this.Set(this.Len() - 1, v)
	} else {
		this.Append(v)
	}
}

func initValues(this function) []float64 {
	values := make([]float64, this.Data().Len())
	for i := 0; i < this.Data().Len(); i++ {
		v := this.BuildValueAt(i)
		values[i] = v
	}
	return values
}
