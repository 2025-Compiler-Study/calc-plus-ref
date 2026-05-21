// Code generated from CalcPlus.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // CalcPlus
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by CalcPlusParser.
type CalcPlusVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by CalcPlusParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by CalcPlusParser#Declare.
	VisitDeclare(ctx *DeclareContext) interface{}

	// Visit a parse tree produced by CalcPlusParser#ExprAssign.
	VisitExprAssign(ctx *ExprAssignContext) interface{}

	// Visit a parse tree produced by CalcPlusParser#ReadAssign.
	VisitReadAssign(ctx *ReadAssignContext) interface{}

	// Visit a parse tree produced by CalcPlusParser#Write.
	VisitWrite(ctx *WriteContext) interface{}

	// Visit a parse tree produced by CalcPlusParser#IfElse.
	VisitIfElse(ctx *IfElseContext) interface{}

	// Visit a parse tree produced by CalcPlusParser#StmtBlock.
	VisitStmtBlock(ctx *StmtBlockContext) interface{}

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

	// Visit a parse tree produced by CalcPlusParser#cond.
	VisitCond(ctx *CondContext) interface{}

	// Visit a parse tree produced by CalcPlusParser#block.
	VisitBlock(ctx *BlockContext) interface{}
}
