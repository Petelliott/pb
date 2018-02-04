package parser

const (
	STMT_DECLARATION = iota
	STMT_CONTROL     = iota
	STMT_EXPRESSION  = iota
)

const (
	EXPR_BINARY     = iota
	EXPR_UNARY      = iota
	EXPR_LITERAL    = iota
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

type Statement interface {
	StatementType() int
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

func (_ Literal) StatementType() int {
	return STMT_EXPRESSION
}

func (_ Call) StatementType() int {
	return STMT_EXPRESSION
}

func (_ Identifier) StatementType() int {
	return STMT_EXPRESSION
}

type Control struct {
	Keyword string
	Expr    string
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

type Literal struct {
	Value string
}

func (_ Literal) ExpressionType() int {
	return EXPR_LITERAL
}

type Identifier struct {
	name string
}

func (_ Identifier) ExpressionType() int {
	return EXPR_IDENTIFIER
}

type Call struct {
	name string
	args []Expression
}

func (_ Call) ExpressionType() int {
	return EXPR_CALL
}
