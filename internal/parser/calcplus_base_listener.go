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

// EnterProgram is called when production program is entered.
func (s *BaseCalcPlusListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseCalcPlusListener) ExitProgram(ctx *ProgramContext) {}

// EnterDeclare is called when production Declare is entered.
func (s *BaseCalcPlusListener) EnterDeclare(ctx *DeclareContext) {}

// ExitDeclare is called when production Declare is exited.
func (s *BaseCalcPlusListener) ExitDeclare(ctx *DeclareContext) {}

// EnterExprAssign is called when production ExprAssign is entered.
func (s *BaseCalcPlusListener) EnterExprAssign(ctx *ExprAssignContext) {}

// ExitExprAssign is called when production ExprAssign is exited.
func (s *BaseCalcPlusListener) ExitExprAssign(ctx *ExprAssignContext) {}

// EnterReadAssign is called when production ReadAssign is entered.
func (s *BaseCalcPlusListener) EnterReadAssign(ctx *ReadAssignContext) {}

// ExitReadAssign is called when production ReadAssign is exited.
func (s *BaseCalcPlusListener) ExitReadAssign(ctx *ReadAssignContext) {}

// EnterWrite is called when production Write is entered.
func (s *BaseCalcPlusListener) EnterWrite(ctx *WriteContext) {}

// ExitWrite is called when production Write is exited.
func (s *BaseCalcPlusListener) ExitWrite(ctx *WriteContext) {}

// EnterIfElse is called when production IfElse is entered.
func (s *BaseCalcPlusListener) EnterIfElse(ctx *IfElseContext) {}

// ExitIfElse is called when production IfElse is exited.
func (s *BaseCalcPlusListener) ExitIfElse(ctx *IfElseContext) {}

// EnterStmtBlock is called when production StmtBlock is entered.
func (s *BaseCalcPlusListener) EnterStmtBlock(ctx *StmtBlockContext) {}

// ExitStmtBlock is called when production StmtBlock is exited.
func (s *BaseCalcPlusListener) ExitStmtBlock(ctx *StmtBlockContext) {}

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

// EnterCond is called when production cond is entered.
func (s *BaseCalcPlusListener) EnterCond(ctx *CondContext) {}

// ExitCond is called when production cond is exited.
func (s *BaseCalcPlusListener) ExitCond(ctx *CondContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseCalcPlusListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseCalcPlusListener) ExitBlock(ctx *BlockContext) {}
