// CAUTION: Generated file - DO NOT EDIT.

package easylang

import (
	"bufio"
	"bytes"
	"log"
	"strconv"
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
}

func addRune(b *bytes.Buffer, c rune) {
	if _, err := b.WriteRune(c); err != nil {
		log.Fatalf("WriteRune: %s", err)
	}
}

func newLexer(src *bufio.Reader) (y *yylexer) {
	y = &yylexer{src: src, buf: &bytes.Buffer{}}
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
	log.Fatal(e)
}

func (y *yylexer) Lex(lval *yySymType) int {
	var err error
	c := y.current
	if y.empty {
		c, y.empty = y.getc(), false
	}

yystate0:

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
		goto yystate3
	case c == '&':
		goto yystate5
	case c == '(':
		goto yystate9
	case c == ')':
		goto yystate10
	case c == '*':
		goto yystate11
	case c == '+':
		goto yystate12
	case c == ',':
		goto yystate13
	case c == '-':
		goto yystate14
	case c == '/':
		goto yystate15
	case c == ':':
		goto yystate19
	case c == ';':
		goto yystate21
	case c == '<':
		goto yystate22
	case c == '=':
		goto yystate24
	case c == '>':
		goto yystate26
	case c == '\'':
		goto yystate7
	case c == '\t' || c == '\n' || c == '\r' || c == ' ':
		goto yystate2
	case c == '{':
		goto yystate29
	case c == '|':
		goto yystate31
	case c >= '0' && c <= '9':
		goto yystate16
	case c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate28
	}

yystate2:
	c = y.getc()
	switch {
	default:
		goto yyrule1
	case c == '\t' || c == '\n' || c == '\r' || c == ' ':
		goto yystate2
	}

yystate3:
	c = y.getc()
	switch {
	default:
		goto yyabort
	case c == '=':
		goto yystate4
	}

yystate4:
	c = y.getc()
	goto yyrule17

yystate5:
	c = y.getc()
	switch {
	default:
		goto yyabort
	case c == '&':
		goto yystate6
	}

yystate6:
	c = y.getc()
	goto yyrule18

yystate7:
	c = y.getc()
	switch {
	default:
		goto yyabort
	case c == '\'':
		goto yystate8
	case c >= '\x01' && c <= '&' || c >= '(' && c <= 'ÿ':
		goto yystate7
	}

yystate8:
	c = y.getc()
	goto yyrule5

yystate9:
	c = y.getc()
	goto yyrule22

yystate10:
	c = y.getc()
	goto yyrule23

yystate11:
	c = y.getc()
	goto yyrule9

yystate12:
	c = y.getc()
	goto yyrule7

yystate13:
	c = y.getc()
	goto yyrule24

yystate14:
	c = y.getc()
	goto yyrule8

yystate15:
	c = y.getc()
	goto yyrule10

yystate16:
	c = y.getc()
	switch {
	default:
		goto yyrule4
	case c == '.':
		goto yystate17
	case c >= '0' && c <= '9':
		goto yystate16
	}

yystate17:
	c = y.getc()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9':
		goto yystate18
	}

yystate18:
	c = y.getc()
	switch {
	default:
		goto yyrule3
	case c >= '0' && c <= '9':
		goto yystate18
	}

yystate19:
	c = y.getc()
	switch {
	default:
		goto yyrule20
	case c == '=':
		goto yystate20
	}

yystate20:
	c = y.getc()
	goto yyrule21

yystate21:
	c = y.getc()
	goto yyrule25

yystate22:
	c = y.getc()
	switch {
	default:
		goto yyrule12
	case c == '=':
		goto yystate23
	}

yystate23:
	c = y.getc()
	goto yyrule11

yystate24:
	c = y.getc()
	switch {
	default:
		goto yyrule16
	case c == '>':
		goto yystate25
	}

yystate25:
	c = y.getc()
	goto yyrule15

yystate26:
	c = y.getc()
	switch {
	default:
		goto yyrule14
	case c == '=':
		goto yystate27
	}

yystate27:
	c = y.getc()
	goto yyrule13

yystate28:
	c = y.getc()
	switch {
	default:
		goto yyrule2
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate28
	}

yystate29:
	c = y.getc()
	switch {
	default:
		goto yyabort
	case c == '}':
		goto yystate30
	case c >= '\x01' && c <= '!' || c >= '#' && c <= '|' || c >= '~' && c <= 'ÿ':
		goto yystate29
	}

yystate30:
	c = y.getc()
	goto yyrule6

yystate31:
	c = y.getc()
	switch {
	default:
		goto yyabort
	case c == '|':
		goto yystate32
	}

yystate32:
	c = y.getc()
	goto yyrule19

yyrule1: // [ \t\r\n]+

	goto yystate0
yyrule2: // {IDENTIFIER}
	{

		s := string(y.buf.Bytes())
		if token, ok := keywords[s]; ok {
			return token
		}
		lval.str = s
		return ID
	}
yyrule3: // {FLOAT}
	{

		if lval.value, err = strconv.ParseFloat(string(y.buf.Bytes()), 64); err != nil {
			y.Error(err.Error())
		}
		return NUM
	}
yyrule4: // {INTEGER}
	{

		if lval.value, err = strconv.ParseFloat(string(y.buf.Bytes()), 64); err != nil {
			y.Error(err.Error())
		}
		return NUM
	}
yyrule5: // {STRING_LITERAL}
	{

		lval.str = string(y.buf.Bytes())
		return STRING
	}
yyrule6: // {COMMENT}

	goto yystate0
yyrule7: // {PLUS}
	{

		lval.str = string(y.buf.Bytes())
		return PLUS
	}
yyrule8: // {MINUS}
	{

		lval.str = string(y.buf.Bytes())
		return MINUS
	}
yyrule9: // {TIMES}
	{

		lval.str = string(y.buf.Bytes())
		return TIMES
	}
yyrule10: // {DIVIDE}
	{

		lval.str = string(y.buf.Bytes())
		return DIVIDE
	}
yyrule11: // {LE}
	{

		lval.str = string(y.buf.Bytes())
		return LE
	}
yyrule12: // {LT}
	{

		lval.str = string(y.buf.Bytes())
		return LT
	}
yyrule13: // {GE}
	{

		lval.str = string(y.buf.Bytes())
		return GE
	}
yyrule14: // {GT}
	{

		lval.str = string(y.buf.Bytes())
		return GT
	}
yyrule15: // {PARAMEQUAL}
	{

		lval.str = string(y.buf.Bytes())
		return PARAMEQUAL
	}
yyrule16: // {EQ}
	{

		lval.str = string(y.buf.Bytes())
		return EQ
	}
yyrule17: // {NE}
	{

		lval.str = string(y.buf.Bytes())
		return NE
	}
yyrule18: // {AND}
	{

		lval.str = string(y.buf.Bytes())
		return AND
	}
yyrule19: // {OR}
	{

		lval.str = string(y.buf.Bytes())
		return OR
	}
yyrule20: // {EQUALS}
	{

		lval.str = string(y.buf.Bytes())
		return EQUALS
	}
yyrule21: // {COLONEQUAL}
	{

		lval.str = string(y.buf.Bytes())
		return COLONEQUAL
	}
yyrule22: // {LPAREN}
	{

		lval.str = string(y.buf.Bytes())
		return LPAREN
	}
yyrule23: // {RPAREN}
	{

		lval.str = string(y.buf.Bytes())
		return RPAREN
	}
yyrule24: // {COMMA}
	{

		lval.str = string(y.buf.Bytes())
		return COMMA
	}
yyrule25: // {SEMI}
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
