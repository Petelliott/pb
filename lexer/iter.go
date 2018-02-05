package lexer

import (
	"fmt"
	"os"
)

var Toktable = map[int]string {
    IDENTIFIER: "identifier",
    LITERAL   : "literal",
    L_PAREN   : "'('",
    R_PAREN   : "')'",
    L_BRACKET : "'['",
    R_BRACKET : "']'",
    L_BRACE   : "'{'",
    R_BRACE   : "'}'",
    SEMICOLON : "';'",
    COMMA     : "','",
    DOT       : "'.'",
    OPERATOR  : "operator",
    EOF       : "EOF",
    ERROR     : "ERROR",
    KW_FOR    : "'for'",
    KW_WHILE  : "'while'",
    KW_IF     : "'if'",
    KW_ELSE   : "'else'",
    KW_RETURN : "'return'",
    TYPE      : "type",
}

type TokenIterator struct {
	tokens []Token
	cursor int
}

func NewTokenIterator(tokens []Token) TokenIterator {
	return TokenIterator{tokens, 0}
}

func (ti *TokenIterator) Lookahead(n int) Token {
	if n < 0 {
		n = 0
	}

	if ti.cursor+n >= len(ti.tokens) {
		return Token{EOF, ""}
	}

	return ti.tokens[ti.cursor+n]
}

func (ti *TokenIterator) Accept(tok int) (Token, bool) {
	token := ti.Lookahead(0)
	if token.Tok == tok {
		ti.cursor++
		return token, true
	}
	return token, false
}

func (ti *TokenIterator) Expect(tok int) Token {
	token := ti.Lookahead(0)
	if token.Tok == tok {
		ti.cursor++
		return token
	}
	fmt.Printf("Expected %s got %s\n", Toktable[tok], Toktable[token.Tok])
	os.Exit(1)
	return token
}
