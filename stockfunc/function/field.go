package function

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/stephenlyu/goformula/function"
)

type fieldbase[T float32 | float64] struct {
	data     RVectorReader[T]
	refCount uint32
}

func (this fieldbase[T]) IsScalar() bool {
	return false
}

func (this *fieldbase[T]) IncRef() {
	this.refCount++
}

func (this *fieldbase[T]) DecRef() {
	this.refCount--
	if this.refCount <= 0 {
		function.GlobalValueCache.Del(this.GetCacheKey())
	}
}

func (this fieldbase[T]) GetCacheKey() string {
	return ""
}

func (this fieldbase[T]) Len() int {
	return this.data.Len()
}

func (this fieldbase[T]) Get(index int) T {
	panic(errors.New("Not Implemented"))
}

func (this fieldbase[T]) Set(index int, v T) {
	panic(errors.New("Not Implemented"))
}

func (this fieldbase[T]) UpdateLastValue() {
}

func (this fieldbase[T]) Append(v T) {
	panic(errors.New("Not Implemented"))
}

func (this fieldbase[T]) LTrim(count int) {
}

func fieldGetCacheKey(value any) string {
	va := reflect.ValueOf(value)
	ptr := va.Pointer()
	return fmt.Sprintf("%x", ptr)
}

// CLOSE

type close[T float32 | float64] struct {
	fieldbase[T]
}

func (this close[T]) Get(index int) T {
	return this.data.Get(index).GetClose()
}

func (this *close[T]) GetCacheKey() string {
	return "close" + fieldGetCacheKey(this.data)
}

func CLOSE[T float32 | float64](data RVectorReader[T]) *close[T] {
	key := "close" + fieldGetCacheKey(data)
	if v := function.GlobalValueCache.Get(key); v != nil {
		return v.(*close[T])
	}
	ret := &close[T]{
		fieldbase[T]{
			data: data,
		},
	}
	ret.IncRef()
	function.GlobalValueCache.Set(key, ret)
	return ret
}

// OPEN

type open[T float32 | float64] struct {
	fieldbase[T]
}

func (this open[T]) Get(index int) T {
	return this.data.Get(index).GetOpen()
}

func (this *open[T]) GetCacheKey() string {
	return "open" + fieldGetCacheKey(this.data)
}

func OPEN[T float32 | float64](data RVectorReader[T]) *open[T] {
	key := "open" + fieldGetCacheKey(data)
	if v := function.GlobalValueCache.Get(key); v != nil {
		return v.(*open[T])
	}
	ret := &open[T]{
		fieldbase[T]{
			data: data,
		},
	}
	ret.IncRef()
	function.GlobalValueCache.Set(key, ret)
	return ret
}

// LOW

type low[T float32 | float64] struct {
	fieldbase[T]
}

func (this low[T]) Get(index int) T {
	return this.data.Get(index).GetLow()
}

func (this *low[T]) GetCacheKey() string {
	return "low" + fieldGetCacheKey(this.data)
}

func LOW[T float32 | float64](data RVectorReader[T]) *low[T] {
	key := "low" + fieldGetCacheKey(data)
	if v := function.GlobalValueCache.Get(key); v != nil {
		return v.(*low[T])
	}
	ret := &low[T]{
		fieldbase[T]{
			data: data,
		},
	}
	ret.IncRef()
	function.GlobalValueCache.Set(key, ret)
	return ret
}

// HIGH

type high[T float32 | float64] struct {
	fieldbase[T]
}

func (this high[T]) Get(index int) T {
	return this.data.Get(index).GetHigh()
}

func (this *high[T]) GetCacheKey() string {
	return "high" + fieldGetCacheKey(this.data)
}

func HIGH[T float32 | float64](data RVectorReader[T]) *high[T] {
	key := "high" + fieldGetCacheKey(data)
	if v := function.GlobalValueCache.Get(key); v != nil {
		return v.(*high[T])
	}
	ret := &high[T]{
		fieldbase[T]{
			data: data,
		},
	}
	ret.IncRef()
	function.GlobalValueCache.Set(key, ret)
	return ret
}

// AMOUNT

type amount[T float32 | float64] struct {
	fieldbase[T]
}

func (this amount[T]) Get(index int) T {
	return this.data.Get(index).GetAmount()
}

func (this *amount[T]) GetCacheKey() string {
	return "amount" + fieldGetCacheKey(this.data)
}

func AMOUNT[T float32 | float64](data RVectorReader[T]) *amount[T] {
	key := "amount" + fieldGetCacheKey(data)
	if v := function.GlobalValueCache.Get(key); v != nil {
		return v.(*amount[T])
	}
	ret := &amount[T]{
		fieldbase[T]{
			data: data,
		},
	}
	return ret
}

// VOLUME

type volume[T float32 | float64] struct {
	fieldbase[T]
}

func (this volume[T]) Get(index int) T {
	return this.data.Get(index).GetVolume()
}

func (this *volume[T]) GetCacheKey() string {
	return "volume" + fieldGetCacheKey(this.data)
}

func VOLUME[T float32 | float64](data RVectorReader[T]) *volume[T] {
	key := "volume" + fieldGetCacheKey(data)
	if v := function.GlobalValueCache.Get(key); v != nil {
		return v.(*volume[T])
	}
	ret := &volume[T]{
		fieldbase[T]{
			data: data,
		},
	}
	ret.IncRef()
	function.GlobalValueCache.Set(key, ret)
	return ret
}

// Period

type fPeriod[T float32 | float64] struct {
	fieldbase[T]

	value T
}

func (this fPeriod[T]) Get(index int) T {
	return this.value
}

func (this *fPeriod[T]) GetCacheKey() string {
	return "period" + fieldGetCacheKey(this.data)
}

func PERIOD[T float32 | float64](data RVectorReader[T]) *fPeriod[T] {
	key := "period" + fieldGetCacheKey(data)
	if v := function.GlobalValueCache.Get(key); v != nil {
		return v.(*fPeriod[T])
	}
	ret := &fPeriod[T]{
		fieldbase: fieldbase[T]{
			data: data,
		},
		value: T(GetPeriodIndex(data.Period())),
	}
	ret.IncRef()
	function.GlobalValueCache.Set(key, ret)
	return ret
}

// ISLASTBAR

type islastbar[T float32 | float64] struct {
	fieldbase[T]
}

func (this islastbar[T]) Get(index int) T {
	if this.data.Len()-1 == index {
		return 1
	}
	return 0
}

func ISLASTBAR[T float32 | float64](data RVectorReader[T]) *islastbar[T] {
	ret := &islastbar[T]{
		fieldbase[T]{
			data: data,
		},
	}
	return ret
}
