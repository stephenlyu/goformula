package easylang

import (
	"strings"
	"fmt"
	"github.com/stephenlyu/goformula/formulalibrary/base/formula"
	"strconv"
)

type LuaGenerator struct {
	context *Context
}

func NewLuaGenerator(context *Context) Generator {
	return &LuaGenerator{context}
}

func (this *LuaGenerator) refFormulaDefineCodes(indent string) string {
	if len(this.context.refFormulas) == 0 {
		return ""
	}

	lines := make([]string, len(this.context.refFormulas))
	for i, f := range this.context.refFormulas {
		name := strings.ToUpper(f.name)
		lines[i] = fmt.Sprintf("%so.%s = FormulaManager.NewFormula('%s', o.%s)", indent, f.String(), name, getRefDataVarName(f.code, f.period))
	}
	return strings.Join(lines, "\n")
}

func (this *LuaGenerator) refDataDefineCodes(indent string) string {
	if len(this.context.refDataList) == 0 {
		return ""
	}

	lines := make([]string, len(this.context.refDataList)*2)
	for i, f := range this.context.refDataList {
		var codeStr, period string
		if f.code == "" {
			codeStr = "data.Code()"
		} else {
			codeStr = "'" + f.code + "'"
		}
		if f.period == "" {
			period = "data.Period().Name()"
		} else {
			period = "'" + f.period + "'"
		}

		lines[2*i] = fmt.Sprintf("%so.%s = DataLibrary.GetData(%s, %s)", indent, getRefDataVarName(f.code, f.period), codeStr, period)
		lines[2*i+1] = fmt.Sprintf("%so.%s = IndexMap(o.%s, o.%s)", indent,
			getIndexMapVarName(f.code, f.period),
			getRefDataVarName("", ""),
			getRefDataVarName(f.code, f.period))
	}
	return strings.Join(lines, "\n")
}

func (this *LuaGenerator) definedCodes(indent string, formulaName string) string {
	var lines []string
	i := 0
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
			lines = append(lines, fmt.Sprintf("%so.%s = %s", indent, expr.DefinedName(), expr.Codes()))
		}
		if DEBUG {
			lines = append(lines, fmt.Sprintf("%sprint('%sClass:New %d...')", indent, formulaName, i))
			i++
		}
	}
	return strings.Join(lines, "\n")
}

func (this *LuaGenerator) updateLastValueCodes(indent string, formulaName string) string {
	lines := []string{}

	i := 0
	if DEBUG {
		lines = append(lines, fmt.Sprintf("%sprint('%sClass:updateLastValue %d')", indent, formulaName, i))
		i++
	}

	// Add Reference Formula UpdateLastValue Calls.
	for _, f := range this.context.refFormulas {
		lines = append(lines, fmt.Sprintf("%so.%s.UpdateLastValue()", indent, f.String()))
		if DEBUG {
			lines = append(lines, fmt.Sprintf("%sprint('%sClass:updateLastValue %d')", indent, formulaName, i))
			i++
		}
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
			lines = append(lines, fmt.Sprintf("%so.%s.UpdateLastValue()", indent, expr.DefinedName()))
			if DEBUG {
				lines = append(lines, fmt.Sprintf("%sprint('%sClass:updateLastValue %d')", indent, formulaName, i))
				i++
			}
		}
	}
	return strings.Join(lines, "\n")
}

func (this *LuaGenerator) refValuesCodes() string {
	items := make([]string, len(this.context.outputVars))
	for i, varName := range this.context.outputVars {
		exp := this.context.definedVarMap[varName]
		items[i] = fmt.Sprintf("o.%s", exp.DefinedName())
	}
	return strings.Join(items, ", ")
}

func (this *LuaGenerator) varNames() string {
	items := make([]string, len(this.context.outputVars))
	for i, varName := range this.context.outputVars {
		exp := this.context.definedVarMap[varName]
		items[i] = fmt.Sprintf("'%s'", exp.DisplayName())
	}
	return strings.Join(items, ", ")
}

func (this *LuaGenerator) translateDescriptions(desciptions []string) (flag int, graphType int, lineThick int, color string, lineStyle int) {
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

func (this *LuaGenerator) varProperties() (flags string, graphTypes string, lineThicks string, colors string, lineStyles string) {
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
			sColors[i] = fmt.Sprintf("{Red=-1, Green=-1, Blue=-1}")
		} else {
			colorObject := ParseColorLiteral(color)
			sColors[i] = fmt.Sprintf("{Red=%d, Green=%d, Blue=%d}", colorObject.Red, colorObject.Green, colorObject.Blue)
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

//o.drawTextActions = {
//{Cond=o.__anonymous_0_div___anonymous_1_gt_const1, Price=o.__anonymous_2, Text=o.string1, Color={Red=-1, Green=-1, Blue=-1}, NoDraw=0}
//}
//
//o.drawIconActions = {
//{Cond=o.__anonymous_19_gt___anonymous_20, Price=o.__anonymous_21, Type=o.const8, NoDraw=0}
//}
//
//o.drawLineActions = {
//{Cond1=o.__anonymous_3_ge_hhv___anonymous_4_const2, Price1=o.__anonymous_5, Cond2=o.__anonymous_6_le_llv___anonymous_7_const3, Price2=o.__anonymous_8, Expand=o.const4, NoDraw=0, Color={Red=-1, Green=-1, Blue=-1}, LineThick=1}
//}
//
//o.drawKLineActions = {
//{High=o.__anonymous_23, Open=o.__anonymous_24, Low=o.__anonymous_25, Close=o.__anonymous_26, NoDraw=0}
//}
//
//o.stickLineActions = {
//{Cond=o.__anonymous_14_gt___anonymous_15, Price1=o.__anonymous_16, Price2=o.__anonymous_17, Width=o.const6, Empty=o.const7, NoDraw=0, Color={Red=-1, Green=-1, Blue=-1}, LineThick=1}
//}
//
//o.ployLineActions = {
//{Cond=o.__anonymous_10_ge_hhv___anonymous_11_const5, Price=o.__anonymous_12, NoDraw=0, Color={Red=-1, Green=-1, Blue=-1}, LineThick=1}
//}
func (this *LuaGenerator) drawFunctionCodes() string {
	var drawTexts []string
	var drawIcons []string
	var drawLines []string
	var drawKLines []string
	var stickLines []string
	var ployLines []string

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
			drawTexts = append(drawTexts, fmt.Sprintf("        {ActionType=%d, Cond=o.%s, Price=o.%s, Text=o.%s, Color={Red=%d, Green=%d, Blue=%d}, NoDraw=%d}",
				formula.FORMULA_DRAW_ACTION_DRAWTEXT,
				expr.arguments[0].DefinedName(),
				expr.arguments[1].DefinedName(),
				expr.arguments[2].DefinedName(),
				color.Red, color.Green, color.Blue,
				noDraw))
		case "DRAWICON":
			drawIcons = append(drawIcons, fmt.Sprintf("        {ActionType=%d, Cond=o.%s, Price=o.%s, Type=%d, NoDraw=%d}",
				formula.FORMULA_DRAW_ACTION_DRAWICON,
				expr.arguments[0].DefinedName(),
				expr.arguments[1].DefinedName(),
				int(expr.arguments[2].(*constantexpr).value),
				noDraw))
		case "DRAWLINE":
			drawLines = append(drawLines, fmt.Sprintf("        {ActionType=%d, Cond1=o.%s, Price1=o.%s, Cond2=o.%s, Price2=o.%s, Expand=%d, NoDraw=%d, Color={Red=%d, Green=%d, Blue=%d}, LineThick=%d, VarIndex=%d}",
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
			drawKLines = append(drawKLines, fmt.Sprintf("        {ActionType=%d, High=o.%s, Open=o.%s, Low=o.%s, Close=o.%s, NoDraw=%d}",
				formula.FORMULA_DRAW_ACTION_DRAWKLINE,
				expr.arguments[0].DefinedName(),
				expr.arguments[1].DefinedName(),
				expr.arguments[2].DefinedName(),
				expr.arguments[3].DefinedName(),
				noDraw))
		case "STICKLINE":
			stickLines = append(stickLines, fmt.Sprintf("        {ActionType=%d, Cond=o.%s, Price1=o.%s, Price2=o.%s, Width=%f, Empty=%d, NoDraw=%d, Color={Red=%d, Green=%d, Blue=%d}, LineThick=%d}",
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
			ployLines = append(ployLines, fmt.Sprintf("        {ActionType=%d, Cond=o.%s, Price=o.%s, NoDraw=%d, Color={Red=%d, Green=%d, Blue=%d}, LineThick=%d, VarIndex=%d}",
				formula.FORMULA_DRAW_ACTION_PLOYLINE,
				expr.arguments[0].DefinedName(),
				expr.arguments[1].DefinedName(),
				noDraw,
				color.Red, color.Green, color.Blue,
				lineThick,
				varIndex))
		}
	}

	return fmt.Sprintf(`    o.drawTextActions = {
%s
    }

    o.drawIconActions = {
%s
    }

    o.drawLineActions = {
%s
    }

    o.drawKLineActions = {
%s
    }

    o.stickLineActions = {
%s
    }

    o.ployLineActions = {
%s
    }`, strings.Join(drawTexts, ",\n"),
		strings.Join(drawIcons, ",\n"),
		strings.Join(drawLines, ",\n"),
		strings.Join(drawKLines, ",\n"),
		strings.Join(stickLines, ",\n"),
		strings.Join(ployLines, ",\n"))
}

func (this *LuaGenerator) getCodes(indent string) string {
	lines := make([]string, len(this.context.outputVars))
	for i, varName := range this.context.outputVars {
		lines[i] = fmt.Sprintf("%so.%s.Get(index),", indent, varName)
	}
	return strings.Join(lines, "\n")
}

func (this *LuaGenerator) paramCodes() string {
	sa := make([]string, len(this.context.params)+1)
	sa[0] = "data"
	for i, p := range this.context.params {
		sa[i+1] = p
	}
	return strings.Join(sa, ", ")
}

func (this *LuaGenerator) paramNames() string {
	sa := make([]string, len(this.context.params))
	for i, p := range this.context.params {
		sa[i] = fmt.Sprintf("'%s'", p)
	}
	return strings.Join(sa, ", ")
}

func (this *LuaGenerator) paramMetaData(name string) string {
	sa := make([]string, len(this.context.params))
	for i, p := range this.context.params {
		exp := this.context.paramMap[p].(*paramexpr)
		sa[i] = fmt.Sprintf("%sClass['%s'] = {%f, %f, %f}", name, p, exp.defaultValue, exp.min, exp.max)
	}
	return strings.Join(sa, "\n")
}

func (this *LuaGenerator) GenerateCode(name string) string {
	name = strings.ToUpper(name)

	flags, graphTypes, lineThicks, colors, lineStyles := this.varProperties()

	const indent = "    "
	code := fmt.Sprintf(`-----------------------------------------------------------
-- GENERATED BY EASYLANG COMPILER.
-- !!!! DON'T MODIFY IT!!!!!!
-----------------------------------------------------------

Sequence = 0
MACDObjects = {}

function SetObject(id, obj)
    MACDObjects[id] = obj
end

function RemoveObject(id)
    MACDObjects[id] = nil
end

function GetObject(id)
    return MACDObjects[id]
end

%sClass = {}

%sClass['name'] = '%s'
%sClass['argName'] = {%s}
%s
%sClass['vars'] = {%s}
%sClass['flags'] = {%s}
%sClass['color'] = {%s}
%sClass['lineThick'] = {%s}
%sClass['lineStyle'] = {%s}
%sClass['graphType'] = {%s}

function %sClass:new(%s)
    o = {}
    setmetatable(o, self)
    self.__index = self
    o.%s = data
%s
%s
%s

%s

    o.ref_values = {%s}

    Sequence = Sequence + 1
    o.__id__ = Sequence
    SetObject(o.__id__, o)

    return o
end

function %sClass:updateLastValue()
%s
end

function %sClass:Len()
    return self.%s.Len()
end


function %sClass:Get(index)
    return {
%s
    }
end

FormulaClass = %sClass
	`, name,
		name,
		name,
		name,
		this.paramNames(),
		this.paramMetaData(name),
		name,
		this.varNames(),
		name,
		flags,
		name,
		colors,
		name,
		lineThicks,
		name,
		lineStyles,
		name,
		graphTypes,
		name,
		this.paramCodes(),
		getRefDataVarName("", ""),
		this.refDataDefineCodes(indent),
		this.refFormulaDefineCodes(indent),
		this.definedCodes(indent, name),
		this.drawFunctionCodes(),
		this.refValuesCodes(),
		name,
		this.updateLastValueCodes(indent, name),
		name,
		getRefDataVarName("", ""),
		name,
		this.getCodes("        "),
		name)

	return code
}
