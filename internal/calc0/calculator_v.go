package calc0

import (
	"calcPlus/internal/parser"
	"github.com/antlr4-go/antlr/v4"
	"strconv"
)

type CalculatorV struct {
	parser.BaseCalcPlusVisitor
	Result int
}

func (c *CalculatorV) Visit(tree antlr.ParseTree) any {
	return tree.Accept(c)
}

func (c *CalculatorV) VisitProg(ctx *parser.ProgContext) any {
	c.Result = c.Visit(ctx.Expr()).(int)
	return c.Result
}

func (c *CalculatorV) VisitMulDiv(ctx *parser.MulDivContext) any {
	left := c.Visit(ctx.Expr(0))
	right := c.Visit(ctx.Expr(1))

	op := ctx.GetChild(1).(antlr.TerminalNode).GetText()

	if op == "*" {
		return left.(int) * right.(int)
	} else {
		return left.(int) / right.(int)
	}
}

func (c *CalculatorV) VisitAddSub(ctx *parser.AddSubContext) any {
	left := c.Visit(ctx.Expr(0))
	right := c.Visit(ctx.Expr(1))

	op := ctx.GetChild(1).(antlr.TerminalNode).GetText()

	if op == "+" {
		return left.(int) + right.(int)
	} else {
		return left.(int) - right.(int)
	}
}

func (c *CalculatorV) VisitParens(ctx *parser.ParensContext) any {
	return c.Visit(ctx.Expr()).(int)
}

func (c *CalculatorV) VisitInt(ctx *parser.IntContext) any {
	val, _ := strconv.Atoi(ctx.GetText())
	return val
}
