package factory

import (
	"github.com/stephenlyu/goformula/formulalibrary/base/formula"
	"github.com/stephenlyu/goformula/stockfunc/function"
)

type FormulaCreator interface {
	CreateFormula(data function.RVectorReader) (error, formula.Formula)
}

type FormulaCreatorFactory interface {
	CreateFormulaCreator(args []float64) FormulaCreator
	GetMeta() formula.FormulaMeta
}
