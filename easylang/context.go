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

func (this *Context) definedCodes(indent string) string {
	lines := make([]string, len(this.definedVarMap))
	i := 0
	for _, varName := range this.definedVars {
		expr, ok := this.definedVarMap[varName]
		if !ok {
			continue
		}
		lines[i] = fmt.Sprintf("%so.%s = %s", indent, varName, expr.Codes())
		i++
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
		default:
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
		this.refValuesCodes(),
		name,
		this.updateLastValueCodes(indent),
		name,
		name,
		this.getCodes("        "),
		name)

	return code
}
