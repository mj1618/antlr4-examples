// Generated from ./Calc.g4 by ANTLR 4.13.1
// noinspection ES6UnusedImports,JSUnusedGlobalSymbols,JSUnusedLocalSymbols
import {
	ATN,
	ATNDeserializer,
	CharStream,
	DecisionState, DFA,
	Lexer,
	LexerATNSimulator,
	RuleContext,
	PredictionContextCache,
	Token
} from "antlr4";
export default class CalcLexer extends Lexer {
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

	public static readonly channelNames: string[] = [ "DEFAULT_TOKEN_CHANNEL", "HIDDEN" ];
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
	public static readonly modeNames: string[] = [ "DEFAULT_MODE", ];

	public static readonly ruleNames: string[] = [
		"IDENTIFIER", "PLUS", "MINUS", "MULT", "DIV", "OP", "CP", "STRING", "NUMBER", 
		"WHITESPACE", "DIGIT", "LETTER", "LOWER", "UPPER",
	];


	constructor(input: CharStream) {
		super(input);
		this._interp = new LexerATNSimulator(this, CalcLexer._ATN, CalcLexer.DecisionsToDFA, new PredictionContextCache());
	}

	public get grammarFileName(): string { return "Calc.g4"; }

	public get literalNames(): (string | null)[] { return CalcLexer.literalNames; }
	public get symbolicNames(): (string | null)[] { return CalcLexer.symbolicNames; }
	public get ruleNames(): string[] { return CalcLexer.ruleNames; }

	public get serializedATN(): number[] { return CalcLexer._serializedATN; }

	public get channelNames(): string[] { return CalcLexer.channelNames; }

	public get modeNames(): string[] { return CalcLexer.modeNames; }

	public static readonly _serializedATN: number[] = [4,0,14,86,6,-1,2,0,7,
	0,2,1,7,1,2,2,7,2,2,3,7,3,2,4,7,4,2,5,7,5,2,6,7,6,2,7,7,7,2,8,7,8,2,9,7,
	9,2,10,7,10,2,11,7,11,2,12,7,12,2,13,7,13,1,0,1,0,1,0,5,0,33,8,0,10,0,12,
	0,36,9,0,1,0,1,0,1,0,1,0,3,0,42,8,0,1,1,1,1,1,2,1,2,1,3,1,3,1,4,1,4,1,5,
	1,5,1,6,1,6,1,7,1,7,1,7,4,7,59,8,7,11,7,12,7,60,1,7,1,7,1,8,4,8,66,8,8,
	11,8,12,8,67,1,9,4,9,71,8,9,11,9,12,9,72,1,9,1,9,1,10,1,10,1,11,1,11,3,
	11,81,8,11,1,12,1,12,1,13,1,13,0,0,14,1,1,3,2,5,3,7,4,9,5,11,6,13,7,15,
	8,17,9,19,10,21,11,23,12,25,13,27,14,1,0,1,3,0,9,10,13,13,32,32,96,0,1,
	1,0,0,0,0,3,1,0,0,0,0,5,1,0,0,0,0,7,1,0,0,0,0,9,1,0,0,0,0,11,1,0,0,0,0,
	13,1,0,0,0,0,15,1,0,0,0,0,17,1,0,0,0,0,19,1,0,0,0,0,21,1,0,0,0,0,23,1,0,
	0,0,0,25,1,0,0,0,0,27,1,0,0,0,1,41,1,0,0,0,3,43,1,0,0,0,5,45,1,0,0,0,7,
	47,1,0,0,0,9,49,1,0,0,0,11,51,1,0,0,0,13,53,1,0,0,0,15,55,1,0,0,0,17,65,
	1,0,0,0,19,70,1,0,0,0,21,76,1,0,0,0,23,80,1,0,0,0,25,82,1,0,0,0,27,84,1,
	0,0,0,29,34,3,23,11,0,30,33,3,23,11,0,31,33,3,21,10,0,32,30,1,0,0,0,32,
	31,1,0,0,0,33,36,1,0,0,0,34,32,1,0,0,0,34,35,1,0,0,0,35,42,1,0,0,0,36,34,
	1,0,0,0,37,42,3,3,1,0,38,42,3,5,2,0,39,42,3,7,3,0,40,42,3,9,4,0,41,29,1,
	0,0,0,41,37,1,0,0,0,41,38,1,0,0,0,41,39,1,0,0,0,41,40,1,0,0,0,42,2,1,0,
	0,0,43,44,5,43,0,0,44,4,1,0,0,0,45,46,5,45,0,0,46,6,1,0,0,0,47,48,5,42,
	0,0,48,8,1,0,0,0,49,50,5,47,0,0,50,10,1,0,0,0,51,52,5,40,0,0,52,12,1,0,
	0,0,53,54,5,41,0,0,54,14,1,0,0,0,55,58,5,34,0,0,56,59,3,23,11,0,57,59,3,
	21,10,0,58,56,1,0,0,0,58,57,1,0,0,0,59,60,1,0,0,0,60,58,1,0,0,0,60,61,1,
	0,0,0,61,62,1,0,0,0,62,63,5,34,0,0,63,16,1,0,0,0,64,66,3,21,10,0,65,64,
	1,0,0,0,66,67,1,0,0,0,67,65,1,0,0,0,67,68,1,0,0,0,68,18,1,0,0,0,69,71,7,
	0,0,0,70,69,1,0,0,0,71,72,1,0,0,0,72,70,1,0,0,0,72,73,1,0,0,0,73,74,1,0,
	0,0,74,75,6,9,0,0,75,20,1,0,0,0,76,77,2,48,57,0,77,22,1,0,0,0,78,81,3,25,
	12,0,79,81,3,27,13,0,80,78,1,0,0,0,80,79,1,0,0,0,81,24,1,0,0,0,82,83,2,
	97,122,0,83,26,1,0,0,0,84,85,2,65,90,0,85,28,1,0,0,0,9,0,32,34,41,58,60,
	67,72,80,1,0,1,0];

	private static __ATN: ATN;
	public static get _ATN(): ATN {
		if (!CalcLexer.__ATN) {
			CalcLexer.__ATN = new ATNDeserializer().deserialize(CalcLexer._serializedATN);
		}

		return CalcLexer.__ATN;
	}


	static DecisionsToDFA = CalcLexer._ATN.decisionToState.map( (ds: DecisionState, index: number) => new DFA(ds, index) );
}