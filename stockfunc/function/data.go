package function

import (
	"errors"
	"sync"

	. "github.com/stephenlyu/tds/period"
)

type Record interface {
	GetUTCDate() uint64
	GetDate() string
	GetOpen() float64
	GetClose() float64
	GetHigh() float64
	GetLow() float64
	GetAmount() float64
	GetVolume() float64
}

type RVectorReader interface {
	Len() int
	Get(index int) Record
	Code() string
	Period() Period
}

type RVectorUpdater interface {
	Set(index int, v Record)
	Append(v Record)
	Update(offset int, values []Record)
}

type RVector interface {
	RVectorReader
	RVectorUpdater
}

type rVector struct {
	Values []Record

	code   string
	period Period

	lock sync.RWMutex
}

func (this *rVector) Len() int {
	this.lock.RLock()
	defer this.lock.RUnlock()
	return len(this.Values)
}

func (this *rVector) Get(index int) Record {
	this.lock.RLock()
	defer this.lock.RUnlock()
	if index < 0 || index >= this.Len() {
		panic(errors.New("index out of range"))
	}
	return this.Values[index]
}

func (this *rVector) Set(index int, v Record) {
	this.lock.Lock()
	defer this.lock.Unlock()
	if index < 0 || index >= this.Len() {
		panic(errors.New("index out of range"))
	}
	this.Values[index] = v
}

func (this *rVector) Append(v Record) {
	this.lock.Lock()
	defer this.lock.Unlock()
	this.Values = append(this.Values, v)
}

func (this *rVector) Update(offset int, values []Record) {
	this.lock.Lock()
	defer this.lock.Unlock()

	for i, v := range values {
		if offset+i < len(this.Values) {
			this.Values[offset+i] = v
		} else {
			this.Values = append(this.Values, v)
		}
	}
}

func (this *rVector) Code() string {
	return this.code
}

func (this *rVector) Period() Period {
	return this.period
}

func RecordVector(v []Record) RVector {
	return &rVector{Values: v}
}

func RecordVectorEx(code string, period Period, v []Record) RVector {
	return &rVector{
		Values: v,
		code:   code,
		period: period,
	}
}
