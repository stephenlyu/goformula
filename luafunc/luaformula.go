package luafunc

import (
	stockfunc "github.com/stephenlyu/goformula/stockfunc/function"
	"github.com/stevedonovan/luar"
	"github.com/aarzilli/golua/lua"
	"github.com/stephenlyu/goformula/function"
	"errors"
	"github.com/stephenlyu/goformula/stockfunc/formula"
)

type Arg struct {
	Min float64
	Max float64
	Default float64
}

type LuaFormula struct {
	L *lua.State

	name string
	ArgNames []string
	Args []Arg
	NoDraws []int
	Colors []string
	LineThicks []int
	Vars []string

	refValues []function.Value
}

func getFormulaDesc(L *lua.State) (name string, argNames []string, args []Arg, noDraw []int, colors []string, lineThick []int, vars []string) {
	L.GetField(-1, "name")
	luar.LuaToGo(L, -1, &name)
	L.Pop(1)

	L.GetField(-1, "argName")
	luar.LuaToGo(L, -1, &argNames)
	L.Pop(1)

	if len(argNames) > 0 {
		args = make([]Arg, len(argNames))
		var values []float64
		for i, argName := range argNames {
			L.GetField(-1, argName)
			luar.LuaToGo(L, -1, &values)
			L.Pop(1)

			args[i].Default = values[0]
			args[i].Min = values[1]
			args[i].Max = values[2]
		}
	}

	L.GetField(-1, "noDraw")
	luar.LuaToGo(L, -1, &noDraw)
	L.Pop(1)

	L.GetField(-1, "colors")
	luar.LuaToGo(L, -1, &colors)
	L.Pop(1)

	L.GetField(-1, "lineThick")
	luar.LuaToGo(L, -1, &lineThick)
	L.Pop(1)

	L.GetField(-1, "vars")
	luar.LuaToGo(L, -1, &vars)
	L.Pop(1)

	return
}

func newFormulaByLuaState(L *lua.State, data *stockfunc.RVector, args []float64) (error, *LuaFormula) {
	L.GetGlobal("FormulaClass")

	name, argNames, argDefs, noDraw, colors, lineThick, vars := getFormulaDesc(L)

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

	return nil, &LuaFormula{
		L: L,

		name: name,
		ArgNames: argNames,
		Args: argDefs,
		NoDraws: noDraw,
		Colors: colors,
		LineThicks: lineThick,
		Vars: vars,

		refValues: values,
	}
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

func (this *LuaFormula) Name() string {
	return this.name
}

// 输出变量

func (this *LuaFormula) VarCount() int {
	return len(this.Vars)
}

func (this *LuaFormula) VarName(index int) string {
	if index < 0 || index >= len(this.Vars) {
		return ""
	}
	return this.Vars[index]
}

func (this *LuaFormula) NoDraw(index int) int {
	if index < 0 || index >= len(this.NoDraws) {
		return 0
	}
	return this.NoDraws[index]
}

func (this *LuaFormula) Color(index int) string {
	if index < 0 || index >= len(this.Colors) {
		return ""
	}
	return this.Colors[index]
}

func (this *LuaFormula) LineThick(index int) int {
	if index < 0 || index >= len(this.LineThicks) {
		return 1
	}
	return this.LineThicks[index]
}

// 公式参数

func (this *LuaFormula) ArgCount() int {
	return len(this.ArgNames)
}

func (this *LuaFormula) ArgRange(index int) (float64, float64) {
	if index < 0 || index >= len(this.Args) {
		return 0, 0
	}
	return this.Args[index].Min, this.Args[index].Max
}

func (this *LuaFormula) ArgDefault(index int) float64 {
	if index < 0 || index >= len(this.Args) {
		return 0
	}
	return this.Args[index].Default
}

func (this *LuaFormula) DrawActions() []formula.DrawAction {
	return nil
}

