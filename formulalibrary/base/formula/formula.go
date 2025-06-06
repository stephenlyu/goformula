package formula

import (
	"github.com/stephenlyu/goformula/function"
	"github.com/stephenlyu/tds/period"
)

const (
	FORMULA_VAR_FLAG_NO_DRAW    = 1
	FORMULA_VAR_FLAG_NO_TEXT    = 2
	FORMULA_VAR_FLAG_DRAW_ABOVE = 4
	FORMULA_VAR_FLAG_NO_FRAME   = 8
)

const (
	FORMULA_GRAPH_NONE = iota
	FORMULA_GRAPH_LINE
	FORMULA_GRAPH_COLOR_STICK
	FORMULA_GRAPH_STICK
	FORMULA_GRAPH_VOL_STICK
	FORMULA_GRAPH_LINE_STICK
)

const (
	FORMULA_LINE_STYLE_SOLID = iota
	FORMULA_LINE_STYLE_DOT
	FORMULA_LINE_STYLE_CROSS_DOT
	FORMULA_LINE_STYLE_CIRCLE_DOT
	FORMULA_LINE_STYLE_POINT_DOT
)

const (
	FORMULA_DRAW_ACTION_PLOYLINE = iota
	FORMULA_DRAW_ACTION_DRAWLINE
	FORMULA_DRAW_ACTION_DRAWKLINE
	FORMULA_DRAW_ACTION_STICKLINE
	FORMULA_DRAW_ACTION_DRAWICON
	FORMULA_DRAW_ACTION_DRAWTEXT
)

type Color struct {
	Red   int
	Green int
	Blue  int
}

type DrawAction interface {
	GetActionType() int
	GetVarIndex() int
	GetColor() *Color
	GetLineThick() int
	IsNoDraw() bool
}

type PloyLine interface {
	DrawAction

	GetCond(index int) float64
	GetPrice(index int) float64
}

type DrawLine interface {
	DrawAction

	GetCond1(index int) float64
	GetPrice1(index int) float64
	GetCond2(index int) float64
	GetPrice2(index int) float64
	GetExpand() int
}

type DrawKLine interface {
	DrawAction

	GetHigh(index int) float64
	GetOpen(index int) float64
	GetLow(index int) float64
	GetClose(index int) float64
}

type StickLine interface {
	DrawAction

	GetCond(index int) float64
	GetPrice1(index int) float64
	GetPrice2(index int) float64
	GetWidth() float64
	GetEmpty() float64
}

type DrawIcon interface {
	DrawAction

	GetCond(index int) float64
	GetPrice(index int) float64
	GetType() int
}

type DrawText interface {
	DrawAction

	GetCond(index int) float64
	GetPrice(index int) float64
	GetText() string
}

type Formula interface {
	FormulaMeta

	Period() period.Period
	Len() int
	UpdateLastValue()
	Get(index int) []float64
	Ref(offset int) []float64
	GetVarValue(varName string) function.Value
	DumpState()

	// 绘制图形
	DrawActions() []DrawAction

	Destroy()
}
