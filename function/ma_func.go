package function

import "github.com/chanxuehong/util/math"

type ma struct {
	funcbase
	N Value
}

func (this ma) BuildValueAt(index int) float64 {
	n := int(this.N.Get(index))
	if index < n -1 {
		return 0
	}

	end := index + 1
	start := int(math.Max(0, int64(end - n)))

	sum := 0.0
	for i := start; i< end; i++ {
		sum += this.data.Get(i)
	}

	return sum / float64(n)
}

func (this *ma) UpdateLastValue() {
	updateLastValue(this)
}

func MA(data Value, N Value) *ma {
	ret := &ma{
		funcbase: funcbase {
			data: data,
		},
		N: N,
	}
	ret.Values = make([]float64, data.Len())
	for i := 0; i < data.Len(); i++ {
		v := ret.BuildValueAt(i)
		ret.Set(i, v)
	}
	return ret
}
