package coder

import (
	"fmt"
	"github.com/Petelliott/pb/lexer"
	"github.com/Petelliott/pb/parser"
	"testing"
)

func TestGenMipsProgram(t *testing.T) {
	ti := lexer.NewTokenIterator(lexer.Lex("func main(word a) { word b; a+b-(3+a); a= 3+3;}"))
	code := GenMipsProgram(parser.ParseProgram(&ti))
	fmt.Print(code)
	t.Fail()
}
