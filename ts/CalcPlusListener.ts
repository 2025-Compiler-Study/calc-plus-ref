// Generated from ../internal/grammar/CalcPlus.g4 by ANTLR 4.9.0-SNAPSHOT


import { ParseTreeListener } from "antlr4ts/tree/ParseTreeListener";

import { MulDivContext } from "./CalcPlusParser";
import { AddSubContext } from "./CalcPlusParser";
import { IntContext } from "./CalcPlusParser";
import { VarContext } from "./CalcPlusParser";
import { ParensContext } from "./CalcPlusParser";
import { ExprAssignContext } from "./CalcPlusParser";
import { ReadAssignContext } from "./CalcPlusParser";
import { IfElseContext } from "./CalcPlusParser";
import { WriteContext } from "./CalcPlusParser";
import { Calc0Context } from "./CalcPlusParser";
import { ExprContext } from "./CalcPlusParser";
import { Calc1Context } from "./CalcPlusParser";
import { StmtContext } from "./CalcPlusParser";
import { Calc2Context } from "./CalcPlusParser";
import { CondContext } from "./CalcPlusParser";
import { BlockContext } from "./CalcPlusParser";
import { Calc3Context } from "./CalcPlusParser";


/**
 * This interface defines a complete listener for a parse tree produced by
 * `CalcPlusParser`.
 */
export interface CalcPlusListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by the `MulDiv`
	 * labeled alternative in `CalcPlusParser.expr`.
	 * @param ctx the parse tree
	 */
	enterMulDiv?: (ctx: MulDivContext) => void;
	/**
	 * Exit a parse tree produced by the `MulDiv`
	 * labeled alternative in `CalcPlusParser.expr`.
	 * @param ctx the parse tree
	 */
	exitMulDiv?: (ctx: MulDivContext) => void;

	/**
	 * Enter a parse tree produced by the `AddSub`
	 * labeled alternative in `CalcPlusParser.expr`.
	 * @param ctx the parse tree
	 */
	enterAddSub?: (ctx: AddSubContext) => void;
	/**
	 * Exit a parse tree produced by the `AddSub`
	 * labeled alternative in `CalcPlusParser.expr`.
	 * @param ctx the parse tree
	 */
	exitAddSub?: (ctx: AddSubContext) => void;

	/**
	 * Enter a parse tree produced by the `Int`
	 * labeled alternative in `CalcPlusParser.expr`.
	 * @param ctx the parse tree
	 */
	enterInt?: (ctx: IntContext) => void;
	/**
	 * Exit a parse tree produced by the `Int`
	 * labeled alternative in `CalcPlusParser.expr`.
	 * @param ctx the parse tree
	 */
	exitInt?: (ctx: IntContext) => void;

	/**
	 * Enter a parse tree produced by the `Var`
	 * labeled alternative in `CalcPlusParser.expr`.
	 * @param ctx the parse tree
	 */
	enterVar?: (ctx: VarContext) => void;
	/**
	 * Exit a parse tree produced by the `Var`
	 * labeled alternative in `CalcPlusParser.expr`.
	 * @param ctx the parse tree
	 */
	exitVar?: (ctx: VarContext) => void;

	/**
	 * Enter a parse tree produced by the `Parens`
	 * labeled alternative in `CalcPlusParser.expr`.
	 * @param ctx the parse tree
	 */
	enterParens?: (ctx: ParensContext) => void;
	/**
	 * Exit a parse tree produced by the `Parens`
	 * labeled alternative in `CalcPlusParser.expr`.
	 * @param ctx the parse tree
	 */
	exitParens?: (ctx: ParensContext) => void;

	/**
	 * Enter a parse tree produced by the `ExprAssign`
	 * labeled alternative in `CalcPlusParser.stmt`.
	 * @param ctx the parse tree
	 */
	enterExprAssign?: (ctx: ExprAssignContext) => void;
	/**
	 * Exit a parse tree produced by the `ExprAssign`
	 * labeled alternative in `CalcPlusParser.stmt`.
	 * @param ctx the parse tree
	 */
	exitExprAssign?: (ctx: ExprAssignContext) => void;

	/**
	 * Enter a parse tree produced by the `ReadAssign`
	 * labeled alternative in `CalcPlusParser.stmt`.
	 * @param ctx the parse tree
	 */
	enterReadAssign?: (ctx: ReadAssignContext) => void;
	/**
	 * Exit a parse tree produced by the `ReadAssign`
	 * labeled alternative in `CalcPlusParser.stmt`.
	 * @param ctx the parse tree
	 */
	exitReadAssign?: (ctx: ReadAssignContext) => void;

	/**
	 * Enter a parse tree produced by the `IfElse`
	 * labeled alternative in `CalcPlusParser.stmt`.
	 * @param ctx the parse tree
	 */
	enterIfElse?: (ctx: IfElseContext) => void;
	/**
	 * Exit a parse tree produced by the `IfElse`
	 * labeled alternative in `CalcPlusParser.stmt`.
	 * @param ctx the parse tree
	 */
	exitIfElse?: (ctx: IfElseContext) => void;

	/**
	 * Enter a parse tree produced by the `Write`
	 * labeled alternative in `CalcPlusParser.stmt`.
	 * @param ctx the parse tree
	 */
	enterWrite?: (ctx: WriteContext) => void;
	/**
	 * Exit a parse tree produced by the `Write`
	 * labeled alternative in `CalcPlusParser.stmt`.
	 * @param ctx the parse tree
	 */
	exitWrite?: (ctx: WriteContext) => void;

	/**
	 * Enter a parse tree produced by `CalcPlusParser.calc0`.
	 * @param ctx the parse tree
	 */
	enterCalc0?: (ctx: Calc0Context) => void;
	/**
	 * Exit a parse tree produced by `CalcPlusParser.calc0`.
	 * @param ctx the parse tree
	 */
	exitCalc0?: (ctx: Calc0Context) => void;

	/**
	 * Enter a parse tree produced by `CalcPlusParser.expr`.
	 * @param ctx the parse tree
	 */
	enterExpr?: (ctx: ExprContext) => void;
	/**
	 * Exit a parse tree produced by `CalcPlusParser.expr`.
	 * @param ctx the parse tree
	 */
	exitExpr?: (ctx: ExprContext) => void;

	/**
	 * Enter a parse tree produced by `CalcPlusParser.calc1`.
	 * @param ctx the parse tree
	 */
	enterCalc1?: (ctx: Calc1Context) => void;
	/**
	 * Exit a parse tree produced by `CalcPlusParser.calc1`.
	 * @param ctx the parse tree
	 */
	exitCalc1?: (ctx: Calc1Context) => void;

	/**
	 * Enter a parse tree produced by `CalcPlusParser.stmt`.
	 * @param ctx the parse tree
	 */
	enterStmt?: (ctx: StmtContext) => void;
	/**
	 * Exit a parse tree produced by `CalcPlusParser.stmt`.
	 * @param ctx the parse tree
	 */
	exitStmt?: (ctx: StmtContext) => void;

	/**
	 * Enter a parse tree produced by `CalcPlusParser.calc2`.
	 * @param ctx the parse tree
	 */
	enterCalc2?: (ctx: Calc2Context) => void;
	/**
	 * Exit a parse tree produced by `CalcPlusParser.calc2`.
	 * @param ctx the parse tree
	 */
	exitCalc2?: (ctx: Calc2Context) => void;

	/**
	 * Enter a parse tree produced by `CalcPlusParser.cond`.
	 * @param ctx the parse tree
	 */
	enterCond?: (ctx: CondContext) => void;
	/**
	 * Exit a parse tree produced by `CalcPlusParser.cond`.
	 * @param ctx the parse tree
	 */
	exitCond?: (ctx: CondContext) => void;

	/**
	 * Enter a parse tree produced by `CalcPlusParser.block`.
	 * @param ctx the parse tree
	 */
	enterBlock?: (ctx: BlockContext) => void;
	/**
	 * Exit a parse tree produced by `CalcPlusParser.block`.
	 * @param ctx the parse tree
	 */
	exitBlock?: (ctx: BlockContext) => void;

	/**
	 * Enter a parse tree produced by `CalcPlusParser.calc3`.
	 * @param ctx the parse tree
	 */
	enterCalc3?: (ctx: Calc3Context) => void;
	/**
	 * Exit a parse tree produced by `CalcPlusParser.calc3`.
	 * @param ctx the parse tree
	 */
	exitCalc3?: (ctx: Calc3Context) => void;
}

