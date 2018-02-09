package coder

import (
	"fmt"
	"github.com/Petelliott/pb/lexer"
	"github.com/Petelliott/pb/parser"
	"strconv"
)

func GenMipsProgram(ast parser.Program) string {
	header := ".text\n"
    prog := ""
	for _, f := range ast.Functions {
		header += ".globl " + f.Name + "\n"
        prog += GenMipsFunction(f)
	}
	return header + "\n" + prog
}

func GenMipsFunction(f parser.Function) string {
	regmap := map[string]string{}
	sreg := 0
	stack := 0

	prog := "\n" + f.Name + ":\n"
    arg_str := ""
	for pos, arg := range f.Args {
		regmap[arg.Name] = "$s" + strconv.Itoa(pos)
		arg_str += fmt.Sprintf("    move $s%d, $a%d\n", pos, pos)
		sreg++
	}

	block_str := GenMipsBlock(f.Body, &regmap, &sreg, stack)
	prog += "    addi $sp, $sp, -4\n    sw $fp, ($sp)\n    move $fp, $sp\n"
	prog += fmt.Sprintf("    addi $sp, $sp, -%d\n", 4*(sreg+1))
	prog += "    sw $ra, -4($fp)\n"
	for i := 1; i < sreg+1; i++ {
		prog += fmt.Sprintf("    sw $s%d, -%d($fp)\n", i-1, (i+1)*4)
	}
    
    prog += "\n"

	prog += arg_str
	prog += block_str

	prog += GenMipsReturn(&sreg, stack)
	return prog
}

func GenMipsReturn(sreg *int, stack int) string {
	prog := "    lw $ra, -4($fp)\n"
	for i := 1; i < *sreg+1; i++ {
		prog += fmt.Sprintf("    lw $s%d, -%d($fp)\n", i-1, (i+1)*4)
	}
	prog += fmt.Sprintf("    addi $sp, $sp, %d\n", 4*(*sreg+1))
	prog += "    lw $fp, ($sp)\n    addi $sp, $sp, 4\n"

	prog += "    jr $ra\n\n"
	return prog
}

func GenMipsBlock(body parser.Block, regmap *map[string]string, sreg *int, stack int) string {
	prog := ""
	for _, stmt := range body.Stmts {
		if stmt.StatementType() == parser.STMT_EXPRESSION {
			expr, _ := GenMipsExpression(stmt.(parser.Expression), regmap, 0, stack)
			prog += expr
			prog += "\n"
		} else if stmt.StatementType() == parser.STMT_CONTROL {
			prog += GenMipsControl(stmt.(parser.Control), regmap, sreg, stack)
		} else if stmt.StatementType() == parser.STMT_DECLARATION {
			decl := stmt.(parser.Declaration)
			(*regmap)[decl.Name] = fmt.Sprintf("$s%d", *sreg)
			(*sreg)++
            stack += 4
		} else if stmt.StatementType() == parser.STMT_RETURN {
			ret := stmt.(parser.Return)
			expr, reg := GenMipsExpression(ret.Expr, regmap, 0, stack)
			prog += expr + fmt.Sprintf("    move $v0, %s\n", reg)
			prog += GenMipsReturn(sreg, stack)
		}
	}
	return prog
}

var ctrl_lbl = 0

func GenMipsControl(ctrl parser.Control, regmap *map[string]string, sreg *int, stack int) string {
	ctrl_prog := fmt.Sprintf("CTRL%d:\n", ctrl_lbl)
	expr_str, reg := GenMipsExpression(ctrl.Expr, regmap, 0, stack)
	ctrl_prog += expr_str
	ctrl_prog += fmt.Sprintf("    beq %s, $zero, ENDCTRL%d\n\n", reg, ctrl_lbl)

	ctrl_prog += GenMipsBlock(ctrl.Body, regmap, sreg, stack)

	if ctrl.Keyword != lexer.KW_IF {
		ctrl_prog += fmt.Sprintf("    j CTRL%d\n", ctrl_lbl)
	}
	ctrl_prog += fmt.Sprintf("ENDCTRL%d:\n", ctrl_lbl)
	ctrl_lbl++
	return ctrl_prog
}

func GenMipsExpression(expr parser.Expression, regmap *map[string]string, treg int, stack int) (string, string) {
	if expr.ExpressionType() == parser.EXPR_BINARY {
		binex := expr.(parser.Binary)
		prog1, reg1 := GenMipsExpression(binex.Arg1, regmap, treg, stack)

		tmpt := treg + 1
		if reg1[1] != 't' {
			tmpt = treg
		}

		prog2, reg2 := GenMipsExpression(binex.Arg2, regmap, tmpt, stack)

		newtreg := fmt.Sprintf("$t%d", treg)
		if binex.Operator == "+" {
			return prog1 + prog2 + fmt.Sprintf("    add %s, %s, %s\n", newtreg, reg1, reg2), newtreg
		} else if binex.Operator == "-" {
			return prog1 + prog2 + fmt.Sprintf("    sub %s, %s, %s\n", newtreg, reg1, reg2), newtreg
		} else if binex.Operator == "=" {
			if binex.Arg1.ExpressionType() == parser.EXPR_IDENTIFIER {
				return prog1 + prog2 + fmt.Sprintf("    move %s, %s\n", reg1, reg2), reg1
			} else {
				fmt.Println("assignment to non identifier expression")
			}
		} else if binex.Operator == "<"{
            return prog1 + prog2 + fmt.Sprintf("    slt %s, %s, %s\n", newtreg, reg1, reg2), newtreg
        } else {
			fmt.Printf("unsupported binary operation '%s'\n", binex.Operator)
		}
	} else if expr.ExpressionType() == parser.EXPR_UNARY {
		unex := expr.(parser.Unary)
		prog1, reg1 := GenMipsExpression(unex.Arg1, regmap, treg, stack)

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
		call_str := "\n"
		for pos, arg := range call.Args {
			arg_str, reg := GenMipsExpression(arg, regmap, treg, stack)
			call_str += arg_str
			call_str += fmt.Sprintf("    move $a%d, %s\n", pos, reg)
		}
		if treg > 0 {
			call_str += fmt.Sprintf("    addi $sp, $sp, -%d\n", 4*(treg))
		}
		for i := 0; i < treg; i++ {
			call_str += fmt.Sprintf("    sw $t%d, -%d($fp)\n", i, stack+(i+1)*4)
		}
		call_str += fmt.Sprintf("    jal %s\n", call.Name)
		for i := 0; i < treg; i++ {
			call_str += fmt.Sprintf("    lw $t%d, -%d($fp)\n", i, 4*(sreg+1)+(i+1)*4)
		}
		if treg > 0 {
			call_str += fmt.Sprintf("    addi $sp, $sp, %d\n", 4*(treg))
		}
        call_str += fmt.Sprintf("    move $t%d, $v0\n", treg)
		return call_str + "\n", fmt.Sprintf("$t%d", treg)
	} else {
		fmt.Println("unsupported expression type")
	}
	return "", ""
}
