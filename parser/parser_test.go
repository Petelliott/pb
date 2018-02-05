package parser

import (
	"fmt"
	"github.com/petelliott/pb/lexer"
	"reflect"
	"testing"
)

func TestParseProgram(t *testing.T) {
	progtok := lexer.NewTokenIterator(lexer.Lex("word a; func main() {}"))
	ast := ParseProgram(&progtok)

	expected := Program{
		[]Function{Function{"main", []Declaration{}, Block{[]Statement{}}}},
		[]Declaration{Declaration{"word", "a"}},
	}

	if !reflect.DeepEqual(ast, expected) {
		fmt.Println("expected:", expected)
		fmt.Println("got     :", ast)
		t.Fail()
	}
}

func TestParseFunction(t *testing.T) {
	progtok := lexer.NewTokenIterator(lexer.Lex("func main(word a, atom b) {1;}"))
	ast := ParseFunction(&progtok)

	expected := Function{
		"main", []Declaration{Declaration{"word", "a"}, Declaration{"atom", "b"}},
		Block{[]Statement{Literal{"1"}}},
	}

	if !reflect.DeepEqual(ast, expected) {
		fmt.Println("expected:", expected)
		fmt.Println("got     :", ast)
		t.Fail()
	}
}

func TestParseDeclaration(t *testing.T) {
	progtok := lexer.NewTokenIterator(lexer.Lex("word abcd"))
	ast := ParseDeclaration(&progtok)

	expected := Declaration{"word", "abcd"}

	if !reflect.DeepEqual(ast, expected) {
		fmt.Println("expected:", expected)
		fmt.Println("got     :", ast)
		t.Fail()
	}
}

func TestParseBlock(t *testing.T) {
	progtok := lexer.NewTokenIterator(lexer.Lex("{1;word a;}"))
	ast := ParseBlock(&progtok)

	expected := Block{[]Statement{Literal{"1"}, Declaration{"word", "a"}}}

	if !reflect.DeepEqual(ast, expected) {
		fmt.Println("expected:", expected)
		fmt.Println("got     :", ast)
		t.Fail()
	}
}

func TestParseControl(t *testing.T) {
	progtok := lexer.NewTokenIterator(lexer.Lex("while (a <= 1) {a+=1;}"))
	ast := ParseControl(&progtok)

	expected := Control{
		lexer.KW_WHILE, Binary{Identifier{"a"}, Literal{"1"}, "<="},
		Block{[]Statement{Binary{Identifier{"a"}, Literal{"1"}, "+="}}},
	}

	if !reflect.DeepEqual(ast, expected) {
		fmt.Println("expected:", expected)
		fmt.Println("got     :", ast)
		t.Fail()
	}
}

func TestParseExpression(t *testing.T) {
	progtok := lexer.NewTokenIterator(lexer.Lex("f(x) == -1 + (1*2) + -1"))
	ast := ParseExpression(&progtok)

	expected := Binary{
		Call{"f", []Expression{Identifier{"x"}}},
		Binary{
			Unary{Literal{"1"}, "-"},
			Binary{
				Binary{Literal{"1"}, Literal{"2"}, "*"},
				Unary{Literal{"1"}, "-"},
				"+",
			},
			"+",
		},
		"==",
	}

	if !reflect.DeepEqual(ast, expected) {
		fmt.Println("expected:", expected)
		fmt.Println("got     :", ast)
		t.Fail()
	}
}
