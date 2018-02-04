package parser

import (
	"fmt"
	"github.com/petelliott/pb/lexer"
	"testing"
)

func TestNothing(t *testing.T) {
	prog := "atom a; func main(word b, word c){word d; if (.) {word c;}} func fib(){.;}"
	toks := lexer.NewTokenIterator(lexer.Lex(prog))
	fmt.Println(ParseProgram(&toks))

	t.Fail()
}
