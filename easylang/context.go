package easylang

import (
	"fmt"
	"github.com/stephenlyu/goformula/stockfunc/formula"
	"strconv"
	"strings"
)

type context interface {
	newAnonymousVarName() string
	define(varName string, expr expression)
	defineParam(varName string, expr expression)
	addDrawFunction(expr *functionexpr)
	isDefined(varName string) bool
	isParamDefined(varName string) bool
}

type Context struct {
	sequence int

	params   []string
	paramMap map[string]expression

	definedVars   []string
	definedVarMap map[string]expression

	outputVars         []string
	outputDescriptions map[string][]string

	notOutputVars         []string
	notOutputDescriptions map[string][]string

	drawFunctions		  []*functionexpr

	// TODO: Handle errors
	errors []SyncError
}

func newContext() *Context {
	return &Context{
		paramMap:              map[string]expression{},
		definedVarMap:         map[string]expression{},
		outputDescriptions:    map[string][]string{},
		notOutputDescriptions: map[string][]string{},
	}
}

func (this *Context) addError(err SyncError) {
	this.errors = append(this.errors, err)
}

func (this *Context) outputErrors() bool {
	for _, err := range this.errors {
		fmt.Println(err.String())
	}
	return len(this.errors) > 0
}

func (this *Context) newAnonymousVarName() string {
	ret := fmt.Sprintf("__anonymous_%d", this.sequence)
	this.sequence++
	return ret
}

func (this *Context) define(varName string, expr expression) {
	varName = strings.ToLower(varName)
	if expr, ok := this.definedVarMap[varName]; ok {
		expr.IncrementRefCount()
		return
	}
	this.definedVarMap[varName] = expr
	this.definedVars = append(this.definedVars, varName)
	expr.IncrementRefCount()
}

func (this *Context) defineParam(varName string, expr expression) {
	varName = strings.ToLower(varName)
	if _, ok := this.paramMap[varName]; ok {
		return
	}
	this.paramMap[varName] = expr
	this.params = append(this.params, varName)

	this.definedVarMap[varName] = expr
	this.definedVars = append(this.definedVars, varName)
	expr.IncrementRefCount()
}

func (this Context) isDefined(varName string) bool {
	_, ok := this.definedVarMap[varName]
	return ok
}

func (this Context) isParamDefined(varName string) bool {
	_, ok := this.paramMap[varName]
	return ok
}

func (this *Context) addOutput(varName string, descriptions []string, line, column int) {
	varName = strings.ToLower(varName)
	for _, v := range this.outputVars {
		if v == varName {
			panic("duplicate output variable")
		}
	}
	this.outputVars = append(this.outputVars, varName)
	this.outputDescriptions[varName] = descriptions
}

func (this *Context) addNotOutputVar(varName string, descriptions []string, line, column int) {
	varName = strings.ToLower(varName)
	for _, v := range this.notOutputVars {
		if v == varName {
			panic("duplicate output variable")
		}
	}
	this.notOutputVars = append(this.notOutputVars, varName)
	this.notOutputDescriptions[varName] = descriptions
}

func (this *Context) defined(varName string) expression {
	expr, _ := this.definedVarMap[strings.ToLower(varName)]
	return expr
}

func (this *Context) definedParam(varName string) expression {
	expr, _ := this.paramMap[strings.ToLower(varName)]
	return expr
}

func (this *Context) addDrawFunction(expr *functionexpr) {
	this.drawFunctions = append(this.drawFunctions, expr)
}

func (this *Context) definedCodes(indent string) string {
	var lines []string
	for _, varName := range this.definedVars {
		expr, ok := this.definedVarMap[varName]
		if !ok {
			continue
		}
		if !expr.IsValid() {
			continue
		}
		if expr.IsVoid() {
			// DO NOTHING
		} else {
			lines = append(lines, fmt.Sprintf("%so.%s = %s", indent, varName, expr.Codes()))
		}
	}
	return strings.Join(lines, "\n")
}

func (this *Context) updateLastValueCodes(indent string) string {
	lines := []string{}
	for _, varName := range this.definedVars {
		expr, ok := this.definedVarMap[varName]
		if !ok {
			continue
		}
		switch expr.(type) {
		case *constantexpr:
		case *assignexpr:
		case *paramexpr:
		case *stringexpr:
		default:
			if !expr.IsValid() || expr.IsVoid() {
				continue
			}
			lines = append(lines, fmt.Sprintf("%so.%s.updateLastValue()", indent, varName))
		}
	}
	return strings.Join(lines, "\n")
}

func (this *Context) refValuesCodes() string {
	items := make([]string, len(this.outputVars))
	for i, varName := range this.outputVars {
		items[i] = fmt.Sprintf("o.%s", varName)
	}
	return strings.Join(items, ", ")
}

func (this *Context) varNames() string {
	items := make([]string, len(this.outputVars))
	for i, varName := range this.outputVars {
		exp := this.definedVarMap[varName]
		items[i] = fmt.Sprintf("'%s'", exp.DisplayName())
	}
	return strings.Join(items, ", ")
}

func (this *Context) translateDescriptions(desciptions []string) (flag int, graphType int, lineThick int, color string, lineStyle int) {
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

func (this *Context) varProperties() (flags string, graphTypes string, lineThicks string, colors string, lineStyles string) {
	sFlags := make([]string, len(this.outputVars))
	sGraphTypes := make([]string, len(this.outputVars))
	sLineThicks := make([]string, len(this.outputVars))
	sLineStyles := make([]string, len(this.outputVars))
	sColors := make([]string, len(this.outputVars))

	for i, varName := range this.outputVars {
		descriptions := this.outputDescriptions[varName]
		flag, graphType, lineThick, color, lineStyle := this.translateDescriptions(descriptions)

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
func (this *Context) drawFunctionCodes() string {
	var drawTexts []string
	var drawIcons []string
	var drawLines []string
	var drawKLines []string
	var stickLines []string
	var ployLines []string

	relatedDescriptions := func (expr expression) []string {
		for varName := range this.outputDescriptions {
			iExpr := this.definedVarMap[varName]
			aExpr, ok := iExpr.(*assignexpr)
			if !ok {
				continue
			}
			if aExpr.operand == expr {
				return this.outputDescriptions[varName]
			}
		}

		for varName := range this.notOutputDescriptions {
			expr := this.definedVarMap[varName]
			aExpr, ok := expr.(*assignexpr)
			if !ok {
				continue
			}

			if aExpr.operand == expr {
				return this.notOutputDescriptions[varName]
			}
		}
		return []string{}
	}

	for _, expr := range this.drawFunctions {
		descriptions := relatedDescriptions(expr)
		flag, _, lineThick, colorStr, _ := this.translateDescriptions(descriptions)

		var color *formula.Color
		if colorStr == "" {
			color = &formula.Color{Red: -1, Green: - 1, Blue: -1}
		} else {
			color = ParseColorLiteral(colorStr)
		}

		var noDraw int
		if flag & formula.FORMULA_VAR_FLAG_NO_DRAW != 0 {
			noDraw = 1
		}

		switch expr.funcName {
		case "DRAWTEXT":
			drawTexts = append(drawTexts, fmt.Sprintf("        {Cond=o.%s, Price=o.%s, Text=o.%s, Color={Red=%d, Green=%d, Blue=%d}, NoDraw=%d}",
				expr.arguments[0].VarName(),
				expr.arguments[1].VarName(),
				expr.arguments[2].VarName(),
				color.Red, color.Green, color.Green,
				noDraw))
		case "DRAWICON":
			drawIcons = append(drawIcons, fmt.Sprintf("        {Cond=o.%s, Price=o.%s, Type=%d, NoDraw=%d}",
				expr.arguments[0].VarName(),
				expr.arguments[1].VarName(),
				int(expr.arguments[2].(*constantexpr).value),
				noDraw))
		case "DRAWLINE":
			drawLines = append(drawLines, fmt.Sprintf("        {Cond1=o.%s, Price1=o.%s, Cond2=o.%s, Price2=o.%s, Expand=%d, NoDraw=%d, Color={Red=%d, Green=%d, Blue=%d}, LineThick=%d}",
				expr.arguments[0].VarName(),
				expr.arguments[1].VarName(),
				expr.arguments[2].VarName(),
				expr.arguments[3].VarName(),
				int(expr.arguments[4].(*constantexpr).value),
				noDraw,
				color.Red, color.Green, color.Green,
				lineThick))
		case "DRAWKLINE":
			drawKLines = append(drawKLines, fmt.Sprintf("        {High=o.%s, Open=o.%s, Low=o.%s, Close=o.%s, NoDraw=%d}",
				expr.arguments[0].VarName(),
				expr.arguments[1].VarName(),
				expr.arguments[2].VarName(),
				expr.arguments[3].VarName(),
				noDraw))
		case "STICKLINE":
			stickLines = append(stickLines, fmt.Sprintf("        {Cond=o.%s, Price1=o.%s, Price2=o.%s, Width=%f, Empty=%d, NoDraw=%d, Color={Red=%d, Green=%d, Blue=%d}, LineThick=%d}",
				expr.arguments[0].VarName(),
				expr.arguments[1].VarName(),
				expr.arguments[2].VarName(),
				expr.arguments[3].(*constantexpr).value,
				int(expr.arguments[4].(*constantexpr).value),
				noDraw,
				color.Red, color.Green, color.Green,
				lineThick))
		case "PLOYLINE":
			ployLines = append(ployLines, fmt.Sprintf("        {Cond=o.%s, Price=o.%s, NoDraw=%d, Color={Red=%d, Green=%d, Blue=%d}, LineThick=%d}",
				expr.arguments[0].VarName(),
				expr.arguments[1].VarName(),
				noDraw,
				color.Red, color.Green, color.Green,
				lineThick))
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

func (this *Context) getCodes(indent string) string {
	lines := make([]string, len(this.outputVars))
	for i, varName := range this.outputVars {
		lines[i] = fmt.Sprintf("%so.%s.Get(index),", indent, varName)
	}
	return strings.Join(lines, "\n")
}

func (this *Context) paramCodes() string {
	sa := make([]string, len(this.params)+1)
	sa[0] = "data"
	for i, p := range this.params {
		sa[i+1] = p
	}
	return strings.Join(sa, ", ")
}

func (this *Context) paramNames() string {
	sa := make([]string, len(this.params))
	for i, p := range this.params {
		sa[i] = fmt.Sprintf("'%s'", p)
	}
	return strings.Join(sa, ", ")
}

func (this *Context) paramMetaData(name string) string {
	sa := make([]string, len(this.params))
	for i, p := range this.params {
		exp := this.paramMap[p].(*paramexpr)
		sa[i] = fmt.Sprintf("%sClass['%s'] = {%f, %f, %f}", name, p, exp.defaultValue, exp.min, exp.max)
	}
	return strings.Join(sa, "\n")
}

func (this *Context) removeUnusedParams() {
}

func (this *Context) Epilog() {
	var outputVars []string
	for _, varName := range this.outputVars {
		exp := this.definedVarMap[varName]
		if !exp.IsValid() {
			continue
		}
		outputVars = append(outputVars, varName)
	}
	this.outputVars = outputVars
}

func (this *Context) generateCode(name string) string {
	name = strings.ToUpper(name)

	this.removeUnusedParams()

	flags, graphTypes, lineThicks, colors, lineStyles := this.varProperties()

	const indent = "    "
	code := fmt.Sprintf(`-----------------------------------------------------------
-- GENERATED BY EASYLANG COMPILER.
-- !!!! DON'T MODIFY IT!!!!!!
-----------------------------------------------------------

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

    o.data = data
%s

%s

    o.ref_values = {%s}
    return o
end

function %sClass:updateLastValue()
%s
end

function %sClass:Len()
    return self.data.Len()
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
		this.definedCodes(indent),
		this.drawFunctionCodes(),
		this.refValuesCodes(),
		name,
		this.updateLastValueCodes(indent),
		name,
		name,
		this.getCodes("        "),
		name)

	return code
}
