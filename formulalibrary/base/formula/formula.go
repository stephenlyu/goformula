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

type PloyLine[T float32 | float64] interface {
	DrawAction

	GetCond(index int) T
	GetPrice(index int) T
}

type DrawLine[T float32 | float64] interface {
	DrawAction

	GetCond1(index int) T
	GetPrice1(index int) T
	GetCond2(index int) T
	GetPrice2(index int) T
	GetExpand() int
}

type DrawKLine[T float32 | float64] interface {
	DrawAction

	GetHigh(index int) T
	GetOpen(index int) T
	GetLow(index int) T
	GetClose(index int) T
}

type StickLine[T float32 | float64] interface {
	DrawAction

	GetCond(index int) T
	GetPrice1(index int) T
	GetPrice2(index int) T
	GetWidth() T
	GetEmpty() T
}

type DrawIcon[T float32 | float64] interface {
	DrawAction

	GetCond(index int) T
	GetPrice(index int) T
	GetType() int
}

type DrawText[T float32 | float64] interface {
	DrawAction

	GetCond(index int) T
	GetPrice(index int) T
	GetText() string
}

type Formula[T float32 | float64] interface {
	FormulaMeta

	SetId(id int64)
	Id() int64

	IncRef()
	DecRef()

	Period() period.Period
	Len() int
	UpdateLastValue()
	Get(index int) []T
	Ref(offset int) []T
	GetVarValue(varName string) function.Value[T]
	DumpState()

	// 绘制图形
	DrawActions() []DrawAction

	Destroy()
}
