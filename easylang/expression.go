package easylang

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strings"
	"unicode/utf8"
)

type funcmap map[string]string
type opmap map[string]string

var unaryFuncMap = funcmap{
	"-":   "MINUS",
	"NOT": "NOT",
}

var binaryFuncMap = funcmap{
	"+":   "ADD",
	"-":   "SUB",
	"*":   "MUL",
	"/":   "DIV",
	"<":   "LT",
	"<=":  "LE",
	">":   "GT",
	">=":  "GE",
	"=":   "EQ",
	"!=":  "NEQ",
	"AND": "AND",
	"OR":  "OR",
}

var noArgFuncMap = funcmap{
	"C":      "CLOSE",
	"O":      "OPEN",
	"L":      "LOW",
	"H":      "HIGH",
	"A":      "AMOUNT",
	"V":      "VOLUME",
	"VOL":    "VOLUME",
	"CLOSE":  "CLOSE",
	"OPEN":   "OPEN",
	"LOW":    "LOW",
	"HIGH":   "HIGH",
	"AMOUNT": "AMOUNT",
	"VOLUME": "VOLUME",
	"PERIOD": "PERIOD",

	"ISLASTBAR": "ISLASTBAR",
}

var funcMap = funcmap{
	"REF":       "REF",
	"BARSCOUNT": "BARSCOUNT",
	"BARSLAST":  "BARSLAST",
	"HHV":       "HHV",
	"LLV":       "LLV",
	"HHVBARS":   "HHVBARS",
	"LLVBARS":   "LLVBARS",
	"ROUND2":    "ROUND2",
	"IF":        "IF",
	"EMA":       "EMA",
	"MA":        "MA",
	"SMA":       "SMA",
	"DMA":       "DMA",
	"EXPMEMA":   "EXPMEMA",
	"COUNT":     "COUNT",
	"EVERY":     "EVERY",
	"CROSS":     "CROSS",
	"MIN":       "MIN",
	"MAX":       "MAX",
	"ABS":       "ABS",
	"AVEDEV":    "AVEDEV",
	"STD":       "STD",
	"SUM":       "SUM",

	// 绘制函数
	"DRAWTEXT":  "DRAWTEXT",
	"DRAWLINE":  "DRAWLINE",
	"PLOYLINE":  "PLOYLINE",
	"DRAWICON":  "DRAWICON",
	"DRAWKLINE": "DRAWKLINE",
	"STICKLINE": "STICKLINE",
}

var voidFuncMap = map[string]bool{
	"DRAWTEXT":  true,
	"DRAWICON":  true,
	"DRAWKLINE": true,
	"STICKLINE": true,
}

var drawFuncMap = map[string]bool{
	"DRAWTEXT":  true,
	"DRAWLINE":  true,
	"PLOYLINE":  true,
	"DRAWICON":  true,
	"DRAWKLINE": true,
	"STICKLINE": true,
}

var (
	CONST_SEQ  = 1
	STRING_SEQ = 1
	VAR_SEQ    = 1
)

var (
	CONST_VAL_NAMES     = make(map[float64]string)
	FORMAT_VAR_NAMES    = make(map[string]string)
	DEFINED_NAMES       = make(map[string]string)
	VAR_NAME_PATTERN, _ = regexp.Compile(`^var[0-9]+$`)
)

func resetAll() {
	CONST_SEQ = 1
	STRING_SEQ = 1
	VAR_SEQ = 1

	CONST_VAL_NAMES = make(map[float64]string)
	FORMAT_VAR_NAMES = make(map[string]string)
	DEFINED_NAMES = make(map[string]string)
}

func dumpVarMapping(filePath string) {
	varRmap := make(map[string]string)
	vars := []string{}
	for k, v := range DEFINED_NAMES {
		varRmap[v] = k
		vars = append(vars, v)
	}

	sort.SliceStable(vars, func(i, j int) bool {
		return vars[i] < vars[j]
	})

	lines := make([]string, len(vars))
	for i, v := range vars {
		lines[i] = fmt.Sprintf("%s: %s", v, varRmap[v])
	}

	ioutil.WriteFile(filePath, []byte(strings.Join(lines, "\n")), 0666)
}

func newConstName(value float64) string {
	ret, ok := CONST_VAL_NAMES[value]
	if ok {
		return ret
	}

	ret = fmt.Sprintf("const%d", CONST_SEQ)
	CONST_SEQ++

	CONST_VAL_NAMES[value] = ret

	return ret
}

func newStringName() string {
	ret := fmt.Sprintf("string%d", STRING_SEQ)
	STRING_SEQ++
	return ret
}

func newVarName() string {
	ret := fmt.Sprintf("var%d", VAR_SEQ)
	VAR_SEQ++
	return ret
}

type expression interface {
	Codes() string
	IncrementRefCount()
	DecrementRefCount()
	RefCount() int
	VarName() string
	DisplayName() string
	DefinedName() string
	DefineName(v string) string

	IsVoid() bool  // DRAWTEXT等没有返回值
	IsValid() bool // 如果一个表达式的子表示IsVoid()或者!IsValid()，则该表达式不合法。不合法的表达式再生成代码过程中会被忽略
}

type baseexpr struct {
	context     context
	refCount    int
	varName     string
	displayName string
}

func isAlpha(c rune) bool {
	return c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z'
}

func isDigit(c rune) bool {
	return c >= '0' && c <= '9'
}

func (this baseexpr) formatVarName(s string) string {
	ret, ok := FORMAT_VAR_NAMES[s]
	if ok {
		return ret
	}

	first := true
	valid := true
	ret = strings.ToLower(s)
	b := []byte(ret)
	for len(b) > 0 {
		c, n := utf8.DecodeRune(b)

		if first {
			if c != '_' && !isAlpha(c) {
				valid = false
				break
			}
			first = false
		} else {
			if c != '_' && !isAlpha(c) && !isDigit(c) {
				valid = false
				break
			}
		}

		b = b[n:]
	}
	if !valid {
		ret = newVarName()
		fmt.Println("newVarName", s, ret)
	}
	FORMAT_VAR_NAMES[s] = ret
	return ret
}

func (this baseexpr) Codes() string {
	return ""
}

func (this baseexpr) VarName() string {
	return this.varName
}

func (this baseexpr) DisplayName() string {
	return this.displayName
}

func (this baseexpr) DefinedName() string {
	return this.varName
}

func (this baseexpr) DefineName(v string) string {
	if !this.context.isNumberingVar() {
		return v
	}

	if VAR_NAME_PATTERN.Match([]byte(v)) {
		return v
	}

	ret, ok := DEFINED_NAMES[v]
	if ok {
		return ret
	}
	ret = newVarName()
	DEFINED_NAMES[v] = ret
	return ret
}

func (this *baseexpr) IncrementRefCount() {
	this.refCount++
}

func (this *baseexpr) DecrementRefCount() {
	this.refCount--
}

func (this *baseexpr) RefCount() int {
	return this.refCount
}

func (this *baseexpr) IsVoid() bool {
	return false
}

func (this *baseexpr) IsValid() bool {
	return true
}

type constantexpr struct {
	baseexpr
	value float64
}

func ConstantExpression(context context, value float64) *constantexpr {
	ret := &constantexpr{
		baseexpr: baseexpr{
			context: context,
			varName: newConstName(value),
		},
		value: value,
	}
	ret.displayName = fmt.Sprintf("const%f", value)
	context.define(ret.varName, ret)
	return ret
}

func (this constantexpr) Codes() string {
	return fmt.Sprintf("Scalar(%f)", this.value)
}

type stringexpr struct {
	baseexpr
	value string
}

func StringExpression(context context, value string) *stringexpr {
	ret := &stringexpr{
		baseexpr: baseexpr{
			context: context,
			varName: newStringName(),
		},
		value: value,
	}
	ret.displayName = fmt.Sprintf("string%s", value)
	context.define(ret.varName, ret)
	return ret
}

func (this stringexpr) Codes() string {
	return fmt.Sprintf("'%s'", this.value)
}

type unaryexpr struct {
	baseexpr
	operator string
	operand  expression
}

func UnaryExpression(context context, operator string, operand expression) *unaryexpr {
	ret := &unaryexpr{
		baseexpr: baseexpr{
			context: context,
		},
		operator: operator,
		operand:  operand,
	}
	opName, ok := unaryFuncMap[operator]
	if !ok {
		panic("Bad operator " + operator)
	}

	opName = strings.ToLower(opName)
	ret.displayName = fmt.Sprintf("%s_%s", opName, operand.VarName())
	ret.varName = ret.formatVarName(ret.displayName)
	context.define(ret.varName, ret)
	return ret
}

func (this unaryexpr) DefinedName() string {
	return this.DefineName(this.varName)
}

func (this unaryexpr) Codes() string {
	funcName, _ := unaryFuncMap[this.operator]
	return fmt.Sprintf("%s(o.%s)", funcName, this.operand.DefinedName())
}

func (this *unaryexpr) IsValid() bool {
	return this.operand.IsValid() && !this.operand.IsVoid()
}

type binaryexpr struct {
	baseexpr
	operator     string
	leftOperand  expression
	rightOperand expression
}

func BinaryExpression(context context, operator string, leftOperand, rightOperand expression) *binaryexpr {
	ret := &binaryexpr{
		baseexpr: baseexpr{
			context: context,
		},
		operator:     operator,
		leftOperand:  leftOperand,
		rightOperand: rightOperand,
	}
	opName, ok := binaryFuncMap[operator]
	if !ok {
		panic("Bad operator " + operator)
	}

	opName = strings.ToLower(opName)
	ret.displayName = fmt.Sprintf("%s_%s_%s", leftOperand.VarName(), opName, rightOperand.VarName())
	ret.varName = ret.formatVarName(ret.displayName)
	context.define(ret.varName, ret)
	return ret
}

func (this binaryexpr) DefinedName() string {
	return this.DefineName(this.varName)
}

func (this binaryexpr) Codes() string {
	funcName, _ := binaryFuncMap[this.operator]
	return fmt.Sprintf("%s(o.%s, o.%s)", funcName, this.leftOperand.DefinedName(), this.rightOperand.DefinedName())
}

func (this *binaryexpr) IsValid() bool {
	return this.leftOperand.IsValid() && !this.leftOperand.IsVoid() && this.rightOperand.IsValid() && !this.rightOperand.IsVoid()
}

type functionexpr struct {
	baseexpr
	funcName  string
	arguments []expression
}

func FunctionExpression(context context, funcName string, arguments []expression) *functionexpr {
	funcName = strings.ToUpper(funcName)
	ret := &functionexpr{
		baseexpr: baseexpr{
			context: context,
		},
		funcName:  funcName,
		arguments: arguments,
	}
	if len(arguments) > 0 {
		sa := make([]string, len(arguments)+1)
		sa[0] = funcName

		for i, arg := range arguments {
			sa[i+1] = arg.VarName()
		}
		ret.displayName = strings.Join(sa, "_")
		ret.varName = ret.formatVarName(ret.displayName)
	} else {
		ret.displayName = funcName
		ret.varName = fmt.Sprintf("__no_arg_func_%s", strings.ToLower(funcName))
	}

	context.define(ret.varName, ret)

	if drawFuncMap[funcName] {
		context.addDrawFunction(ret)
	}
	return ret
}

func (this functionexpr) DefinedName() string {
	return this.DefineName(this.varName)
}

func (this functionexpr) Codes() string {
	if len(this.arguments) > 0 {
		sa := make([]string, len(this.arguments))
		for i, arg := range this.arguments {
			sa[i] = "o." + arg.DefinedName()
		}
		return fmt.Sprintf("%s(%s)", this.funcName, strings.Join(sa, ", "))
	} else {
		return fmt.Sprintf("%s(o.%s)", this.funcName, getRefDataVarName("", ""))
	}
}

func (this functionexpr) IsVoid() bool {
	return voidFuncMap[this.funcName]
}

func (this *functionexpr) IsValid() bool {
	for _, e := range this.arguments {
		if !e.IsValid() || e.IsVoid() {
			return false
		}
	}
	return true
}

type referenceexpr struct {
	baseexpr
	formulaName string
	refVarName  string
}

func ReferenceExpression(context context, formulaName string, refVarName string) *referenceexpr {
	formulaName = strings.ToUpper(formulaName)
	context.refFormula(formulaName, "", "")

	ret := &referenceexpr{
		baseexpr: baseexpr{
			context: context,
		},
		formulaName: formulaName,
		refVarName:  refVarName,
	}
	varName := fmt.Sprintf("%s.%s", formulaName, refVarName)
	ret.varName = ret.formatVarName(varName)
	ret.displayName = varName
	context.define(ret.varName, ret)
	return ret
}

func (this referenceexpr) DefinedName() string {
	return this.DefineName(this.varName)
}

func (this referenceexpr) Codes() string {
	return fmt.Sprintf("o.%s.GetVarValue('%s')", getRefFormulaVarName(this.formulaName, "", ""), strings.ToUpper(this.refVarName))
}

type crossreferenceexpr struct {
	baseexpr
	formulaName string
	refVarName  string
	code        string
	period      string
}

func CrossReferenceExpression(context context, formulaName string, refVarName string, code string, period string) *crossreferenceexpr {
	context.refData(code, period)
	context.refFormula(formulaName, code, period)
	ret := &crossreferenceexpr{
		baseexpr: baseexpr{
			context: context,
		},
		formulaName: formulaName,
		refVarName:  refVarName,
		code:        code,
		period:      period,
	}
	varName := fmt.Sprintf("%s.%s_code%d_%s", formulaName, refVarName, getCodeId(code), strings.ToLower(period))
	ret.varName = varName
	ret.displayName = varName
	context.define(ret.varName, ret)
	return ret
}

func (this crossreferenceexpr) DefinedName() string {
	return this.DefineName(this.varName)
}

func (this crossreferenceexpr) Codes() string {
	return fmt.Sprintf("CrossValue(o.%s.GetVarValue('%s'), o.%s)",
		getRefFormulaVarName(this.formulaName, this.code, this.period),
		strings.ToUpper(this.refVarName),
		getIndexMapVarName(this.code, this.period))
}

type crossfunctionexpr struct {
	baseexpr
	funcName string
	code     string
	period   string
}

func CrossFunctionExpression(context context, funcName string, code string, period string) *crossfunctionexpr {
	context.refData(code, period)
	funcName = strings.ToUpper(funcName)
	ret := &crossfunctionexpr{
		baseexpr: baseexpr{
			context: context,
		},
		funcName: funcName,
		code:     code,
		period:   period,
	}
	varName := fmt.Sprintf("%s_code%d_%s", strings.ToLower(funcName), getCodeId(code), strings.ToLower(period))
	ret.varName = varName
	ret.displayName = varName
	context.define(ret.varName, ret)
	return ret
}

func (this crossfunctionexpr) DefinedName() string {
	return this.DefineName(this.varName)
}

func (this crossfunctionexpr) Codes() string {
	return fmt.Sprintf("CrossValue(o.%s(%s), o.%s)",
		this.funcName,
		getRefDataVarName(this.code, this.period),
		getIndexMapVarName(this.code, this.period))
}

type assignexpr struct {
	baseexpr
	operand expression
}

func AssignmentExpression(context context, varName string, operand expression, isAnonymous bool) *assignexpr {
	ret := &assignexpr{
		baseexpr: baseexpr{
			context: context,
		},
		operand: operand,
	}
	ret.varName = ret.formatVarName(varName)
	if !isAnonymous {
		ret.displayName = varName
	}
	if context.isDefined(ret.varName) {
		panic(fmt.Sprintf("duplicate definition %s", ret.varName))
	}
	context.define(ret.varName, ret)
	return ret
}

func (this assignexpr) Codes() string {
	return fmt.Sprintf("o.%s", this.operand.DefinedName())
}

func (this assignexpr) IsValid() bool {
	return this.operand.IsValid() && !this.operand.IsVoid()
}

type paramexpr struct {
	baseexpr
	defaultValue, min, max float64
}

func ParamExpression(context context, varName string, defaultValue float64, min float64, max float64) *paramexpr {
	ret := &paramexpr{
		baseexpr: baseexpr{
			context:     context,
			displayName: varName,
		},
		defaultValue: defaultValue,
		min:          min,
		max:          max,
	}
	ret.varName = ret.formatVarName(ret.displayName)
	if context.isParamDefined(ret.varName) {
		panic(fmt.Sprintf("duplicate param %s", ret.varName))
	}
	context.defineParam(ret.varName, ret)
	return ret
}

func (this paramexpr) Codes() string {
	return fmt.Sprintf("Scalar(%s)", strings.ToLower(this.DefinedName()))
}

type errorexpr struct {
	baseexpr
}

func ErrorExpression(context context, varName string) *errorexpr {
	return &errorexpr{
		baseexpr{
			context: context,
			varName: strings.ToLower(varName),
		},
	}
}
