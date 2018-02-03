package lexer

import (
	"fmt"
	"testing"
)

func TestLex(t *testing.T) {
	fmt.Println(Lex("int main()  {a=\"a=242\"};  "))
	t.Fail()
}
