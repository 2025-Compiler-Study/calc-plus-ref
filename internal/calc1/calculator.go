package calc1

import (
	"calcPlus/internal/parser"
	"github.com/antlr4-go/antlr/v4"
	"strconv"
)

type Calculator struct {
	parser.BaseCalcPlusVisitor
	Variables map[string]int
}

func NewCalculator() *Calculator {
	return &Calculator{Variables: make(map[string]int)}
}

func (c *Calculator) Visit(tree antlr.ParseTree) any {
	return tree.Accept(c)
}

func (c *Calculator) VisitCalc1(ctx *parser.Calc1Context) any {
	for _, stmt := range ctx.AllStmt() {
		c.Visit(stmt)
	}
	return nil
}

func (c *Calculator) VisitExprAssign(ctx *parser.ExprAssignContext) any {
	result := c.Visit(ctx.Expr()).(int)
	c.Variables[ctx.VAR().GetText()] = result

	return nil
}

func (c *Calculator) VisitMulDiv(ctx *parser.MulDivContext) any {
	left := c.Visit(ctx.Expr(0))
	right := c.Visit(ctx.Expr(1))

	op := ctx.GetChild(1).(antlr.TerminalNode).GetText()

	if op == "*" {
		return left.(int) * right.(int)
	} else {
		return left.(int) / right.(int)
	}
}

func (c *Calculator) VisitAddSub(ctx *parser.AddSubContext) any {
	left := c.Visit(ctx.Expr(0))
	right := c.Visit(ctx.Expr(1))

	op := ctx.GetChild(1).(antlr.TerminalNode).GetText()

	if op == "+" {
		return left.(int) + right.(int)
	} else {
		return left.(int) - right.(int)
	}
}

func (c *Calculator) VisitParens(ctx *parser.ParensContext) any {
	return c.Visit(ctx.Expr()).(int)
}

func (c *Calculator) VisitVar(ctx *parser.VarContext) any {
	return c.Variables[ctx.GetText()]
}

func (c *Calculator) VisitInt(ctx *parser.IntContext) any {
	val, _ := strconv.Atoi(ctx.GetText())
	return val
}
