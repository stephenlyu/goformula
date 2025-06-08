package luaformula

import (
	"errors"
	"fmt"
	"math"
	"strings"

	"github.com/aarzilli/golua/lua"
	. "github.com/stephenlyu/goformula/formulalibrary/base/formula"
	"github.com/stephenlyu/goformula/function"
	stockfunc "github.com/stephenlyu/goformula/stockfunc/function"
	"github.com/stephenlyu/tds/period"
	"github.com/stephenlyu/tds/util"
	"github.com/stevedonovan/luar"
)

type LuaFormula struct {
	*FormulaMetaImpl
	L *lua.State

	period period.Period
	code   string
	data   stockfunc.RVectorReader

	args      []float64
	refValues []function.Value

	id int64

	// Draw Actions
	drawActions []DrawAction
}

func newFormulaByLuaState(L *lua.State, meta *FormulaMetaImpl, data stockfunc.RVectorReader, args []float64) (error, *LuaFormula) {
	L.GetGlobal("FormulaClass")
	if meta == nil {
		meta = &FormulaMetaImpl{}
		GetMetaFromLuaState(L, meta)
	}

	L.GetField(-1, "new")
	L.PushValue(-2)
	luar.GoToLuaProxy(L, data)
	for _, arg := range args {
		luar.GoToLua(L, arg)
	}
	L.Call(2+len(args), 1)
	if L.IsNil(-1) {
		return errors.New("Create formula fail"), nil
	}

	L.GetField(-1, "ref_values")
	var values []function.Value
	luar.LuaToGo(L, -1, &values)
	L.Pop(1)

	L.GetField(-1, "__id__")
	var id int64
	luar.LuaToGo(L, -1, &id)
	L.Pop(1)

	L.GetField(-1, "drawTextActions")
	var drawTexts []DrawTextAction
	luar.LuaToGo(L, -1, &drawTexts)
	L.Pop(1)

	L.GetField(-1, "drawIconActions")
	var drawIcons []DrawIconAction
	luar.LuaToGo(L, -1, &drawIcons)
	L.Pop(1)

	L.GetField(-1, "drawLineActions")
	var drawLines []DrawLineAction
	luar.LuaToGo(L, -1, &drawLines)
	L.Pop(1)

	L.GetField(-1, "drawKLineActions")
	var drawKLines []DrawKLineAction
	luar.LuaToGo(L, -1, &drawKLines)
	L.Pop(1)

	L.GetField(-1, "stickLineActions")
	var stickLines []StickLineAction
	luar.LuaToGo(L, -1, &stickLines)
	L.Pop(1)

	L.GetField(-1, "ployLineActions")
	var ployLines []PloyLineAction
	luar.LuaToGo(L, -1, &ployLines)
	L.Pop(1)

	drawActions := make([]DrawAction, len(drawTexts)+len(drawIcons)+len(drawLines)+len(drawKLines)+len(stickLines)+len(ployLines))
	i := 0
	for j := range drawTexts {
		action := &drawTexts[j]
		if action.Color != nil && action.Color.Red == -1 {
			action.Color = nil
		}
		drawActions[i] = action
		i++
	}
	for j := range drawIcons {
		action := &drawIcons[j]
		drawActions[i] = action
		i++
	}
	for j := range drawLines {
		action := &drawLines[j]
		if action.Color != nil && action.Color.Red == -1 {
			action.Color = nil
		}
		drawActions[i] = action
		i++
	}
	for j := range drawKLines {
		action := &drawKLines[j]
		drawActions[i] = action
		i++
	}
	for j := range stickLines {
		action := &stickLines[j]
		if action.Color != nil && action.Color.Red == -1 {
			action.Color = nil
		}
		drawActions[i] = action
		i++
	}
	for j := range ployLines {
		action := &ployLines[j]
		if action.Color != nil && action.Color.Red == -1 {
			action.Color = nil
		}
		drawActions[i] = action
		i++
	}

	L.Remove(1) // Remove FormulaClass from stack
	L.Pop(1)    // Remove formula object from stack
	util.Assert(L.GetTop() == 0, "")

	formula := &LuaFormula{
		FormulaMetaImpl: meta,
		L:               L,

		period: data.Period(),
		code:   data.Code(),
		data:   data,

		args:      args,
		refValues: values,

		id: id,

		drawActions: drawActions,
	}

	return nil, formula
}

func NewFormula(luaFile string, data stockfunc.RVectorReader, args []float64) (error, *LuaFormula) {
	L := luar.Init()

	luar.Register(L, "", GetFunctionMap(luar.Map{}))

	err := L.DoFile(luaFile)
	if err != nil {
		return err, nil
	}

	return newFormulaByLuaState(L, nil, data, args)
}

func NewFormulaFromState(L *lua.State, meta *FormulaMetaImpl, data stockfunc.RVectorReader, args []float64) (error, *LuaFormula) {
	return newFormulaByLuaState(L, meta, data, args)
}

func NewFormulaFromCode(luaCode string, data stockfunc.RVectorReader, args []float64) (error, *LuaFormula) {
	L := luar.Init()

	luar.Register(L, "", GetFunctionMap(luar.Map{}))

	err := L.DoString(luaCode)
	if err != nil {
		return err, nil
	}

	return newFormulaByLuaState(L, nil, data, args)
}

func (this *LuaFormula) Destroy() {
	this.L.GetGlobal("RemoveObject")
	this.L.PushInteger(this.id)
	this.L.Call(1, 1)
}

func (this *LuaFormula) Period() period.Period {
	return this.period
}

func (this *LuaFormula) Len() int {
	return this.data.Len()
}

func (this *LuaFormula) getObject() {
	this.L.GetGlobal("GetObject")
	this.L.PushInteger(this.id)
	this.L.Call(1, 1)
}

func (this *LuaFormula) UpdateLastValue() {
	//util.Assert(this.L.GetTop() == 0, "")
	this.getObject()
	this.L.GetField(-1, "updateLastValue")
	this.L.PushValue(-2)
	this.L.Call(1, 0)
	this.L.Pop(1)
	//util.Assert(this.L.GetTop() == 0, "")
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

func (this *LuaFormula) GetVarValue(varName string) function.Value {
	for i, v := range this.FormulaMetaImpl.Vars {
		if varName == v {
			return this.refValues[i]
		}
	}
	return nil
}

func (this *LuaFormula) Name() string {
	if len(this.ArgMeta) == 0 {
		return this.GetName()
	}

	formatValue := func(v float64) string {
		if float64(int(v)) == v {
			return fmt.Sprintf("%d", int(v))
		}

		n := math.Log(math.Abs(v))
		switch {
		case n > 4:
			return fmt.Sprintf("%.0f", v)
		case n == 3:
			return fmt.Sprintf("%.1f", v)
		case n > -2:
			return fmt.Sprintf("%.2f", v)
		case n == -3:
			return fmt.Sprintf("%.3f", v)
		default:
			return fmt.Sprintf("%.4f", v)
		}
	}

	items := make([]string, len(this.ArgMeta))
	for i, arg := range this.args {
		items[i] = fmt.Sprintf(formatValue(arg))
	}

	return fmt.Sprintf("%s(%s)", this.GetName(), strings.Join(items, ", "))
}

func (this *LuaFormula) DrawActions() []DrawAction {
	return this.drawActions
}

func (this *LuaFormula) DumpState() {
}
