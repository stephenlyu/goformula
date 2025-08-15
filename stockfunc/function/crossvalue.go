package function

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/stephenlyu/goformula/function"
	"github.com/stephenlyu/tds/util"
)

// 将周期为data.Period生成的指标映射到destData.Period所表示的周期上
// 实现function.Value接口
type crossValue[T float32 | float64] struct {
	value     function.Value[T] // 跨周期的数据
	indexMap  *IndexMap[T]
	refCount1 uint32
}

func (this *crossValue[T]) IncRef() {
	this.value.IncRef()
	this.refCount1++
}

func (this *crossValue[T]) DecRef() {
	this.value.DecRef()
	this.refCount1--
	if this.refCount1 <= 0 {
		function.GlobalValueCache.Del(this.GetCacheKey())
	}
}

func crossValueGetCacheKey[T float32 | float64](value function.Value[T], indexMap *IndexMap[T]) string {
	var parts []string
	va := reflect.ValueOf(value)
	ptr := va.Pointer()
	parts = append(parts, fmt.Sprintf("%x", ptr))
	va = reflect.ValueOf(indexMap)
	ptr = va.Pointer()
	parts = append(parts, fmt.Sprintf("%x", ptr))
	return strings.Join(parts, "_")
}

func (this *crossValue[T]) GetCacheKey() string {
	return "crossValue_" + crossValueGetCacheKey(this.value, this.indexMap)
}

func CrossValue[T float32 | float64](value function.Value[T], indexMap *IndexMap[T]) function.Value[T] {
	util.Assert(!value.IsScalar(), "!value.IsScalar()")
	key := "crossValue_" + crossValueGetCacheKey(value, indexMap)
	if v := function.GlobalValueCache.Get(key); v != nil {
		return v.(function.Value[T])
	}
	ret := &crossValue[T]{value: value, indexMap: indexMap}
	ret.IncRef()
	function.GlobalValueCache.Set(key, ret)
	return ret
}

func (this *crossValue[T]) IsScalar() bool {
	return false
}

func (this *crossValue[T]) Len() int {
	return this.indexMap.srcData.Len()
}

func (this *crossValue[T]) Get(index int) T {
	index = this.indexMap.Get(index)
	if index < 0 {
		return T(function.NaN)
	}

	return this.value.Get(index)
}

func (this *crossValue[T]) Set(index int, v T) {
	panic(errors.New("Set not supported!\n"))
}

func (this *crossValue[T]) UpdateLastValue() {
	this.value.UpdateLastValue()
}

func (this *crossValue[T]) LTrim(count int) {
	this.value.LTrim(count)
}

func (this *crossValue[T]) Append(v T) {
	panic(errors.New("Append not supported!\n"))
}
