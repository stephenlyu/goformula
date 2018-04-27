package easylang

import (
	"fmt"
	"github.com/stephenlyu/goformula/formulalibrary/base/formula"
	"strings"
)

var (
	PERIOD_MAP = map[string]string{
		"DAY":   "D1",
		"WEEK":  "W1",
		"MIN1":  "M1",
		"MIN5":  "M5",
		"MIN15": "M15",
		"MIN30": "M30",
		"MIN60": "M60",
	}

	CODE_MAP = map[string]int{
		"": 0,
	}
	DEBUG = false
)

func translatePeriod(period string) string {
	period1, ok := PERIOD_MAP[period]
	if ok {
		return period1
	}
	return ""
}

type context interface {
	newAnonymousVarName() string
	define(varName string, expr expression)
	defineParam(varName string, expr expression)
	refFormula(formulaName string, code string, period string)
	refData(code string, period string)
	addDrawFunction(expr *functionexpr)
	isDefined(varName string) bool
	isParamDefined(varName string) bool
	isNumberingVar() bool
}

func getCodeId(code string) int {
	ret, ok := CODE_MAP[code]
	if ok {
		return ret
	}
	CODE_MAP[code] = len(CODE_MAP)
	return CODE_MAP[code]
}

func formatCode(code string) string {
	return fmt.Sprintf("code%d", getCodeId(code))
}

func getRefFormulaVarName(name string, code string, period string) string {
	return fmt.Sprintf("formula_%s_%s_%s", formatCode(code), strings.ToLower(period), strings.ToLower(name))
}

func getRefDataVarName(code string, period string) string {
	return fmt.Sprintf("__data_code%d_%s__", getCodeId(code), strings.ToLower(period))
}

func getIndexMapVarName(code string, period string) string {
	return fmt.Sprintf("__index_map_code%d_%s__", getCodeId(code), strings.ToLower(period))
}

type RefFormula struct {
	name   string
	code   string
	period string
}

func (this *RefFormula) String() string {
	return getRefFormulaVarName(this.name, this.code, this.period)
}

type RefData struct {
	code   string
	period string
}

type Context struct {
	sequence int

	formulaManager formula.FormulaManager
	numberingVar   bool

	params   []string
	paramMap map[string]expression

	definedVars   []string
	definedVarMap map[string]expression

	refFormulas []RefFormula
	refDataList []RefData

	outputVars         []string
	outputDescriptions map[string][]string

	notOutputVars         []string
	notOutputDescriptions map[string][]string

	drawFunctions []*functionexpr

	// TODO: Handle errors
	errors []SyncError
}

func newContext() *Context {
	CODE_MAP = map[string]int{
		"": 0,
	}

	resetAll()
	return &Context{
		paramMap:              map[string]expression{},
		definedVarMap:         map[string]expression{},
		outputDescriptions:    map[string][]string{},
		notOutputDescriptions: map[string][]string{},
	}
}

func (this *Context) SetFormulaManager(formulaManager formula.FormulaManager) {
	this.formulaManager = formulaManager
}

func (this *Context) SetNumberingVar(v bool) {
	this.numberingVar = v
}

func (this *Context) isNumberingVar() bool {
	return this.numberingVar
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

func (this *Context) isReferenceSupport(formulaName string, refVarName string) bool {
	if this.formulaManager == nil {
		return true
	}
	formulaName = strings.ToUpper(formulaName)
	refVarName = strings.ToUpper(refVarName)
	return this.formulaManager.CanSupportVar(formulaName, refVarName)
}

func (this *Context) isPeriodSupport(period string) bool {
	if this.formulaManager == nil {
		return true
	}
	return this.formulaManager.CanSupportPeriod(period)
}

func (this *Context) isSecuritySupport(code string) bool {
	if this.formulaManager == nil {
		return true
	}
	return this.formulaManager.CanSupportSecurity(code)
}

func (this *Context) refFormula(formulaName string, code string, period string) {
	for _, r := range this.refFormulas {
		if r.name == formulaName && r.code == code && r.period == period {
			return
		}
	}
	this.refFormulas = append(this.refFormulas, RefFormula{formulaName, code, period})
}

func (this *Context) refData(code string, period string) {
	for _, r := range this.refDataList {
		if r.code == code && r.period == period {
			return
		}
	}
	this.refDataList = append(this.refDataList, RefData{code, period})
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
