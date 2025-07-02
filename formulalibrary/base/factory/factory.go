package factory

import (
	"github.com/stephenlyu/goformula/formulalibrary/base/formula"
	"github.com/stephenlyu/goformula/stockfunc/function"
)

type FormulaCreator[T float32 | float64] interface {
	CreateFormula(data function.RVectorReader[T]) (error, formula.Formula[T])
}

type FormulaCreatorFactory[T float32 | float64] interface {
	CreateFormulaCreator(args []float64) FormulaCreator[T]
	GetMeta() formula.FormulaMeta
}
