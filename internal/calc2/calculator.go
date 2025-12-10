package calc2

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

func (c *Calculator) VisitCalc2(ctx *parser.Calc2Context) any {
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

func (c *Calculator) VisitIfElse(ctx *parser.IfElseContext) any {
	if c.Visit(ctx.Cond()).(bool) {
		c.Visit(ctx.GetThenBlock())
	} else if ctx.GetElseBlock() != nil {
		c.Visit(ctx.GetElseBlock())
	}

	return nil
}

func (c *Calculator) VisitMulDiv(ctx *parser.MulDivContext) any {
	left := c.Visit(ctx.Expr(0)).(int)
	right := c.Visit(ctx.Expr(1)).(int)

	op := ctx.GetChild(1).(antlr.TerminalNode).GetText()
	if op == "*" {
		return left * right
	} else {
		return left / right
	}
}

func (c *Calculator) VisitAddSub(ctx *parser.AddSubContext) any {
	left := c.Visit(ctx.Expr(0)).(int)
	right := c.Visit(ctx.Expr(1)).(int)

	op := ctx.GetChild(1).(antlr.TerminalNode).GetText()

	if op == "+" {
		return left + right
	} else {
		return left - right
	}
}

func (c *Calculator) VisitInt(ctx *parser.IntContext) any {
	val, _ := strconv.Atoi(ctx.GetText())
	return val
}

func (c *Calculator) VisitVar(ctx *parser.VarContext) any {
	return c.Variables[ctx.GetText()]
}

func (c *Calculator) VisitParens(ctx *parser.ParensContext) any {
	return c.Visit(ctx.Expr()).(int)
}

func (c *Calculator) VisitCond(ctx *parser.CondContext) any {
	left := c.Visit(ctx.Expr(0)).(int)
	right := c.Visit(ctx.Expr(1)).(int)

	op := ctx.GetChild(1).(antlr.TerminalNode).GetText()
	switch op {
	case "==":
		return left == right
	case "!=":
		return left != right
	case ">":
		return left > right
	case ">=":
		return left >= right
	case "<":
		return left < right
	case "<=":
		return left <= right
	}
	return false
}

func (c *Calculator) VisitBlock(ctx *parser.BlockContext) any {
	for _, stmt := range ctx.AllStmt() {
		c.Visit(stmt)
	}
	return nil
}
