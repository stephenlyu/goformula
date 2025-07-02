package function

import (
	"errors"

	"github.com/stephenlyu/goformula/function"
	"github.com/stephenlyu/tds/util"
)

// 将周期为data.Period生成的指标映射到destData.Period所表示的周期上
// 实现function.Value接口
type crossValue[T float32 | float64] struct {
	value    function.Value[T] // 跨周期的数据
	indexMap *IndexMap[T]
}

func CrossValue[T float32 | float64](value function.Value[T], indexMap *IndexMap[T]) function.Value[T] {
	util.Assert(!value.IsScalar(), "!value.IsScalar()")
	return &crossValue[T]{value: value, indexMap: indexMap}
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

func (this *crossValue[T]) Append(v T) {
	panic(errors.New("Append not supported!\n"))
}
