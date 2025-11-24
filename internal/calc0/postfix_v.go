package calc0

import (
	"calcPlus/internal/parser"
	"github.com/antlr4-go/antlr/v4"
)

type PostfixV struct {
	parser.BaseCalcPlusVisitor
	PostfixExpr string
}

func (p *PostfixV) Visit(tree antlr.ParseTree) any {
	return tree.Accept(p)
}

func (p *PostfixV) VisitProg(ctx *parser.ProgContext) any {
	p.PostfixExpr = p.Visit(ctx.Expr()).(string)
	return p.PostfixExpr
}

func (p *PostfixV) VisitInt(ctx *parser.IntContext) any {
	return ctx.GetText() + " "
}

func (p *PostfixV) VisitParens(ctx *parser.ParensContext) any {
	return p.Visit(ctx.Expr()).(string)
}

func (p *PostfixV) VisitAddSub(ctx *parser.AddSubContext) any {
	left := p.Visit(ctx.Expr(0)).(string)
	right := p.Visit(ctx.Expr(1)).(string)

	op := ctx.GetChild(1).(antlr.TerminalNode).GetText()
	return left + right + op + " "
}

func (p *PostfixV) VisitMulDiv(ctx *parser.MulDivContext) any {
	left := p.Visit(ctx.Expr(0)).(string)
	right := p.Visit(ctx.Expr(1)).(string)

	op := ctx.GetChild(1).(antlr.TerminalNode).GetText()
	return left + right + op + " "
}
