// Code generated from CalcPlus.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // CalcPlus
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by CalcPlusParser.
type CalcPlusVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by CalcPlusParser#calc0.
	VisitCalc0(ctx *Calc0Context) interface{}

	// Visit a parse tree produced by CalcPlusParser#MulDiv.
	VisitMulDiv(ctx *MulDivContext) interface{}

	// Visit a parse tree produced by CalcPlusParser#AddSub.
	VisitAddSub(ctx *AddSubContext) interface{}

	// Visit a parse tree produced by CalcPlusParser#Var.
	VisitVar(ctx *VarContext) interface{}

	// Visit a parse tree produced by CalcPlusParser#Parens.
	VisitParens(ctx *ParensContext) interface{}

	// Visit a parse tree produced by CalcPlusParser#Int.
	VisitInt(ctx *IntContext) interface{}

	// Visit a parse tree produced by CalcPlusParser#calc1.
	VisitCalc1(ctx *Calc1Context) interface{}

	// Visit a parse tree produced by CalcPlusParser#ExprAssign.
	VisitExprAssign(ctx *ExprAssignContext) interface{}

	// Visit a parse tree produced by CalcPlusParser#ReadAssign.
	VisitReadAssign(ctx *ReadAssignContext) interface{}

	// Visit a parse tree produced by CalcPlusParser#IfElse.
	VisitIfElse(ctx *IfElseContext) interface{}

	// Visit a parse tree produced by CalcPlusParser#Write.
	VisitWrite(ctx *WriteContext) interface{}

	// Visit a parse tree produced by CalcPlusParser#Declare.
	VisitDeclare(ctx *DeclareContext) interface{}

	// Visit a parse tree produced by CalcPlusParser#StmtBlock.
	VisitStmtBlock(ctx *StmtBlockContext) interface{}

	// Visit a parse tree produced by CalcPlusParser#calc2.
	VisitCalc2(ctx *Calc2Context) interface{}

	// Visit a parse tree produced by CalcPlusParser#cond.
	VisitCond(ctx *CondContext) interface{}

	// Visit a parse tree produced by CalcPlusParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by CalcPlusParser#calc3.
	VisitCalc3(ctx *Calc3Context) interface{}

	// Visit a parse tree produced by CalcPlusParser#calc4.
	VisitCalc4(ctx *Calc4Context) interface{}
}
