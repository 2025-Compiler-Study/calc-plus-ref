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
}
