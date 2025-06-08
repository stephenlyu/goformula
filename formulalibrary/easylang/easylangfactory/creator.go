package easylangfactory

import (
	"github.com/stephenlyu/goformula/formulalibrary/base/formula"
	"github.com/stephenlyu/goformula/formulalibrary/lua/luaformula"
	"github.com/stephenlyu/goformula/stockfunc/function"
)

type easyLangFormulaCreator struct {
	factory *easylangFormulaCreatorFactory
	args    []float64
}

func (this *easyLangFormulaCreator) CreateFormula(data function.RVectorReader) (error, formula.Formula) {
	return luaformula.NewFormulaFromState(this.factory.L, this.factory.Meta, data, this.args)
}
