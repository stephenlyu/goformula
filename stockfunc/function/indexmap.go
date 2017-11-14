package function

import (
	"github.com/stephenlyu/tds/util"
	"github.com/stephenlyu/tds/period"
)


const DAY_MILLIS = 24 * 60 * 60 * 1000

// 比如在分钟线中引用日线数据时，srcData为分钟线数据，destData为日线数据。
// IndexMap负责将分钟线的索引映射为日线数据的索引，然后存取日线数据
type indexMap struct {
	srcData *RVector
	destData *RVector

	indexMap map[int]int
}

func IndexMap(srcData *RVector, destData *RVector) *indexMap {
	this := &indexMap{srcData: srcData, destData: destData}
	this.buildIndexMap()
	return this
}

func (this *indexMap) buildIndexMap() {
	needTrimDate := this.destData.period.Unit() != period.PERIOD_UNIT_MINUTE

	m := make(map[int]int)

	for i, j := 0, 0; i < this.srcData.Len() && j < this.destData.Len(); {
		srcDate := this.srcData.Get(i).GetUTCDate()
		if needTrimDate {
			srcDate = srcDate / DAY_MILLIS * DAY_MILLIS
		}

		destDate := this.srcData.Get(j).GetUTCDate()

		if srcDate <= destDate {
			m[i] = j
			i++
		} else {
			j++
		}
	}

	this.indexMap = m
}

func (this *indexMap) Get(index int) int {
	// 品种相同且周期相同时，直接从value中取值
	if this.srcData.code == this.destData.code && this.srcData.period.Eq(this.destData.period) {
		util.Assert(this.srcData.Len() == this.destData.Len(), "")
		return index
	}
	// 大周期引用小周期数据时，返回-1（不支持大周期引用小周期）
	if this.srcData.period.Gt(this.destData.period) {
		return -1
	}

	ret, ok := this.indexMap[index]
	if !ok {
		return -1
	}

	return ret
}

func (this *indexMap) UpdateLastValue() {
	i := this.srcData.Len() - 1
	j := this.destData.Len() - 1

	if i < 0 || j < 0 {
		return
	}
	needTrimDate := this.destData.period.Unit() != period.PERIOD_UNIT_MINUTE

	srcDate := this.srcData.Get(i).GetUTCDate()
	if needTrimDate {
		srcDate = srcDate / DAY_MILLIS * DAY_MILLIS
	}

	for j >= 0 {
		destDate := this.srcData.Get(j).GetUTCDate()
		if destDate < srcDate {
			break
		}
		j--
	}

	if j + 1 < this.destData.Len() {
		this.indexMap[i] = j + 1
	}
}
