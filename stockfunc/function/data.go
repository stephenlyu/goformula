package function

import (
	"errors"
	"sync"

	. "github.com/stephenlyu/tds/period"
)

type Record[T float32 | float64] interface {
	GetUTCDate() uint64
	GetDate() string
	GetOpen() T
	GetClose() T
	GetHigh() T
	GetLow() T
	GetAmount() T
	GetVolume() T
}

type RVectorReader[T float32 | float64] interface {
	Len() int
	Get(index int) Record[T]
	Code() string
	Period() Period
}

type RVectorUpdater[T float32 | float64] interface {
	Set(index int, v Record[T])
	Append(v Record[T])
	Update(offset int, values []Record[T])
}

type RVector[T float32 | float64] interface {
	RVectorReader[T]
	RVectorUpdater[T]
}

type rVector[T float32 | float64] struct {
	Values []Record[T]

	code   string
	period Period

	lock sync.RWMutex
}

func (this *rVector[T]) Len() int {
	this.lock.RLock()
	defer this.lock.RUnlock()
	return len(this.Values)
}

func (this *rVector[T]) Get(index int) Record[T] {
	this.lock.RLock()
	defer this.lock.RUnlock()
	if index < 0 || index >= this.Len() {
		panic(errors.New("index out of range"))
	}
	return this.Values[index]
}

func (this *rVector[T]) Set(index int, v Record[T]) {
	this.lock.Lock()
	defer this.lock.Unlock()
	if index < 0 || index >= this.Len() {
		panic(errors.New("index out of range"))
	}
	this.Values[index] = v
}

func (this *rVector[T]) Append(v Record[T]) {
	this.lock.Lock()
	defer this.lock.Unlock()
	this.Values = append(this.Values, v)
}

func (this *rVector[T]) Update(offset int, values []Record[T]) {
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

func (this *rVector[T]) Code() string {
	return this.code
}

func (this *rVector[T]) Period() Period {
	return this.period
}

func RecordVector[T float32 | float64](v []Record[T]) RVector[T] {
	return &rVector[T]{Values: v}
}

func RecordVectorEx[T float32 | float64](code string, period Period, v []Record[T]) RVector[T] {
	return &rVector[T]{
		Values: v,
		code:   code,
		period: period,
	}
}
