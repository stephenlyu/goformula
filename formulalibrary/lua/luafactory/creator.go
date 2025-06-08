package luafactory

import (
	"github.com/stephenlyu/goformula/formulalibrary/base/formula"
	"github.com/stephenlyu/goformula/formulalibrary/lua/luaformula"
	"github.com/stephenlyu/goformula/stockfunc/function"
)

type luaFormulaCreator struct {
	factory *luaFormulaCreatorFactory
	args    []float64
}

func (this *luaFormulaCreator) CreateFormula(data function.RVectorReader) (error, formula.Formula) {
	return luaformula.NewFormulaFromState(this.factory.L, this.factory.Meta, data, this.args)
}
