package easylang

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"github.com/stephenlyu/goformula/formulalibrary/base/formula"
)

func CompileFile(sourceFile string, formulaManager formula.FormulaManager) (error, string) {
	file, err := os.Open(sourceFile)
	if err != nil {
		return err, ""
	}
	defer file.Close()

	_context = newContext()
	_context.SetFormulaManager(formulaManager)
	ret := yyParse(newLexer(bufio.NewReader(file)))
	if ret == 1 {
		return errors.New("compile failure"), ""
	}

	if _context.outputErrors() {
		return errors.New("compile failure"), ""
	}

	_context.Epilog()

	baseName := filepath.Base(sourceFile)
	mainName := strings.Split(baseName, ".")[0]

	return nil, _context.generateCode(mainName)
}

func Compile(sourceFile string, destFile string, formulaManager formula.FormulaManager) error {
	err, code := CompileFile(sourceFile, formulaManager)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(destFile, []byte(code), 0666)
}

func Tokenizer(sourceFile string) error {
	file, err := os.Open(sourceFile)
	if err != nil {
		return err
	}
	defer file.Close()

	lexer := newLexer(bufio.NewReader(file))
	lval := &yySymType{}
	for {
		char := lexer.Lex(lval)
		if char <= 0 {
			break
		}

		if char == NUM {
			fmt.Println(lval.value)
		} else {
			fmt.Println(lval.str)
		}
	}

	return nil
}
