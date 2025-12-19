// Generated from ../internal/grammar/CalcPlus.g4 by ANTLR 4.9.0-SNAPSHOT


import { ATN } from "antlr4ts/atn/ATN";
import { ATNDeserializer } from "antlr4ts/atn/ATNDeserializer";
import { CharStream } from "antlr4ts/CharStream";
import { Lexer } from "antlr4ts/Lexer";
import { LexerATNSimulator } from "antlr4ts/atn/LexerATNSimulator";
import { NotNull } from "antlr4ts/Decorators";
import { Override } from "antlr4ts/Decorators";
import { RuleContext } from "antlr4ts/RuleContext";
import { Vocabulary } from "antlr4ts/Vocabulary";
import { VocabularyImpl } from "antlr4ts/VocabularyImpl";

import * as Utils from "antlr4ts/misc/Utils";


export class CalcPlusLexer extends Lexer {
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

	// tslint:disable:no-trailing-whitespace
	public static readonly channelNames: string[] = [
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	];

	// tslint:disable:no-trailing-whitespace
	public static readonly modeNames: string[] = [
		"DEFAULT_MODE",
	];

	public static readonly ruleNames: string[] = [
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8", 
		"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16", 
		"T__17", "T__18", "T__19", "WS", "INT", "VAR",
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
	public static readonly VOCABULARY: Vocabulary = new VocabularyImpl(CalcPlusLexer._LITERAL_NAMES, CalcPlusLexer._SYMBOLIC_NAMES, []);

	// @Override
	// @NotNull
	public get vocabulary(): Vocabulary {
		return CalcPlusLexer.VOCABULARY;
	}
	// tslint:enable:no-trailing-whitespace


	constructor(input: CharStream) {
		super(input);
		this._interp = new LexerATNSimulator(CalcPlusLexer._ATN, this);
	}

	// @Override
	public get grammarFileName(): string { return "CalcPlus.g4"; }

	// @Override
	public get ruleNames(): string[] { return CalcPlusLexer.ruleNames; }

	// @Override
	public get serializedATN(): string { return CalcPlusLexer._serializedATN; }

	// @Override
	public get channelNames(): string[] { return CalcPlusLexer.channelNames; }

	// @Override
	public get modeNames(): string[] { return CalcPlusLexer.modeNames; }

	public static readonly _serializedATN: string =
		"\x03\uC91D\uCABA\u058D\uAFBA\u4F53\u0607\uEA8B\uC241\x02\x19y\b\x01\x04" +
		"\x02\t\x02\x04\x03\t\x03\x04\x04\t\x04\x04\x05\t\x05\x04\x06\t\x06\x04" +
		"\x07\t\x07\x04\b\t\b\x04\t\t\t\x04\n\t\n\x04\v\t\v\x04\f\t\f\x04\r\t\r" +
		"\x04\x0E\t\x0E\x04\x0F\t\x0F\x04\x10\t\x10\x04\x11\t\x11\x04\x12\t\x12" +
		"\x04\x13\t\x13\x04\x14\t\x14\x04\x15\t\x15\x04\x16\t\x16\x04\x17\t\x17" +
		"\x04\x18\t\x18\x03\x02\x03\x02\x03\x03\x03\x03\x03\x04\x03\x04\x03\x05" +
		"\x03\x05\x03\x06\x03\x06\x03\x07\x03\x07\x03\b\x03\b\x03\t\x03\t\x03\n" +
		"\x03\n\x03\n\x03\n\x03\n\x03\v\x03\v\x03\v\x03\f\x03\f\x03\f\x03\f\x03" +
		"\f\x03\r\x03\r\x03\r\x03\r\x03\r\x03\r\x03\x0E\x03\x0E\x03\x0E\x03\x0F" +
		"\x03\x0F\x03\x0F\x03\x10\x03\x10\x03\x11\x03\x11\x03\x11\x03\x12\x03\x12" +
		"\x03\x13\x03\x13\x03\x13\x03\x14\x03\x14\x03\x15\x03\x15\x03\x16\x06\x16" +
		"j\n\x16\r\x16\x0E\x16k\x03\x16\x03\x16\x03\x17\x06\x17q\n\x17\r\x17\x0E" +
		"\x17r\x03\x18\x06\x18v\n\x18\r\x18\x0E\x18w\x02\x02\x02\x19\x03\x02\x03" +
		"\x05\x02\x04\x07\x02\x05\t\x02\x06\v\x02\x07\r\x02\b\x0F\x02\t\x11\x02" +
		"\n\x13\x02\v\x15\x02\f\x17\x02\r\x19\x02\x0E\x1B\x02\x0F\x1D\x02\x10\x1F" +
		"\x02\x11!\x02\x12#\x02\x13%\x02\x14\'\x02\x15)\x02\x16+\x02\x17-\x02\x18" +
		"/\x02\x19\x03\x02\x05\x05\x02\v\f\x0F\x0F\"\"\x03\x022;\x04\x02C\\c|\x02" +
		"{\x02\x03\x03\x02\x02\x02\x02\x05\x03\x02\x02\x02\x02\x07\x03\x02\x02" +
		"\x02\x02\t\x03\x02\x02\x02\x02\v\x03\x02\x02\x02\x02\r\x03\x02\x02\x02" +
		"\x02\x0F\x03\x02\x02\x02\x02\x11\x03\x02\x02\x02\x02\x13\x03\x02\x02\x02" +
		"\x02\x15\x03\x02\x02\x02\x02\x17\x03\x02\x02\x02\x02\x19\x03\x02\x02\x02" +
		"\x02\x1B\x03\x02\x02\x02\x02\x1D\x03\x02\x02\x02\x02\x1F\x03\x02\x02\x02" +
		"\x02!\x03\x02\x02\x02\x02#\x03\x02\x02\x02\x02%\x03\x02\x02\x02\x02\'" +
		"\x03\x02\x02\x02\x02)\x03\x02\x02\x02\x02+\x03\x02\x02\x02\x02-\x03\x02" +
		"\x02\x02\x02/\x03\x02\x02\x02\x031\x03\x02\x02\x02\x053\x03\x02\x02\x02" +
		"\x075\x03\x02\x02\x02\t7\x03\x02\x02\x02\v9\x03\x02\x02\x02\r;\x03\x02" +
		"\x02\x02\x0F=\x03\x02\x02\x02\x11?\x03\x02\x02\x02\x13A\x03\x02\x02\x02" +
		"\x15F\x03\x02\x02\x02\x17I\x03\x02\x02\x02\x19N\x03\x02\x02\x02\x1BT\x03" +
		"\x02\x02\x02\x1DW\x03\x02\x02\x02\x1FZ\x03\x02\x02\x02!\\\x03\x02\x02" +
		"\x02#_\x03\x02\x02\x02%a\x03\x02\x02\x02\'d\x03\x02\x02\x02)f\x03\x02" +
		"\x02\x02+i\x03\x02\x02\x02-p\x03\x02\x02\x02/u\x03\x02\x02\x0212\x07," +
		"\x02\x022\x04\x03\x02\x02\x0234\x071\x02\x024\x06\x03\x02\x02\x0256\x07" +
		"-\x02\x026\b\x03\x02\x02\x0278\x07/\x02\x028\n\x03\x02\x02\x029:\x07*" +
		"\x02\x02:\f\x03\x02\x02\x02;<\x07+\x02\x02<\x0E\x03\x02\x02\x02=>\x07" +
		"?\x02\x02>\x10\x03\x02\x02\x02?@\x07=\x02\x02@\x12\x03\x02\x02\x02AB\x07" +
		"t\x02\x02BC\x07g\x02\x02CD\x07c\x02\x02DE\x07f\x02\x02E\x14\x03\x02\x02" +
		"\x02FG\x07k\x02\x02GH\x07h\x02\x02H\x16\x03\x02\x02\x02IJ\x07g\x02\x02" +
		"JK\x07n\x02\x02KL\x07u\x02\x02LM\x07g\x02\x02M\x18\x03\x02\x02\x02NO\x07" +
		"y\x02\x02OP\x07t\x02\x02PQ\x07k\x02\x02QR\x07v\x02\x02RS\x07g\x02\x02" +
		"S\x1A\x03\x02\x02\x02TU\x07?\x02\x02UV\x07?\x02\x02V\x1C\x03\x02\x02\x02" +
		"WX\x07#\x02\x02XY\x07?\x02\x02Y\x1E\x03\x02\x02\x02Z[\x07@\x02\x02[ \x03" +
		"\x02\x02\x02\\]\x07@\x02\x02]^\x07?\x02\x02^\"\x03\x02\x02\x02_`\x07>" +
		"\x02\x02`$\x03\x02\x02\x02ab\x07>\x02\x02bc\x07?\x02\x02c&\x03\x02\x02" +
		"\x02de\x07}\x02\x02e(\x03\x02\x02\x02fg\x07\x7F\x02\x02g*\x03\x02\x02" +
		"\x02hj\t\x02\x02\x02ih\x03\x02\x02\x02jk\x03\x02\x02\x02ki\x03\x02\x02" +
		"\x02kl\x03\x02\x02\x02lm\x03\x02\x02\x02mn\b\x16\x02\x02n,\x03\x02\x02" +
		"\x02oq\t\x03\x02\x02po\x03\x02\x02\x02qr\x03\x02\x02\x02rp\x03\x02\x02" +
		"\x02rs\x03\x02\x02\x02s.\x03\x02\x02\x02tv\t\x04\x02\x02ut\x03\x02\x02" +
		"\x02vw\x03\x02\x02\x02wu\x03\x02\x02\x02wx\x03\x02\x02\x02x0\x03\x02\x02" +
		"\x02\x06\x02krw\x03\b\x02\x02";
	public static __ATN: ATN;
	public static get _ATN(): ATN {
		if (!CalcPlusLexer.__ATN) {
			CalcPlusLexer.__ATN = new ATNDeserializer().deserialize(Utils.toCharArray(CalcPlusLexer._serializedATN));
		}

		return CalcPlusLexer.__ATN;
	}

}

