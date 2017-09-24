/*

This file is a modified excerpt from the GNU Bison Manual examples originally found here:
http://www.gnu.org/software/bison/manual/html_node/Infix-Calc.htmlInfix-Calc

The Copyright License for the GNU Bison Manual can be found in the "fdl-1.3" file.

*/

/* Infix notation calculator.  */

%{

package easylang

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
%token  EQUALS
%token  PARAMEQUAL
%token  COLONEQUAL
%token  LPAREN
%token  RPAREN
%token  COMMA
%token  SEMI
%token  NOT

%left   OR
%left   AND
%left   EQ NE
%left   GT GE LT LE
%left	MINUS PLUS
%left	TIMES DIVIDE
%left	UNARY

%type	<value>	NUM
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
                $$ = AssignmentExpression(_context, $1, $3)
                _context.addOutput($$.VarName(), $4, 0, 0)
           }
           | ID COLONEQUAL expression statement_suffix { $$ = AssignmentExpression(_context, $1, $3) }
           | ID PARAMEQUAL LPAREN NUM COMMA NUM COMMA NUM RPAREN SEMI {
                $$ = ParamExpression(_context, $1, $4, $6, $8)
           }
           | expression statement_suffix  {
                varName := _context.newAnonymousVarName()
                $$ = AssignmentExpression(_context, varName, $1)
                _context.addOutput(varName, $2, 0, 0)
           }
;

statement_suffix: SEMI {}
                  | graph_description_list SEMI { $$ = $1 }

graph_description_list: COMMA graph_description { $$ = append($$, $2) }
                        | graph_description_list COMMA graph_description { $$ = append($1, $3) }

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
                    | STRING {}
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
