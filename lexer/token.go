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
	OPERATOR   = iota
	KEYWORD    = iota
	EOF        = iota
)

type Token struct {
	Tok   int
	Value string
}
