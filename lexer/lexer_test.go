package lexer

import (
	"fmt"
	"reflect"
	"testing"
)

func TestControl(t *testing.T) {
	result := Lex("()}{[]")
	expected := []Token{
		Token{L_PAREN, ""}, Token{R_PAREN, ""}, Token{R_BRACE, ""},
		Token{L_BRACE, ""}, Token{L_BRACKET, ""}, Token{R_BRACKET, ""},
	}

	if !reflect.DeepEqual(result, expected) {
		fmt.Println("expected:", expected)
		fmt.Println("got:     ", result)
		t.Fail()
	}
}

func TestString(t *testing.T) {
	result := Lex("'\" ' \"''''\"")
	expected := []Token{
		Token{LITERAL, "'\" '"}, Token{LITERAL, "\"''''\""},
	}

	if !reflect.DeepEqual(result, expected) {
		fmt.Println("expected:", expected)
		fmt.Println("got:     ", result)
		t.Fail()
	}
}

func TestKeyword(t *testing.T) {
	result := Lex("for bat \nif   =while")
	expected := []Token{
		Token{KEYWORD, "for"}, Token{IDENTIFIER, "bat"},
		Token{KEYWORD, "if"}, Token{OPERATOR, "="}, Token{KEYWORD, "while"},
	}

	if !reflect.DeepEqual(result, expected) {
		fmt.Println("expected:", expected)
		fmt.Println("got:     ", result)
		t.Fail()
	}
}

func TestOperators(t *testing.T) {
	result := Lex("+==-=*>=+ ==")
	expected := []Token{
		Token{OPERATOR, "+="}, Token{OPERATOR, "="}, Token{OPERATOR, "-="},
		Token{OPERATOR, "*"}, Token{OPERATOR, ">="}, Token{OPERATOR, "+"},
		Token{OPERATOR, "=="},
	}

	if !reflect.DeepEqual(result, expected) {
		fmt.Println("expected:", expected)
		fmt.Println("got:     ", result)
		t.Fail()
	}
}

func TestIntLiteral(t *testing.T) {
	result := Lex(" (=123 57)")
	expected := []Token{
		Token{L_PAREN, ""}, Token{OPERATOR, "="}, Token{LITERAL, "123"},
		Token{LITERAL, "57"}, Token{R_PAREN, ""},
	}

	if !reflect.DeepEqual(result, expected) {
		fmt.Println("expected:", expected)
		fmt.Println("got:     ", result)
		t.Fail()
	}
}

func TestEx1(t *testing.T) {
	result := Lex("func main() {\n\tword abc=5;\n\tabc += abc*6;\n\treturn \"abc=\"+abc;\n}")
	expected := []Token{
		Token{KEYWORD, "func"}, Token{IDENTIFIER, "main"}, Token{L_PAREN, ""},
		Token{R_PAREN, ""}, Token{L_BRACE, ""}, Token{KEYWORD, "word"},
		Token{IDENTIFIER, "abc"}, Token{OPERATOR, "="}, Token{LITERAL, "5"},
		Token{SEMICOLON, ""}, Token{IDENTIFIER, "abc"}, Token{OPERATOR, "+="},
		Token{IDENTIFIER, "abc"}, Token{OPERATOR, "*"}, Token{LITERAL, "6"},
		Token{SEMICOLON, ""}, Token{KEYWORD, "return"}, Token{LITERAL, "\"abc=\""},
		Token{OPERATOR, "+"}, Token{IDENTIFIER, "abc"}, Token{SEMICOLON, ""},
		Token{R_BRACE, ""},
	}

	if !reflect.DeepEqual(result, expected) {
		fmt.Println("expected:", expected)
		fmt.Println("got:     ", result)
		t.Fail()
	}
}
