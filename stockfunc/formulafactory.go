package stockfunc

import (
	"github.com/stephenlyu/goformula/stockfunc/function"
	"github.com/stephenlyu/goformula/stockfunc/formula"
	"github.com/stephenlyu/goformula/easylang"
	"github.com/stephenlyu/goformula/luafunc"
	"io/ioutil"
)

type FormulaFactory struct {
	debug bool
}

func NewFormulaFactory(debug bool) *FormulaFactory {
	return &FormulaFactory{debug}
}

func (this FormulaFactory) NewLuaFormula(luaFile string, data *function.RVector, args []float64) (error, formula.Formula) {
	return luafunc.NewFormula(luaFile, data, args)
}

func (this FormulaFactory) NewEasyLangFormula(easyLangFile string, data *function.RVector, args []float64) (error, formula.Formula) {
	err, code := easylang.CompileFile(easyLangFile)
	if err != nil {
		return err, nil
	}
	if this.debug {
		ioutil.WriteFile(easyLangFile + ".lua", []byte(code), 0666)
	}
	return luafunc.NewFormulaFromCode(code, data, args)
}
