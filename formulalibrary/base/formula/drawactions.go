package formula

import "github.com/stephenlyu/goformula/function"

// Draw Text Action
type DrawTextAction[T float32 | float64] struct {
	ActionType int
	NoDraw     int
	Color      *Color
	LineThick  int

	Cond  function.Value[T]
	Price function.Value[T]
	Text  string
}

func (this *DrawTextAction[T]) GetActionType() int {
	return this.ActionType
}

func (this *DrawTextAction[T]) GetColor() *Color {
	return this.Color
}

func (this *DrawTextAction[T]) GetLineThick() int {
	return this.LineThick
}

func (this *DrawTextAction[T]) IsNoDraw() bool {
	return this.NoDraw != 0
}

func (this *DrawTextAction[T]) GetVarIndex() int {
	return -1
}

func (this *DrawTextAction[T]) GetCond(index int) T {
	return this.Cond.Get(index)
}

func (this *DrawTextAction[T]) GetPrice(index int) T {
	return this.Price.Get(index)
}

func (this *DrawTextAction[T]) GetText() string {
	return this.Text
}

// Draw Icon Action
type DrawIconAction[T float32 | float64] struct {
	ActionType int
	NoDraw     int
	Color      *Color
	LineThick  int

	Cond  function.Value[T]
	Price function.Value[T]
	Type  int
}

func (this *DrawIconAction[T]) GetActionType() int {
	return this.ActionType
}

func (this *DrawIconAction[T]) GetColor() *Color {
	return this.Color
}

func (this *DrawIconAction[T]) GetLineThick() int {
	return this.LineThick
}

func (this *DrawIconAction[T]) IsNoDraw() bool {
	return this.NoDraw != 0
}

func (this *DrawIconAction[T]) GetVarIndex() int {
	return -1
}

func (this *DrawIconAction[T]) GetCond(index int) T {
	return this.Cond.Get(index)
}

func (this *DrawIconAction[T]) GetPrice(index int) T {
	return this.Price.Get(index)
}

func (this *DrawIconAction[T]) GetType() int {
	return this.Type
}

// Draw Line Action

type DrawLineAction[T float32 | float64] struct {
	ActionType int
	NoDraw     int
	Color      *Color
	LineThick  int
	VarIndex   int

	Cond1  function.Value[T]
	Price1 function.Value[T]
	Cond2  function.Value[T]
	Price2 function.Value[T]
	Expand int
}

func (this *DrawLineAction[T]) GetActionType() int {
	return this.ActionType
}

func (this *DrawLineAction[T]) GetColor() *Color {
	return this.Color
}

func (this *DrawLineAction[T]) GetLineThick() int {
	return this.LineThick
}

func (this *DrawLineAction[T]) IsNoDraw() bool {
	return this.NoDraw != 0
}

func (this *DrawLineAction[T]) GetVarIndex() int {
	return this.VarIndex
}

func (this *DrawLineAction[T]) GetCond1(index int) T {
	return this.Cond1.Get(index)
}

func (this *DrawLineAction[T]) GetPrice1(index int) T {
	return this.Price1.Get(index)
}

func (this *DrawLineAction[T]) GetCond2(index int) T {
	return this.Cond2.Get(index)
}

func (this *DrawLineAction[T]) GetPrice2(index int) T {
	if index < 0 || index >= this.Price2.Len() {
		return T(function.NaN)
	}
	return this.Price2.Get(index)
}

func (this *DrawLineAction[T]) GetExpand() int {
	return this.Expand
}

// Stick Line Action

type StickLineAction[T float32 | float64] struct {
	ActionType int
	NoDraw     int
	Color      *Color
	LineThick  int

	Cond   function.Value[T]
	Price1 function.Value[T]
	Price2 function.Value[T]
	Width  T
	Empty  T
}

func (this *StickLineAction[T]) GetActionType() int {
	return this.ActionType
}

func (this *StickLineAction[T]) GetColor() *Color {
	return this.Color
}

func (this *StickLineAction[T]) GetLineThick() int {
	return this.LineThick
}

func (this *StickLineAction[T]) IsNoDraw() bool {
	return this.NoDraw != 0
}

func (this *StickLineAction[T]) GetVarIndex() int {
	return -1
}

func (this *StickLineAction[T]) GetCond(index int) T {
	return this.Cond.Get(index)
}

func (this *StickLineAction[T]) GetPrice1(index int) T {
	return this.Price1.Get(index)
}

func (this *StickLineAction[T]) GetPrice2(index int) T {
	return this.Price2.Get(index)
}

func (this *StickLineAction[T]) GetWidth() T {
	return this.Width
}

func (this *StickLineAction[T]) GetEmpty() T {
	return this.Empty
}

// Ploy Line Action

type PloyLineAction[T float32 | float64] struct {
	ActionType int
	NoDraw     int
	Color      *Color
	LineThick  int
	VarIndex   int

	Cond  function.Value[T]
	Price function.Value[T]
}

func (this *PloyLineAction[T]) GetActionType() int {
	return this.ActionType
}

func (this *PloyLineAction[T]) GetColor() *Color {
	return this.Color
}

func (this *PloyLineAction[T]) GetLineThick() int {
	return this.LineThick
}

func (this *PloyLineAction[T]) IsNoDraw() bool {
	return this.NoDraw != 0
}

func (this *PloyLineAction[T]) GetVarIndex() int {
	return this.VarIndex
}

func (this *PloyLineAction[T]) GetCond(index int) T {
	return this.Cond.Get(index)
}

func (this *PloyLineAction[T]) GetPrice(index int) T {
	return this.Price.Get(index)
}

// Draw KLine Action

type DrawKLineAction[T float32 | float64] struct {
	ActionType int
	NoDraw     int
	Color      *Color
	LineThick  int

	High  function.Value[T]
	Open  function.Value[T]
	Low   function.Value[T]
	Close function.Value[T]
}

func (this *DrawKLineAction[T]) GetActionType() int {
	return this.ActionType
}

func (this *DrawKLineAction[T]) GetColor() *Color {
	return this.Color
}

func (this *DrawKLineAction[T]) GetLineThick() int {
	return this.LineThick
}

func (this *DrawKLineAction[T]) IsNoDraw() bool {
	return this.NoDraw != 0
}

func (this *DrawKLineAction[T]) GetVarIndex() int {
	return -1
}

func (this *DrawKLineAction[T]) GetHigh(index int) T {
	return this.High.Get(index)
}

func (this *DrawKLineAction[T]) GetOpen(index int) T {
	return this.Open.Get(index)
}

func (this *DrawKLineAction[T]) GetLow(index int) T {
	return this.Low.Get(index)
}

func (this *DrawKLineAction[T]) GetClose(index int) T {
	return this.Close.Get(index)
}
