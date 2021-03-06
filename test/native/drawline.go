//
// GENERATED BY EASYLANG COMPILER.
// !!!! DON'T MODIFY IT!!!!!!
//

package native

import (
	. "github.com/stephenlyu/goformula/stockfunc/function"
	. "github.com/stephenlyu/goformula/function"
	. "github.com/stephenlyu/goformula/formulalibrary/base/formula"
	. "github.com/stephenlyu/goformula/formulalibrary/native/nativeformulas"
)

type drawline struct {
	BaseNativeFormula

	// Data of all referenced period


	// Referenced Formulas


	// Vectors
    var1 Value
    a Value
    const1 Value
    var2 Value
    b Value
    var3 Value
    bb Value
    var4 Value
    c Value
    var5 Value
    d Value
    var6 Value
    dd Value
    const2 Value
    var7 Value
    __anonymous_0 Value
}

var (
	DRAWLINE_META = &FormulaMetaImpl{
		Name: "DRAWLINE",
		ArgNames: []string{},
		ArgMeta: []Arg {

		},
		Flags: []int{0x00000000, 0x00000000, 0x00000000, 0x00000000, 0x00000000, 0x00000000, 0x00000001},
		Colors: []*Color{{Red:-1, Green:-1, Blue:-1}, {Red:-1, Green:-1, Blue:-1}, {Red:-1, Green:-1, Blue:-1}, {Red:-1, Green:-1, Blue:-1}, {Red:-1, Green:-1, Blue:-1}, {Red:-1, Green:-1, Blue:-1}, {Red:0, Green:255, Blue:0}},
		LineThicks: []int{1, 1, 1, 1, 1, 1, 8},
		LineStyles: []int{0, 0, 0, 0, 0, 0, 0},
		GraphTypes: []int{1, 1, 1, 1, 1, 1, 0},
		Vars: []string{"A", "B", "BB", "C", "D", "DD", ""},
	}
)

func NewDRAWLINE(data *RVector, args []float64) Formula {
	o := &drawline{
		BaseNativeFormula: BaseNativeFormula{
			FormulaMetaImpl: DRAWLINE_META,
			Data__: data,
		},
	}

	// Data of all referenced period


	// Referenced Formulas


	// Vectors
    o.var1 = HIGH(o.Data__)
    o.a = o.var1
    o.const1 = Scalar(5.000000)
    o.var2 = HHV(o.var1, o.const1)
    o.b = o.var2
    o.var3 = GE(o.var1, o.var2)
    o.bb = o.var3
    o.var4 = LOW(o.Data__)
    o.c = o.var4
    o.var5 = LLV(o.var4, o.const1)
    o.d = o.var5
    o.var6 = LE(o.var4, o.var5)
    o.dd = o.var6
    o.const2 = Scalar(1.000000)
    o.var7 = DRAWLINE(o.var3, o.var1, o.var6, o.var4, o.const2)
    o.__anonymous_0 = o.var7

	// Actions

    o.DrawActions__ = []DrawAction{
        &DrawLineAction{ActionType:1, Cond1:o.var3, Price1:o.var1, Cond2:o.var6, Price2:o.var4, Expand:1, NoDraw:1, Color:&Color{Red:0, Green:255, Blue:0}, LineThick:8, VarIndex:6},
    }

	o.RefValues__ = []Value {o.a, o.b, o.bb, o.c, o.d, o.dd, o.__anonymous_0}
	return o
}

func (this *drawline) UpdateLastValue() {
    this.var1.UpdateLastValue()
    this.var2.UpdateLastValue()
    this.var3.UpdateLastValue()
    this.var4.UpdateLastValue()
    this.var5.UpdateLastValue()
    this.var6.UpdateLastValue()
    this.var7.UpdateLastValue()
}

func init() {
	RegisterNativeFormula(NewDRAWLINE, DRAWLINE_META)
}

	