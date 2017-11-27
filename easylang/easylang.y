/*

This file is a modified excerpt from the GNU Bison Manual examples originally found here:
http://www.gnu.org/software/bison/manual/html_node/Infix-Calc.htmlInfix-Calc

The Copyright License for the GNU Bison Manual can be found in the "fdl-1.3" file.

*/

/* Infix notation calculator.  */

/* TODO:
    1. 支持字符串                [DONE]
    2. 支持绘制函数               [DONE]
    3. 支持跨公式引用 KDJ.K        [DONE]
    4. 支持跨周期引用
        a. CLOSE#WEEK
        b. KDJ.K#WEEK
        c. MIN1, MIN5, MIN15, MIN30, MIN60, WEEK, SEASON, DAY
    5. 支持跨品种数据引用  "000001$CLOSE"
    6. 支持更多函数
 */

%{

package easylang

import "strings"

var _context = newContext()

%}

%union{
    value float64
    str string
    expr expression
    descriptions []string
    arguments []expression
}

%token  ID
%token	NUM
%token  STRING
%token  STRING_EXPR
%token  EQUALS
%token  PARAMEQUAL
%token  COLONEQUAL
%token  LPAREN
%token  RPAREN
%token  COMMA
%token  SEMI
%token  DOT
%token  POUND
%token  DOLLAR
%token  NOT

%left   OR
%left   AND
%left   EQ NE
%left   GT GE LT LE
%left	MINUS PLUS
%left	TIMES DIVIDE
%left	UNARY

%type	<value>	NUM
%type   <str> STRING
%type   <str> STRING_EXPR
%type   <str> ID graph_description OR AND EQ NE GT GE LT LE MINUS PLUS TIMES DIVIDE NOT
%type   <expr> statement expression primary_expression postfix_expression unary_expression multiplicative_expression additive_expression relational_expression equality_expression logical_and_expression
%type   <descriptions> statement_suffix graph_description_list
%type   <arguments> argument_expression_list


%% /* The grammar follows.  */

formula : statement_list
;

statement_list: statement
                | statement_list statement
;

statement: ID EQUALS expression statement_suffix {
                if _, ok := $3.(*stringexpr); ok {
                    lexer, _ := yylex.(*yylexer)
                    _context.addError(GeneralError(lexer.lineno, lexer.column, "string can't be right value"))
                }
                $$ = AssignmentExpression(_context, $1, $3, false)
                _context.addOutput($$.VarName(), $4, 0, 0)
           }
           | ID COLONEQUAL expression statement_suffix {
                if _, ok := $3.(*stringexpr); ok {
                    lexer, _ := yylex.(*yylexer)
                    _context.addError(GeneralError(lexer.lineno, lexer.column, "string can't be right value"))
                }
                $$ = AssignmentExpression(_context, $1, $3, false)
                _context.addNotOutputVar($$.VarName(), $4, 0, 0)
           }
           | ID PARAMEQUAL LPAREN NUM COMMA NUM COMMA NUM RPAREN SEMI {
                $$ = ParamExpression(_context, $1, $4, $6, $8)
           }
           | expression statement_suffix  {
                if _, ok := $1.(*stringexpr); ok {
                    lexer, _ := yylex.(*yylexer)
                    _context.addError(GeneralError(lexer.lineno, lexer.column, "string can't be right value"))
                }
                varName := _context.newAnonymousVarName()
                $$ = AssignmentExpression(_context, varName, $1, true)
                _context.addOutput(varName, $2, 0, 0)
           }
;

statement_suffix: SEMI {}
                  | graph_description_list SEMI { $$ = $1 }

graph_description_list: COMMA graph_description {
                            if !IsValidDescription($2) {
                                lexer, _ := yylex.(*yylexer)
                                _context.addError(BadGraphDescError(lexer.lineno, lexer.column, $2))
                            }
                            $$ = append($$, $2)
                        }
                        | graph_description_list COMMA graph_description {
                            if !IsValidDescription($3) {
                                lexer, _ := yylex.(*yylexer)
                                _context.addError(BadGraphDescError(lexer.lineno, lexer.column, $3))
                            }
                            $$ = append($1, $3)
                        }

graph_description : ID  { $$ = $1 }

primary_expression: ID  {
                        expr := _context.defined($1)
                        if expr == nil {
                            expr = _context.definedParam($1)
                        }
                        if expr != nil {
                        } else if funcName, ok := noArgFuncMap[$1]; ok {
                            expr = FunctionExpression(_context, funcName, nil)
                        } else {
                            lexer, _ := yylex.(*yylexer)
                            _context.addError(UndefinedVarError(lexer.lineno, lexer.column, $1))
                            expr = ErrorExpression(_context, $1)
                        }
                        $$ = expr
                    }
                    | NUM { $$ = ConstantExpression(_context, $1) }
                    | STRING { $$ = StringExpression(_context, $1[1:len($1) - 1]) }
                    | LPAREN expression RPAREN { $$ = $2 }

postfix_expression: primary_expression  { $$ = $1 }
                    | ID LPAREN argument_expression_list RPAREN {
                        if _, ok := funcMap[$1]; !ok {
                            lexer, _ := yylex.(*yylexer)
                            _context.addError(UndefinedFunctionError(lexer.lineno, lexer.column, $1))
                            $$ = ErrorExpression(_context, $1)
                        } else {
                            $$ = FunctionExpression(_context, $1, $3)
                        }
                    }
                    | ID DOT ID {
                        if !_context.isReferenceSupport($1, $3) {
                            lexer, _ := yylex.(*yylexer)
                            _context.addError(GeneralError(lexer.lineno, lexer.column, __yyfmt__.Sprintf("%s.%s not supported", $1, $3)))
                            $$ = ErrorExpression(_context, $3)
                        } else {
                            $$ = ReferenceExpression(_context, $1, $3)
                        }
                    }
                    | ID POUND ID {
                        period:=translatePeriod($3)
                        if funcName, ok := noArgFuncMap[$1]; ok && _context.isPeriodSupport(period) {
                            $$ = CrossFunctionExpression(_context, funcName, "", period)
                        } else {
                            lexer, _ := yylex.(*yylexer)
                            _context.addError(GeneralError(lexer.lineno, lexer.column, __yyfmt__.Sprintf("%s#%s not supported", $1, $3)))
                            $$ = ErrorExpression(_context, $3)
                        }
                    }
                    | ID DOT ID POUND ID {
                        period:=translatePeriod($5)
                        if !_context.isReferenceSupport($1, $3) || !_context.isPeriodSupport(period) {
                            lexer, _ := yylex.(*yylexer)
                            _context.addError(GeneralError(lexer.lineno, lexer.column, __yyfmt__.Sprintf("%s.%s#%s not supported", $1, $3, $5)))
                            $$ = ErrorExpression(_context, $3)
                        } else {
                            $$ = CrossReferenceExpression(_context, $1, $3, "", period)
                        }
                    }
                    | STRING_EXPR {
                        parts := strings.Split($1[1:len($1)-1], "$")
                        lexer, _ := yylex.(*yylexer)

                        var reportError = func(msg string) {
                            if msg == "" {
                                msg = __yyfmt__.Sprintf("\"%s\" not supported", $1)
                            }
                            _context.addError(GeneralError(lexer.lineno, lexer.column, msg))
                            $$ = ErrorExpression(_context, $1)
                        }

                        if len(parts) != 2 {
                            reportError("")
                            break
                        }

                        code := parts[0]
                        if !_context.isSecuritySupport(code) {
                            reportError(__yyfmt__.Sprintf("code %s not supported", code))
                            break
                        }

                        expr := parts[1]

                        var period string

                        parts = strings.Split(expr, "#")
                        if len(parts) > 2 {
                            reportError("")
                            break
                        } else if len(parts) == 2 {
                            period = translatePeriod(parts[1])
                            if !_context.isPeriodSupport(period) {
                                reportError(__yyfmt__.Sprintf("period %s not supported", parts[1]))
                                break
                            }
                        }

                        parts = strings.Split(parts[0], ".")
                        switch len(parts) {
                        case 1:
                            if funcName, ok := noArgFuncMap[parts[0]]; !ok {
                                reportError(__yyfmt__.Sprintf("function %s not supported", parts[0]))
                            } else {
                                $$ = CrossFunctionExpression(_context, funcName, code, period)
                            }
                        case 2:
                            if !_context.isReferenceSupport(parts[0], parts[1]) {
                                reportError("")
                            } else {
                                $$ = CrossReferenceExpression(_context, parts[0], parts[1], code, period)
                            }
                        default:
                            reportError("")
                        }
                    }

argument_expression_list: expression  { $$ = []expression{$1} }
                          | argument_expression_list COMMA expression { $$ = append($$, $3) }

unary_expression
	: postfix_expression  { $$ = $1 }
	| NOT unary_expression %prec UNARY { $$ = UnaryExpression(_context, $1, $2) }
	| MINUS unary_expression %prec UNARY { $$ = UnaryExpression(_context, $1, $2) }

multiplicative_expression
	: unary_expression { $$ = $1 }
	| multiplicative_expression TIMES unary_expression { $$ = BinaryExpression(_context, $2, $1, $3) }
	| multiplicative_expression DIVIDE unary_expression { $$ = BinaryExpression(_context, $2, $1, $3) }

additive_expression
	: multiplicative_expression { $$ = $1 }
	| additive_expression PLUS multiplicative_expression { $$ = BinaryExpression(_context, $2, $1, $3) }
	| additive_expression MINUS multiplicative_expression { $$ = BinaryExpression(_context, $2, $1, $3) }

relational_expression
	: additive_expression { $$ = $1 }
	| relational_expression LT additive_expression { $$ = BinaryExpression(_context, $2, $1, $3) }
	| relational_expression GT additive_expression { $$ = BinaryExpression(_context, $2, $1, $3) }
	| relational_expression LE additive_expression { $$ = BinaryExpression(_context, $2, $1, $3) }
	| relational_expression GE additive_expression { $$ = BinaryExpression(_context, $2, $1, $3) }

equality_expression
	: relational_expression { $$ = $1 }
	| equality_expression EQ relational_expression { $$ = BinaryExpression(_context, $2, $1, $3) }
	| equality_expression NE relational_expression { $$ = BinaryExpression(_context, $2, $1, $3) }

logical_and_expression
	: equality_expression { $$ = $1 }
	| logical_and_expression AND equality_expression { $$ = BinaryExpression(_context, $2, $1, $3) }

expression
	: logical_and_expression { $$ = $1 }
	| expression OR logical_and_expression { $$ = BinaryExpression(_context, $2, $1, $3) }

%%
