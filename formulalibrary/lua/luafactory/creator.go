package luafactory

import (
	"github.com/stephenlyu/goformula/formulalibrary/base/formula"
	"github.com/stephenlyu/goformula/stockfunc/function"
	"github.com/stephenlyu/goformula/formulalibrary/lua/luaformula"
)

type luaFormulaCreator struct {
	factory *luaFormulaCreatorFactory
	args []float64
}

func (this *luaFormulaCreator) CreateFormula(data *function.RVector) (error, formula.Formula) {
	return luaformula.NewFormulaFromCode(this.factory.luaCode, this.factory.Meta, data, this.args)
}
