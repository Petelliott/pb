package lexer

const (
	IDENTIFIER = iota
	LITERAL    = iota
	L_PAREN    = iota
	R_PAREN    = iota
	L_BRACKET  = iota
	R_BRACKET  = iota
	L_BRACE    = iota
	R_BRACE    = iota
	SEMICOLON  = iota
	COMMA      = iota
	DOT        = iota
	OPERATOR   = iota
	EOF        = iota
	ERROR      = iota
	KW_FOR     = iota
	KW_WHILE   = iota
	KW_IF      = iota
	KW_ELSE    = iota
	KW_FUNC    = iota
	KW_RETURN  = iota
	TYPE       = iota
)

type Token struct {
	Tok   int
	Value string
}
