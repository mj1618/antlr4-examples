// Code generated from ./Calc.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Calc

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type CalcParser struct {
	*antlr.BaseParser
}

var CalcParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func calcParserInit() {
	staticData := &CalcParserStaticData
	staticData.LiteralNames = []string{
		"", "", "'+'", "'-'", "'*'", "'/'", "'('", "')'",
	}
	staticData.SymbolicNames = []string{
		"", "IDENTIFIER", "PLUS", "MINUS", "MULT", "DIV", "OP", "CP", "STRING",
		"NUMBER", "WHITESPACE", "DIGIT", "LETTER", "LOWER", "UPPER",
	}
	staticData.RuleNames = []string{
		"program", "expression",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 14, 26, 2, 0, 7, 0, 2, 1, 7, 1, 1, 0, 5, 0, 6, 8, 0, 10, 0, 12, 0,
		9, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 19, 8, 1,
		10, 1, 12, 1, 22, 9, 1, 1, 1, 1, 1, 1, 1, 0, 0, 2, 0, 2, 0, 0, 28, 0, 7,
		1, 0, 0, 0, 2, 12, 1, 0, 0, 0, 4, 6, 3, 2, 1, 0, 5, 4, 1, 0, 0, 0, 6, 9,
		1, 0, 0, 0, 7, 5, 1, 0, 0, 0, 7, 8, 1, 0, 0, 0, 8, 10, 1, 0, 0, 0, 9, 7,
		1, 0, 0, 0, 10, 11, 5, 0, 0, 1, 11, 1, 1, 0, 0, 0, 12, 13, 5, 6, 0, 0,
		13, 20, 5, 1, 0, 0, 14, 19, 5, 1, 0, 0, 15, 19, 5, 8, 0, 0, 16, 19, 5,
		9, 0, 0, 17, 19, 3, 2, 1, 0, 18, 14, 1, 0, 0, 0, 18, 15, 1, 0, 0, 0, 18,
		16, 1, 0, 0, 0, 18, 17, 1, 0, 0, 0, 19, 22, 1, 0, 0, 0, 20, 18, 1, 0, 0,
		0, 20, 21, 1, 0, 0, 0, 21, 23, 1, 0, 0, 0, 22, 20, 1, 0, 0, 0, 23, 24,
		5, 7, 0, 0, 24, 3, 1, 0, 0, 0, 3, 7, 18, 20,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// CalcParserInit initializes any static state used to implement CalcParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewCalcParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func CalcParserInit() {
	staticData := &CalcParserStaticData
	staticData.once.Do(calcParserInit)
}

// NewCalcParser produces a new parser instance for the optional input antlr.TokenStream.
func NewCalcParser(input antlr.TokenStream) *CalcParser {
	CalcParserInit()
	this := new(CalcParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &CalcParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Calc.g4"

	return this
}

// CalcParser tokens.
const (
	CalcParserEOF        = antlr.TokenEOF
	CalcParserIDENTIFIER = 1
	CalcParserPLUS       = 2
	CalcParserMINUS      = 3
	CalcParserMULT       = 4
	CalcParserDIV        = 5
	CalcParserOP         = 6
	CalcParserCP         = 7
	CalcParserSTRING     = 8
	CalcParserNUMBER     = 9
	CalcParserWHITESPACE = 10
	CalcParserDIGIT      = 11
	CalcParserLETTER     = 12
	CalcParserLOWER      = 13
	CalcParserUPPER      = 14
)

// CalcParser rules.
const (
	CalcParserRULE_program    = 0
	CalcParserRULE_expression = 1
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(CalcParserEOF, 0)
}

func (s *ProgramContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CalcVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CalcParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CalcParserRULE_program)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(7)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == CalcParserOP {
		{
			p.SetState(4)
			p.Expression()
		}

		p.SetState(9)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(10)
		p.Match(CalcParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOperand returns the operand token.
	GetOperand() antlr.Token

	// SetOperand sets the operand token.
	SetOperand(antlr.Token)

	// Getter signatures
	OP() antlr.TerminalNode
	CP() antlr.TerminalNode
	AllIDENTIFIER() []antlr.TerminalNode
	IDENTIFIER(i int) antlr.TerminalNode
	AllSTRING() []antlr.TerminalNode
	STRING(i int) antlr.TerminalNode
	AllNUMBER() []antlr.TerminalNode
	NUMBER(i int) antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser  antlr.Parser
	operand antlr.Token
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) GetOperand() antlr.Token { return s.operand }

func (s *ExpressionContext) SetOperand(v antlr.Token) { s.operand = v }

func (s *ExpressionContext) OP() antlr.TerminalNode {
	return s.GetToken(CalcParserOP, 0)
}

func (s *ExpressionContext) CP() antlr.TerminalNode {
	return s.GetToken(CalcParserCP, 0)
}

func (s *ExpressionContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(CalcParserIDENTIFIER)
}

func (s *ExpressionContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(CalcParserIDENTIFIER, i)
}

func (s *ExpressionContext) AllSTRING() []antlr.TerminalNode {
	return s.GetTokens(CalcParserSTRING)
}

func (s *ExpressionContext) STRING(i int) antlr.TerminalNode {
	return s.GetToken(CalcParserSTRING, i)
}

func (s *ExpressionContext) AllNUMBER() []antlr.TerminalNode {
	return s.GetTokens(CalcParserNUMBER)
}

func (s *ExpressionContext) NUMBER(i int) antlr.TerminalNode {
	return s.GetToken(CalcParserNUMBER, i)
}

func (s *ExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CalcVisitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CalcParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CalcParserRULE_expression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(12)
		p.Match(CalcParserOP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(13)

		var _m = p.Match(CalcParserIDENTIFIER)

		localctx.(*ExpressionContext).operand = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(20)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&834) != 0 {
		p.SetState(18)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case CalcParserIDENTIFIER:
			{
				p.SetState(14)
				p.Match(CalcParserIDENTIFIER)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case CalcParserSTRING:
			{
				p.SetState(15)
				p.Match(CalcParserSTRING)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case CalcParserNUMBER:
			{
				p.SetState(16)
				p.Match(CalcParserNUMBER)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case CalcParserOP:
			{
				p.SetState(17)
				p.Expression()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(22)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(23)
		p.Match(CalcParserCP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
