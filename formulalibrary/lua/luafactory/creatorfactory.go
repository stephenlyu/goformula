package luafactory

import (
	"github.com/stephenlyu/goformula/formulalibrary/base/factory"
	"github.com/stevedonovan/luar"
	"github.com/stephenlyu/goformula/formulalibrary/lua/luaformula"
	"github.com/stephenlyu/goformula/formulalibrary/base/formula"
	"io/ioutil"
)

type luaFormulaCreatorFactory struct {
	luaCode string

	Meta *formula.FormulaMetaImpl
}

func NewLuaFormulaCreatorFactory(luaFile string) (error, factory.FormulaCreatorFactory) {
	bytes, err := ioutil.ReadFile(luaFile)
	if err != nil {
		return err, nil
	}
	luaCode := string(bytes)

	ret := &luaFormulaCreatorFactory{luaCode:luaCode}
	err = ret.init()
	if err != nil {
		return err, nil
	}

	return nil, ret
}

func (this *luaFormulaCreatorFactory) init() error {
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
