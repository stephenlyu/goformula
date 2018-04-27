package luafactory

import (
	"github.com/stephenlyu/goformula/formulalibrary/base/factory"
	"github.com/aarzilli/golua/lua"
	"github.com/stevedonovan/luar"
	"github.com/stephenlyu/goformula/formulalibrary/lua/luaformula"
	"github.com/stephenlyu/goformula/formulalibrary/base/formula"
)

type luaFormulaCreatorFactory struct {
	luaFile string

	L *lua.State
	Meta *formula.FormulaMetaImpl
}

func NewLuaFormulaCreatorFactory(luaFile string) (error, factory.FormulaCreatorFactory) {
	L := luar.Init()

	luar.Register(L, "", luaformula.GetFunctionMap(luar.Map{}))

	err := L.DoFile(luaFile)
	if err != nil {
		return err, nil
	}

	L.GetGlobal("FormulaClass")
	meta := &formula.FormulaMetaImpl{}
	luaformula.GetMetaFromLuaState(L, meta)
	L.Pop(1)

	return nil, &luaFormulaCreatorFactory{luaFile:luaFile, L: L, Meta: meta}
}
func (this *luaFormulaCreatorFactory) GetMeta() formula.FormulaMeta {
	return this.Meta
}

func (this *luaFormulaCreatorFactory) CreateFormulaCreator(args []float64) factory.FormulaCreator {
	if args == nil {
		args = this.Meta.DefaultArgs()
	}

	return &luaFormulaCreator{
		factory: this,
		args: args,
	}
}
