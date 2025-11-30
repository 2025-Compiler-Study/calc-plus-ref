// Code generated from CalcPlus.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // CalcPlus
import "github.com/antlr4-go/antlr/v4"

// BaseCalcPlusListener is a complete listener for a parse tree produced by CalcPlusParser.
type BaseCalcPlusListener struct{}

var _ CalcPlusListener = &BaseCalcPlusListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCalcPlusListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCalcPlusListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCalcPlusListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCalcPlusListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterCalc0 is called when production calc0 is entered.
func (s *BaseCalcPlusListener) EnterCalc0(ctx *Calc0Context) {}

// ExitCalc0 is called when production calc0 is exited.
func (s *BaseCalcPlusListener) ExitCalc0(ctx *Calc0Context) {}

// EnterMulDiv is called when production MulDiv is entered.
func (s *BaseCalcPlusListener) EnterMulDiv(ctx *MulDivContext) {}

// ExitMulDiv is called when production MulDiv is exited.
func (s *BaseCalcPlusListener) ExitMulDiv(ctx *MulDivContext) {}

// EnterAddSub is called when production AddSub is entered.
func (s *BaseCalcPlusListener) EnterAddSub(ctx *AddSubContext) {}

// ExitAddSub is called when production AddSub is exited.
func (s *BaseCalcPlusListener) ExitAddSub(ctx *AddSubContext) {}

// EnterVar is called when production Var is entered.
func (s *BaseCalcPlusListener) EnterVar(ctx *VarContext) {}

// ExitVar is called when production Var is exited.
func (s *BaseCalcPlusListener) ExitVar(ctx *VarContext) {}

// EnterParens is called when production Parens is entered.
func (s *BaseCalcPlusListener) EnterParens(ctx *ParensContext) {}

// ExitParens is called when production Parens is exited.
func (s *BaseCalcPlusListener) ExitParens(ctx *ParensContext) {}

// EnterInt is called when production Int is entered.
func (s *BaseCalcPlusListener) EnterInt(ctx *IntContext) {}

// ExitInt is called when production Int is exited.
func (s *BaseCalcPlusListener) ExitInt(ctx *IntContext) {}

// EnterCalc1 is called when production calc1 is entered.
func (s *BaseCalcPlusListener) EnterCalc1(ctx *Calc1Context) {}

// ExitCalc1 is called when production calc1 is exited.
func (s *BaseCalcPlusListener) ExitCalc1(ctx *Calc1Context) {}

// EnterExprAssign is called when production ExprAssign is entered.
func (s *BaseCalcPlusListener) EnterExprAssign(ctx *ExprAssignContext) {}

// ExitExprAssign is called when production ExprAssign is exited.
func (s *BaseCalcPlusListener) ExitExprAssign(ctx *ExprAssignContext) {}
