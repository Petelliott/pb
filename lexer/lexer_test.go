package lexer

import (
	"fmt"
	"testing"
)

func TestLex(t *testing.T) {
	fmt.Println(Lex("atom main()  {word a+=\"a=242\"};  "))
	t.Fail()
}
