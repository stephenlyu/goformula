package easylangfactory

import (
	"github.com/stephenlyu/goformula/formulalibrary/base/factory"
	"github.com/stevedonovan/luar"
	"github.com/stephenlyu/goformula/formulalibrary/lua/luaformula"
	"github.com/stephenlyu/goformula/formulalibrary/base/formula"
	"github.com/stephenlyu/goformula/easylang"
	"io/ioutil"
)

type easylangFormulaCreatorFactory struct {
	luaCode string

	Meta *formula.FormulaMetaImpl
}

func NewEasyLangFormulaCreatorFactory(easyLangFile string, formulaManager formula.FormulaManager, debug bool) (error, factory.FormulaCreatorFactory) {
	// Compile Easy Lang File
	err, code := easylang.CompileFile(easyLangFile, formulaManager, true, debug)
	if err != nil {
		return err, nil
	}
	if debug {
		ioutil.WriteFile(easyLangFile + ".lua", []byte(code), 0666)
	}

	ret := &easylangFormulaCreatorFactory{luaCode: code}
	err = ret.init()
	if err != nil {
		return err, nil
	}

	return nil, ret
}

func (this *easylangFormulaCreatorFactory) init() error {
	L := luar.Init()

	luar.Register(L, "", luaformula.GetFunctionMap(luar.Map{}))

	err := L.DoString(this.luaCode)
	if err != nil {
		return err
	}

	L.GetGlobal("FormulaClass")
	meta := &formula.FormulaMetaImpl{}
	luaformula.GetMetaFromLuaState(L, meta)
	L.Pop(1)
	L.Close()
	this.Meta = meta

	return nil
}

func (this *easylangFormulaCreatorFactory) GetMeta() formula.FormulaMeta {
	return this.Meta
}

func (this *easylangFormulaCreatorFactory) CreateFormulaCreator(args []float64) factory.FormulaCreator {
	if args == nil {
		args = this.Meta.DefaultArgs()
	}

	return &easyLangFormulaCreator{
		factory: this,
		args: args,
	}
}
