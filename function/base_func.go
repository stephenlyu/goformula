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
	if this.Data().Len() < this.Len() || this.Data().Len() == 0 {
		return
	}

	if this.Len() == this.Data().Len() {
		v := this.BuildValueAt(this.Data().Len() - 1)
		this.Set(this.Len() - 1, v)
	} else {
		for i := this.Len(); i < this.Data().Len(); i++ {
			v := this.BuildValueAt(i)
			this.Append(v)
		}
	}
}

func initValues(this function, values []float64) {
	for i := 0; i < this.Data().Len(); i++ {
		v := this.BuildValueAt(i)
		values[i] = v
	}
}
