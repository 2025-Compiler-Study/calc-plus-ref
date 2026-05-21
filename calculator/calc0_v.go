package calculator

import (
	"calcPlus/internal/parser"
	"github.com/antlr4-go/antlr/v4"
	"strconv"
)

type Calc0V struct {
	parser.BaseCalcPlusVisitor
	Result int
}

func (c *Calc0V) Visit(tree antlr.ParseTree) any {
	return tree.Accept(c)
}

func (c *Calc0V) VisitMulDiv(ctx *parser.MulDivContext) any {
	left := c.Visit(ctx.Expr(0))
	right := c.Visit(ctx.Expr(1))

	op := ctx.GetChild(1).(antlr.TerminalNode).GetText()

	if op == "*" {
		return left.(int) * right.(int)
	} else {
		return left.(int) / right.(int)
	}
}

func (c *Calc0V) VisitAddSub(ctx *parser.AddSubContext) any {
	left := c.Visit(ctx.Expr(0))
	right := c.Visit(ctx.Expr(1))

	op := ctx.GetChild(1).(antlr.TerminalNode).GetText()

	if op == "+" {
		return left.(int) + right.(int)
	} else {
		return left.(int) - right.(int)
	}
}

func (c *Calc0V) VisitParens(ctx *parser.ParensContext) any {
	return c.Visit(ctx.Expr()).(int)
}

func (c *Calc0V) VisitInt(ctx *parser.IntContext) any {
	val, _ := strconv.Atoi(ctx.GetText())
	return val
}
