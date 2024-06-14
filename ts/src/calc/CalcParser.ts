// Generated from ./Calc.g4 by ANTLR 4.13.1
// noinspection ES6UnusedImports,JSUnusedGlobalSymbols,JSUnusedLocalSymbols

import {
	ATN,
	ATNDeserializer, DecisionState, DFA, FailedPredicateException,
	RecognitionException, NoViableAltException, BailErrorStrategy,
	Parser, ParserATNSimulator,
	RuleContext, ParserRuleContext, PredictionMode, PredictionContextCache,
	TerminalNode, RuleNode,
	Token, TokenStream,
	Interval, IntervalSet
} from 'antlr4';
import CalcListener from "./CalcListener.js";
import CalcVisitor from "./CalcVisitor.js";

// for running tests with parameters, TODO: discuss strategy for typed parameters in CI
// eslint-disable-next-line no-unused-vars
type int = number;

export default class CalcParser extends Parser {
	public static readonly IDENTIFIER = 1;
	public static readonly PLUS = 2;
	public static readonly MINUS = 3;
	public static readonly MULT = 4;
	public static readonly DIV = 5;
	public static readonly OP = 6;
	public static readonly CP = 7;
	public static readonly STRING = 8;
	public static readonly NUMBER = 9;
	public static readonly WHITESPACE = 10;
	public static readonly DIGIT = 11;
	public static readonly LETTER = 12;
	public static readonly LOWER = 13;
	public static readonly UPPER = 14;
	public static readonly EOF = Token.EOF;
	public static readonly RULE_program = 0;
	public static readonly RULE_expression = 1;
	public static readonly literalNames: (string | null)[] = [ null, null, 
                                                            "'+'", "'-'", 
                                                            "'*'", "'/'", 
                                                            "'('", "')'" ];
	public static readonly symbolicNames: (string | null)[] = [ null, "IDENTIFIER", 
                                                             "PLUS", "MINUS", 
                                                             "MULT", "DIV", 
                                                             "OP", "CP", 
                                                             "STRING", "NUMBER", 
                                                             "WHITESPACE", 
                                                             "DIGIT", "LETTER", 
                                                             "LOWER", "UPPER" ];
	// tslint:disable:no-trailing-whitespace
	public static readonly ruleNames: string[] = [
		"program", "expression",
	];
	public get grammarFileName(): string { return "Calc.g4"; }
	public get literalNames(): (string | null)[] { return CalcParser.literalNames; }
	public get symbolicNames(): (string | null)[] { return CalcParser.symbolicNames; }
	public get ruleNames(): string[] { return CalcParser.ruleNames; }
	public get serializedATN(): number[] { return CalcParser._serializedATN; }

	protected createFailedPredicateException(predicate?: string, message?: string): FailedPredicateException {
		return new FailedPredicateException(this, predicate, message);
	}

	constructor(input: TokenStream) {
		super(input);
		this._interp = new ParserATNSimulator(this, CalcParser._ATN, CalcParser.DecisionsToDFA, new PredictionContextCache());
	}
	// @RuleVersion(0)
	public program(): ProgramContext {
		let localctx: ProgramContext = new ProgramContext(this, this._ctx, this.state);
		this.enterRule(localctx, 0, CalcParser.RULE_program);
		let _la: number;
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 7;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			while (_la===6) {
				{
				{
				this.state = 4;
				this.expression();
				}
				}
				this.state = 9;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
			}
			this.state = 10;
			this.match(CalcParser.EOF);
			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
				throw re;
			}
		}
		finally {
			this.exitRule();
		}
		return localctx;
	}
	// @RuleVersion(0)
	public expression(): ExpressionContext {
		let localctx: ExpressionContext = new ExpressionContext(this, this._ctx, this.state);
		this.enterRule(localctx, 2, CalcParser.RULE_expression);
		let _la: number;
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 12;
			this.match(CalcParser.OP);
			this.state = 13;
			localctx._operand = this.match(CalcParser.IDENTIFIER);
			this.state = 20;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			while ((((_la) & ~0x1F) === 0 && ((1 << _la) & 834) !== 0)) {
				{
				this.state = 18;
				this._errHandler.sync(this);
				switch (this._input.LA(1)) {
				case 1:
					{
					this.state = 14;
					this.match(CalcParser.IDENTIFIER);
					}
					break;
				case 8:
					{
					this.state = 15;
					this.match(CalcParser.STRING);
					}
					break;
				case 9:
					{
					this.state = 16;
					this.match(CalcParser.NUMBER);
					}
					break;
				case 6:
					{
					this.state = 17;
					this.expression();
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				}
				this.state = 22;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
			}
			this.state = 23;
			this.match(CalcParser.CP);
			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
				throw re;
			}
		}
		finally {
			this.exitRule();
		}
		return localctx;
	}

	public static readonly _serializedATN: number[] = [4,1,14,26,2,0,7,0,2,
	1,7,1,1,0,5,0,6,8,0,10,0,12,0,9,9,0,1,0,1,0,1,1,1,1,1,1,1,1,1,1,1,1,5,1,
	19,8,1,10,1,12,1,22,9,1,1,1,1,1,1,1,0,0,2,0,2,0,0,28,0,7,1,0,0,0,2,12,1,
	0,0,0,4,6,3,2,1,0,5,4,1,0,0,0,6,9,1,0,0,0,7,5,1,0,0,0,7,8,1,0,0,0,8,10,
	1,0,0,0,9,7,1,0,0,0,10,11,5,0,0,1,11,1,1,0,0,0,12,13,5,6,0,0,13,20,5,1,
	0,0,14,19,5,1,0,0,15,19,5,8,0,0,16,19,5,9,0,0,17,19,3,2,1,0,18,14,1,0,0,
	0,18,15,1,0,0,0,18,16,1,0,0,0,18,17,1,0,0,0,19,22,1,0,0,0,20,18,1,0,0,0,
	20,21,1,0,0,0,21,23,1,0,0,0,22,20,1,0,0,0,23,24,5,7,0,0,24,3,1,0,0,0,3,
	7,18,20];

	private static __ATN: ATN;
	public static get _ATN(): ATN {
		if (!CalcParser.__ATN) {
			CalcParser.__ATN = new ATNDeserializer().deserialize(CalcParser._serializedATN);
		}

		return CalcParser.__ATN;
	}


	static DecisionsToDFA = CalcParser._ATN.decisionToState.map( (ds: DecisionState, index: number) => new DFA(ds, index) );

}

export class ProgramContext extends ParserRuleContext {
	constructor(parser?: CalcParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public EOF(): TerminalNode {
		return this.getToken(CalcParser.EOF, 0);
	}
	public expression_list(): ExpressionContext[] {
		return this.getTypedRuleContexts(ExpressionContext) as ExpressionContext[];
	}
	public expression(i: number): ExpressionContext {
		return this.getTypedRuleContext(ExpressionContext, i) as ExpressionContext;
	}
    public get ruleIndex(): number {
    	return CalcParser.RULE_program;
	}
	public enterRule(listener: CalcListener): void {
	    if(listener.enterProgram) {
	 		listener.enterProgram(this);
		}
	}
	public exitRule(listener: CalcListener): void {
	    if(listener.exitProgram) {
	 		listener.exitProgram(this);
		}
	}
	// @Override
	public accept<Result>(visitor: CalcVisitor<Result>): Result {
		if (visitor.visitProgram) {
			return visitor.visitProgram(this);
		} else {
			return visitor.visitChildren(this);
		}
	}
}


export class ExpressionContext extends ParserRuleContext {
	public _operand!: Token;
	constructor(parser?: CalcParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public OP(): TerminalNode {
		return this.getToken(CalcParser.OP, 0);
	}
	public CP(): TerminalNode {
		return this.getToken(CalcParser.CP, 0);
	}
	public IDENTIFIER_list(): TerminalNode[] {
	    	return this.getTokens(CalcParser.IDENTIFIER);
	}
	public IDENTIFIER(i: number): TerminalNode {
		return this.getToken(CalcParser.IDENTIFIER, i);
	}
	public STRING_list(): TerminalNode[] {
	    	return this.getTokens(CalcParser.STRING);
	}
	public STRING(i: number): TerminalNode {
		return this.getToken(CalcParser.STRING, i);
	}
	public NUMBER_list(): TerminalNode[] {
	    	return this.getTokens(CalcParser.NUMBER);
	}
	public NUMBER(i: number): TerminalNode {
		return this.getToken(CalcParser.NUMBER, i);
	}
	public expression_list(): ExpressionContext[] {
		return this.getTypedRuleContexts(ExpressionContext) as ExpressionContext[];
	}
	public expression(i: number): ExpressionContext {
		return this.getTypedRuleContext(ExpressionContext, i) as ExpressionContext;
	}
    public get ruleIndex(): number {
    	return CalcParser.RULE_expression;
	}
	public enterRule(listener: CalcListener): void {
	    if(listener.enterExpression) {
	 		listener.enterExpression(this);
		}
	}
	public exitRule(listener: CalcListener): void {
	    if(listener.exitExpression) {
	 		listener.exitExpression(this);
		}
	}
	// @Override
	public accept<Result>(visitor: CalcVisitor<Result>): Result {
		if (visitor.visitExpression) {
			return visitor.visitExpression(this);
		} else {
			return visitor.visitChildren(this);
		}
	}
}
