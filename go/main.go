// example1.go
package main

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr4-go/antlr/v4"

	parser "example/antlr-lisp/calc"
)

var lispParser *parser.CalcParser

type CalcListener struct {
	parser.BaseCalcListener
}

type CalcVisitor struct {
	parser.BaseCalcVisitor
}

func (v *CalcVisitor) VisitProgram(ctx *parser.ProgramContext) interface{} {
	// fmt.Println("VisitProgram")
	// ctx.Child
	for eIdx := range ctx.AllExpression() {
		res := v.VisitExpression(ctx.Expression(eIdx).(*parser.ExpressionContext))
		// res := ctx.Expression(eIdx).Accept(v)
		fmt.Println(res)
	}

	// return v.Visit(ctx.Expression(0))
	return nil
}

func reduce(numbers []int, operand string, acc int) int {
	sum := acc
	for _, num := range numbers {
		switch operand {
		case "+":
			sum += num
		case "-":
			sum -= num
		case "*":
			sum *= num
		case "/":
			sum /= num
		}
	}
	return sum
}

func (v *CalcVisitor) VisitExpression(ctx *parser.ExpressionContext) interface{} {
	// fmt.Println("VisitExpression")

	numberList := []int{}
	children := ctx.GetChildren()
	for idx := range children {
		child := ctx.GetChild(idx)
		// fmt.Println(reflect.TypeOf(child))

		switch child := child.(type) {
		case *parser.ExpressionContext:
			res := v.VisitExpression(child)
			// fmt.Println(res.(int))
			numberList = append(numberList, res.(int))
		case *antlr.TerminalNodeImpl:
			// fmt.Println(child.(*antlr.TerminalNodeImpl).GetSymbol().GetText())
			if child.GetSymbol().GetTokenType() == parser.CalcParserNUMBER {

				number, _ := strconv.Atoi(child.GetSymbol().GetText())

				numberList = append(numberList, number)
			}
		default:
			panic(fmt.Sprintf("unexpected type %s", reflect.TypeOf(child)))

		}
	}

	operand := ctx.GetOperand().GetText()
	fmt.Println(numberList, operand)
	switch operand {
	case "+":
		return reduce(numberList, operand, 0)
	case "-":
		return reduce(numberList, operand, 0)
	case "*":
		return reduce(numberList, operand, 1)
	case "/":
		return reduce(numberList[1:], operand, numberList[0])

	default:
		return nil
	}

}

func main() {

	is := antlr.NewInputStream("(/ 48 (* 4 (+ 1 2 3)))")
	lexer := parser.NewCalcLexer(is)
	lispParser = parser.NewCalcParser(antlr.NewCommonTokenStream(lexer, 0))
	v := &CalcVisitor{}
	v.VisitProgram(lispParser.Program().(*parser.ProgramContext))

}
