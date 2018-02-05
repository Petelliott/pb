package parser

import (
	"fmt"
	"github.com/petelliott/pb/lexer"
	"os"
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
			break
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

	for {
		if _, ok := tokens.Accept(lexer.R_BRACE); ok {
			break
		}
		block.Stmts = append(block.Stmts, ParseStatement(tokens))
	}
	return block
}

func ParseStatement(tokens *lexer.TokenIterator) Statement {
	tok := tokens.Lookahead(0)
	if tok.Tok == lexer.KW_WHILE || tok.Tok == lexer.KW_IF {
		return ParseControl(tokens)
	} else if tok.Tok == lexer.TYPE {
		stmnt := ParseDeclaration(tokens)
		tokens.Expect(lexer.SEMICOLON)
		return stmnt
	} else {
		stmnt := ParseExpression(tokens)
		tokens.Expect(lexer.SEMICOLON)
		if stmnt.ExpressionType() == EXPR_BINARY {
			return stmnt.(Binary)
		} else if stmnt.ExpressionType() == EXPR_UNARY {
			return stmnt.(Unary)
		} else if stmnt.ExpressionType() == EXPR_LITERAL {
			return stmnt.(Literal)
		} else if stmnt.ExpressionType() == EXPR_IDENTIFIER {
			return stmnt.(Identifier)
		} else {
			return stmnt.(Call)
		}
	}
}

func ParseControl(tokens *lexer.TokenIterator) Control {
	var cont Control

	if tok, ok := tokens.Accept(lexer.KW_WHILE); ok {
		cont.Keyword = tok.Tok
	} else {
		cont.Keyword = tokens.Expect(lexer.KW_IF).Tok
	}

	cont.Expr = ParseExpression(tokens)
	cont.Body = ParseBlock(tokens)
	return cont
}

func ParseExpression(tokens *lexer.TokenIterator) Expression {
	left := ParseExpression1(tokens)
	if tok, ok := tokens.Accept(lexer.OPERATOR); ok {
		var exp Binary
		exp.Arg1 = left
		exp.Operator = tok.Value
		exp.Arg2 = ParseExpression(tokens)
		return exp
	} else {
		return left
	}
}

func ParseExpression1(tokens *lexer.TokenIterator) Expression {
	if tok, ok := tokens.Accept(lexer.OPERATOR); ok {
		var exp Unary
		exp.Operator = tok.Value
		exp.Arg1 = ParseExpression2(tokens)
		return exp
	} else {
		return ParseExpression2(tokens)
	}
}

func ParseExpression2(tokens *lexer.TokenIterator) Expression {
	if tok, ok := tokens.Accept(lexer.IDENTIFIER); ok {
		if _, ok2 := tokens.Accept(lexer.L_PAREN); ok2 {
			tokens.Expect(lexer.R_PAREN) // TODO: function args
			return Call{tok.Value, make([]Expression, 0)}
		} else {
			return Identifier{tok.Value}
		}
	} else if tok, ok := tokens.Accept(lexer.LITERAL); ok {
		return Literal{tok.Value}
	} else if _, ok := tokens.Accept(lexer.L_PAREN); ok {
		exp := ParseExpression(tokens)
		tokens.Expect(lexer.R_PAREN)
		return exp
	} else {
		fmt.Println("error parsing expression")
		os.Exit(1)
		return Literal{""}
	}
}
