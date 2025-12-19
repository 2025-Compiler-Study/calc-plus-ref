// Code generated from CalcPlus.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // CalcPlus
import "github.com/antlr4-go/antlr/v4"

type BaseCalcPlusVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseCalcPlusVisitor) VisitCalc0(ctx *Calc0Context) interface{} {
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

func (v *BaseCalcPlusVisitor) VisitCalc1(ctx *Calc1Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalcPlusVisitor) VisitExprAssign(ctx *ExprAssignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalcPlusVisitor) VisitReadAssign(ctx *ReadAssignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalcPlusVisitor) VisitIfElse(ctx *IfElseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalcPlusVisitor) VisitWrite(ctx *WriteContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalcPlusVisitor) VisitCalc2(ctx *Calc2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalcPlusVisitor) VisitCond(ctx *CondContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalcPlusVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalcPlusVisitor) VisitCalc3(ctx *Calc3Context) interface{} {
	return v.VisitChildren(ctx)
}
