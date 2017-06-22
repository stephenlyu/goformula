package easylang

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func Compile(sourceFile string, destFile string) error {
	file, err := os.Open(sourceFile)
	if err != nil {
		return err
	}
	defer file.Close()

	ret := yyParse(newLexer(bufio.NewReader(file)))
	if ret == 1 {
		return errors.New("compile failure")
	}

	if _context.outputErrors() {
		return errors.New("compile failure")
	}

	baseName := filepath.Base(sourceFile)
	mainName := strings.Split(baseName, ".")[0]

	err = _context.generateCode(mainName, destFile)

	return nil
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
