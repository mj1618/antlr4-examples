// Code generated from ./Calc.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Calc

import "github.com/antlr4-go/antlr/v4"

// BaseCalcListener is a complete listener for a parse tree produced by CalcParser.
type BaseCalcListener struct{}

var _ CalcListener = &BaseCalcListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCalcListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCalcListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCalcListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCalcListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseCalcListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseCalcListener) ExitProgram(ctx *ProgramContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseCalcListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseCalcListener) ExitExpression(ctx *ExpressionContext) {}
