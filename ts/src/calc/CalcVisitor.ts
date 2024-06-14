// Generated from ./Calc.g4 by ANTLR 4.13.1

import {ParseTreeVisitor} from 'antlr4';


import { ProgramContext } from "./CalcParser";
import { ExpressionContext } from "./CalcParser";


/**
 * This interface defines a complete generic visitor for a parse tree produced
 * by `CalcParser`.
 *
 * @param <Result> The return type of the visit operation. Use `void` for
 * operations with no return type.
 */
export default class CalcVisitor<Result> extends ParseTreeVisitor<Result> {
	/**
	 * Visit a parse tree produced by `CalcParser.program`.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	visitProgram?: (ctx: ProgramContext) => Result;
	/**
	 * Visit a parse tree produced by `CalcParser.expression`.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	visitExpression?: (ctx: ExpressionContext) => Result;
}

