package calc0

import (
	"calcPlus/internal/parser"
	"github.com/antlr4-go/antlr/v4"
)

type PostfixL struct {
	parser.BaseCalcPlusListener
	PostfixExpr string
}

func (p *PostfixL) ExitInt(ctx *parser.IntContext) {
	p.PostfixExpr += ctx.GetText() + " "
}

func (p *PostfixL) ExitAddSub(ctx *parser.AddSubContext) {
	op := ctx.GetChild(1).(antlr.TerminalNode).GetText()
	p.PostfixExpr += op + " "
}

func (p *PostfixL) ExitMulDiv(ctx *parser.MulDivContext) {
	op := ctx.GetChild(1).(antlr.TerminalNode).GetText()
	p.PostfixExpr += op + " "
}
