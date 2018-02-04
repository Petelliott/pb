package parser

import (
	"fmt"
	"github.com/petelliott/pb/lexer"
)

func ParseProgram(tokens *lexer.TokenIterator) Program {
	prog := Program{make([]Function, 0), make([]Declaration, 0)}

	for tokens.Lookahead(0).Tok != lexer.EOF {
		if tokens.Lookahead(0).Tok == lexer.KW_FUNC {
			prog.Functions = append(prog.Functions, ParseFunction(tokens))
		} else if tokens.Lookahead(0).Tok == lexer.TYPE {
			prog.Declarations = append(prog.Declarations, ParseDeclaration(tokens))
			tokens.Expect(lexer.SEMICOLON)
		} else {
			fmt.Println("Error while parsing program.")
		}
	}

	return prog
}

func ParseFunction(tokens *lexer.TokenIterator) Function {
	var funct Function

	tokens.Expect(lexer.KW_FUNC)
	funct.Name = tokens.Expect(lexer.IDENTIFIER).Value

	funct.Args = make([]Declaration, 0)
	tokens.Expect(lexer.L_PAREN)
	if _, ok := tokens.Accept(lexer.R_PAREN); !ok {
		for {
			funct.Args = append(funct.Args, ParseDeclaration(tokens))
			if _, ok := tokens.Accept(lexer.R_PAREN); ok {
				break
			}
			tokens.Expect(lexer.COMMA)
		}
	}

	funct.Body = ParseBlock(tokens)
	return funct
}

func ParseDeclaration(tokens *lexer.TokenIterator) Declaration {
	var decl Declaration

	decl.Typ = tokens.Expect(lexer.TYPE).Value
	decl.Name = tokens.Expect(lexer.IDENTIFIER).Value

	return decl
}

func ParseBlock(tokens *lexer.TokenIterator) Block {
	block := Block{make([]Statement, 0)}

	tokens.Expect(lexer.L_BRACE)
	tokens.Expect(lexer.R_BRACE)

	return block
}
