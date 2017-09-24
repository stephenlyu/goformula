package formula

const (
	FORMULA_VAR_FLAG_NO_DRAW = 1
)

const (
	FORMULA_NO_COLOR = ""
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
	NoDraw(index int) int 						// 是否绘制
	Color(index int) string						// 变量颜色, 形如black或FFFFFF
	LineThick(index int) int 					// 线宽，1-9

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
