package postfix

import (
	"calcPlus/internal/parser"
	"github.com/antlr4-go/antlr/v4"
)

type Calc0PostfixL struct {
	parser.BaseCalcPlusListener
	PostfixExpr string
}

func (p *Calc0PostfixL) ExitInt(ctx *parser.IntContext) {
	p.PostfixExpr += ctx.GetText() + " "
}

func (p *Calc0PostfixL) ExitAddSub(ctx *parser.AddSubContext) {
	op := ctx.GetChild(1).(antlr.TerminalNode).GetText()
	p.PostfixExpr += op + " "
}

func (p *Calc0PostfixL) ExitMulDiv(ctx *parser.MulDivContext) {
	op := ctx.GetChild(1).(antlr.TerminalNode).GetText()
	p.PostfixExpr += op + " "
}
