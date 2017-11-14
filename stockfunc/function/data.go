package function

import (
	"errors"
	. "github.com/stephenlyu/tds/period"
)


type Record interface {
	GetUTCDate() uint64
	GetDate() string
	GetOpen() float32
	GetClose() float32
	GetHigh() float32
	GetLow() float32
	GetAmount() float32
	GetVolume() float32
}

type RVector struct {
	Values 	[]Record

	code string
	period Period
}

func (this *RVector) Len() int {
	return len(this.Values)
}

func (this *RVector) Get(index int) Record {
	if index < 0 || index >= this.Len() {
		panic(errors.New("index out of range"))
	}
	return this.Values[index]
}

func (this *RVector) Set(index int, v Record) {
	if index < 0 || index >= this.Len() {
		panic(errors.New("index out of range"))
	}
	this.Values[index] = v
}

func (this *RVector) Append(v Record) {
	this.Values = append(this.Values, v)
}

func (this *RVector) Code() string {
	return this.code
}

func (this *RVector) Period() Period {
	return this.period
}

func RecordVector(v []Record) *RVector {
	return &RVector{Values: v}
}

func RecordVectorEx(code string, period Period, v []Record) *RVector {
	return &RVector{
		Values: v,
		code: code,
		period: period,
	}
}
