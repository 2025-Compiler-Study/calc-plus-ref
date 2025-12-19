// Generated from ../internal/grammar/CalcPlus.g4 by ANTLR 4.9.0-SNAPSHOT


import { ATN } from "antlr4ts/atn/ATN";
import { ATNDeserializer } from "antlr4ts/atn/ATNDeserializer";
import { FailedPredicateException } from "antlr4ts/FailedPredicateException";
import { NotNull } from "antlr4ts/Decorators";
import { NoViableAltException } from "antlr4ts/NoViableAltException";
import { Override } from "antlr4ts/Decorators";
import { Parser } from "antlr4ts/Parser";
import { ParserRuleContext } from "antlr4ts/ParserRuleContext";
import { ParserATNSimulator } from "antlr4ts/atn/ParserATNSimulator";
import { ParseTreeListener } from "antlr4ts/tree/ParseTreeListener";
import { ParseTreeVisitor } from "antlr4ts/tree/ParseTreeVisitor";
import { RecognitionException } from "antlr4ts/RecognitionException";
import { RuleContext } from "antlr4ts/RuleContext";
//import { RuleVersion } from "antlr4ts/RuleVersion";
import { TerminalNode } from "antlr4ts/tree/TerminalNode";
import { Token } from "antlr4ts/Token";
import { TokenStream } from "antlr4ts/TokenStream";
import { Vocabulary } from "antlr4ts/Vocabulary";
import { VocabularyImpl } from "antlr4ts/VocabularyImpl";

import * as Utils from "antlr4ts/misc/Utils";

import { CalcPlusListener } from "./CalcPlusListener";
import { CalcPlusVisitor } from "./CalcPlusVisitor";


export class CalcPlusParser extends Parser {
	public static readonly T__0 = 1;
	public static readonly T__1 = 2;
	public static readonly T__2 = 3;
	public static readonly T__3 = 4;
	public static readonly T__4 = 5;
	public static readonly T__5 = 6;
	public static readonly T__6 = 7;
	public static readonly T__7 = 8;
	public static readonly T__8 = 9;
	public static readonly T__9 = 10;
	public static readonly T__10 = 11;
	public static readonly T__11 = 12;
	public static readonly T__12 = 13;
	public static readonly T__13 = 14;
	public static readonly T__14 = 15;
	public static readonly T__15 = 16;
	public static readonly T__16 = 17;
	public static readonly T__17 = 18;
	public static readonly T__18 = 19;
	public static readonly T__19 = 20;
	public static readonly WS = 21;
	public static readonly INT = 22;
	public static readonly VAR = 23;
	public static readonly RULE_calc0 = 0;
	public static readonly RULE_expr = 1;
	public static readonly RULE_calc1 = 2;
	public static readonly RULE_stmt = 3;
	public static readonly RULE_calc2 = 4;
	public static readonly RULE_cond = 5;
	public static readonly RULE_block = 6;
	public static readonly RULE_calc3 = 7;
	// tslint:disable:no-trailing-whitespace
	public static readonly ruleNames: string[] = [
		"calc0", "expr", "calc1", "stmt", "calc2", "cond", "block", "calc3",
	];

	private static readonly _LITERAL_NAMES: Array<string | undefined> = [
		undefined, "'*'", "'/'", "'+'", "'-'", "'('", "')'", "'='", "';'", "'read'", 
		"'if'", "'else'", "'write'", "'=='", "'!='", "'>'", "'>='", "'<'", "'<='", 
		"'{'", "'}'",
	];
	private static readonly _SYMBOLIC_NAMES: Array<string | undefined> = [
		undefined, undefined, undefined, undefined, undefined, undefined, undefined, 
		undefined, undefined, undefined, undefined, undefined, undefined, undefined, 
		undefined, undefined, undefined, undefined, undefined, undefined, undefined, 
		"WS", "INT", "VAR",
	];
	public static readonly VOCABULARY: Vocabulary = new VocabularyImpl(CalcPlusParser._LITERAL_NAMES, CalcPlusParser._SYMBOLIC_NAMES, []);

	// @Override
	// @NotNull
	public get vocabulary(): Vocabulary {
		return CalcPlusParser.VOCABULARY;
	}
	// tslint:enable:no-trailing-whitespace

	// @Override
	public get grammarFileName(): string { return "CalcPlus.g4"; }

	// @Override
	public get ruleNames(): string[] { return CalcPlusParser.ruleNames; }

	// @Override
	public get serializedATN(): string { return CalcPlusParser._serializedATN; }

	protected createFailedPredicateException(predicate?: string, message?: string): FailedPredicateException {
		return new FailedPredicateException(this, predicate, message);
	}

	constructor(input: TokenStream) {
		super(input);
		this._interp = new ParserATNSimulator(CalcPlusParser._ATN, this);
	}
	// @RuleVersion(0)
	public calc0(): Calc0Context {
		let _localctx: Calc0Context = new Calc0Context(this._ctx, this.state);
		this.enterRule(_localctx, 0, CalcPlusParser.RULE_calc0);
		try {
			this.enterOuterAlt(_localctx, 1);
			{
			this.state = 16;
			this.expr(0);
			this.state = 17;
			this.match(CalcPlusParser.EOF);
			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				_localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
				throw re;
			}
		}
		finally {
			this.exitRule();
		}
		return _localctx;
	}

	public expr(): ExprContext;
	public expr(_p: number): ExprContext;
	// @RuleVersion(0)
	public expr(_p?: number): ExprContext {
		if (_p === undefined) {
			_p = 0;
		}

		let _parentctx: ParserRuleContext = this._ctx;
		let _parentState: number = this.state;
		let _localctx: ExprContext = new ExprContext(this._ctx, _parentState);
		let _prevctx: ExprContext = _localctx;
		let _startState: number = 2;
		this.enterRecursionRule(_localctx, 2, CalcPlusParser.RULE_expr, _p);
		let _la: number;
		try {
			let _alt: number;
			this.enterOuterAlt(_localctx, 1);
			{
			this.state = 26;
			this._errHandler.sync(this);
			switch (this._input.LA(1)) {
			case CalcPlusParser.INT:
				{
				_localctx = new IntContext(_localctx);
				this._ctx = _localctx;
				_prevctx = _localctx;

				this.state = 20;
				this.match(CalcPlusParser.INT);
				}
				break;
			case CalcPlusParser.VAR:
				{
				_localctx = new VarContext(_localctx);
				this._ctx = _localctx;
				_prevctx = _localctx;
				this.state = 21;
				this.match(CalcPlusParser.VAR);
				}
				break;
			case CalcPlusParser.T__4:
				{
				_localctx = new ParensContext(_localctx);
				this._ctx = _localctx;
				_prevctx = _localctx;
				this.state = 22;
				this.match(CalcPlusParser.T__4);
				this.state = 23;
				this.expr(0);
				this.state = 24;
				this.match(CalcPlusParser.T__5);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			this._ctx._stop = this._input.tryLT(-1);
			this.state = 36;
			this._errHandler.sync(this);
			_alt = this.interpreter.adaptivePredict(this._input, 2, this._ctx);
			while (_alt !== 2 && _alt !== ATN.INVALID_ALT_NUMBER) {
				if (_alt === 1) {
					if (this._parseListeners != null) {
						this.triggerExitRuleEvent();
					}
					_prevctx = _localctx;
					{
					this.state = 34;
					this._errHandler.sync(this);
					switch ( this.interpreter.adaptivePredict(this._input, 1, this._ctx) ) {
					case 1:
						{
						_localctx = new MulDivContext(new ExprContext(_parentctx, _parentState));
						this.pushNewRecursionContext(_localctx, _startState, CalcPlusParser.RULE_expr);
						this.state = 28;
						if (!(this.precpred(this._ctx, 5))) {
							throw this.createFailedPredicateException("this.precpred(this._ctx, 5)");
						}
						this.state = 29;
						_la = this._input.LA(1);
						if (!(_la === CalcPlusParser.T__0 || _la === CalcPlusParser.T__1)) {
						this._errHandler.recoverInline(this);
						} else {
							if (this._input.LA(1) === Token.EOF) {
								this.matchedEOF = true;
							}

							this._errHandler.reportMatch(this);
							this.consume();
						}
						this.state = 30;
						this.expr(6);
						}
						break;

					case 2:
						{
						_localctx = new AddSubContext(new ExprContext(_parentctx, _parentState));
						this.pushNewRecursionContext(_localctx, _startState, CalcPlusParser.RULE_expr);
						this.state = 31;
						if (!(this.precpred(this._ctx, 4))) {
							throw this.createFailedPredicateException("this.precpred(this._ctx, 4)");
						}
						this.state = 32;
						_la = this._input.LA(1);
						if (!(_la === CalcPlusParser.T__2 || _la === CalcPlusParser.T__3)) {
						this._errHandler.recoverInline(this);
						} else {
							if (this._input.LA(1) === Token.EOF) {
								this.matchedEOF = true;
							}

							this._errHandler.reportMatch(this);
							this.consume();
						}
						this.state = 33;
						this.expr(5);
						}
						break;
					}
					}
				}
				this.state = 38;
				this._errHandler.sync(this);
				_alt = this.interpreter.adaptivePredict(this._input, 2, this._ctx);
			}
			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				_localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
				throw re;
			}
		}
		finally {
			this.unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}
	// @RuleVersion(0)
	public calc1(): Calc1Context {
		let _localctx: Calc1Context = new Calc1Context(this._ctx, this.state);
		this.enterRule(_localctx, 4, CalcPlusParser.RULE_calc1);
		let _la: number;
		try {
			this.enterOuterAlt(_localctx, 1);
			{
			this.state = 40;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			do {
				{
				{
				this.state = 39;
				this.stmt();
				}
				}
				this.state = 42;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
			} while ((((_la) & ~0x1F) === 0 && ((1 << _la) & ((1 << CalcPlusParser.T__9) | (1 << CalcPlusParser.T__11) | (1 << CalcPlusParser.VAR))) !== 0));
			this.state = 44;
			this.match(CalcPlusParser.EOF);
			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				_localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
				throw re;
			}
		}
		finally {
			this.exitRule();
		}
		return _localctx;
	}
	// @RuleVersion(0)
	public stmt(): StmtContext {
		let _localctx: StmtContext = new StmtContext(this._ctx, this.state);
		this.enterRule(_localctx, 6, CalcPlusParser.RULE_stmt);
		let _la: number;
		try {
			this.state = 72;
			this._errHandler.sync(this);
			switch ( this.interpreter.adaptivePredict(this._input, 5, this._ctx) ) {
			case 1:
				_localctx = new ExprAssignContext(_localctx);
				this.enterOuterAlt(_localctx, 1);
				{
				this.state = 46;
				this.match(CalcPlusParser.VAR);
				this.state = 47;
				this.match(CalcPlusParser.T__6);
				this.state = 48;
				this.expr(0);
				this.state = 49;
				this.match(CalcPlusParser.T__7);
				}
				break;

			case 2:
				_localctx = new ReadAssignContext(_localctx);
				this.enterOuterAlt(_localctx, 2);
				{
				this.state = 51;
				this.match(CalcPlusParser.VAR);
				this.state = 52;
				this.match(CalcPlusParser.T__6);
				this.state = 53;
				this.match(CalcPlusParser.T__8);
				this.state = 54;
				this.match(CalcPlusParser.T__4);
				this.state = 55;
				this.match(CalcPlusParser.T__5);
				this.state = 56;
				this.match(CalcPlusParser.T__7);
				}
				break;

			case 3:
				_localctx = new IfElseContext(_localctx);
				this.enterOuterAlt(_localctx, 3);
				{
				this.state = 57;
				this.match(CalcPlusParser.T__9);
				this.state = 58;
				this.match(CalcPlusParser.T__4);
				this.state = 59;
				this.cond();
				this.state = 60;
				this.match(CalcPlusParser.T__5);
				this.state = 61;
				(_localctx as IfElseContext)._thenBlock = this.block();
				this.state = 64;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
				if (_la === CalcPlusParser.T__10) {
					{
					this.state = 62;
					this.match(CalcPlusParser.T__10);
					this.state = 63;
					(_localctx as IfElseContext)._elseBlock = this.block();
					}
				}

				}
				break;

			case 4:
				_localctx = new WriteContext(_localctx);
				this.enterOuterAlt(_localctx, 4);
				{
				this.state = 66;
				this.match(CalcPlusParser.T__11);
				this.state = 67;
				this.match(CalcPlusParser.T__4);
				this.state = 68;
				this.expr(0);
				this.state = 69;
				this.match(CalcPlusParser.T__5);
				this.state = 70;
				this.match(CalcPlusParser.T__7);
				}
				break;
			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				_localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
				throw re;
			}
		}
		finally {
			this.exitRule();
		}
		return _localctx;
	}
	// @RuleVersion(0)
	public calc2(): Calc2Context {
		let _localctx: Calc2Context = new Calc2Context(this._ctx, this.state);
		this.enterRule(_localctx, 8, CalcPlusParser.RULE_calc2);
		let _la: number;
		try {
			this.enterOuterAlt(_localctx, 1);
			{
			this.state = 75;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			do {
				{
				{
				this.state = 74;
				this.stmt();
				}
				}
				this.state = 77;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
			} while ((((_la) & ~0x1F) === 0 && ((1 << _la) & ((1 << CalcPlusParser.T__9) | (1 << CalcPlusParser.T__11) | (1 << CalcPlusParser.VAR))) !== 0));
			this.state = 79;
			this.match(CalcPlusParser.EOF);
			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				_localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
				throw re;
			}
		}
		finally {
			this.exitRule();
		}
		return _localctx;
	}
	// @RuleVersion(0)
	public cond(): CondContext {
		let _localctx: CondContext = new CondContext(this._ctx, this.state);
		this.enterRule(_localctx, 10, CalcPlusParser.RULE_cond);
		let _la: number;
		try {
			this.enterOuterAlt(_localctx, 1);
			{
			this.state = 81;
			this.expr(0);
			this.state = 82;
			_la = this._input.LA(1);
			if (!((((_la) & ~0x1F) === 0 && ((1 << _la) & ((1 << CalcPlusParser.T__12) | (1 << CalcPlusParser.T__13) | (1 << CalcPlusParser.T__14) | (1 << CalcPlusParser.T__15) | (1 << CalcPlusParser.T__16) | (1 << CalcPlusParser.T__17))) !== 0))) {
			this._errHandler.recoverInline(this);
			} else {
				if (this._input.LA(1) === Token.EOF) {
					this.matchedEOF = true;
				}

				this._errHandler.reportMatch(this);
				this.consume();
			}
			this.state = 83;
			this.expr(0);
			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				_localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
				throw re;
			}
		}
		finally {
			this.exitRule();
		}
		return _localctx;
	}
	// @RuleVersion(0)
	public block(): BlockContext {
		let _localctx: BlockContext = new BlockContext(this._ctx, this.state);
		this.enterRule(_localctx, 12, CalcPlusParser.RULE_block);
		let _la: number;
		try {
			this.enterOuterAlt(_localctx, 1);
			{
			this.state = 85;
			this.match(CalcPlusParser.T__18);
			this.state = 89;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			while ((((_la) & ~0x1F) === 0 && ((1 << _la) & ((1 << CalcPlusParser.T__9) | (1 << CalcPlusParser.T__11) | (1 << CalcPlusParser.VAR))) !== 0)) {
				{
				{
				this.state = 86;
				this.stmt();
				}
				}
				this.state = 91;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
			}
			this.state = 92;
			this.match(CalcPlusParser.T__19);
			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				_localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
				throw re;
			}
		}
		finally {
			this.exitRule();
		}
		return _localctx;
	}
	// @RuleVersion(0)
	public calc3(): Calc3Context {
		let _localctx: Calc3Context = new Calc3Context(this._ctx, this.state);
		this.enterRule(_localctx, 14, CalcPlusParser.RULE_calc3);
		let _la: number;
		try {
			this.enterOuterAlt(_localctx, 1);
			{
			this.state = 95;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			do {
				{
				{
				this.state = 94;
				this.stmt();
				}
				}
				this.state = 97;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
			} while ((((_la) & ~0x1F) === 0 && ((1 << _la) & ((1 << CalcPlusParser.T__9) | (1 << CalcPlusParser.T__11) | (1 << CalcPlusParser.VAR))) !== 0));
			this.state = 99;
			this.match(CalcPlusParser.EOF);
			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				_localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
				throw re;
			}
		}
		finally {
			this.exitRule();
		}
		return _localctx;
	}

	public sempred(_localctx: RuleContext, ruleIndex: number, predIndex: number): boolean {
		switch (ruleIndex) {
		case 1:
			return this.expr_sempred(_localctx as ExprContext, predIndex);
		}
		return true;
	}
	private expr_sempred(_localctx: ExprContext, predIndex: number): boolean {
		switch (predIndex) {
		case 0:
			return this.precpred(this._ctx, 5);

		case 1:
			return this.precpred(this._ctx, 4);
		}
		return true;
	}

	public static readonly _serializedATN: string =
		"\x03\uC91D\uCABA\u058D\uAFBA\u4F53\u0607\uEA8B\uC241\x03\x19h\x04\x02" +
		"\t\x02\x04\x03\t\x03\x04\x04\t\x04\x04\x05\t\x05\x04\x06\t\x06\x04\x07" +
		"\t\x07\x04\b\t\b\x04\t\t\t\x03\x02\x03\x02\x03\x02\x03\x03\x03\x03\x03" +
		"\x03\x03\x03\x03\x03\x03\x03\x03\x03\x05\x03\x1D\n\x03\x03\x03\x03\x03" +
		"\x03\x03\x03\x03\x03\x03\x03\x03\x07\x03%\n\x03\f\x03\x0E\x03(\v\x03\x03" +
		"\x04\x06\x04+\n\x04\r\x04\x0E\x04,\x03\x04\x03\x04\x03\x05\x03\x05\x03" +
		"\x05\x03\x05\x03\x05\x03\x05\x03\x05\x03\x05\x03\x05\x03\x05\x03\x05\x03" +
		"\x05\x03\x05\x03\x05\x03\x05\x03\x05\x03\x05\x03\x05\x05\x05C\n\x05\x03" +
		"\x05\x03\x05\x03\x05\x03\x05\x03\x05\x03\x05\x05\x05K\n\x05\x03\x06\x06" +
		"\x06N\n\x06\r\x06\x0E\x06O\x03\x06\x03\x06\x03\x07\x03\x07\x03\x07\x03" +
		"\x07\x03\b\x03\b\x07\bZ\n\b\f\b\x0E\b]\v\b\x03\b\x03\b\x03\t\x06\tb\n" +
		"\t\r\t\x0E\tc\x03\t\x03\t\x03\t\x02\x02\x03\x04\n\x02\x02\x04\x02\x06" +
		"\x02\b\x02\n\x02\f\x02\x0E\x02\x10\x02\x02\x05\x03\x02\x03\x04\x03\x02" +
		"\x05\x06\x03\x02\x0F\x14\x02k\x02\x12\x03\x02\x02\x02\x04\x1C\x03\x02" +
		"\x02\x02\x06*\x03\x02\x02\x02\bJ\x03\x02\x02\x02\nM\x03\x02\x02\x02\f" +
		"S\x03\x02\x02\x02\x0EW\x03\x02\x02\x02\x10a\x03\x02\x02\x02\x12\x13\x05" +
		"\x04\x03\x02\x13\x14\x07\x02\x02\x03\x14\x03\x03\x02\x02\x02\x15\x16\b" +
		"\x03\x01\x02\x16\x1D\x07\x18\x02\x02\x17\x1D\x07\x19\x02\x02\x18\x19\x07" +
		"\x07\x02\x02\x19\x1A\x05\x04\x03\x02\x1A\x1B\x07\b\x02\x02\x1B\x1D\x03" +
		"\x02\x02\x02\x1C\x15\x03\x02\x02\x02\x1C\x17\x03\x02\x02\x02\x1C\x18\x03" +
		"\x02\x02\x02\x1D&\x03\x02\x02\x02\x1E\x1F\f\x07\x02\x02\x1F \t\x02\x02" +
		"\x02 %\x05\x04\x03\b!\"\f\x06\x02\x02\"#\t\x03\x02\x02#%\x05\x04\x03\x07" +
		"$\x1E\x03\x02\x02\x02$!\x03\x02\x02\x02%(\x03\x02\x02\x02&$\x03\x02\x02" +
		"\x02&\'\x03\x02\x02\x02\'\x05\x03\x02\x02\x02(&\x03\x02\x02\x02)+\x05" +
		"\b\x05\x02*)\x03\x02\x02\x02+,\x03\x02\x02\x02,*\x03\x02\x02\x02,-\x03" +
		"\x02\x02\x02-.\x03\x02\x02\x02./\x07\x02\x02\x03/\x07\x03\x02\x02\x02" +
		"01\x07\x19\x02\x0212\x07\t\x02\x0223\x05\x04\x03\x0234\x07\n\x02\x024" +
		"K\x03\x02\x02\x0256\x07\x19\x02\x0267\x07\t\x02\x0278\x07\v\x02\x0289" +
		"\x07\x07\x02\x029:\x07\b\x02\x02:K\x07\n\x02\x02;<\x07\f\x02\x02<=\x07" +
		"\x07\x02\x02=>\x05\f\x07\x02>?\x07\b\x02\x02?B\x05\x0E\b\x02@A\x07\r\x02" +
		"\x02AC\x05\x0E\b\x02B@\x03\x02\x02\x02BC\x03\x02\x02\x02CK\x03\x02\x02" +
		"\x02DE\x07\x0E\x02\x02EF\x07\x07\x02\x02FG\x05\x04\x03\x02GH\x07\b\x02" +
		"\x02HI\x07\n\x02\x02IK\x03\x02\x02\x02J0\x03\x02\x02\x02J5\x03\x02\x02" +
		"\x02J;\x03\x02\x02\x02JD\x03\x02\x02\x02K\t\x03\x02\x02\x02LN\x05\b\x05" +
		"\x02ML\x03\x02\x02\x02NO\x03\x02\x02\x02OM\x03\x02\x02\x02OP\x03\x02\x02" +
		"\x02PQ\x03\x02\x02\x02QR\x07\x02\x02\x03R\v\x03\x02\x02\x02ST\x05\x04" +
		"\x03\x02TU\t\x04\x02\x02UV\x05\x04\x03\x02V\r\x03\x02\x02\x02W[\x07\x15" +
		"\x02\x02XZ\x05\b\x05\x02YX\x03\x02\x02\x02Z]\x03\x02\x02\x02[Y\x03\x02" +
		"\x02\x02[\\\x03\x02\x02\x02\\^\x03\x02\x02\x02][\x03\x02\x02\x02^_\x07" +
		"\x16\x02\x02_\x0F\x03\x02\x02\x02`b\x05\b\x05\x02a`\x03\x02\x02\x02bc" +
		"\x03\x02\x02\x02ca\x03\x02\x02\x02cd\x03\x02\x02\x02de\x03\x02\x02\x02" +
		"ef\x07\x02\x02\x03f\x11\x03\x02\x02\x02\v\x1C$&,BJO[c";
	public static __ATN: ATN;
	public static get _ATN(): ATN {
		if (!CalcPlusParser.__ATN) {
			CalcPlusParser.__ATN = new ATNDeserializer().deserialize(Utils.toCharArray(CalcPlusParser._serializedATN));
		}

		return CalcPlusParser.__ATN;
	}

}

export class Calc0Context extends ParserRuleContext {
	public expr(): ExprContext {
		return this.getRuleContext(0, ExprContext);
	}
	public EOF(): TerminalNode { return this.getToken(CalcPlusParser.EOF, 0); }
	constructor(parent: ParserRuleContext | undefined, invokingState: number) {
		super(parent, invokingState);
	}
	// @Override
	public get ruleIndex(): number { return CalcPlusParser.RULE_calc0; }
	// @Override
	public enterRule(listener: CalcPlusListener): void {
		if (listener.enterCalc0) {
			listener.enterCalc0(this);
		}
	}
	// @Override
	public exitRule(listener: CalcPlusListener): void {
		if (listener.exitCalc0) {
			listener.exitCalc0(this);
		}
	}
	// @Override
	public accept<Result>(visitor: CalcPlusVisitor<Result>): Result {
		if (visitor.visitCalc0) {
			return visitor.visitCalc0(this);
		} else {
			return visitor.visitChildren(this);
		}
	}
}


export class ExprContext extends ParserRuleContext {
	constructor(parent: ParserRuleContext | undefined, invokingState: number) {
		super(parent, invokingState);
	}
	// @Override
	public get ruleIndex(): number { return CalcPlusParser.RULE_expr; }
	public copyFrom(ctx: ExprContext): void {
		super.copyFrom(ctx);
	}
}
export class MulDivContext extends ExprContext {
	public expr(): ExprContext[];
	public expr(i: number): ExprContext;
	public expr(i?: number): ExprContext | ExprContext[] {
		if (i === undefined) {
			return this.getRuleContexts(ExprContext);
		} else {
			return this.getRuleContext(i, ExprContext);
		}
	}
	constructor(ctx: ExprContext) {
		super(ctx.parent, ctx.invokingState);
		this.copyFrom(ctx);
	}
	// @Override
	public enterRule(listener: CalcPlusListener): void {
		if (listener.enterMulDiv) {
			listener.enterMulDiv(this);
		}
	}
	// @Override
	public exitRule(listener: CalcPlusListener): void {
		if (listener.exitMulDiv) {
			listener.exitMulDiv(this);
		}
	}
	// @Override
	public accept<Result>(visitor: CalcPlusVisitor<Result>): Result {
		if (visitor.visitMulDiv) {
			return visitor.visitMulDiv(this);
		} else {
			return visitor.visitChildren(this);
		}
	}
}
export class AddSubContext extends ExprContext {
	public expr(): ExprContext[];
	public expr(i: number): ExprContext;
	public expr(i?: number): ExprContext | ExprContext[] {
		if (i === undefined) {
			return this.getRuleContexts(ExprContext);
		} else {
			return this.getRuleContext(i, ExprContext);
		}
	}
	constructor(ctx: ExprContext) {
		super(ctx.parent, ctx.invokingState);
		this.copyFrom(ctx);
	}
	// @Override
	public enterRule(listener: CalcPlusListener): void {
		if (listener.enterAddSub) {
			listener.enterAddSub(this);
		}
	}
	// @Override
	public exitRule(listener: CalcPlusListener): void {
		if (listener.exitAddSub) {
			listener.exitAddSub(this);
		}
	}
	// @Override
	public accept<Result>(visitor: CalcPlusVisitor<Result>): Result {
		if (visitor.visitAddSub) {
			return visitor.visitAddSub(this);
		} else {
			return visitor.visitChildren(this);
		}
	}
}
export class IntContext extends ExprContext {
	public INT(): TerminalNode { return this.getToken(CalcPlusParser.INT, 0); }
	constructor(ctx: ExprContext) {
		super(ctx.parent, ctx.invokingState);
		this.copyFrom(ctx);
	}
	// @Override
	public enterRule(listener: CalcPlusListener): void {
		if (listener.enterInt) {
			listener.enterInt(this);
		}
	}
	// @Override
	public exitRule(listener: CalcPlusListener): void {
		if (listener.exitInt) {
			listener.exitInt(this);
		}
	}
	// @Override
	public accept<Result>(visitor: CalcPlusVisitor<Result>): Result {
		if (visitor.visitInt) {
			return visitor.visitInt(this);
		} else {
			return visitor.visitChildren(this);
		}
	}
}
export class VarContext extends ExprContext {
	public VAR(): TerminalNode { return this.getToken(CalcPlusParser.VAR, 0); }
	constructor(ctx: ExprContext) {
		super(ctx.parent, ctx.invokingState);
		this.copyFrom(ctx);
	}
	// @Override
	public enterRule(listener: CalcPlusListener): void {
		if (listener.enterVar) {
			listener.enterVar(this);
		}
	}
	// @Override
	public exitRule(listener: CalcPlusListener): void {
		if (listener.exitVar) {
			listener.exitVar(this);
		}
	}
	// @Override
	public accept<Result>(visitor: CalcPlusVisitor<Result>): Result {
		if (visitor.visitVar) {
			return visitor.visitVar(this);
		} else {
			return visitor.visitChildren(this);
		}
	}
}
export class ParensContext extends ExprContext {
	public expr(): ExprContext {
		return this.getRuleContext(0, ExprContext);
	}
	constructor(ctx: ExprContext) {
		super(ctx.parent, ctx.invokingState);
		this.copyFrom(ctx);
	}
	// @Override
	public enterRule(listener: CalcPlusListener): void {
		if (listener.enterParens) {
			listener.enterParens(this);
		}
	}
	// @Override
	public exitRule(listener: CalcPlusListener): void {
		if (listener.exitParens) {
			listener.exitParens(this);
		}
	}
	// @Override
	public accept<Result>(visitor: CalcPlusVisitor<Result>): Result {
		if (visitor.visitParens) {
			return visitor.visitParens(this);
		} else {
			return visitor.visitChildren(this);
		}
	}
}


export class Calc1Context extends ParserRuleContext {
	public EOF(): TerminalNode { return this.getToken(CalcPlusParser.EOF, 0); }
	public stmt(): StmtContext[];
	public stmt(i: number): StmtContext;
	public stmt(i?: number): StmtContext | StmtContext[] {
		if (i === undefined) {
			return this.getRuleContexts(StmtContext);
		} else {
			return this.getRuleContext(i, StmtContext);
		}
	}
	constructor(parent: ParserRuleContext | undefined, invokingState: number) {
		super(parent, invokingState);
	}
	// @Override
	public get ruleIndex(): number { return CalcPlusParser.RULE_calc1; }
	// @Override
	public enterRule(listener: CalcPlusListener): void {
		if (listener.enterCalc1) {
			listener.enterCalc1(this);
		}
	}
	// @Override
	public exitRule(listener: CalcPlusListener): void {
		if (listener.exitCalc1) {
			listener.exitCalc1(this);
		}
	}
	// @Override
	public accept<Result>(visitor: CalcPlusVisitor<Result>): Result {
		if (visitor.visitCalc1) {
			return visitor.visitCalc1(this);
		} else {
			return visitor.visitChildren(this);
		}
	}
}


export class StmtContext extends ParserRuleContext {
	constructor(parent: ParserRuleContext | undefined, invokingState: number) {
		super(parent, invokingState);
	}
	// @Override
	public get ruleIndex(): number { return CalcPlusParser.RULE_stmt; }
	public copyFrom(ctx: StmtContext): void {
		super.copyFrom(ctx);
	}
}
export class ExprAssignContext extends StmtContext {
	public VAR(): TerminalNode { return this.getToken(CalcPlusParser.VAR, 0); }
	public expr(): ExprContext {
		return this.getRuleContext(0, ExprContext);
	}
	constructor(ctx: StmtContext) {
		super(ctx.parent, ctx.invokingState);
		this.copyFrom(ctx);
	}
	// @Override
	public enterRule(listener: CalcPlusListener): void {
		if (listener.enterExprAssign) {
			listener.enterExprAssign(this);
		}
	}
	// @Override
	public exitRule(listener: CalcPlusListener): void {
		if (listener.exitExprAssign) {
			listener.exitExprAssign(this);
		}
	}
	// @Override
	public accept<Result>(visitor: CalcPlusVisitor<Result>): Result {
		if (visitor.visitExprAssign) {
			return visitor.visitExprAssign(this);
		} else {
			return visitor.visitChildren(this);
		}
	}
}
export class ReadAssignContext extends StmtContext {
	public VAR(): TerminalNode { return this.getToken(CalcPlusParser.VAR, 0); }
	constructor(ctx: StmtContext) {
		super(ctx.parent, ctx.invokingState);
		this.copyFrom(ctx);
	}
	// @Override
	public enterRule(listener: CalcPlusListener): void {
		if (listener.enterReadAssign) {
			listener.enterReadAssign(this);
		}
	}
	// @Override
	public exitRule(listener: CalcPlusListener): void {
		if (listener.exitReadAssign) {
			listener.exitReadAssign(this);
		}
	}
	// @Override
	public accept<Result>(visitor: CalcPlusVisitor<Result>): Result {
		if (visitor.visitReadAssign) {
			return visitor.visitReadAssign(this);
		} else {
			return visitor.visitChildren(this);
		}
	}
}
export class IfElseContext extends StmtContext {
	public _thenBlock!: BlockContext;
	public _elseBlock!: BlockContext;
	public cond(): CondContext {
		return this.getRuleContext(0, CondContext);
	}
	public block(): BlockContext[];
	public block(i: number): BlockContext;
	public block(i?: number): BlockContext | BlockContext[] {
		if (i === undefined) {
			return this.getRuleContexts(BlockContext);
		} else {
			return this.getRuleContext(i, BlockContext);
		}
	}
	constructor(ctx: StmtContext) {
		super(ctx.parent, ctx.invokingState);
		this.copyFrom(ctx);
	}
	// @Override
	public enterRule(listener: CalcPlusListener): void {
		if (listener.enterIfElse) {
			listener.enterIfElse(this);
		}
	}
	// @Override
	public exitRule(listener: CalcPlusListener): void {
		if (listener.exitIfElse) {
			listener.exitIfElse(this);
		}
	}
	// @Override
	public accept<Result>(visitor: CalcPlusVisitor<Result>): Result {
		if (visitor.visitIfElse) {
			return visitor.visitIfElse(this);
		} else {
			return visitor.visitChildren(this);
		}
	}
}
export class WriteContext extends StmtContext {
	public expr(): ExprContext {
		return this.getRuleContext(0, ExprContext);
	}
	constructor(ctx: StmtContext) {
		super(ctx.parent, ctx.invokingState);
		this.copyFrom(ctx);
	}
	// @Override
	public enterRule(listener: CalcPlusListener): void {
		if (listener.enterWrite) {
			listener.enterWrite(this);
		}
	}
	// @Override
	public exitRule(listener: CalcPlusListener): void {
		if (listener.exitWrite) {
			listener.exitWrite(this);
		}
	}
	// @Override
	public accept<Result>(visitor: CalcPlusVisitor<Result>): Result {
		if (visitor.visitWrite) {
			return visitor.visitWrite(this);
		} else {
			return visitor.visitChildren(this);
		}
	}
}


export class Calc2Context extends ParserRuleContext {
	public EOF(): TerminalNode { return this.getToken(CalcPlusParser.EOF, 0); }
	public stmt(): StmtContext[];
	public stmt(i: number): StmtContext;
	public stmt(i?: number): StmtContext | StmtContext[] {
		if (i === undefined) {
			return this.getRuleContexts(StmtContext);
		} else {
			return this.getRuleContext(i, StmtContext);
		}
	}
	constructor(parent: ParserRuleContext | undefined, invokingState: number) {
		super(parent, invokingState);
	}
	// @Override
	public get ruleIndex(): number { return CalcPlusParser.RULE_calc2; }
	// @Override
	public enterRule(listener: CalcPlusListener): void {
		if (listener.enterCalc2) {
			listener.enterCalc2(this);
		}
	}
	// @Override
	public exitRule(listener: CalcPlusListener): void {
		if (listener.exitCalc2) {
			listener.exitCalc2(this);
		}
	}
	// @Override
	public accept<Result>(visitor: CalcPlusVisitor<Result>): Result {
		if (visitor.visitCalc2) {
			return visitor.visitCalc2(this);
		} else {
			return visitor.visitChildren(this);
		}
	}
}


export class CondContext extends ParserRuleContext {
	public expr(): ExprContext[];
	public expr(i: number): ExprContext;
	public expr(i?: number): ExprContext | ExprContext[] {
		if (i === undefined) {
			return this.getRuleContexts(ExprContext);
		} else {
			return this.getRuleContext(i, ExprContext);
		}
	}
	constructor(parent: ParserRuleContext | undefined, invokingState: number) {
		super(parent, invokingState);
	}
	// @Override
	public get ruleIndex(): number { return CalcPlusParser.RULE_cond; }
	// @Override
	public enterRule(listener: CalcPlusListener): void {
		if (listener.enterCond) {
			listener.enterCond(this);
		}
	}
	// @Override
	public exitRule(listener: CalcPlusListener): void {
		if (listener.exitCond) {
			listener.exitCond(this);
		}
	}
	// @Override
	public accept<Result>(visitor: CalcPlusVisitor<Result>): Result {
		if (visitor.visitCond) {
			return visitor.visitCond(this);
		} else {
			return visitor.visitChildren(this);
		}
	}
}


export class BlockContext extends ParserRuleContext {
	public stmt(): StmtContext[];
	public stmt(i: number): StmtContext;
	public stmt(i?: number): StmtContext | StmtContext[] {
		if (i === undefined) {
			return this.getRuleContexts(StmtContext);
		} else {
			return this.getRuleContext(i, StmtContext);
		}
	}
	constructor(parent: ParserRuleContext | undefined, invokingState: number) {
		super(parent, invokingState);
	}
	// @Override
	public get ruleIndex(): number { return CalcPlusParser.RULE_block; }
	// @Override
	public enterRule(listener: CalcPlusListener): void {
		if (listener.enterBlock) {
			listener.enterBlock(this);
		}
	}
	// @Override
	public exitRule(listener: CalcPlusListener): void {
		if (listener.exitBlock) {
			listener.exitBlock(this);
		}
	}
	// @Override
	public accept<Result>(visitor: CalcPlusVisitor<Result>): Result {
		if (visitor.visitBlock) {
			return visitor.visitBlock(this);
		} else {
			return visitor.visitChildren(this);
		}
	}
}


export class Calc3Context extends ParserRuleContext {
	public EOF(): TerminalNode { return this.getToken(CalcPlusParser.EOF, 0); }
	public stmt(): StmtContext[];
	public stmt(i: number): StmtContext;
	public stmt(i?: number): StmtContext | StmtContext[] {
		if (i === undefined) {
			return this.getRuleContexts(StmtContext);
		} else {
			return this.getRuleContext(i, StmtContext);
		}
	}
	constructor(parent: ParserRuleContext | undefined, invokingState: number) {
		super(parent, invokingState);
	}
	// @Override
	public get ruleIndex(): number { return CalcPlusParser.RULE_calc3; }
	// @Override
	public enterRule(listener: CalcPlusListener): void {
		if (listener.enterCalc3) {
			listener.enterCalc3(this);
		}
	}
	// @Override
	public exitRule(listener: CalcPlusListener): void {
		if (listener.exitCalc3) {
			listener.exitCalc3(this);
		}
	}
	// @Override
	public accept<Result>(visitor: CalcPlusVisitor<Result>): Result {
		if (visitor.visitCalc3) {
			return visitor.visitCalc3(this);
		} else {
			return visitor.visitChildren(this);
		}
	}
}


