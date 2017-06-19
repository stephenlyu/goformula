package function

import "errors"

type Record struct {
	Date string				`json:"date"`
	Open float32			`json:"open"`
	Close float32			`json:"close"`
	High float32			`json:"high"`
	Low float32				`json:"low"`
	Volume float32			`json:"volume"`
	Amount float32			`json:"amount"`
}


type RVector struct {
	Values 	[]Record
}

func (this RVector) Len() int {
	return len(this.Values)
}

func (this RVector) Get(index int) *Record {
	if index < 0 || index >= this.Len() {
		panic(errors.New("index out of range"))
	}
	return &this.Values[index]
}

func (this RVector) Set(index int, v *Record) {
	if index < 0 || index >= this.Len() {
		panic(errors.New("index out of range"))
	}
	this.Values[index] = *v
}

func (this RVector) Append(v *Record) {
	this.Values = append(this.Values, *v)
}

func RecordVector(v []Record) *RVector {
	return &RVector{Values: v}
}
