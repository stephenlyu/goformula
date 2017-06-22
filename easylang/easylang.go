//line easylang.y:13

package easylang

import __yyfmt__ "fmt"

//line easylang.y:16
var _context = newContext()

//line easylang.y:22
type yySymType struct {
	yys          int
	value        float64
	str          string
	expr         expression
	descriptions []string
	arguments    []expression
}

const ID = 57346
const NUM = 57347
const STRING = 57348
const EQUALS = 57349
const PARAMEQUAL = 57350
const COLONEQUAL = 57351
const LPAREN = 57352
const RPAREN = 57353
const COMMA = 57354
const SEMI = 57355
const NOT = 57356
const OR = 57357
const AND = 57358
const EQ = 57359
const NE = 57360
const GT = 57361
const GE = 57362
const LT = 57363
const LE = 57364
const MINUS = 57365
const PLUS = 57366
const TIMES = 57367
const DIVIDE = 57368
const UNARY = 57369

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"ID",
	"NUM",
	"STRING",
	"EQUALS",
	"PARAMEQUAL",
	"COLONEQUAL",
	"LPAREN",
	"RPAREN",
	"COMMA",
	"SEMI",
	"NOT",
	"OR",
	"AND",
	"EQ",
	"NE",
	"GT",
	"GE",
	"LT",
	"LE",
	"MINUS",
	"PLUS",
	"TIMES",
	"DIVIDE",
	"UNARY",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line easylang.y:149

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 81

var yyAct = [...]int{

	5, 52, 24, 11, 9, 7, 6, 10, 38, 39,
	37, 36, 33, 35, 32, 34, 29, 40, 42, 43,
	8, 44, 45, 46, 48, 30, 31, 28, 26, 65,
	25, 25, 49, 25, 23, 54, 53, 57, 58, 59,
	60, 2, 63, 64, 61, 62, 1, 66, 67, 68,
	47, 55, 56, 71, 41, 16, 17, 4, 16, 17,
	18, 51, 50, 18, 13, 69, 70, 13, 27, 3,
	12, 72, 19, 14, 15, 0, 14, 20, 22, 21,
	23,
}
var yyPact = [...]int{

	53, -1000, 53, -1000, 70, 15, 0, 8, -7, -13,
	-17, -1000, -1000, 50, 50, -1000, -1000, -1000, 50, -1000,
	50, 50, 50, 50, -1000, 50, -1000, 49, 32, 50,
	50, 50, 50, 50, 50, 50, 50, 50, 50, 50,
	-1000, 24, -1000, 18, 15, 15, 15, 54, 16, 0,
	-1000, 32, -1000, -1000, 8, -7, -7, -13, -13, -13,
	-13, -17, -17, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	50, -1000, 16,
}
var yyPgo = [...]int{

	0, 1, 69, 0, 74, 70, 3, 7, 4, 20,
	5, 6, 2, 68, 50, 46, 41,
}
var yyR1 = [...]int{

	0, 15, 16, 16, 2, 2, 2, 2, 12, 12,
	13, 13, 1, 4, 4, 4, 4, 5, 5, 14,
	14, 6, 6, 6, 7, 7, 7, 8, 8, 8,
	9, 9, 9, 9, 9, 10, 10, 10, 11, 11,
	3, 3,
}
var yyR2 = [...]int{

	0, 1, 1, 2, 4, 4, 4, 2, 1, 2,
	2, 3, 1, 1, 1, 1, 3, 1, 4, 1,
	3, 1, 2, 2, 1, 3, 3, 1, 3, 3,
	1, 3, 3, 3, 3, 1, 3, 3, 1, 3,
	1, 3,
}
var yyChk = [...]int{

	-1000, -15, -16, -2, 4, -3, -11, -10, -9, -8,
	-7, -6, -5, 14, 23, -4, 5, 6, 10, -2,
	7, 9, 8, 10, -12, 15, 13, -13, 12, 16,
	17, 18, 21, 19, 22, 20, 24, 23, 25, 26,
	-6, 4, -6, -3, -3, -3, -3, -14, -3, -11,
	13, 12, -1, 4, -10, -9, -9, -8, -8, -8,
	-8, -7, -7, -6, -6, 11, -12, -12, -12, 11,
	12, -1, -3,
}
var yyDef = [...]int{

	0, -2, 1, 2, 13, 0, 40, 38, 35, 30,
	27, 24, 21, 0, 0, 17, 14, 15, 0, 3,
	0, 0, 0, 0, 7, 0, 8, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	22, 13, 23, 0, 0, 0, 0, 0, 19, 41,
	9, 0, 10, 12, 39, 36, 37, 31, 32, 33,
	34, 28, 29, 25, 26, 16, 4, 5, 6, 18,
	0, 11, 20,
}
var yyTok1 = [...]int{

	1,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 4:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line easylang.y:66
		{
			yyVAL.expr = AssignmentExpression(_context, yyDollar[1].str, yyDollar[3].expr)
			_context.addOutput(yyVAL.expr.VarName(), yyDollar[4].descriptions, 0, 0)
		}
	case 5:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line easylang.y:70
		{
			yyVAL.expr = AssignmentExpression(_context, yyDollar[1].str, yyDollar[3].expr)
		}
	case 6:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line easylang.y:71
		{
			yyVAL.expr = ParamExpression(_context, yyDollar[1].str, yyDollar[3].expr)
		}
	case 7:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line easylang.y:72
		{
			varName := _context.newAnonymousVarName()
			yyVAL.expr = AssignmentExpression(_context, varName, yyDollar[1].expr)
			_context.addOutput(varName, yyDollar[2].descriptions, 0, 0)
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line easylang.y:79
		{
		}
	case 9:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line easylang.y:80
		{
			yyVAL.descriptions = yyDollar[1].descriptions
		}
	case 10:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line easylang.y:82
		{
			yyVAL.descriptions = append(yyVAL.descriptions, yyDollar[2].str)
		}
	case 11:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line easylang.y:83
		{
			yyVAL.descriptions = append(yyDollar[1].descriptions, yyDollar[3].str)
		}
	case 12:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line easylang.y:85
		{
			yyVAL.str = yyDollar[1].str
		}
	case 13:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line easylang.y:87
		{
			expr := _context.defined(yyDollar[1].str)
			if expr == nil {
				expr = _context.definedParam(yyDollar[1].str)
			}
			if expr != nil {
			} else if funcName, ok := noArgFuncMap[yyDollar[1].str]; ok {
				expr = FunctionExpression(_context, funcName, nil)
			} else {
				// TODO: handle error
				expr = ErrorExpression(_context, yyDollar[1].str)
			}
			yyVAL.expr = expr
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line easylang.y:101
		{
			yyVAL.expr = ConstantExpression(_context, yyDollar[1].value)
		}
	case 15:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line easylang.y:102
		{
		}
	case 16:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line easylang.y:103
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 17:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line easylang.y:105
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 18:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line easylang.y:106
		{
			// TODO: handle error
			yyVAL.expr = FunctionExpression(_context, yyDollar[1].str, yyDollar[3].arguments)
		}
	case 19:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line easylang.y:111
		{
			yyVAL.arguments = []expression{yyDollar[1].expr}
		}
	case 20:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line easylang.y:112
		{
			yyVAL.arguments = append(yyVAL.arguments, yyDollar[3].expr)
		}
	case 21:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line easylang.y:115
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 22:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line easylang.y:116
		{
			yyVAL.expr = UnaryExpression(_context, yyDollar[1].str, yyDollar[2].expr)
		}
	case 23:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line easylang.y:117
		{
			yyVAL.expr = UnaryExpression(_context, yyDollar[1].str, yyDollar[2].expr)
		}
	case 24:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line easylang.y:120
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 25:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line easylang.y:121
		{
			yyVAL.expr = BinaryExpression(_context, yyDollar[2].str, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 26:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line easylang.y:122
		{
			yyVAL.expr = BinaryExpression(_context, yyDollar[2].str, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 27:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line easylang.y:125
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 28:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line easylang.y:126
		{
			yyVAL.expr = BinaryExpression(_context, yyDollar[2].str, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 29:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line easylang.y:127
		{
			yyVAL.expr = BinaryExpression(_context, yyDollar[2].str, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 30:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line easylang.y:130
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 31:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line easylang.y:131
		{
			yyVAL.expr = BinaryExpression(_context, yyDollar[2].str, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 32:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line easylang.y:132
		{
			yyVAL.expr = BinaryExpression(_context, yyDollar[2].str, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 33:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line easylang.y:133
		{
			yyVAL.expr = BinaryExpression(_context, yyDollar[2].str, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 34:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line easylang.y:134
		{
			yyVAL.expr = BinaryExpression(_context, yyDollar[2].str, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 35:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line easylang.y:137
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 36:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line easylang.y:138
		{
			yyVAL.expr = BinaryExpression(_context, yyDollar[2].str, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 37:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line easylang.y:139
		{
			yyVAL.expr = BinaryExpression(_context, yyDollar[2].str, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 38:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line easylang.y:142
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 39:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line easylang.y:143
		{
			yyVAL.expr = BinaryExpression(_context, yyDollar[2].str, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line easylang.y:146
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 41:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line easylang.y:147
		{
			yyVAL.expr = BinaryExpression(_context, yyDollar[2].str, yyDollar[1].expr, yyDollar[3].expr)
		}
	}
	goto yystack /* stack new state and value */
}
