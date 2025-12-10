// Code generated from CalcPlus.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // CalcPlus
import "github.com/antlr4-go/antlr/v4"

// CalcPlusListener is a complete listener for a parse tree produced by CalcPlusParser.
type CalcPlusListener interface {
	antlr.ParseTreeListener

	// EnterCalc0 is called when entering the calc0 production.
	EnterCalc0(c *Calc0Context)

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

	// EnterCalc1 is called when entering the calc1 production.
	EnterCalc1(c *Calc1Context)

	// EnterExprAssign is called when entering the ExprAssign production.
	EnterExprAssign(c *ExprAssignContext)

	// EnterIfElse is called when entering the IfElse production.
	EnterIfElse(c *IfElseContext)

	// EnterCalc2 is called when entering the calc2 production.
	EnterCalc2(c *Calc2Context)

	// EnterCond is called when entering the cond production.
	EnterCond(c *CondContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// ExitCalc0 is called when exiting the calc0 production.
	ExitCalc0(c *Calc0Context)

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

	// ExitCalc1 is called when exiting the calc1 production.
	ExitCalc1(c *Calc1Context)

	// ExitExprAssign is called when exiting the ExprAssign production.
	ExitExprAssign(c *ExprAssignContext)

	// ExitIfElse is called when exiting the IfElse production.
	ExitIfElse(c *IfElseContext)

	// ExitCalc2 is called when exiting the calc2 production.
	ExitCalc2(c *Calc2Context)

	// ExitCond is called when exiting the cond production.
	ExitCond(c *CondContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)
}
