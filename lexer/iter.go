package lexer

import (
	"fmt"
	"os"
)

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
	fmt.Println("expected:", tok, ", got:", token.Tok)
	os.Exit(1)
	return token
}
