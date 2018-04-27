package nativefactory

import (
	"github.com/stephenlyu/goformula/formulalibrary/base/formula"
	"github.com/stephenlyu/goformula/stockfunc/function"
)

type nativeFormulaCreator struct {
	factory *nativeFormulaCreatorFactory
	args []float64
}

func (this *nativeFormulaCreator) CreateFormula(data *function.RVector) (error, formula.Formula) {
	return nil, this.factory.nativeFormula.Creator(data, this.args)
}
