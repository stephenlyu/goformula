%{

package easylang

import (
    "bufio"
    "log"
    "strconv"
    "bytes"
	"unicode/utf8"
	"fmt"
)

var keywords map[string]int = map[string]int{
    "AND": AND,
    "OR": OR,
    "NOT": NOT,
}

type yylexer struct{
    src     *bufio.Reader
    buf     *bytes.Buffer
    empty   bool
    current rune

    lineno  int
    column  int
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
%}

%yyc c
%yyn c = y.getc()


IDENTIFIER      [a-zA-Z_][a-zA-Z0-9_]*
INTEGER         [0-9]+
FLOAT           [0-9]+"."[0-9]+
STRING_LITERAL  '[^\']*'
COMMENT         "{"[^"}"]*"}"
PLUS            "+"
MINUS           "-"
TIMES           "*"
DIVIDE          "/"
LT              <
GT              >
LE              <=
GE              >=
EQ              =
NE              !=
AND             &&
OR              "||"
EQUALS          :
COLONEQUAL      :=
PARAMEQUAL      =>

LPAREN          "("
RPAREN          ")"
COMMA           ,
SEMI            ;
DOT             "."
POUND           "#"
DOLLAR          "$"

%%
    y.column += utf8.RuneCount(y.buf.Bytes())
    y.buf.Reset()

\r\n
    y.lineno++
    y.column = 0

[\r\n]
   y.lineno++
   y.column = 0

[ \t]+

{IDENTIFIER}
    s := string(y.buf.Bytes())
    lval.str = s
    if token, ok := keywords[s]; ok {
        return token
    }
    return ID

{FLOAT}
    if lval.value, err = strconv.ParseFloat(string(y.buf.Bytes()), 64); err != nil {
        y.Error(err.Error())
    }
    return NUM

{INTEGER}
    if lval.value, err = strconv.ParseFloat(string(y.buf.Bytes()), 64); err != nil {
        y.Error(err.Error())
    }
    return NUM

{STRING_LITERAL}
    lval.str = string(y.buf.Bytes())
    return STRING

{COMMENT}

{PLUS}
    lval.str = string(y.buf.Bytes())
    return PLUS

{MINUS}
    lval.str = string(y.buf.Bytes())
    return MINUS

{TIMES}
    lval.str = string(y.buf.Bytes())
    return TIMES

{DIVIDE}
    lval.str = string(y.buf.Bytes())
    return DIVIDE

{LE}
    lval.str = string(y.buf.Bytes())
    return LE

{LT}
    lval.str = string(y.buf.Bytes())
    return LT

{GE}
    lval.str = string(y.buf.Bytes())
    return GE

{GT}
    lval.str = string(y.buf.Bytes())
    return GT

{PARAMEQUAL}
    lval.str = string(y.buf.Bytes())
    return PARAMEQUAL

{EQ}
    lval.str = string(y.buf.Bytes())
    return EQ

{NE}
    lval.str = string(y.buf.Bytes())
    return NE

{AND}
    lval.str = string(y.buf.Bytes())
    return AND

{OR}
    lval.str = string(y.buf.Bytes())
    return OR

{EQUALS}
    lval.str = string(y.buf.Bytes())
    return EQUALS

{COLONEQUAL}
    lval.str = string(y.buf.Bytes())
    return COLONEQUAL

{LPAREN}
    lval.str = string(y.buf.Bytes())
    return LPAREN

{RPAREN}
    lval.str = string(y.buf.Bytes())
    return RPAREN

{COMMA}
    lval.str = string(y.buf.Bytes())
    return COMMA

{SEMI}
    lval.str = string(y.buf.Bytes())
    return SEMI

{POUND}
    lval.str = string(y.buf.Bytes())
    return POUND

{DOLLAR}
    lval.str = string(y.buf.Bytes())
    return DOLLAR

{DOT}
    lval.str = string(y.buf.Bytes())
    return DOT

%%
    y.empty = true
    return int(c)
}
