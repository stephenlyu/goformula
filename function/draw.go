package function

// TODO: 处理未来函数。所有的函数需要增加一个withFutureData参数，如果withFutureData为true， 则UpdateLastValue时，需要更新整个数据

// DRAWLINE
// TODO: 不支持延长线

type drawline struct {
	funcbase
	Cond1            Value
	Price1           Value
	Cond2            Value
	Price2           Value
	Expand           Value

	lastCond1Index   int
	lastCond2Indices []int
}

func (this *drawline) BuildValueAt(index int) {
	cond1 := isTrue(this.Cond1.Get(index))
	cond2 := isTrue(this.Cond2.Get(index))

	if cond1 {
		this.lastCond1Index = index
		this.lastCond2Indices = []int{}
		this.Values[index] = NaN
	} else if cond2 {
		if this.lastCond1Index >= 0 {
			Interpolate(this.Values, this.lastCond1Index, index, this.Price1.Get(this.lastCond1Index), this.Price2.Get(index))
			this.lastCond2Indices = append(this.lastCond2Indices, index)
		} else {
			this.Values[index] = NaN
		}
	} else {
		if this.lastCond1Index >= 0 {
			n := len(this.lastCond2Indices)
			if n > 0 {
				prevCond2Index := this.lastCond2Indices[n - 1]
				if prevCond2Index == index {
					// 信号出现了又消失的情况
					this.lastCond2Indices = this.lastCond2Indices[:n-1]
					n--

					var from int
					if len(this.lastCond2Indices) > 0 {
						// 一个起点，多个终点时，仅将上一个终点之后的数值修改成NaN
						from = this.lastCond2Indices[n - 1] + 1
					} else {
						// 修改为未找到终点时的情况，即起点后的所有index修改为NaN
						from = this.lastCond1Index
					}
					for i := from; i < index; i++ {
						this.Values[i] = NaN
					}
				} else {
					this.Values[index] = NaN
				}
			} else {
				this.Values[index] = NaN
			}
		} else {
			this.Values[index] = NaN
		}
	}
}

func (this *drawline) UpdateLastValue() {
	if this.Data().Len() < this.Len() {
		return
	}
	if this.Len() != this.Data().Len() {
		Assert(this.Len() + 1 == this.Data().Len(), "")
		this.Append(0)
	}
	this.BuildValueAt(this.Data().Len() - 1)
}

func (this *drawline) initValues() {
	for i := 0; i < this.Data().Len(); i++ {
		this.BuildValueAt(i)
	}
}

func DRAWLINE(cond1 Value, price1 Value, cond2 Value, price2 Value, expand Value) *drawline {
	if expand == nil {
		expand = Scalar(0)
	}

	ret := &drawline{
		funcbase: funcbase {
			data: cond1,
		},
		Cond1: cond1,
		Price1: price1,
		Cond2: cond2,
		Price2: price2,
		Expand: expand,

		lastCond1Index: -1,
	}
	ret.Values = make([]float64, cond1.Len())
	ret.initValues()
	return ret
}

// PLOYLINE

type ployline struct {
	funcbase
	Cond Value
	Price Value

	lastSegFromIndex int		// 最后一条线段的起始位置
	lastSegToIndex int			// 最后一条线段结束位置
}

func (this *ployline) BuildValueAt(index int) {
	yes := isTrue(this.Cond.Get(index))

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
					this.Values[i] = NaN
				}
			}
		} else {
			this.Values[index] = NaN
		}
	}
}

func (this *ployline) UpdateLastValue() {
	if this.Data().Len() < this.Len() {
		return
	}
	if this.Len() != this.Data().Len() {
		Assert(this.Len() + 1 == this.Data().Len(), "")
		this.Append(0)
	}
	this.BuildValueAt(this.Data().Len() - 1)
}

func (this *ployline) initValues() {
	for i := 0; i < this.Data().Len(); i++ {
		this.BuildValueAt(i)
	}
}

func PLOYLINE(cond Value, price Value) *ployline {
	ret := &ployline{
		funcbase: funcbase {
			data: cond,
		},
		Cond: cond,
		Price: price,
		lastSegFromIndex: -1,
		lastSegToIndex: -1,
	}
	ret.Values = make([]float64, cond.Len())
	ret.initValues()
	return ret
}
