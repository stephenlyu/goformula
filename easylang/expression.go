package easylang

import (
	"fmt"
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
	"DRAWLINE":  "DRAWLINE",
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
}

var (
	CONST_SEQ = 1
	VAR_SEQ   = 1
)

func newConstName() string {
	ret := fmt.Sprintf("const%d", CONST_SEQ)
	CONST_SEQ++
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
	first := true
	valid := true
	s = strings.ToLower(s)
	b := []byte(s)
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
		return newVarName()
	}
	return s
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

func (this *baseexpr) IncrementRefCount() {
	this.refCount++
}

func (this *baseexpr) DecrementRefCount() {
	this.refCount--
}

func (this *baseexpr) RefCount() int {
	return this.refCount
}

type constantexpr struct {
	baseexpr
	value float64
}

func ConstantExpression(context context, value float64) *constantexpr {
	ret := &constantexpr{
		baseexpr: baseexpr{
			context: context,
			varName: newConstName(),
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

func (this unaryexpr) Codes() string {
	funcName, _ := unaryFuncMap[this.operator]
	return fmt.Sprintf("%s(o.%s)", funcName, this.operand.VarName())
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

func (this binaryexpr) Codes() string {
	funcName, _ := binaryFuncMap[this.operator]
	return fmt.Sprintf("%s(o.%s, o.%s)", funcName, this.leftOperand.VarName(), this.rightOperand.VarName())
}

type functionexpr struct {
	baseexpr
	funcName  string
	arguments []expression
}

func FunctionExpression(context context, funcName string, arguments []expression) *functionexpr {
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
	} else {
		ret.displayName = funcName
	}
	ret.varName = ret.formatVarName(ret.displayName)

	context.define(ret.varName, ret)
	return ret
}

func (this functionexpr) Codes() string {
	if len(this.arguments) > 0 {
		sa := make([]string, len(this.arguments))
		for i, arg := range this.arguments {
			sa[i] = "o." + arg.VarName()
		}
		return fmt.Sprintf("%s(%s)", this.funcName, strings.Join(sa, ", "))
	} else {
		return fmt.Sprintf("%s(o.data)", this.funcName)
	}
}

type assignexpr struct {
	baseexpr
	operand expression
}

func AssignmentExpression(context context, varName string, operand expression) *assignexpr {
	ret := &assignexpr{
		baseexpr: baseexpr{
			context:     context,
			displayName: varName,
		},
		operand: operand,
	}
	ret.varName = ret.formatVarName(ret.displayName)
	if context.isDefined(ret.varName) {
		panic("duplicate definition")
	}
	context.define(ret.varName, ret)
	return ret
}

func (this assignexpr) Codes() string {
	return fmt.Sprintf("o.%s", this.operand.VarName())
}

type paramexpr struct {
	baseexpr
	operand expression
}

func ParamExpression(context context, varName string, operand expression) *paramexpr {
	ret := &paramexpr{
		baseexpr: baseexpr{
			context:     context,
			displayName: varName,
		},
		operand: operand,
	}
	ret.varName = ret.formatVarName(ret.displayName)
	if context.isParamDefined(ret.varName) {
		panic("duplicate param")
	}
	context.defineParam(ret.varName, ret)
	return ret
}

func (this paramexpr) Codes() string {
	return fmt.Sprintf("Scalar(%s)", strings.ToLower(this.varName))
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
