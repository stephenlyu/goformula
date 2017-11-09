package easylangfactory

import (
	"github.com/stephenlyu/goformula/formulalibrary/base/factory"
	"github.com/aarzilli/golua/lua"
	"github.com/stevedonovan/luar"
	"github.com/stephenlyu/goformula/formulalibrary/lua/luaformula"
	"github.com/stephenlyu/goformula/formulalibrary/base/formula"
	"github.com/stephenlyu/goformula/easylang"
	"io/ioutil"
)

type easylangFormulaCreatorFactory struct {
	easyLangFile string

	L *lua.State
	Meta *formula.FormulaMetaImpl
}

func NewEasyLangFormulaCreatorFactory(easyLangFile string, formulaManager formula.FormulaManager, debug bool) (error, factory.FormulaCreatorFactory) {
	// Compile Easy Lang File
	err, code := easylang.CompileFile(easyLangFile, formulaManager)
	if err != nil {
		return err, nil
	}
	if debug {
		ioutil.WriteFile(easyLangFile + ".lua", []byte(code), 0666)
	}

	// Init Lua Engine
	L := luar.Init()
	luar.Register(L, "", luaformula.GetFunctionMap(luar.Map{}))

	err = L.DoString(code)
	if err != nil {
		return err, nil
	}

	L.GetGlobal("FormulaClass")
	meta := &formula.FormulaMetaImpl{}
	luaformula.GetMetaFromLuaState(L, meta)
	L.Pop(1)

	return nil, &easylangFormulaCreatorFactory{easyLangFile:easyLangFile, L: L, Meta: meta}
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
