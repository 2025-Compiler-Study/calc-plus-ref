// Code generated from CalcPlus.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // CalcPlus
import "github.com/antlr4-go/antlr/v4"

type BaseCalcPlusVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseCalcPlusVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalcPlusVisitor) VisitDeclare(ctx *DeclareContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalcPlusVisitor) VisitExprAssign(ctx *ExprAssignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalcPlusVisitor) VisitReadAssign(ctx *ReadAssignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalcPlusVisitor) VisitWrite(ctx *WriteContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalcPlusVisitor) VisitIfElse(ctx *IfElseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalcPlusVisitor) VisitStmtBlock(ctx *StmtBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalcPlusVisitor) VisitMulDiv(ctx *MulDivContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalcPlusVisitor) VisitAddSub(ctx *AddSubContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalcPlusVisitor) VisitVar(ctx *VarContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalcPlusVisitor) VisitParens(ctx *ParensContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalcPlusVisitor) VisitInt(ctx *IntContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalcPlusVisitor) VisitCond(ctx *CondContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalcPlusVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}
