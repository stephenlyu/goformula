package nativefactory

import (
	"github.com/stephenlyu/goformula/formulalibrary/base/factory"
	"github.com/stephenlyu/goformula/formulalibrary/base/formula"
	"github.com/stephenlyu/goformula/formulalibrary/native"
)

type nativeFormulaCreatorFactory[T float32 | float64] struct {
	provider      native.FormulaProvider[T]
	nativeFormula *native.NativeFormula[T]
}

func NewNativeFormulaCreatorFactory[T float32 | float64](provider native.FormulaProvider[T], nativeFormula *native.NativeFormula[T]) (error, factory.FormulaCreatorFactory[T]) {
	return nil, &nativeFormulaCreatorFactory[T]{
		provider:      provider,
		nativeFormula: nativeFormula,
	}
}
func (this *nativeFormulaCreatorFactory[T]) GetMeta() formula.FormulaMeta {
	return this.nativeFormula.Meta
}

func (this *nativeFormulaCreatorFactory[T]) CreateFormulaCreator(args []float64) factory.FormulaCreator[T] {
	if args == nil {
		args = this.nativeFormula.Meta.DefaultArgs()
	}

	return &nativeFormulaCreator[T]{
		provider: this.provider,
		factory:  this,
		args:     args,
	}
}
