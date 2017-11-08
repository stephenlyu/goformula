package native

import (
	"github.com/stephenlyu/goformula/stockfunc/function"
	"github.com/stephenlyu/goformula/formulalibrary/base/formula"
)

type CreateFunc func (meta *formula.FormulaMetaImpl, data *function.RVector, args []float64) formula.Formula

type NativeFormula struct {
	Creator CreateFunc
	Meta *formula.FormulaMetaImpl
}
