package calc1

import "calcPlus/internal/parser"

type TokenPosition struct {
	Line  int
	Col   int
	Value string
}

type Linter struct {
	parser.BaseCalcPlusListener
	Variables map[string]bool
	Errors    []TokenPosition
}

func NewLinter() *Linter {
	return &Linter{
		BaseCalcPlusListener: parser.BaseCalcPlusListener{},
		Variables:            make(map[string]bool),
		Errors:               make([]TokenPosition, 0),
	}
}

func (l *Linter) ExitVar(ctx *parser.VarContext) {
	if !l.Variables[ctx.GetText()] {
		// When the variable are not registered
		errorPosition := TokenPosition{
			Line:  ctx.GetStart().GetLine(),
			Col:   ctx.GetStart().GetColumn(),
			Value: ctx.GetText(),
		}
		l.Errors = append(l.Errors, errorPosition)
	}
}

func (l *Linter) ExitExprAssign(ctx *parser.ExprAssignContext) {
	// Register variable
	l.Variables[ctx.VAR().GetText()] = true
}
