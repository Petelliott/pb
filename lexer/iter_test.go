package lexer

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewTokenIterator(t *testing.T) {
	tokens := []Token{Token{1, ""}, Token{2, ""}}

	if !reflect.DeepEqual(NewTokenIterator(tokens), TokenIterator{tokens, 0}) {
		t.Fail()
	}
}

func TestLookahead(t *testing.T) {
	tokens := NewTokenIterator([]Token{Token{1, ""}, Token{2, ""}})

	if tokens.Lookahead(-1) != (Token{1, ""}) {
		t.Fail()
	}

	if tokens.Lookahead(0) != (Token{1, ""}) {
		t.Fail()
	}

	if tokens.Lookahead(1) != (Token{2, ""}) {
		t.Fail()
	}

	if tokens.Lookahead(2) != (Token{EOF, ""}) {
		t.Fail()
	}

	if tokens.Lookahead(500) != (Token{EOF, ""}) {
		t.Fail()
	}
}

func TestAccept(t *testing.T) {
	tokens := NewTokenIterator([]Token{Token{1, ""}, Token{2, ""}})

	if !tokens.Accept(Token{1, ""}) {
		fmt.Println("did not accept 1")
		t.Fail()
	}

	if tokens.Accept(Token{1, ""}) {
		fmt.Println("accepted 1")
		t.Fail()
	}

	if !tokens.Accept(Token{2, ""}) {
		fmt.Println("did not accept 2")
		t.Fail()
	}

	if !tokens.Accept(Token{EOF, ""}) {
		fmt.Println("did not accept EOF")
		t.Fail()
	}
}
