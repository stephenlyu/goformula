package native

import (
	"github.com/stephenlyu/goformula/formulalibrary/base/formula"
	"github.com/stephenlyu/goformula/stockfunc/function"
)

type CreateFunc func(data function.RVectorReader, args []float64) formula.Formula

type NativeFormula struct {
	Creator CreateFunc
	Meta    *formula.FormulaMetaImpl
}
