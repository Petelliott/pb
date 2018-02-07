package parser

const (
	STMT_DECLARATION = iota
	STMT_CONTROL     = iota
	STMT_EXPRESSION  = iota
	STMT_RETURN      = iota
)

const (
	EXPR_BINARY     = iota
	EXPR_UNARY      = iota
	EXPR_STRLITERAL = iota
	EXPR_INTLITERAL = iota
	EXPR_IDENTIFIER = iota
	EXPR_CALL       = iota
)

type Program struct {
	Functions    []Function
	Declarations []Declaration
}

type Function struct {
	Name string
	Args []Declaration
	Body Block
}

type Declaration struct {
	Typ  string
	Name string
}

type Block struct {
	Stmts []Statement
}

type Return struct {
	Expr Expression
}

type Statement interface {
	StatementType() int
}

func (_ Return) StatementType() int {
	return STMT_RETURN
}

func (_ Declaration) StatementType() int {
	return STMT_DECLARATION
}

func (_ Control) StatementType() int {
	return STMT_CONTROL
}

func (_ Binary) StatementType() int {
	return STMT_EXPRESSION
}

func (_ Unary) StatementType() int {
	return STMT_EXPRESSION
}

func (_ IntLiteral) StatementType() int {
	return STMT_EXPRESSION
}

func (_ StrLiteral) StatementType() int {
	return STMT_EXPRESSION
}

func (_ Call) StatementType() int {
	return STMT_EXPRESSION
}

func (_ Identifier) StatementType() int {
	return STMT_EXPRESSION
}

type Control struct {
	Keyword int
	Expr    Expression
	Body    Block
}

type Expression interface {
	ExpressionType() int
}

type Binary struct {
	Arg1     Expression
	Arg2     Expression
	Operator string
}

func (_ Binary) ExpressionType() int {
	return EXPR_BINARY
}

type Unary struct {
	Arg1     Expression
	Operator string
}

func (_ Unary) ExpressionType() int {
	return EXPR_UNARY
}

type IntLiteral struct {
	Value int
}

func (_ IntLiteral) ExpressionType() int {
	return EXPR_INTLITERAL
}

type StrLiteral struct {
	Value string
}

func (_ StrLiteral) ExpressionType() int {
	return EXPR_STRLITERAL
}

type Identifier struct {
	Name string
}

func (_ Identifier) ExpressionType() int {
	return EXPR_IDENTIFIER
}

type Call struct {
	Name string
	Args []Expression
}

func (_ Call) ExpressionType() int {
	return EXPR_CALL
}
