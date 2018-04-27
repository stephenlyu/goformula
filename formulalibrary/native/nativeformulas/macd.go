package nativeformulas

import (
	stockfunc "github.com/stephenlyu/goformula/stockfunc/function"
	"github.com/stephenlyu/goformula/function"
	. "github.com/stephenlyu/goformula/formulalibrary/base/formula"
)

type macd struct {
	BaseNativeFormula

	// Data of all referenced period

	// Referenced Formulas

	// Vectors
	short           function.Value
	long            function.Value
	mid             function.Value
	close           function.Value
	ema_close_short function.Value
	ema_close_long  function.Value
	dif             function.Value
	dea             function.Value
	dif_sub_dea     function.Value
	const2          function.Value
	macd            function.Value
	enter_long      function.Value
	enter_short     function.Value
}

var (
	MACD_META = &FormulaMetaImpl{
		Name: "MACD",
		ArgNames: []string{"SHORT", "LONG", "MID"},
		ArgMeta: []Arg {
			Arg{2, 200, 12},
			Arg{2, 200, 26},
			Arg{2, 200, 9},
		},
		Flags: []int{0, 0, 0},
		Colors: []*Color{nil, nil, nil},
		LineThicks: []int{1, 1,1},
		LineStyles: []int{FORMULA_LINE_STYLE_SOLID, FORMULA_LINE_STYLE_SOLID, FORMULA_LINE_STYLE_SOLID},
		GraphTypes: []int{FORMULA_GRAPH_LINE, FORMULA_GRAPH_LINE, FORMULA_GRAPH_COLOR_STICK},
		Vars: []string{"DIF", "DEA", "MACD", "", ""},
	}
)

func MACD(data *stockfunc.RVector, args []float64) Formula {
	if len(args) != 3 {
		panic("bad args")
	}
	ret := &macd{
		BaseNativeFormula: BaseNativeFormula{
			FormulaMetaImpl: MACD_META,
			Data__: data,
		},
	}

	ret.short = function.Scalar(args[0])
	ret.long = function.Scalar(args[1])
	ret.mid = function.Scalar(args[2])
	ret.close = stockfunc.CLOSE(data)
	ret.const2 = function.Scalar(2)

	ret.ema_close_short = function.EMA(ret.close, ret.short)
	ret.ema_close_long = function.EMA(ret.close, ret.long)
	ret.dif = function.SUB(ret.ema_close_short, ret.ema_close_long)
	ret.dea = function.EMA(ret.dif, ret.mid)

	ret.dif_sub_dea = function.SUB(ret.dif, ret.dea)
	ret.macd = function.MUL(ret.dif_sub_dea, ret.const2)
	ret.enter_long = function.CROSS(ret.dif, ret.dea)
	ret.enter_short = function.CROSS(ret.dea, ret.dif)

	ret.RefValues__ = []function.Value {
		ret.dif,
		ret.dea,
		ret.macd,
		ret.enter_long,
		ret.enter_short,
	}
	return ret
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
