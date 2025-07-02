package native

import (
	"github.com/stephenlyu/goformula/datalibrary"
	"github.com/stephenlyu/goformula/formulalibrary/base/formula"
	"github.com/stephenlyu/goformula/stockfunc/function"
)

type FormulaProvider[T float32 | float64] interface {
	GetDataLibrary() datalibrary.DataLibrary[T]
	NewFormula(name string, data function.RVectorReader[T]) formula.Formula[T]
}

type CreateFunc[T float32 | float64] func(
	provider FormulaProvider[T],
	data function.RVectorReader[T],
	args []float64) formula.Formula[T]

type NativeFormula[T float32 | float64] struct {
	Creator CreateFunc[T]
	Meta    *formula.FormulaMetaImpl
}
