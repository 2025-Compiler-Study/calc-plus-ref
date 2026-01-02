package postfix

import (
	"calcPlus/internal/parser"
	"github.com/antlr4-go/antlr/v4"
)

type Calc0PostfixV struct {
	parser.BaseCalcPlusVisitor
	PostfixExpr string
}

func (p *Calc0PostfixV) Visit(tree antlr.ParseTree) any {
	return tree.Accept(p)
}

func (p *Calc0PostfixV) VisitInt(ctx *parser.IntContext) any {
	return ctx.GetText() + " "
}

func (p *Calc0PostfixV) VisitParens(ctx *parser.ParensContext) any {
	return p.Visit(ctx.Expr()).(string)
}

func (p *Calc0PostfixV) VisitAddSub(ctx *parser.AddSubContext) any {
	left := p.Visit(ctx.Expr(0)).(string)
	right := p.Visit(ctx.Expr(1)).(string)

	op := ctx.GetChild(1).(antlr.TerminalNode).GetText()
	return left + right + op + " "
}

func (p *Calc0PostfixV) VisitMulDiv(ctx *parser.MulDivContext) any {
	left := p.Visit(ctx.Expr(0)).(string)
	right := p.Visit(ctx.Expr(1)).(string)

	op := ctx.GetChild(1).(antlr.TerminalNode).GetText()
	return left + right + op + " "
}
