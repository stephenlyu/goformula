package nativeformulas

import (
	. "github.com/stephenlyu/goformula/formulalibrary/base/formula"
	"github.com/stephenlyu/goformula/function"
	stockfunc "github.com/stephenlyu/goformula/stockfunc/function"
	"github.com/stephenlyu/tds/period"
)

type BaseNativeFormula struct {
	*FormulaMetaImpl
	RefValues__   []function.Value
	DrawActions__ []DrawAction
	Data__        stockfunc.RVectorReader
}

func (this *BaseNativeFormula) Period() period.Period {
	return this.Data__.Period()
}

func (this *BaseNativeFormula) Len() int {
	return this.Data__.Len()
}

func (this BaseNativeFormula) Get(index int) []float64 {
	ret := make([]float64, len(this.RefValues__))
	for i, v := range this.RefValues__ {
		ret[i] = v.Get(index)
	}

	return ret
}

func (this BaseNativeFormula) Ref(offset int) []float64 {
	index := this.Data__.Len() - 1 - offset
	return this.Get(index)
}

func (this BaseNativeFormula) GetVarValue(varName string) function.Value {
	for i, v := range this.FormulaMetaImpl.Vars {
		if varName == v {
			return this.RefValues__[i]
		}
	}
	return nil
}

func (this *BaseNativeFormula) Destroy() {
}

func (this *BaseNativeFormula) Name() string {
	return this.GetName()
}

func (this *BaseNativeFormula) DrawActions() []DrawAction {
	return this.DrawActions__
}

func (this *BaseNativeFormula) DumpState() {
}
