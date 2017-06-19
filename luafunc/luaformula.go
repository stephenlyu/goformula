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
}

func NewFormula(luaFile string, data *stockfunc.RVector, args []float64) (error, *LuaFormula) {
	L := luar.Init()

	luar.Register(L, "", GetFunctionMap(luar.Map{}))

	err := L.DoFile(luaFile)
	if err != nil {
		return err, nil
	}

	L.GetGlobal("FormulaClass")
	L.GetField(-1, "new")
	L.PushValue(-2)
	luar.GoToLuaProxy(L, data)
	for _, arg := range args {
		luar.GoToLuaProxy(L, function.Scalar(arg))
	}
	L.Call(2 + len(args), 1)
	if L.IsNil(-1) {
		return errors.New("Create formula fail"), nil
	}

	L.Remove(1)

	return nil, &LuaFormula{L: L}
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

func (this *LuaFormula) UpdateLateValue() {
	this.L.GetField(-1, "updateLastValue")
	this.L.PushValue(-2)
	this.L.Call(1, 0)
}

func (this *LuaFormula) Get(index int) []float64 {
	this.L.GetField(-1, "Get")
	this.L.PushValue(-2)
	this.L.PushInteger(int64(index))
	this.L.Call(2, 1)

	var ret []float64
	luar.LuaToGo(this.L, -1, &ret)
	this.L.Pop(1)
	return ret
}

func (this *LuaFormula) Ref(offset int) []float64 {
	index := this.Len() - 1 - offset
	return this.Get(index)
}
