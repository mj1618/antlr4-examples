// Code generated from ./Calc.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Calc

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by CalcParser.
type CalcVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by CalcParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by CalcParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}
}
