package function

import "math"

// TODO: 处理未来函数。所有的函数需要增加一个withFutureData参数，如果withFutureData为true， 则UpdateLastValue时，需要更新整个数据

// DRAWLINE
// TODO: 不支持延长线

type drawline[T float32 | float64] struct {
	funcbase[T]
	Cond1  Value[T]
	Price1 Value[T]
	Cond2  Value[T]
	Price2 Value[T]
	Expand Value[T]

	lastCond1Index   int
	lastCond2Indices []int
}

func (this *drawline[T]) BuildValueAt(index int) {
	cond1 := IsTrue(this.Cond1.Get(index))
	cond2 := IsTrue(this.Cond2.Get(index))

	if cond1 {
		this.lastCond1Index = index
		this.lastCond2Indices = []int{}
		this.Values[index] = T(NaN)
	} else if cond2 {
		if this.lastCond1Index >= 0 {
			Interpolate(this.Values, this.lastCond1Index, index, this.Price1.Get(this.lastCond1Index), this.Price2.Get(index))
			this.lastCond2Indices = append(this.lastCond2Indices, index)
		} else {
			this.Values[index] = T(NaN)
		}
	} else {
		if this.lastCond1Index >= 0 {
			n := len(this.lastCond2Indices)
			if n > 0 {
				prevCond2Index := this.lastCond2Indices[n-1]
				if prevCond2Index == index {
					// 信号出现了又消失的情况
					this.lastCond2Indices = this.lastCond2Indices[:n-1]
					n--

					var from int
					if len(this.lastCond2Indices) > 0 {
						// 一个起点，多个终点时，仅将上一个终点之后的数值修改成NaN
						from = this.lastCond2Indices[n-1] + 1
					} else {
						// 修改为未找到终点时的情况，即起点后的所有index修改为NaN
						from = this.lastCond1Index
					}
					for i := from; i < index; i++ {
						this.Values[i] = T(NaN)
					}
				} else {
					this.Values[index] = T(NaN)
				}
			} else {
				this.Values[index] = T(NaN)
			}
		} else {
			this.Values[index] = T(NaN)
		}
	}
}

func (this *drawline[T]) UpdateLastValue() {
	if this.Data().Len() < this.Len() {
		return
	}
	if this.Data().Len() == this.Len() {
		this.BuildValueAt(this.Data().Len() - 1)
	} else {
		start := this.Len()
		for i := this.Len(); i < this.Data().Len(); i++ {
			this.Append(0)
		}

		for i := start; i < this.Len(); i++ {
			this.BuildValueAt(i)
		}
	}
}

func (this *drawline[T]) initValues() {
	for i := 0; i < this.Data().Len(); i++ {
		this.BuildValueAt(i)
	}
}

func DRAWLINE[T float32 | float64](cond1 Value[T], price1 Value[T], cond2 Value[T], price2 Value[T], expand Value[T]) Value[T] {
	if cond1.IsScalar() && price1.IsScalar() && cond2.IsScalar() && price2.IsScalar() && expand.IsScalar() {
		return Scalar(T(math.NaN()))
	}

	if expand == nil {
		expand = Scalar(T(0))
	}

	var length int
	switch {
	case !cond1.IsScalar():
		length = cond1.Len()
	case !price1.IsScalar():
		length = price1.Len()
	case !cond2.IsScalar():
		length = cond2.Len()
	case !price2.IsScalar():
		length = price2.Len()
	default:
		length = expand.Len()
	}

	ret := &drawline[T]{
		funcbase: funcbase[T]{
			data: cond1,
		},
		Cond1:  cond1,
		Price1: price1,
		Cond2:  cond2,
		Price2: price2,
		Expand: expand,

		lastCond1Index: -1,
	}
	ret.Values = make([]T, length)
	ret.initValues()
	return ret
}

// PLOYLINE

type ployline[T float32 | float64] struct {
	funcbase[T]
	Cond  Value[T]
	Price Value[T]

	lastSegFromIndex int // 最后一条线段的起始位置
	lastSegToIndex   int // 最后一条线段结束位置
}

func (this *ployline[T]) BuildValueAt(index int) {
	yes := IsTrue(this.Cond.Get(index))

	if yes {
		this.Values[index] = this.Price.Get(index)

		if this.lastSegToIndex >= 0 {
			Assert(index >= this.lastSegToIndex, "")
			if index > this.lastSegToIndex {
				Interpolate(this.Values, this.lastSegToIndex, index, this.Price.Get(this.lastSegToIndex), this.Price.Get(index))
				this.lastSegFromIndex = this.lastSegToIndex
				this.lastSegToIndex = index
			} else {
				Interpolate(this.Values, this.lastSegFromIndex, index, this.Price.Get(this.lastSegFromIndex), this.Price.Get(index))
			}
		} else {
			this.Values[index] = this.Price.Get(index)
			this.lastSegFromIndex = index
			this.lastSegToIndex = index
		}
	} else {
		if this.lastSegToIndex == index {
			// 信号忽闪时，需要恢复前面的值
			if this.lastSegFromIndex >= 0 {
				for i := this.lastSegFromIndex + 1; i < index; i++ {
					this.Values[i] = T(NaN)
				}
			}
		} else {
			this.Values[index] = T(NaN)
		}
	}
}

func (this *ployline[T]) UpdateLastValue() {
	if this.Data().Len() < this.Len() {
		return
	}
	if this.Data().Len() == this.Len() {
		this.BuildValueAt(this.Data().Len() - 1)
	} else {
		start := this.Len()
		for i := this.Len(); i < this.Data().Len(); i++ {
			this.Append(0)
		}

		for i := start; i < this.Len(); i++ {
			this.BuildValueAt(i)
		}
	}
}

func (this *ployline[T]) initValues() {
	for i := 0; i < this.Data().Len(); i++ {
		this.BuildValueAt(i)
	}
}

func PLOYLINE[T float32 | float64](cond Value[T], price Value[T]) Value[T] {
	if cond.IsScalar() && price.IsScalar() {
		return Scalar(T(math.NaN()))
	}

	var length int
	if !cond.IsScalar() {
		length = cond.Len()
	} else {
		length = price.Len()
	}

	ret := &ployline[T]{
		funcbase: funcbase[T]{
			data: cond,
		},
		Cond:             cond,
		Price:            price,
		lastSegFromIndex: -1,
		lastSegToIndex:   -1,
	}
	ret.Values = make([]T, length)
	ret.initValues()
	return ret
}
