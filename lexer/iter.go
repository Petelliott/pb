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

func (ti *TokenIterator) Accept(tok Token) bool {
	if ti.Lookahead(0) == tok {
		ti.cursor++
		return true
	}
	return false
}

func (ti *TokenIterator) Expect(tok Token) bool {
	if ti.Lookahead(0) == tok {
		ti.cursor++
		return true
	}
	fmt.Println("Unexpected token.")
	os.Exit(1)
	return false
}
