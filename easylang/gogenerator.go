package easylang

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/stephenlyu/goformula/formulalibrary/base/formula"
)

type GoGenerator struct {
	context     *Context
	packageFull string
}

func NewGoGenerator(context *Context, packageFull string) Generator {
	return &GoGenerator{context, packageFull}
}

func (this *GoGenerator) translateDescriptions(desciptions []string) (flag int, graphType int, lineThick int, color string, lineStyle int) {
	flag = 0
	graphType = formula.FORMULA_GRAPH_LINE
	lineThick = 1
	lineStyle = formula.FORMULA_LINE_STYLE_SOLID
	color = ""

	for _, desc := range desciptions {
		switch {
		case desc == "DRAWABOVE":
			flag |= formula.FORMULA_VAR_FLAG_DRAW_ABOVE
		case desc == "NOFRAME":
			flag |= formula.FORMULA_VAR_FLAG_NO_FRAME
		case desc == "NODRAW":
			flag |= formula.FORMULA_VAR_FLAG_NO_DRAW
		case desc == "NOTEXT":
			flag |= formula.FORMULA_VAR_FLAG_NO_TEXT
		case desc == "COLORSTICK":
			graphType = formula.FORMULA_GRAPH_COLOR_STICK
		case desc == "STICK":
			graphType = formula.FORMULA_GRAPH_STICK
		case desc == "LINESTICK":
			graphType = formula.FORMULA_GRAPH_LINE_STICK
		case desc == "VOLSTICK":
			graphType = formula.FORMULA_GRAPH_VOL_STICK
		case desc == "DOTLINE":
			lineStyle = formula.FORMULA_LINE_STYLE_DOT
		case desc == "CROSSDOT":
			lineStyle = formula.FORMULA_LINE_STYLE_CROSS_DOT
		case desc == "CIRCLEDOT":
			lineStyle = formula.FORMULA_LINE_STYLE_CIRCLE_DOT
		case desc == "POINTDOT":
			lineStyle = formula.FORMULA_LINE_STYLE_POINT_DOT
		case strings.HasPrefix(desc, "COLOR"):
			color = desc
		case strings.HasPrefix(desc, "LINETHICK"):
			lineThick, _ = strconv.Atoi(desc[len("LINETHICK"):])
		}
	}
	return
}

func (this *GoGenerator) varProperties() (flags string, graphTypes string, lineThicks string, colors string, lineStyles string) {
	sFlags := make([]string, len(this.context.outputVars))
	sGraphTypes := make([]string, len(this.context.outputVars))
	sLineThicks := make([]string, len(this.context.outputVars))
	sLineStyles := make([]string, len(this.context.outputVars))
	sColors := make([]string, len(this.context.outputVars))

	isDrawFunc := func(expr expression) bool {
		aExpr := expr.(*assignexpr)

		for _, f := range this.context.drawFunctions {
			if aExpr.operand == expression(f) {
				return true
			}
		}
		return false
	}

	for i, varName := range this.context.outputVars {
		descriptions := this.context.outputDescriptions[varName]
		flag, graphType, lineThick, color, lineStyle := this.translateDescriptions(descriptions)
		exp := this.context.definedVarMap[varName]
		if isDrawFunc(exp) {
			graphType = formula.FORMULA_GRAPH_NONE
		}

		sFlags[i] = fmt.Sprintf("0x%08x", flag)
		sGraphTypes[i] = fmt.Sprintf("%d", graphType)
		sLineThicks[i] = fmt.Sprintf("%d", lineThick)
		if color == "" {
			sColors[i] = fmt.Sprintf("{Red:-1, Green:-1, Blue:-1}")
		} else {
			colorObject := ParseColorLiteral(color)
			sColors[i] = fmt.Sprintf("{Red:%d, Green:%d, Blue:%d}", colorObject.Red, colorObject.Green, colorObject.Blue)
		}
		sLineStyles[i] = fmt.Sprintf("%d", lineStyle)
	}
	flags = strings.Join(sFlags, ", ")
	graphTypes = strings.Join(sGraphTypes, ", ")
	lineThicks = strings.Join(sLineThicks, ", ")
	colors = strings.Join(sColors, ", ")
	lineStyles = strings.Join(sLineStyles, ", ")

	return
}

func (this *GoGenerator) getReferencedDataDeclarations(indent string) string {
	if len(this.context.refDataList) == 0 {
		return ""
	}

	lines := make([]string, len(this.context.refDataList)*2)
	for i, f := range this.context.refDataList {
		lines[2*i] = fmt.Sprintf("%s%s RVectorReader", indent, getRefDataVarName(f.code, f.period))
		lines[2*i+1] = fmt.Sprintf("%s%s *IndexMap", indent,
			getIndexMapVarName(f.code, f.period))
	}
	return strings.Join(lines, "\n")
}

func (this *GoGenerator) getReferencedFormulaDeclarations(indent string) string {
	if len(this.context.refFormulas) == 0 {
		return ""
	}

	lines := make([]string, len(this.context.refFormulas))
	for i, f := range this.context.refFormulas {
		lines[i] = fmt.Sprintf("%s%s Formula", indent, f.String())
	}
	return strings.Join(lines, "\n")
}

func (this *GoGenerator) getVectorDeclarations(indent string) string {
	var lines []string
	for _, varName := range this.context.definedVars {
		expr, ok := this.context.definedVarMap[varName]
		if !ok {
			continue
		}
		if !expr.IsValid() {
			continue
		}
		if expr.IsVoid() {
			// DO NOTHING
		} else {
			switch expr.(type) {
			case *stringexpr:
				lines = append(lines, fmt.Sprintf("%s%s string", indent, expr.DefinedName()))
			default:
				lines = append(lines, fmt.Sprintf("%s%s Value", indent, expr.DefinedName()))
			}
		}
	}
	return strings.Join(lines, "\n")
}

func (this *GoGenerator) getDataLibraryImport() string {
	if len(this.context.refDataList) == 0 {
		return ""
	}
	return `
	. "github.com/stephenlyu/goformula/datalibrary"`
}

func (this *GoGenerator) getFormulaLibraryImport() string {
	if len(this.context.refFormulas) == 0 {
		return ""
	}
	return `
	. "github.com/stephenlyu/goformula/formulalibrary"`
}

func (this *GoGenerator) getReferencedDataDefinitions(indent string) string {
	if len(this.context.refDataList) == 0 {
		return ""
	}

	lines := make([]string, len(this.context.refDataList)*2)
	for i, f := range this.context.refDataList {
		var codeStr, period string
		if f.code == "" {
			codeStr = "data.Code()"
		} else {
			codeStr = fmt.Sprintf(`"%s"`, f.code)
		}
		if f.period == "" {
			period = "data.Period().Name()"
		} else {
			period = fmt.Sprintf(`"%s"`, f.period)
		}

		lines[2*i] = fmt.Sprintf("%so.%s = GlobalDataLibrary.GetData(%s, %s)", indent, getRefDataVarName(f.code, f.period), codeStr, period)
		lines[2*i+1] = fmt.Sprintf("%so.%s = NewIndexMap(o.Data__, o.%s)", indent,
			getIndexMapVarName(f.code, f.period),
			getRefDataVarName(f.code, f.period))
	}
	return strings.Join(lines, "\n")
}

func (this *GoGenerator) getReferencedFormulaDefinitions(indent string) string {
	if len(this.context.refFormulas) == 0 {
		return ""
	}

	lines := make([]string, len(this.context.refFormulas))
	for i, f := range this.context.refFormulas {
		name := strings.ToUpper(f.name)
		varName := getRefDataVarName(f.code, f.period)
		lines[i] = fmt.Sprintf(`%so.%s = GlobalLibrary.NewFormula("%s", o.%s)`, indent, f.String(), name, strings.Replace(varName, "__data_code0___", "Data__", -1))
	}
	return strings.Join(lines, "\n")
}

func (this *GoGenerator) getVectorDefinitions(indent string, formulaName string) string {
	var lines []string
	for _, varName := range this.context.definedVars {
		expr, ok := this.context.definedVarMap[varName]
		if !ok {
			continue
		}
		if !expr.IsValid() {
			continue
		}
		if expr.IsVoid() {
			// DO NOTHING
		} else if _, ok := this.context.paramMap[varName]; ok {
			for i, p := range this.context.params {
				if p == varName {
					lines = append(lines, fmt.Sprintf("%so.%s = Scalar(args[%d])", indent, expr.DefinedName(), i))
					break
				}
			}
		} else {
			codes := strings.Replace(expr.Codes(), "__data_code0___", "Data__", -1)
			lines = append(lines, fmt.Sprintf("%so.%s = %s", indent, expr.DefinedName(), codes))
		}
	}
	return strings.Join(lines, "\n")
}

func (this *GoGenerator) paramNames() string {
	sa := make([]string, len(this.context.params))
	for i, p := range this.context.params {
		sa[i] = fmt.Sprintf(`"%s"`, p)
	}
	return strings.Join(sa, ", ")
}

func (this *GoGenerator) paramMetaData(name string) string {
	sa := make([]string, len(this.context.params))
	for i, p := range this.context.params {
		exp := this.context.paramMap[p].(*paramexpr)
		sa[i] = fmt.Sprintf("			Arg{Default:%f, Min:%f, Max:%f},", exp.defaultValue, exp.min, exp.max)
	}
	return strings.Join(sa, "\n")
}

func (this *GoGenerator) varNames() string {
	items := make([]string, len(this.context.outputVars))
	for i, varName := range this.context.outputVars {
		exp := this.context.definedVarMap[varName]
		items[i] = fmt.Sprintf(`"%s"`, exp.DisplayName())
	}
	return strings.Join(items, ", ")
}

func (this *GoGenerator) drawFunctionCodes() string {
	var drawActions []string

	relatedDescriptions := func(expr expression) ([]string, string) {
		for varName := range this.context.outputDescriptions {
			iExpr := this.context.definedVarMap[varName]
			aExpr, ok := iExpr.(*assignexpr)
			if !ok {
				continue
			}
			if aExpr.operand == expr {
				return this.context.outputDescriptions[varName], varName
			}
		}

		for varName := range this.context.notOutputDescriptions {
			expr := this.context.definedVarMap[varName]
			aExpr, ok := expr.(*assignexpr)
			if !ok {
				continue
			}

			if aExpr.operand == expr {
				return this.context.notOutputDescriptions[varName], varName
			}
		}
		return []string{}, ""
	}

	outputVarIndexOf := func(varName string) int {
		for i, n := range this.context.outputVars {
			if n == varName {
				return i
			}
		}
		return -1
	}

	for _, expr := range this.context.drawFunctions {
		descriptions, leftVarName := relatedDescriptions(expr)
		flag, _, lineThick, colorStr, _ := this.translateDescriptions(descriptions)

		varIndex := outputVarIndexOf(leftVarName)

		var color *formula.Color
		if colorStr == "" {
			color = &formula.Color{Red: -1, Green: -1, Blue: -1}
		} else {
			color = ParseColorLiteral(colorStr)
		}

		var noDraw int
		if flag&formula.FORMULA_VAR_FLAG_NO_DRAW != 0 {
			noDraw = 1
		}

		switch expr.funcName {
		case "DRAWTEXT":
			drawActions = append(drawActions, fmt.Sprintf("        &DrawTextAction{ActionType:%d, Cond:o.%s, Price:o.%s, Text:o.%s, Color:&Color{Red:%d, Green:%d, Blue:%d}, NoDraw:%d},",
				formula.FORMULA_DRAW_ACTION_DRAWTEXT,
				expr.arguments[0].DefinedName(),
				expr.arguments[1].DefinedName(),
				expr.arguments[2].DefinedName(),
				color.Red, color.Green, color.Blue,
				noDraw))
		case "DRAWICON":
			drawActions = append(drawActions, fmt.Sprintf("        &DrawIconAction{ActionType:%d, Cond:o.%s, Price:o.%s, Type:%d, NoDraw:%d},",
				formula.FORMULA_DRAW_ACTION_DRAWICON,
				expr.arguments[0].DefinedName(),
				expr.arguments[1].DefinedName(),
				int(expr.arguments[2].(*constantexpr).value),
				noDraw))
		case "DRAWLINE":
			drawActions = append(drawActions, fmt.Sprintf("        &DrawLineAction{ActionType:%d, Cond1:o.%s, Price1:o.%s, Cond2:o.%s, Price2:o.%s, Expand:%d, NoDraw:%d, Color:&Color{Red:%d, Green:%d, Blue:%d}, LineThick:%d, VarIndex:%d},",
				formula.FORMULA_DRAW_ACTION_DRAWLINE,
				expr.arguments[0].DefinedName(),
				expr.arguments[1].DefinedName(),
				expr.arguments[2].DefinedName(),
				expr.arguments[3].DefinedName(),
				int(expr.arguments[4].(*constantexpr).value),
				noDraw,
				color.Red, color.Green, color.Blue,
				lineThick,
				varIndex))
		case "DRAWKLINE":
			drawActions = append(drawActions, fmt.Sprintf("        &DrawKLineAction{ActionType:%d, High:o.%s, Open:o.%s, Low:o.%s, Close:o.%s, NoDraw:%d},",
				formula.FORMULA_DRAW_ACTION_DRAWKLINE,
				expr.arguments[0].DefinedName(),
				expr.arguments[1].DefinedName(),
				expr.arguments[2].DefinedName(),
				expr.arguments[3].DefinedName(),
				noDraw))
		case "STICKLINE":
			drawActions = append(drawActions, fmt.Sprintf("        &StickLineAction{ActionType:%d, Cond:o.%s, Price1:o.%s, Price2:o.%s, Width:%f, Empty:%d, NoDraw:%d, Color:&Color{Red:%d, Green:%d, Blue:%d}, LineThick:%d},",
				formula.FORMULA_DRAW_ACTION_STICKLINE,
				expr.arguments[0].DefinedName(),
				expr.arguments[1].DefinedName(),
				expr.arguments[2].DefinedName(),
				expr.arguments[3].(*constantexpr).value,
				int(expr.arguments[4].(*constantexpr).value),
				noDraw,
				color.Red, color.Green, color.Blue,
				lineThick))
		case "PLOYLINE":
			drawActions = append(drawActions, fmt.Sprintf("        &PloyLineAction{ActionType:%d, Cond:o.%s, Price:o.%s, NoDraw:%d, Color:&Color{Red:%d, Green:%d, Blue:%d}, LineThick:%d, VarIndex:%d},",
				formula.FORMULA_DRAW_ACTION_PLOYLINE,
				expr.arguments[0].DefinedName(),
				expr.arguments[1].DefinedName(),
				noDraw,
				color.Red, color.Green, color.Blue,
				lineThick,
				varIndex))
		}
	}

	return fmt.Sprintf(`
    o.DrawActions__ = []DrawAction{
%s
    }
`, strings.Join(drawActions, "\n"))
}

func (this *GoGenerator) refValuesCodes() string {
	items := make([]string, len(this.context.outputVars))
	for i, varName := range this.context.outputVars {
		exp := this.context.definedVarMap[varName]
		items[i] = fmt.Sprintf("o.%s", exp.DefinedName())
	}
	return strings.Join(items, ", ")
}

func (this *GoGenerator) updateLastValueCodes(indent string, formulaName string) string {
	lines := []string{}

	// Add Reference Formula UpdateLastValue Calls.
	for _, f := range this.context.refFormulas {
		lines = append(lines, fmt.Sprintf("%sthis.%s.UpdateLastValue()", indent, f.String()))
	}

	// Add Var UpdateLastValue Calls
	for _, varName := range this.context.definedVars {
		expr, ok := this.context.definedVarMap[varName]
		if !ok {
			continue
		}
		switch expr.(type) {
		case *constantexpr:
		case *assignexpr:
		case *paramexpr:
		case *stringexpr:
		case *referenceexpr:
		default:
			if !expr.IsValid() || expr.IsVoid() {
				continue
			}
			lines = append(lines, fmt.Sprintf("%sthis.%s.UpdateLastValue()", indent, expr.DefinedName()))
		}
	}
	return strings.Join(lines, "\n")
}

func (this *GoGenerator) dumpStateCodes(indent string, formulaName string) string {
	lines := []string{}

	// Add Var UpdateLastValue Calls
	for _, varName := range this.context.definedVars {
		expr, ok := this.context.definedVarMap[varName]
		if !ok {
			continue
		}
		if !expr.IsValid() {
			continue
		}
		if !expr.IsVoid() {
			// DO NOTHING
		}

		switch expr.(type) {
		case *constantexpr:
		case *assignexpr:
			lines = append(lines, fmt.Sprintf(`%s%sfmt.Printf("%s: %%.03f ", this.%s.Get(i))`,
				indent, indent, expr.DefinedName(), expr.DefinedName()))
		case *paramexpr:
		case *stringexpr:
		case *referenceexpr:
		default:
		}
	}
	return strings.Join(lines, "\n")
}

func (this *GoGenerator) GenerateCode(name string) string {
	name = strings.ToUpper(name)

	parts := filepath.SplitList(this.packageFull)
	packageName := parts[len(parts)-1]

	flags, graphTypes, lineThicks, colors, lineStyles := this.varProperties()

	const indent = "    "
	code := fmt.Sprintf(`//
// GENERATED BY EASYLANG COMPILER.
// !!!! DON'T MODIFY IT!!!!!!
//

package %s

import (
	"fmt"
	. "github.com/stephenlyu/goformula/stockfunc/function"
	. "github.com/stephenlyu/goformula/function"
	. "github.com/stephenlyu/goformula/formulalibrary/base/formula"
	. "github.com/stephenlyu/goformula/formulalibrary/native/nativeformulas"%s%s
)

type %s struct {
	BaseNativeFormula

	// Data of all referenced period
%s

	// Referenced Formulas
%s

	// Vectors
%s
}

var (
	%s_META = &FormulaMetaImpl{
		Name: "%s",
		ArgNames: []string{%s},
		ArgMeta: []Arg {
%s
		},
		Flags: []int{%s},
		Colors: []*Color{%s},
		LineThicks: []int{%s},
		LineStyles: []int{%s},
		GraphTypes: []int{%s},
		Vars: []string{%s},
	}
)

func New%s(data RVectorReader, args []float64) Formula {
	o := &%s{
		BaseNativeFormula: BaseNativeFormula{
			FormulaMetaImpl: %s_META,
			Data__: data,
		},
	}

	// Data of all referenced period
%s

	// Referenced Formulas
%s

	// Vectors
%s

	// Actions
%s
	o.RefValues__ = []Value {%s}
	return o
}

func (this *%s) UpdateLastValue() {
%s
}

func (this *%s) DumpState() {
	for i := 0; i < this.Data__.Len(); i++ {
		r := this.Data__.Get(i)
		fmt.Printf("date: %%s ", r.GetDate())
		// Dump Var Start
%s
		// Dump Var End
		fmt.Println("")
	}
}

func init() {
	RegisterNativeFormula(New%s, %s_META)
}

	`,
		packageName,
		this.getDataLibraryImport(),
		this.getFormulaLibraryImport(),
		strings.ToLower(name),
		this.getReferencedDataDeclarations(indent),
		this.getReferencedFormulaDeclarations(indent),
		this.getVectorDeclarations(indent),
		name,
		name,
		this.paramNames(),
		this.paramMetaData(name),
		flags,
		colors,
		lineThicks,
		lineStyles,
		graphTypes,
		this.varNames(),
		name,
		strings.ToLower(name),
		name,
		this.getReferencedDataDefinitions(indent),
		this.getReferencedFormulaDefinitions(indent),
		this.getVectorDefinitions(indent, name),
		this.drawFunctionCodes(),
		this.refValuesCodes(),
		strings.ToLower(name),
		this.updateLastValueCodes(indent, name),
		strings.ToLower(name),
		this.dumpStateCodes(indent, name),
		name,
		name,
	)

	return code
}
