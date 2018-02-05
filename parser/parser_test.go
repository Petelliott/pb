package parser

import (
	"fmt"
	"github.com/petelliott/pb/lexer"
	"testing"
)

func TestNothing(t *testing.T) {
	prog := "atom a; func main(word b, word c){word d;} func fib(){a=a+1; b=(b+6)*((-6));}"
	toks := lexer.NewTokenIterator(lexer.Lex(prog))
	fmt.Println(ParseProgram(&toks))

	t.Fail()
}
