// CAUTION: Generated file - DO NOT EDIT.

package easylang

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"strconv"
	"unicode/utf8"
)

var keywords map[string]int = map[string]int{
	"AND": AND,
	"OR":  OR,
	"NOT": NOT,
}

type yylexer struct {
	src     *bufio.Reader
	buf     *bytes.Buffer
	empty   bool
	current rune

	lineno int
	column int
}

func addRune(b *bytes.Buffer, c rune) {
	if _, err := b.WriteRune(c); err != nil {
		log.Fatalf("WriteRune: %s", err)
	}
}

func newLexer(src *bufio.Reader) (y *yylexer) {
	y = &yylexer{src: src, buf: &bytes.Buffer{}, lineno: 1, column: 0}
	if r, _, err := src.ReadRune(); err == nil {
		y.current = r
	}
	return
}

func (y *yylexer) getc() rune {
	if y.current != 0 {
		addRune(y.buf, y.current)
	}
	y.current = 0
	if r, _, err := y.src.ReadRune(); err == nil {
		y.current = r
	}
	return y.current
}

func (y yylexer) Error(e string) {
	fmt.Printf("Error: %s at %d:%d\n", e, y.lineno, y.column)
}

func (y *yylexer) Lex(lval *yySymType) int {
	var err error
	c := y.current
	if y.empty {
		c, y.empty = y.getc(), false
	}

yystate0:

	y.column += utf8.RuneCount(y.buf.Bytes())
	y.buf.Reset()

	goto yystart1

	goto yystate0 // silence unused label error
	goto yystate1 // silence unused label error
yystate1:
	c = y.getc()
yystart1:
	switch {
	default:
		goto yyabort
	case c == '!':
		goto yystate6
	case c == '&':
		goto yystate8
	case c == '(':
		goto yystate12
	case c == ')':
		goto yystate13
	case c == '*':
		goto yystate14
	case c == '+':
		goto yystate15
	case c == ',':
		goto yystate16
	case c == '-':
		goto yystate17
	case c == '/':
		goto yystate18
	case c == ':':
		goto yystate22
	case c == ';':
		goto yystate24
	case c == '<':
		goto yystate25
	case c == '=':
		goto yystate27
	case c == '>':
		goto yystate29
	case c == '\'':
		goto yystate10
	case c == '\n':
		goto yystate3
	case c == '\r':
		goto yystate4
	case c == '\t' || c == ' ':
		goto yystate2
	case c == '{':
		goto yystate32
	case c == '|':
		goto yystate34
	case c >= '0' && c <= '9':
		goto yystate19
	case c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c >= 0x4e00 && c <= 0x9fcb:
		goto yystate31
	}

yystate2:
	c = y.getc()
	switch {
	default:
		goto yyrule3
	case c == '\t' || c == ' ':
		goto yystate2
	}

yystate3:
	c = y.getc()
	goto yyrule2

yystate4:
	c = y.getc()
	switch {
	default:
		goto yyrule2
	case c == '\n':
		goto yystate5
	}

yystate5:
	c = y.getc()
	goto yyrule1

yystate6:
	c = y.getc()
	switch {
	default:
		goto yyabort
	case c == '=':
		goto yystate7
	}

yystate7:
	c = y.getc()
	goto yyrule19

yystate8:
	c = y.getc()
	switch {
	default:
		goto yyabort
	case c == '&':
		goto yystate9
	}

yystate9:
	c = y.getc()
	goto yyrule20

yystate10:
	c = y.getc()
	switch {
	default:
		goto yyabort
	case c == '\'':
		goto yystate11
	case c >= '\x01' && c <= '&' || c >= '(' && c <= 'ÿ' || c >= 0x4e00 && c <= 0x9fcb:
		goto yystate10
	}

yystate11:
	c = y.getc()
	goto yyrule7

yystate12:
	c = y.getc()
	goto yyrule24

yystate13:
	c = y.getc()
	goto yyrule25

yystate14:
	c = y.getc()
	goto yyrule11

yystate15:
	c = y.getc()
	goto yyrule9

yystate16:
	c = y.getc()
	goto yyrule26

yystate17:
	c = y.getc()
	goto yyrule10

yystate18:
	c = y.getc()
	goto yyrule12

yystate19:
	c = y.getc()
	switch {
	default:
		goto yyrule6
	case c == '.':
		goto yystate20
	case c >= '0' && c <= '9':
		goto yystate19
	}

yystate20:
	c = y.getc()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9':
		goto yystate21
	}

yystate21:
	c = y.getc()
	switch {
	default:
		goto yyrule5
	case c >= '0' && c <= '9':
		goto yystate21
	}

yystate22:
	c = y.getc()
	switch {
	default:
		goto yyrule22
	case c == '=':
		goto yystate23
	}

yystate23:
	c = y.getc()
	goto yyrule23

yystate24:
	c = y.getc()
	goto yyrule27

yystate25:
	c = y.getc()
	switch {
	default:
		goto yyrule14
	case c == '=':
		goto yystate26
	}

yystate26:
	c = y.getc()
	goto yyrule13

yystate27:
	c = y.getc()
	switch {
	default:
		goto yyrule18
	case c == '>':
		goto yystate28
	}

yystate28:
	c = y.getc()
	goto yyrule17

yystate29:
	c = y.getc()
	switch {
	default:
		goto yyrule16
	case c == '=':
		goto yystate30
	}

yystate30:
	c = y.getc()
	goto yyrule15

yystate31:
	c = y.getc()
	switch {
	default:
		goto yyrule4
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c >= 0x4e00 && c <= 0x9fcb:
		goto yystate31
	}

yystate32:
	c = y.getc()
	switch {
	default:
		goto yyabort
	case c == '}':
		goto yystate33
	case c >= '\x01' && c <= '!' || c >= '#' && c <= '|' || c >= '~' && c <= 'ÿ' || c >= 0x4e00 && c <= 0x9fcb:
		goto yystate32
	}

yystate33:
	c = y.getc()
	goto yyrule8

yystate34:
	c = y.getc()
	switch {
	default:
		goto yyabort
	case c == '|':
		goto yystate35
	}

yystate35:
	c = y.getc()
	goto yyrule21

yyrule1: // \r\n
	{

		y.lineno++
		y.column = 0
		goto yystate0
	}
yyrule2: // [\r\n]
	{

		y.lineno++
		y.column = 0
		goto yystate0
	}
yyrule3: // [ \t]+

	goto yystate0
yyrule4: // {IDENTIFIER}
	{

		s := string(y.buf.Bytes())
		if token, ok := keywords[s]; ok {
			return token
		}
		lval.str = s
		return ID
	}
yyrule5: // {FLOAT}
	{

		if lval.value, err = strconv.ParseFloat(string(y.buf.Bytes()), 64); err != nil {
			y.Error(err.Error())
		}
		return NUM
	}
yyrule6: // {INTEGER}
	{

		if lval.value, err = strconv.ParseFloat(string(y.buf.Bytes()), 64); err != nil {
			y.Error(err.Error())
		}
		return NUM
	}
yyrule7: // {STRING_LITERAL}
	{

		lval.str = string(y.buf.Bytes())
		return STRING
	}
yyrule8: // {COMMENT}

	goto yystate0
yyrule9: // {PLUS}
	{

		lval.str = string(y.buf.Bytes())
		return PLUS
	}
yyrule10: // {MINUS}
	{

		lval.str = string(y.buf.Bytes())
		return MINUS
	}
yyrule11: // {TIMES}
	{

		lval.str = string(y.buf.Bytes())
		return TIMES
	}
yyrule12: // {DIVIDE}
	{

		lval.str = string(y.buf.Bytes())
		return DIVIDE
	}
yyrule13: // {LE}
	{

		lval.str = string(y.buf.Bytes())
		return LE
	}
yyrule14: // {LT}
	{

		lval.str = string(y.buf.Bytes())
		return LT
	}
yyrule15: // {GE}
	{

		lval.str = string(y.buf.Bytes())
		return GE
	}
yyrule16: // {GT}
	{

		lval.str = string(y.buf.Bytes())
		return GT
	}
yyrule17: // {PARAMEQUAL}
	{

		lval.str = string(y.buf.Bytes())
		return PARAMEQUAL
	}
yyrule18: // {EQ}
	{

		lval.str = string(y.buf.Bytes())
		return EQ
	}
yyrule19: // {NE}
	{

		lval.str = string(y.buf.Bytes())
		return NE
	}
yyrule20: // {AND}
	{

		lval.str = string(y.buf.Bytes())
		return AND
	}
yyrule21: // {OR}
	{

		lval.str = string(y.buf.Bytes())
		return OR
	}
yyrule22: // {EQUALS}
	{

		lval.str = string(y.buf.Bytes())
		return EQUALS
	}
yyrule23: // {COLONEQUAL}
	{

		lval.str = string(y.buf.Bytes())
		return COLONEQUAL
	}
yyrule24: // {LPAREN}
	{

		lval.str = string(y.buf.Bytes())
		return LPAREN
	}
yyrule25: // {RPAREN}
	{

		lval.str = string(y.buf.Bytes())
		return RPAREN
	}
yyrule26: // {COMMA}
	{

		lval.str = string(y.buf.Bytes())
		return COMMA
	}
yyrule27: // {SEMI}
	{

		lval.str = string(y.buf.Bytes())
		return SEMI
	}
	panic("unreachable")

	goto yyabort // silence unused label error

yyabort: // no lexem recognized
	y.empty = true
	return int(c)
}
