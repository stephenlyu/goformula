package formula

const (
	FORMULA_VAR_FLAG_NO_DRAW = 1
	FORMULA_VAR_FLAG_NO_TEXT = 2
	FORMULA_VAR_FLAG_DRAW_ABOVE = 4
	FORMULA_VAR_FLAG_NO_FRAME = 8
)

const (
	FORMULA_NO_COLOR = ""
)

const (
	FORMULA_GRAPH_LINE = iota
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

type DrawAction interface {
	Color() string
	LineThick() int
	NoDraw() bool
}

type PolyLine interface {
	DrawAction

	Cond(index int) float64
	Price(index int) float64
}

type DrawLine interface {
	DrawAction

	Cond1(index int) float64
	Price1(index int) float64
	Cond2(index int) float64
	Price2(index int) float64
	Expand() int
}

type DrawKLine interface {
	DrawAction

	High(index int) float64
	Open(index int) float64
	Low(index int) float64
	Close(index int) float64
}

type StickLine interface {
	DrawAction

	Cond(index int) float64
	Price1(index int) float64
	Price2(index int) float64
	Width() float64
	Empty() float64
}

type DrawIcon interface {
	DrawAction

	Cond(index int) float64
	Price(index int) float64
	Type() int
}

type DrawText interface {
	DrawAction

	Cond(index int) float64
	Price(index int) float64
	Text() string
}

type Formula interface {
	Name() string								// 公式名

	// 输出变量
	VarCount() int
	VarName(index int) string					// Ref变量名列表
	NoDraw(index int) bool 						// 是否绘制图形
	NoText(index int) bool 						// 是否绘制文本
	DrawAbove(index int) bool
	NoFrame(index int) bool
	Color(index int) string						// 变量颜色, 形如black或FFFFFF
	LineThick(index int) int 					// 线宽，1-9
	LineStyle(index int) int 					// 线宽，1-9
	GraphType(index int) int

	// 公式参数
	ArgCount() int
	ArgRange(index int) (float64, float64)		// 参数范围
	ArgDefault(index int) float64				// 参数默认值

	Len() int
	UpdateLastValue()
	Get(index int) []float64
	Ref(offset int) []float64

	// 绘制图形
	DrawActions() []DrawAction

	Destroy()
}
