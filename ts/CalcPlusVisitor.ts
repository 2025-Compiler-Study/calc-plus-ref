// Generated from ../internal/grammar/CalcPlus.g4 by ANTLR 4.9.0-SNAPSHOT


import { ParseTreeVisitor } from "antlr4ts/tree/ParseTreeVisitor";

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
 * This interface defines a complete generic visitor for a parse tree produced
 * by `CalcPlusParser`.
 *
 * @param <Result> The return type of the visit operation. Use `void` for
 * operations with no return type.
 */
export interface CalcPlusVisitor<Result> extends ParseTreeVisitor<Result> {
	/**
	 * Visit a parse tree produced by the `MulDiv`
	 * labeled alternative in `CalcPlusParser.expr`.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	visitMulDiv?: (ctx: MulDivContext) => Result;

	/**
	 * Visit a parse tree produced by the `AddSub`
	 * labeled alternative in `CalcPlusParser.expr`.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	visitAddSub?: (ctx: AddSubContext) => Result;

	/**
	 * Visit a parse tree produced by the `Int`
	 * labeled alternative in `CalcPlusParser.expr`.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	visitInt?: (ctx: IntContext) => Result;

	/**
	 * Visit a parse tree produced by the `Var`
	 * labeled alternative in `CalcPlusParser.expr`.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	visitVar?: (ctx: VarContext) => Result;

	/**
	 * Visit a parse tree produced by the `Parens`
	 * labeled alternative in `CalcPlusParser.expr`.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	visitParens?: (ctx: ParensContext) => Result;

	/**
	 * Visit a parse tree produced by the `ExprAssign`
	 * labeled alternative in `CalcPlusParser.stmt`.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	visitExprAssign?: (ctx: ExprAssignContext) => Result;

	/**
	 * Visit a parse tree produced by the `ReadAssign`
	 * labeled alternative in `CalcPlusParser.stmt`.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	visitReadAssign?: (ctx: ReadAssignContext) => Result;

	/**
	 * Visit a parse tree produced by the `IfElse`
	 * labeled alternative in `CalcPlusParser.stmt`.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	visitIfElse?: (ctx: IfElseContext) => Result;

	/**
	 * Visit a parse tree produced by the `Write`
	 * labeled alternative in `CalcPlusParser.stmt`.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	visitWrite?: (ctx: WriteContext) => Result;

	/**
	 * Visit a parse tree produced by `CalcPlusParser.calc0`.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	visitCalc0?: (ctx: Calc0Context) => Result;

	/**
	 * Visit a parse tree produced by `CalcPlusParser.expr`.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	visitExpr?: (ctx: ExprContext) => Result;

	/**
	 * Visit a parse tree produced by `CalcPlusParser.calc1`.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	visitCalc1?: (ctx: Calc1Context) => Result;

	/**
	 * Visit a parse tree produced by `CalcPlusParser.stmt`.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	visitStmt?: (ctx: StmtContext) => Result;

	/**
	 * Visit a parse tree produced by `CalcPlusParser.calc2`.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	visitCalc2?: (ctx: Calc2Context) => Result;

	/**
	 * Visit a parse tree produced by `CalcPlusParser.cond`.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	visitCond?: (ctx: CondContext) => Result;

	/**
	 * Visit a parse tree produced by `CalcPlusParser.block`.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	visitBlock?: (ctx: BlockContext) => Result;

	/**
	 * Visit a parse tree produced by `CalcPlusParser.calc3`.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	visitCalc3?: (ctx: Calc3Context) => Result;
}

