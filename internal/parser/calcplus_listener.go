// Code generated from CalcPlus.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // CalcPlus
import "github.com/antlr4-go/antlr/v4"

// CalcPlusListener is a complete listener for a parse tree produced by CalcPlusParser.
type CalcPlusListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterDeclare is called when entering the Declare production.
	EnterDeclare(c *DeclareContext)

	// EnterExprAssign is called when entering the ExprAssign production.
	EnterExprAssign(c *ExprAssignContext)

	// EnterReadAssign is called when entering the ReadAssign production.
	EnterReadAssign(c *ReadAssignContext)

	// EnterWrite is called when entering the Write production.
	EnterWrite(c *WriteContext)

	// EnterIfElse is called when entering the IfElse production.
	EnterIfElse(c *IfElseContext)

	// EnterStmtBlock is called when entering the StmtBlock production.
	EnterStmtBlock(c *StmtBlockContext)

	// EnterMulDiv is called when entering the MulDiv production.
	EnterMulDiv(c *MulDivContext)

	// EnterAddSub is called when entering the AddSub production.
	EnterAddSub(c *AddSubContext)

	// EnterVar is called when entering the Var production.
	EnterVar(c *VarContext)

	// EnterParens is called when entering the Parens production.
	EnterParens(c *ParensContext)

	// EnterInt is called when entering the Int production.
	EnterInt(c *IntContext)

	// EnterCond is called when entering the cond production.
	EnterCond(c *CondContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitDeclare is called when exiting the Declare production.
	ExitDeclare(c *DeclareContext)

	// ExitExprAssign is called when exiting the ExprAssign production.
	ExitExprAssign(c *ExprAssignContext)

	// ExitReadAssign is called when exiting the ReadAssign production.
	ExitReadAssign(c *ReadAssignContext)

	// ExitWrite is called when exiting the Write production.
	ExitWrite(c *WriteContext)

	// ExitIfElse is called when exiting the IfElse production.
	ExitIfElse(c *IfElseContext)

	// ExitStmtBlock is called when exiting the StmtBlock production.
	ExitStmtBlock(c *StmtBlockContext)

	// ExitMulDiv is called when exiting the MulDiv production.
	ExitMulDiv(c *MulDivContext)

	// ExitAddSub is called when exiting the AddSub production.
	ExitAddSub(c *AddSubContext)

	// ExitVar is called when exiting the Var production.
	ExitVar(c *VarContext)

	// ExitParens is called when exiting the Parens production.
	ExitParens(c *ParensContext)

	// ExitInt is called when exiting the Int production.
	ExitInt(c *IntContext)

	// ExitCond is called when exiting the cond production.
	ExitCond(c *CondContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)
}
