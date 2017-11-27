package function

import (
	"github.com/stephenlyu/goformula/function"
	"github.com/stephenlyu/tds/util"
	"errors"
)

// 将周期为data.Period生成的指标映射到destData.Period所表示的周期上
// 实现function.Value接口
type crossValue struct {
	value function.Value			// 跨周期的数据
	indexMap *indexMap
}

func CrossValue(value function.Value, indexMap *indexMap) function.Value {
	util.Assert(!value.IsScalar(), "!value.IsScalar()")
	return &crossValue{value: value, indexMap: indexMap}
}

func (this *crossValue) IsScalar() bool {
	return false
}

func (this *crossValue) Len() int {
	return this.indexMap.srcData.Len()
}

func (this *crossValue) Get(index int) float64 {
	index = this.indexMap.Get(index)
	if index < 0 {
		return function.NaN
	}

	return this.value.Get(index)
}

func (this *crossValue) Set(index int, v float64) {
	panic(errors.New("Set not supported!"))
}

func (this *crossValue) UpdateLastValue() {
	this.value.UpdateLastValue()
}

func (this *crossValue) Append(v float64) {
	panic(errors.New("Append not supported!"))
}
