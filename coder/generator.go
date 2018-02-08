package coder

import (
    "github.com/Petelliott/parser"
)

func GenTACprogram(prog parser.Program) TACprogram {
    var tac TACprogram

    for _, funct := range prog.Functions {
        tac.Funcs = append(tac.Funcs, GenTACfunction(funct))
    }
    return tac
}

func GenTACfunction(funct parser.Function) TACfunction {
    var tacfunc TACfunction

    int sreg = 0;
    tacfunc.Args = make([]string, len(funct.Args))
    tacfunc.Variables = map[string]int{}

    for i, arg := range funct.Args {
        tacfunc.Args[i] = arg.Name
        tacfunv.Variables[arg.Name] = TACaddr{TAC_VARIABLE, sreg}
        sreg++
    }
    return tacfunc
}

func GenTACblock(block parser.Block, regmap *map[string]int, sreg *int) []TACstatement {
    stmts := []TACstatement{}

    for _, stmt := range block.Stmts {
        if stmt.StatementType() == parser.STMT_DECLARATION {
            decl := stmt.(parser.Declaration)
            (*regmap)[decl.Name] = TACaddr{TAC_VARIABLE, sreg}
            (*sreg)++
        } else if stmt.StatementType() == parser.STMT_CONTROL {
            //TODO
        } else if stmt.StatementType() == parser.STMT_EXPRESSION {
            stmts = append(stmts, GenTACExpression()...)
        }
    }
    return stmts
}

func GenTACexpression(expr parser.Expression, regmap *map[string]int, target TACaddr) []TACstatement {
    if expr.ExpressionType() == parser.EXPR_BINARY {
        return GentTACbinary(expr.(parser.Binary), regmap, target)
    }
}

func GenTACbinary(expr parser.Binary, regmap, target TACaddr) {
    stmts := []TACstatement{}
    
    var arg1 TACaddr
    treg_carry := 0

    if expr.Arg1.ExpressionType() == EXPR_INTLITERAL {
        arg1.TACtype = TAC_INTLITERAL
        arg1.Index = expr.Arg1.(parser.IntLiteral).Value
    } else if expr.Arg1.EspressionType() == EXPR_IDENTIFIER {
        arg1.TACtype = TAC_VARIABLE
        arg1.Index = (*regmap)[expr.Arg1.(parser.Identifier).Name]
    } else {
        if target.TACtype == TAC_TEMPORARY {
            arg1 = target
            treg_carry = 1
        } else {
            arg1 = TACaddr{TAC_TEMPORARY, 0}
        }
        stmts = append(stmts, GenTACexpression(expr.Arg1, regmap, arg1))
    }
}
