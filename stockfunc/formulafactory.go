package stockfunc

import (
	"github.com/stephenlyu/goformula/stockfunc/function"
	"github.com/stephenlyu/goformula/stockfunc/formula"
	"github.com/stephenlyu/goformula/easylang"
	"github.com/stephenlyu/goformula/luafunc"
)

type FormulaFactory struct {
}

func NewFormulaFactory() *FormulaFactory {
	return &FormulaFactory{}
}

func (this FormulaFactory) NewLuaFormula(luaFile string, data *function.RVector, args []float64) (error, formula.Formula) {
	return luafunc.NewFormula(luaFile, data, args)
}

func (this FormulaFactory) NewEasyLangFormula(easyLangFile string, data *function.RVector, args []float64) (error, formula.Formula) {
	err, code := easylang.CompileFile(easyLangFile)
	if err != nil {
		return err, nil
	}

	return luafunc.NewFormulaFromCode(code, data, args)
}
