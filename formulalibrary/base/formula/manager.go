package formula

import "github.com/stephenlyu/goformula/stockfunc/function"

type FormulaManager interface {
	// 是否支持名为name的公式
	CanSupportVar(name string, varName string) bool

	// 使用默认参数创建公式
	NewFormula(name string, data *function.RVector) Formula

	// 使用指定参数创建公式
	NewFormulaWithArgs(name string, data *function.RVector, args []float64) Formula
}
