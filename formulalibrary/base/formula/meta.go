package formula

type FormulaMeta interface {
	GetName() string							// 公式名

	// 输出变量
	VarCount() int
	VarName(index int) string					// Ref变量名列表
	HasVar(name string) bool
	NoDraw(index int) bool 						// 是否绘制图形
	NoText(index int) bool 						// 是否绘制文本
	DrawAbove(index int) bool
	NoFrame(index int) bool
	Color(index int) *Color						// 变量颜色
	LineThick(index int) int 					// 线宽，1-9
	LineStyle(index int) int 					// 线宽，1-9
	GraphType(index int) int

	// 公式参数
	ArgCount() int
	ArgRange(index int) (float64, float64)		// 参数范围
	ArgDefault(index int) float64				// 参数默认值
	DefaultArgs() []float64
}

// Formula Meta Implementations

type Arg struct {
	Min float64
	Max float64
	Default float64
}

type FormulaMetaImpl struct {
	Name       string
	ArgNames   []string
	ArgMeta    []Arg
	Flags      []int
	Colors     []*Color
	LineThicks []int
	LineStyles []int
	GraphTypes []int
	Vars       []string
}

func (this *FormulaMetaImpl) GetName() string {
	return this.Name
}

// 输出变量

func (this *FormulaMetaImpl) VarCount() int {
	return len(this.Vars)
}

func (this *FormulaMetaImpl) VarName(index int) string {
	if index < 0 || index >= len(this.Vars) {
		return ""
	}
	return this.Vars[index]
}

func (this *FormulaMetaImpl) HasVar(varName string) bool {
	for _, v := range this.Vars {
		if v == varName {
			return true
		}
	}
	return false
}

func (this *FormulaMetaImpl) NoDraw(index int) bool {
	if index < 0 || index >= len(this.Flags) {
		return false
	}
	return (this.Flags[index] & FORMULA_VAR_FLAG_NO_DRAW) != 0
}

func (this *FormulaMetaImpl) NoText(index int) bool {
	if index < 0 || index >= len(this.Flags) {
		return false
	}
	return (this.Flags[index] & FORMULA_VAR_FLAG_NO_TEXT) != 0
}

func (this *FormulaMetaImpl) DrawAbove(index int) bool {
	if index < 0 || index >= len(this.Flags) {
		return false
	}
	return this.Flags[index] & FORMULA_VAR_FLAG_DRAW_ABOVE != 0
}

func (this *FormulaMetaImpl) NoFrame(index int) bool {
	if index < 0 || index >= len(this.Flags) {
		return false
	}
	return this.Flags[index] & FORMULA_VAR_FLAG_NO_FRAME != 0
}

func (this *FormulaMetaImpl) Color(index int) *Color {
	if index < 0 || index >= len(this.Colors) {
		return nil
	}
	return this.Colors[index]
}

func (this *FormulaMetaImpl) LineThick(index int) int {
	if index < 0 || index >= len(this.LineThicks) {
		return 1
	}
	return this.LineThicks[index]
}

func (this *FormulaMetaImpl) LineStyle(index int) int {
	if index < 0 || index >= len(this.LineStyles) {
		return 1
	}
	return this.LineStyles[index]
}

func (this *FormulaMetaImpl) GraphType(index int) int {
	if index < 0 || index >= len(this.GraphTypes) {
		return 1
	}
	return this.GraphTypes[index]
}

// 公式参数

func (this *FormulaMetaImpl) ArgCount() int {
	return len(this.ArgNames)
}

func (this *FormulaMetaImpl) ArgRange(index int) (float64, float64) {
	if index < 0 || index >= len(this.ArgMeta) {
		return 0, 0
	}
	return this.ArgMeta[index].Min, this.ArgMeta[index].Max
}

func (this *FormulaMetaImpl) ArgDefault(index int) float64 {
	if index < 0 || index >= len(this.ArgMeta) {
		return 0
	}
	return this.ArgMeta[index].Default
}

func (this *FormulaMetaImpl) DefaultArgs() []float64 {
	ret := make([]float64, this.ArgCount())
	for i := range this.ArgMeta {
		ret[i] = this.ArgMeta[i].Default
	}
	return ret
}
