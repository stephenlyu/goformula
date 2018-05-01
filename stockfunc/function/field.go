package function

import "errors"

type fieldbase struct {
	data *RVector
}

func (this fieldbase) IsScalar() bool {
	return false
}

func (this fieldbase ) Len() int {
	return this.data.Len()
}

func (this fieldbase) Get(index int) float64 {
	panic(errors.New("Not Implemented"))
}

func (this fieldbase) Set(index int, v float64) {
	panic(errors.New("Not Implemented"))
}

func (this fieldbase) UpdateLastValue() {
}

func (this fieldbase) Append(v float64) {
	panic(errors.New("Not Implemented"))
}

// CLOSE

type close struct {
	fieldbase
}

func (this close) Get(index int) float64 {
	return float64(this.data.Get(index).GetClose())
}

func CLOSE(data *RVector) *close {
	ret := &close{
		fieldbase {
			data: data,
		},
	}
	return ret
}

// OPEN

type open struct {
	fieldbase
}

func (this open) Get(index int) float64 {
	return float64(this.data.Get(index).GetOpen())
}

func OPEN(data *RVector) *open {
	ret := &open{
		fieldbase {
			data: data,
		},
	}
	return ret
}

// LOW

type low struct {
	fieldbase
}

func (this low) Get(index int) float64 {
	return float64(this.data.Get(index).GetLow())
}

func LOW(data *RVector) *low {
	ret := &low{
		fieldbase {
			data: data,
		},
	}
	return ret
}

// HIGH

type high struct {
	fieldbase
}

func (this high) Get(index int) float64 {
	return float64(this.data.Get(index).GetHigh())
}

func HIGH(data *RVector) *high {
	ret := &high{
		fieldbase {
			data: data,
		},
	}
	return ret
}

// AMOUNT

type amount struct {
	fieldbase
}

func (this amount) Get(index int) float64 {
	return float64(this.data.Get(index).GetAmount())
}

func AMOUNT(data *RVector) *amount {
	ret := &amount{
		fieldbase {
			data: data,
		},
	}
	return ret
}

// VOLUME

type volume struct {
	fieldbase
}

func (this volume) Get(index int) float64 {
	return float64(this.data.Get(index).GetVolume())
}

func VOLUME(data *RVector) *volume {
	ret := &volume{
		fieldbase {
			data: data,
		},
	}
	return ret
}

// Period

type fPeriod struct {
	fieldbase

	value float64
}

func (this fPeriod) Get(index int) float64 {
	return this.value
}

func PERIOD(data *RVector) *fPeriod {
	ret := &fPeriod{
		fieldbase: fieldbase {
			data: data,
		},
		value: float64(GetPeriodIndex(data.period)),
	}
	return ret
}

// ISLASTBAR

type islastbar struct {
	fieldbase
}

func (this islastbar) Get(index int) float64 {
	if this.data.Len() - 1 == index {
		return 1
	}
	return 0
}

func ISLASTBAR(data *RVector) *islastbar {
	ret := &islastbar{
		fieldbase{
			data: data,
		},
	}
	return ret
}
