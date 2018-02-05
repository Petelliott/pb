package optimizer

import (
	"fmt"
	"github.com/Petelliott/pb/lexer"
	"github.com/Petelliott/pb/parser"
	"reflect"
	"testing"
)

func TestOptSimpleExpr(t *testing.T) {
	ti := lexer.NewTokenIterator(lexer.Lex("(1+5)*(-6)"))
	ast := parser.ParseExpression(&ti)
	res := OptSimpleExpr(ast)

	if !reflect.DeepEqual(res, parser.IntLiteral{-36}) {
		fmt.Println("expected: {-36}")
		fmt.Println("got     :", res)
		t.Fail()
	}
}
