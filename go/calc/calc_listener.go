// Code generated from ./Calc.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Calc

import "github.com/antlr4-go/antlr/v4"

// CalcListener is a complete listener for a parse tree produced by CalcParser.
type CalcListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)
}
