// Generated from ./Calc.g4 by ANTLR 4.13.1

import {ParseTreeListener} from "antlr4";


import { ProgramContext } from "./CalcParser";
import { ExpressionContext } from "./CalcParser";


/**
 * This interface defines a complete listener for a parse tree produced by
 * `CalcParser`.
 */
export default class CalcListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by `CalcParser.program`.
	 * @param ctx the parse tree
	 */
	enterProgram?: (ctx: ProgramContext) => void;
	/**
	 * Exit a parse tree produced by `CalcParser.program`.
	 * @param ctx the parse tree
	 */
	exitProgram?: (ctx: ProgramContext) => void;
	/**
	 * Enter a parse tree produced by `CalcParser.expression`.
	 * @param ctx the parse tree
	 */
	enterExpression?: (ctx: ExpressionContext) => void;
	/**
	 * Exit a parse tree produced by `CalcParser.expression`.
	 * @param ctx the parse tree
	 */
	exitExpression?: (ctx: ExpressionContext) => void;
}

