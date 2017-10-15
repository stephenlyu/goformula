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

var (
	vars = []string{"DIF", "DEA", "MACD"}
	colors = []*Color{nil, nil, nil}
	flags = []int{0, 0, 0}
	lineThicks = []int{1, 1,1}
	lineStyles = []int{FORMULA_LINE_STYLE_SOLID, FORMULA_LINE_STYLE_SOLID, FORMULA_LINE_STYLE_SOLID}
	graphTypes = []int{FORMULA_GRAPH_LINE, FORMULA_GRAPH_LINE, FORMULA_GRAPH_COLOR_STICK}
	args = [][]float64{
		[]float64{12, 2, 200},
		[]float64{26, 2, 200},
		[]float64{9, 2, 200},
	}
)

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

func (this *macd) Destroy() {
}


func (this *macd) Name() string {
	return "MACD"
}

// 输出变量

func (this *macd) VarCount() int {
	return len(vars)
}

func (this *macd) VarName(index int) string {
	if index < 0 || index >= len(vars) {
		return ""
	}
	return vars[index]
}

func (this *macd) NoDraw(index int) bool {
	if index < 0 || index >= len(flags) {
		return false
	}
	return (flags[index] & FORMULA_VAR_FLAG_NO_DRAW) != 0
}

func (this *macd) NoText(index int) bool {
	if index < 0 || index >= len(flags) {
		return false
	}
	return (flags[index] & FORMULA_VAR_FLAG_NO_TEXT) != 0
}

func (this *macd) DrawAbove(index int) bool {
	if index < 0 || index >= len(flags) {
		return false
	}
	return flags[index] & FORMULA_VAR_FLAG_DRAW_ABOVE != 0
}

func (this *macd) NoFrame(index int) bool {
	if index < 0 || index >= len(flags) {
		return false
	}
	return flags[index] & FORMULA_VAR_FLAG_NO_FRAME != 0
}


func (this *macd) Color(index int) *Color {
	if index < 0 || index >= len(colors) {
		return nil
	}
	return colors[index]
}

func (this *macd) LineThick(index int) int {
	if index < 0 || index >= len(lineThicks) {
		return 1
	}
	return lineThicks[index]
}

func (this *macd) LineStyle(index int) int {
	if index < 0 || index >= len(lineStyles) {
		return 1
	}
	return lineStyles[index]
}

func (this *macd) GraphType(index int) int {
	if index < 0 || index >= len(graphTypes) {
		return 1
	}
	return graphTypes[index]
}

// 公式参数

func (this *macd) ArgCount() int {
	return len(args)
}

func (this *macd) ArgRange(index int) (float64, float64) {
	if index < 0 || index >= len(args) {
		return 0, 0
	}
	return args[index][1], args[index][2]
}

func (this *macd) ArgDefault(index int) float64 {
	if index < 0 || index >= len(args) {
		return 0
	}
	return args[index][0]
}

func (this *macd) DrawActions() []DrawAction {
	return nil
}
