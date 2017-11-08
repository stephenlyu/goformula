package formula

import "github.com/stephenlyu/goformula/function"

// Draw Text Action
type DrawTextAction struct {
	ActionType int
	NoDraw    int
	Color     *Color
	LineThick int

	Cond  function.Value
	Price function.Value
	Text  string
}

func (this *DrawTextAction) GetActionType() int {
	return this.ActionType
}

func (this *DrawTextAction) GetColor() *Color {
	return this.Color
}

func (this *DrawTextAction) GetLineThick() int {
	return this.LineThick
}

func (this *DrawTextAction) IsNoDraw() bool {
	return this.NoDraw != 0
}

func (this *DrawTextAction) GetVarIndex() int {
	return -1
}

func (this *DrawTextAction) GetCond(index int) float64 {
	if index < 0 || index >= this.Cond.Len() {
		return function.NaN
	}
	return this.Cond.Get(index)
}

func (this *DrawTextAction) GetPrice(index int) float64 {
	if index < 0 || index >= this.Price.Len() {
		return function.NaN
	}
	return this.Price.Get(index)
}

func (this *DrawTextAction) GetText() string {
	return this.Text
}

// Draw Icon Action
type DrawIconAction struct {
	ActionType int
	NoDraw    int
	Color     *Color
	LineThick int

	Cond  function.Value
	Price function.Value
	Type  int
}

func (this *DrawIconAction) GetActionType() int {
	return this.ActionType
}

func (this *DrawIconAction) GetColor() *Color {
	return this.Color
}

func (this *DrawIconAction) GetLineThick() int {
	return this.LineThick
}

func (this *DrawIconAction) IsNoDraw() bool {
	return this.NoDraw != 0
}

func (this *DrawIconAction) GetVarIndex() int {
	return -1
}

func (this *DrawIconAction) GetCond(index int) float64 {
	if index < 0 || index >= this.Cond.Len() {
		return function.NaN
	}
	return this.Cond.Get(index)
}

func (this *DrawIconAction) GetPrice(index int) float64 {
	if index < 0 || index >= this.Price.Len() {
		return function.NaN
	}
	return this.Price.Get(index)
}

func (this *DrawIconAction) GetType() int {
	return this.Type
}

// Draw Line Action

type DrawLineAction struct {
	ActionType int
	NoDraw    int
	Color     *Color
	LineThick int
	VarIndex int

	Cond1  function.Value
	Price1 function.Value
	Cond2  function.Value
	Price2 function.Value
	Expand int
}

func (this *DrawLineAction) GetActionType() int {
	return this.ActionType
}

func (this *DrawLineAction) GetColor() *Color {
	return this.Color
}

func (this *DrawLineAction) GetLineThick() int {
	return this.LineThick
}

func (this *DrawLineAction) IsNoDraw() bool {
	return this.NoDraw != 0
}

func (this *DrawLineAction) GetVarIndex() int {
	return this.VarIndex
}

func (this *DrawLineAction) GetCond1(index int) float64 {
	if index < 0 || index >= this.Cond1.Len() {
		return function.NaN
	}
	return this.Cond1.Get(index)
}

func (this *DrawLineAction) GetPrice1(index int) float64 {
	if index < 0 || index >= this.Price1.Len() {
		return function.NaN
	}
	return this.Price1.Get(index)
}

func (this *DrawLineAction) GetCond2(index int) float64 {
	if index < 0 || index >= this.Cond2.Len() {
		return function.NaN
	}
	return this.Cond2.Get(index)
}

func (this *DrawLineAction) GetPrice2(index int) float64 {
	if index < 0 || index >= this.Price2.Len() {
		return function.NaN
	}
	return this.Price2.Get(index)
}

func (this *DrawLineAction) GetExpand() int {
	return this.Expand
}

// Stick Line Action

type StickLineAction struct {
	ActionType int
	NoDraw    int
	Color     *Color
	LineThick int

	Cond   function.Value
	Price1 function.Value
	Price2 function.Value
	Width  float64
	Empty  float64
}

func (this *StickLineAction) GetActionType() int {
	return this.ActionType
}

func (this *StickLineAction) GetColor() *Color {
	return this.Color
}

func (this *StickLineAction) GetLineThick() int {
	return this.LineThick
}

func (this *StickLineAction) IsNoDraw() bool {
	return this.NoDraw != 0
}

func (this *StickLineAction) GetVarIndex() int {
	return -1
}

func (this *StickLineAction) GetCond(index int) float64 {
	if index < 0 || index >= this.Cond.Len() {
		return function.NaN
	}
	return this.Cond.Get(index)
}

func (this *StickLineAction) GetPrice1(index int) float64 {
	if index < 0 || index >= this.Price1.Len() {
		return function.NaN
	}
	return this.Price1.Get(index)
}

func (this *StickLineAction) GetPrice2(index int) float64 {
	if index < 0 || index >= this.Price2.Len() {
		return function.NaN
	}
	return this.Price2.Get(index)
}

func (this *StickLineAction) GetWidth() float64 {
	return this.Width
}

func (this *StickLineAction) GetEmpty() float64 {
	return this.Empty
}

// Ploy Line Action

type PloyLineAction struct {
	ActionType int
	NoDraw    int
	Color     *Color
	LineThick int
	VarIndex int

	Cond  function.Value
	Price function.Value
}

func (this *PloyLineAction) GetActionType() int {
	return this.ActionType
}

func (this *PloyLineAction) GetColor() *Color {
	return this.Color
}

func (this *PloyLineAction) GetLineThick() int {
	return this.LineThick
}

func (this *PloyLineAction) IsNoDraw() bool {
	return this.NoDraw != 0
}

func (this *PloyLineAction) GetVarIndex() int {
	return this.VarIndex
}

func (this *PloyLineAction) GetCond(index int) float64 {
	if index < 0 || index >= this.Cond.Len() {
		return function.NaN
	}
	return this.Cond.Get(index)
}

func (this *PloyLineAction) GetPrice(index int) float64 {
	if index < 0 || index >= this.Price.Len() {
		return function.NaN
	}
	return this.Price.Get(index)
}

// Draw KLine Action

type DrawKLineAction struct {
	ActionType int
	NoDraw    int
	Color     *Color
	LineThick int

	High  function.Value
	Open  function.Value
	Low   function.Value
	Close function.Value
}

func (this *DrawKLineAction) GetActionType() int {
	return this.ActionType
}

func (this *DrawKLineAction) GetColor() *Color {
	return this.Color
}

func (this *DrawKLineAction) GetLineThick() int {
	return this.LineThick
}

func (this *DrawKLineAction) IsNoDraw() bool {
	return this.NoDraw != 0
}

func (this *DrawKLineAction) GetVarIndex() int {
	return -1
}

func (this *DrawKLineAction) GetHigh(index int) float64 {
	if index < 0 || index >= this.High.Len() {
		return function.NaN
	}
	return this.High.Get(index)
}

func (this *DrawKLineAction) GetOpen(index int) float64 {
	if index < 0 || index >= this.Open.Len() {
		return function.NaN
	}
	return this.Open.Get(index)
}

func (this *DrawKLineAction) GetLow(index int) float64 {
	if index < 0 || index >= this.Low.Len() {
		return function.NaN
	}
	return this.Low.Get(index)
}

func (this *DrawKLineAction) GetClose(index int) float64 {
	if index < 0 || index >= this.Close.Len() {
		return function.NaN
	}
	return this.Close.Get(index)
}
