package luafunc

import (
	"github.com/stevedonovan/luar"
	"github.com/stephenlyu/goformula/function"
	stockfunc "github.com/stephenlyu/goformula/stockfunc/function"
)

var functionMap luar.Map = luar.Map{
	// Scalar & Vector
	"Scalar": function.Scalar,
	"Vector": function.Vector,

	// Op functions
	"NOT": function.NOT,
	"MINUS": function.MINUS,
	"ADD": function.ADD,
	"SUB": function.SUB,
	"MUL": function.MUL,
	"DIV": function.DIV,
	"GT": function.GT,
	"GE": function.GE,
	"LT": function.LT,
	"LE": function.LE,
	"EQ": function.EQ,
	"NEQ": function.NEQ,
	"AND": function.AND,
	"OR": function.OR,

	// MA functions

	"MA": function.MA,
	"SMA": function.SMA,
	"DMA": function.DMA,
	"EMA": function.EMA,
	"EXPEMA": function.EXPMEMA,

	// Stat functionA

	"LLV": function.LLV,
	"LLVBARS": function.LLVBARS,
	"HHV": function.HHV,
	"HHVBARS": function.HHVBARS,
	"STD": function.STD,
	"AVEDEV": function.AVEDEV,
	"SUM": function.SUM,
	"CROSS": function.CROSS,
	"COUNT": function.COUNT,
	"IF": function.IF,
	"EVERY": function.EVERY,
	"BARSLAST": function.BARSLAST,
	"BARSCOUNT": function.BARSCOUNT,
	"ISLASTBAR": function.ISLASTBAR,
	"ROUND2": function.ROUND2,
	"REF": function.REF,
	"MIN": function.MIN,
	"MAX": function.MAX,
	"ABS": function.ABS,
	"SLOPE": function.SLOPE,

	// Stock functions

	"RecordVector": stockfunc.RecordVector,
	"OPEN": stockfunc.OPEN,
	"CLOSE": stockfunc.CLOSE,
	"LOW": stockfunc.LOW,
	"HIGH": stockfunc.HIGH,
	"AMOUNT": stockfunc.AMOUNT,
	"VOLUME": stockfunc.VOLUME,
}

func GetFunctionMap(inMap luar.Map) luar.Map {
	result := luar.Map{}
	if inMap != nil {
		for k, v := range inMap {
			result[k] = v
		}
	}

	for k, v := range functionMap {
		result[k] = v
	}
	return result
}
