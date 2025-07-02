package nativefactory

import (
	"github.com/stephenlyu/goformula/formulalibrary/base/formula"
	"github.com/stephenlyu/goformula/formulalibrary/native"
	"github.com/stephenlyu/goformula/stockfunc/function"
)

type nativeFormulaCreator[T float32 | float64] struct {
	provider native.FormulaProvider[T]
	factory  *nativeFormulaCreatorFactory[T]
	args     []float64
}

func (this *nativeFormulaCreator[T]) CreateFormula(data function.RVectorReader[T]) (error, formula.Formula[T]) {
	return nil, this.factory.nativeFormula.Creator(this.provider, data, this.args)
}
