package coder

import (
	"fmt"
	"github.com/Petelliott/pb/parser"
	"strconv"
)

func GenMipsProgram(ast parser.Program) string {
	prog := ".text\n"
	for _, f := range ast.Functions {
		prog += GenMipsFunction(f)
	}
	return prog
}

func GenMipsFunction(f parser.Function) string {
	regmap := map[string]string{}
	sreg := 0

	prog := "\n.globl " + f.Name + "\n" + f.Name + ":\n"

	for pos, arg := range f.Args {
		regmap[arg.Name] = "$s" + strconv.Itoa(pos)
		prog += fmt.Sprintf("    move $s%d, $a%d\n", pos, pos)
		sreg++
	}

	prog += GenMipsBlock(f.Body, &regmap, &sreg)
	prog += "    jr $ra\n"
	return prog
}

func GenMipsBlock(body parser.Block, regmap *map[string]string, sreg *int) string {
	prog := ""
	for _, stmt := range body.Stmts {
		if stmt.StatementType() == parser.STMT_EXPRESSION {
			expr, _ := GenMipsExpression(stmt.(parser.Expression), regmap, 0)
			prog += expr
			prog += "\n"
		} else if stmt.StatementType() == parser.STMT_CONTROL {
			// prog += GenMipsControl(stmt, regmap, sreg)
		} else if stmt.StatementType() == parser.STMT_DECLARATION {
			decl := stmt.(parser.Declaration)
			(*regmap)[decl.Name] = fmt.Sprintf("$s%d", *sreg)
			(*sreg)++
		}
	}
	return prog
}

func GenMipsExpression(expr parser.Expression, regmap *map[string]string, treg int) (string, string) {
	if expr.ExpressionType() == parser.EXPR_BINARY {
		binex := expr.(parser.Binary)
		prog1, reg1 := GenMipsExpression(binex.Arg1, regmap, treg)

		tmpt := treg + 1
		if reg1[1] != 't' {
			tmpt = treg
		}

		prog2, reg2 := GenMipsExpression(binex.Arg2, regmap, tmpt)

		newtreg := fmt.Sprintf("$t%d", treg)
		if binex.Operator == "+" {
			return prog1 + prog2 + fmt.Sprintf("    add %s, %s, %s\n", newtreg, reg1, reg2), newtreg
		} else if binex.Operator == "-" {
			return prog1 + prog2 + fmt.Sprintf("    sub %s, %s, %s\n", newtreg, reg1, reg2), newtreg
		} else if binex.Operator == "=" {
			if binex.Arg1.ExpressionType() == parser.EXPR_IDENTIFIER {
				return prog1 + prog2 + fmt.Sprintf("    move %s, %s\n", reg1, reg2), reg2
			} else {
				fmt.Println("assignment to non identifier expression")
			}
		} else {
			fmt.Printf("unsupported binary operation '%s'\n", binex.Operator)
		}
	} else if expr.ExpressionType() == parser.EXPR_UNARY {
		unex := expr.(parser.Unary)
		prog1, reg1 := GenMipsExpression(unex.Arg1, regmap, treg)

		newtreg := fmt.Sprintf("$t%d", treg)
		if unex.Operator == "-" {
			return prog1 + fmt.Sprintf("    sub %s, $zero, %s\n", newtreg, reg1), newtreg
		} else if unex.Operator == "*" {
			return prog1 + fmt.Sprintf("    lw %s, (%s)\n", newtreg, reg1), newtreg
		} else {
			fmt.Printf("unsupported unary operation '%s'\n", unex.Operator)
		}
	} else if expr.ExpressionType() == parser.EXPR_INTLITERAL {
		litr := expr.(parser.IntLiteral)
		newtreg := fmt.Sprintf("$t%d", treg)
		return fmt.Sprintf("    li %s, %d\n", newtreg, litr.Value), newtreg
	} else if expr.ExpressionType() == parser.EXPR_IDENTIFIER {
		ident := expr.(parser.Identifier)
		return "", (*regmap)[ident.Name]
	} else if expr.ExpressionType() == parser.EXPR_CALL {
		call := expr.(parser.Call)
		call_str := ""
		for pos, arg := range call.Args {
			arg_str, reg := GenMipsExpression(arg, regmap, treg)
			call_str += arg_str
			call_str += fmt.Sprintf("    move $a%d, %s\n\n", pos, reg)
		}
		call_str += fmt.Sprintf("    jal %s\n\n", call.Name)
		return call_str, "$v0"
	} else {
		fmt.Println("unsupported expression type")
	}
	return "", ""
}
