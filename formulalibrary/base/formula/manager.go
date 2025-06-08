package formula

import "github.com/stephenlyu/goformula/stockfunc/function"

type FormulaManager interface {
	// 是否支持名为name的公式
	CanSupportVar(name string, varName string) bool

	// 是否支持周期
	CanSupportPeriod(period string) bool

	// 是否支持证券品种
	CanSupportSecurity(code string) bool

	// 使用默认参数创建公式
	NewFormula(name string, data function.RVectorReader) Formula

	// 使用指定参数创建公式
	NewFormulaWithArgs(name string, data function.RVectorReader, args []float64) Formula
}
