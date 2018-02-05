package optimizer

import "github.com/Petelliott/pb/parser"

func OptSimpleExpr(ast parser.Expression) parser.Expression {
	if ast.ExpressionType() == parser.EXPR_UNARY {
		unast := ast.(parser.Unary)
		expr := OptSimpleExpr(unast.Arg1)

		if expr.ExpressionType() == parser.EXPR_INTLITERAL && unast.Operator == "-" {
			return parser.IntLiteral{-(expr.(parser.IntLiteral).Value)}
		} else {
			return parser.Unary{expr, unast.Operator}
		}
	} else if ast.ExpressionType() == parser.EXPR_BINARY {
		binast := ast.(parser.Binary)
		arg1 := OptSimpleExpr(binast.Arg1)
		arg2 := OptSimpleExpr(binast.Arg2)

		if arg1.ExpressionType() == parser.EXPR_INTLITERAL && arg2.ExpressionType() == parser.EXPR_INTLITERAL {
			if binast.Operator == "+" {
				return parser.IntLiteral{arg1.(parser.IntLiteral).Value + arg2.(parser.IntLiteral).Value}
			} else if binast.Operator == "-" {
				return parser.IntLiteral{arg1.(parser.IntLiteral).Value - arg2.(parser.IntLiteral).Value}
			} else if binast.Operator == "*" {
				return parser.IntLiteral{arg1.(parser.IntLiteral).Value * arg2.(parser.IntLiteral).Value}
			} else if binast.Operator == "/" {
				return parser.IntLiteral{arg1.(parser.IntLiteral).Value / arg2.(parser.IntLiteral).Value}
			} else {
				return parser.Binary{arg1, arg2, binast.Operator}
			}
		} else {
			return parser.Binary{arg1, arg2, binast.Operator}
		}
	} else {
		return ast
	}
}
