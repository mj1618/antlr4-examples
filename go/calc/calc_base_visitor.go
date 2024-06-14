// Code generated from ./Calc.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Calc

import "github.com/antlr4-go/antlr/v4"

type BaseCalcVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseCalcVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalcVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}
