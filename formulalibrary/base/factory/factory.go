package factory

import (
	"github.com/stephenlyu/goformula/stockfunc/function"
	"github.com/stephenlyu/goformula/formulalibrary/base/formula"
)

type FormulaCreator interface {
	CreateFormula(data *function.RVector) (error, formula.Formula)
}

type FormulaCreatorFactory interface {
	CreateFormulaCreator(args []float64) FormulaCreator
	GetDefaultArgs() []float64
}
