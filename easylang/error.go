package easylang

import "fmt"

type SyncError interface {
	String() string
}

type synxerror struct {
	line   int
	column int
}

type undefinedVarError struct {
	synxerror
	varName string
}

func UndefinedVarError(line, column int, varName string) *undefinedVarError {
	return &undefinedVarError{
		synxerror: synxerror{
			line: line, column: column,
		},
		varName: varName,
	}
}

func (this undefinedVarError) String() string {
	return fmt.Sprintf("undefined variable '%s' at line %d column %d", this.varName, this.line, this.column)
}

type undefinedFunctionError struct {
	synxerror
	varName string
}

func UndefinedFunctionError(line, column int, varName string) *undefinedFunctionError {
	return &undefinedFunctionError{
		synxerror: synxerror{
			line: line, column: column,
		},
		varName: varName,
	}
}

func (this undefinedFunctionError) String() string {
	return fmt.Sprintf("undefined function '%s' at line %d column %d", this.varName, this.line, this.column)
}

type badGraphDescError struct {
	synxerror
	desc string
}

func BadGraphDescError(line, column int, desc string) *badGraphDescError {
	return &badGraphDescError{
		synxerror: synxerror{
			line: line, column: column,
		},
		desc: desc,
	}
}

func (this badGraphDescError) String() string {
	return fmt.Sprintf("bad graph description '%s' at line %d column %d", this.desc, this.line, this.column)
}
