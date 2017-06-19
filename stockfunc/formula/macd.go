package formula

import (
	stockfunc "github.com/stephenlyu/goformula/stockfunc/function"
	"github.com/stephenlyu/goformula/function"
)

type macd struct {
	data *stockfunc.RVector
	short function.Value
	long function.Value
	mid function.Value
	close function.Value
	ema_close_short function.Value
	ema_close_long function.Value
	dif function.Value
	dea function.Value
	dif_sub_dea function.Value
	const2 function.Value
	macd function.Value
	enter_long function.Value
	enter_short function.Value
}

func MACD(data *stockfunc.RVector, short, long, mid function.Value) *macd {
	if short == nil {
		short = function.Scalar(12)
	}
	if long == nil {
		long = function.Scalar(26)
	}
	if mid == nil {
		mid = function.Scalar(9)
	}

	ret := &macd{
		data: data,
		short: short,
		long: long,
		mid: mid,
		close: stockfunc.CLOSE(data),
		const2: function.Scalar(2)}

	ret.ema_close_short = function.EMA(ret.close, short)
	ret.ema_close_long = function.EMA(ret.close, long)
	ret.dif = function.SUB(ret.ema_close_short, ret.ema_close_long)
	ret.dea = function.EMA(ret.dif, mid)

	ret.dif_sub_dea = function.SUB(ret.dif, ret.dea)
	ret.macd = function.MUL(ret.dif_sub_dea, ret.const2)
	ret.enter_long = function.CROSS(ret.dif, ret.dea)
	ret.enter_short = function.CROSS(ret.dea, ret.dif)
	return ret
}

func (this macd) Len() int {
	return this.data.Len()
}

func (this *macd) UpdateLastValue() {
	this.short.UpdateLastValue()
	this.long.UpdateLastValue()
	this.mid.UpdateLastValue()
	this.close.UpdateLastValue()
	this.ema_close_short.UpdateLastValue()
	this.ema_close_long.UpdateLastValue()
	this.dif.UpdateLastValue()
	this.dea.UpdateLastValue()
	this.dif_sub_dea.UpdateLastValue()
	this.macd.UpdateLastValue()
	this.enter_long.UpdateLastValue()
	this.enter_short.UpdateLastValue()
}

func (this macd) Get(index int) []float64 {
	return []float64 {
		this.dif.Get(index),
		this.dea.Get(index),
		this.macd.Get(index),
		this.enter_long.Get(index),
		this.enter_short.Get(index),
	}
}

func (this macd) Ref(offset int) []float64 {
	index := this.data.Len() - 1 - offset
	return this.Get(index)
}
