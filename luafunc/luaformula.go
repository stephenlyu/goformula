package luafunc

import (
	stockfunc "github.com/stephenlyu/goformula/stockfunc/function"
	"github.com/stevedonovan/luar"
	"github.com/aarzilli/golua/lua"
	"github.com/stephenlyu/goformula/function"
	"errors"
)

type LuaFormula struct {
	L *lua.State

	refValues []function.Value
}

func newFormulaByLuaState(L *lua.State, data *stockfunc.RVector, args []float64) (error, *LuaFormula) {
	L.GetGlobal("FormulaClass")
	L.GetField(-1, "new")
	L.PushValue(-2)
	luar.GoToLuaProxy(L, data)
	for _, arg := range args {
		luar.GoToLua(L, arg)
	}
	L.Call(2 + len(args), 1)
	if L.IsNil(-1) {
		return errors.New("Create formula fail"), nil
	}

	L.GetField(-1, "ref_values")
	var values []function.Value
	luar.LuaToGo(L, -1, &values)
	L.Pop(1)

	L.Remove(1)

	return nil, &LuaFormula{L: L, refValues: values}
}

func NewFormula(luaFile string, data *stockfunc.RVector, args []float64) (error, *LuaFormula) {
	L := luar.Init()

	luar.Register(L, "", GetFunctionMap(luar.Map{}))

	err := L.DoFile(luaFile)
	if err != nil {
		return err, nil
	}

	return newFormulaByLuaState(L, data, args)
}

func NewFormulaFromCode(luaCode string, data *stockfunc.RVector, args []float64) (error, *LuaFormula) {
	L := luar.Init()

	luar.Register(L, "", GetFunctionMap(luar.Map{}))

	err := L.DoString(luaCode)
	if err != nil {
		return err, nil
	}

	return newFormulaByLuaState(L, data, args)
}

func (this *LuaFormula) Destroy() {
	this.L.Close()
}

func (this *LuaFormula) Len() int {
	this.L.GetField(-1, "Len")
	this.L.PushValue(-2)
	this.L.Call(1, 1)
	ret := this.L.ToInteger(-1)
	this.L.Pop(1)
	return ret
}

func (this *LuaFormula) UpdateLastValue() {
	this.L.GetField(-1, "updateLastValue")
	this.L.PushValue(-2)
	this.L.Call(1, 0)
}

func (this *LuaFormula) Get(index int) []float64 {
	ret := make([]float64, len(this.refValues))

	for i, refValue := range this.refValues {
		ret[i] = refValue.Get(index)
	}

	return ret
}

func (this *LuaFormula) Ref(offset int) []float64 {
	index := this.Len() - 1 - offset
	return this.Get(index)
}
