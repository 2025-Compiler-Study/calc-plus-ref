// Code generated from CalcPlus.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type CalcPlusLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var CalcPlusLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func calcpluslexerLexerInit() {
	staticData := &CalcPlusLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'int'", "','", "';'", "'='", "'read'", "'('", "')'", "'write'",
		"'if'", "'else'", "'*'", "'/'", "'+'", "'-'", "'=='", "'!='", "'>'",
		"'>='", "'<'", "'<='", "'{'", "'}'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "WS", "INT", "VAR",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
		"T__17", "T__18", "T__19", "T__20", "T__21", "WS", "INT", "VAR",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 25, 129, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 1, 0, 1, 0,
		1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8,
		1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1,
		12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 16,
		1, 16, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 20, 1,
		20, 1, 21, 1, 21, 1, 22, 4, 22, 114, 8, 22, 11, 22, 12, 22, 115, 1, 22,
		1, 22, 1, 23, 4, 23, 121, 8, 23, 11, 23, 12, 23, 122, 1, 24, 4, 24, 126,
		8, 24, 11, 24, 12, 24, 127, 0, 0, 25, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11,
		6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15,
		31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24,
		49, 25, 1, 0, 3, 3, 0, 9, 10, 13, 13, 32, 32, 1, 0, 48, 57, 2, 0, 65, 90,
		97, 122, 131, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0,
		7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0,
		0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0,
		0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0,
		0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1,
		0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45,
		1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 1, 51, 1, 0, 0, 0, 3,
		55, 1, 0, 0, 0, 5, 57, 1, 0, 0, 0, 7, 59, 1, 0, 0, 0, 9, 61, 1, 0, 0, 0,
		11, 66, 1, 0, 0, 0, 13, 68, 1, 0, 0, 0, 15, 70, 1, 0, 0, 0, 17, 76, 1,
		0, 0, 0, 19, 79, 1, 0, 0, 0, 21, 84, 1, 0, 0, 0, 23, 86, 1, 0, 0, 0, 25,
		88, 1, 0, 0, 0, 27, 90, 1, 0, 0, 0, 29, 92, 1, 0, 0, 0, 31, 95, 1, 0, 0,
		0, 33, 98, 1, 0, 0, 0, 35, 100, 1, 0, 0, 0, 37, 103, 1, 0, 0, 0, 39, 105,
		1, 0, 0, 0, 41, 108, 1, 0, 0, 0, 43, 110, 1, 0, 0, 0, 45, 113, 1, 0, 0,
		0, 47, 120, 1, 0, 0, 0, 49, 125, 1, 0, 0, 0, 51, 52, 5, 105, 0, 0, 52,
		53, 5, 110, 0, 0, 53, 54, 5, 116, 0, 0, 54, 2, 1, 0, 0, 0, 55, 56, 5, 44,
		0, 0, 56, 4, 1, 0, 0, 0, 57, 58, 5, 59, 0, 0, 58, 6, 1, 0, 0, 0, 59, 60,
		5, 61, 0, 0, 60, 8, 1, 0, 0, 0, 61, 62, 5, 114, 0, 0, 62, 63, 5, 101, 0,
		0, 63, 64, 5, 97, 0, 0, 64, 65, 5, 100, 0, 0, 65, 10, 1, 0, 0, 0, 66, 67,
		5, 40, 0, 0, 67, 12, 1, 0, 0, 0, 68, 69, 5, 41, 0, 0, 69, 14, 1, 0, 0,
		0, 70, 71, 5, 119, 0, 0, 71, 72, 5, 114, 0, 0, 72, 73, 5, 105, 0, 0, 73,
		74, 5, 116, 0, 0, 74, 75, 5, 101, 0, 0, 75, 16, 1, 0, 0, 0, 76, 77, 5,
		105, 0, 0, 77, 78, 5, 102, 0, 0, 78, 18, 1, 0, 0, 0, 79, 80, 5, 101, 0,
		0, 80, 81, 5, 108, 0, 0, 81, 82, 5, 115, 0, 0, 82, 83, 5, 101, 0, 0, 83,
		20, 1, 0, 0, 0, 84, 85, 5, 42, 0, 0, 85, 22, 1, 0, 0, 0, 86, 87, 5, 47,
		0, 0, 87, 24, 1, 0, 0, 0, 88, 89, 5, 43, 0, 0, 89, 26, 1, 0, 0, 0, 90,
		91, 5, 45, 0, 0, 91, 28, 1, 0, 0, 0, 92, 93, 5, 61, 0, 0, 93, 94, 5, 61,
		0, 0, 94, 30, 1, 0, 0, 0, 95, 96, 5, 33, 0, 0, 96, 97, 5, 61, 0, 0, 97,
		32, 1, 0, 0, 0, 98, 99, 5, 62, 0, 0, 99, 34, 1, 0, 0, 0, 100, 101, 5, 62,
		0, 0, 101, 102, 5, 61, 0, 0, 102, 36, 1, 0, 0, 0, 103, 104, 5, 60, 0, 0,
		104, 38, 1, 0, 0, 0, 105, 106, 5, 60, 0, 0, 106, 107, 5, 61, 0, 0, 107,
		40, 1, 0, 0, 0, 108, 109, 5, 123, 0, 0, 109, 42, 1, 0, 0, 0, 110, 111,
		5, 125, 0, 0, 111, 44, 1, 0, 0, 0, 112, 114, 7, 0, 0, 0, 113, 112, 1, 0,
		0, 0, 114, 115, 1, 0, 0, 0, 115, 113, 1, 0, 0, 0, 115, 116, 1, 0, 0, 0,
		116, 117, 1, 0, 0, 0, 117, 118, 6, 22, 0, 0, 118, 46, 1, 0, 0, 0, 119,
		121, 7, 1, 0, 0, 120, 119, 1, 0, 0, 0, 121, 122, 1, 0, 0, 0, 122, 120,
		1, 0, 0, 0, 122, 123, 1, 0, 0, 0, 123, 48, 1, 0, 0, 0, 124, 126, 7, 2,
		0, 0, 125, 124, 1, 0, 0, 0, 126, 127, 1, 0, 0, 0, 127, 125, 1, 0, 0, 0,
		127, 128, 1, 0, 0, 0, 128, 50, 1, 0, 0, 0, 4, 0, 115, 122, 127, 1, 6, 0,
		0,
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

// CalcPlusLexerInit initializes any static state used to implement CalcPlusLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewCalcPlusLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func CalcPlusLexerInit() {
	staticData := &CalcPlusLexerLexerStaticData
	staticData.once.Do(calcpluslexerLexerInit)
}

// NewCalcPlusLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewCalcPlusLexer(input antlr.CharStream) *CalcPlusLexer {
	CalcPlusLexerInit()
	l := new(CalcPlusLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &CalcPlusLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "CalcPlus.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// CalcPlusLexer tokens.
const (
	CalcPlusLexerT__0  = 1
	CalcPlusLexerT__1  = 2
	CalcPlusLexerT__2  = 3
	CalcPlusLexerT__3  = 4
	CalcPlusLexerT__4  = 5
	CalcPlusLexerT__5  = 6
	CalcPlusLexerT__6  = 7
	CalcPlusLexerT__7  = 8
	CalcPlusLexerT__8  = 9
	CalcPlusLexerT__9  = 10
	CalcPlusLexerT__10 = 11
	CalcPlusLexerT__11 = 12
	CalcPlusLexerT__12 = 13
	CalcPlusLexerT__13 = 14
	CalcPlusLexerT__14 = 15
	CalcPlusLexerT__15 = 16
	CalcPlusLexerT__16 = 17
	CalcPlusLexerT__17 = 18
	CalcPlusLexerT__18 = 19
	CalcPlusLexerT__19 = 20
	CalcPlusLexerT__20 = 21
	CalcPlusLexerT__21 = 22
	CalcPlusLexerWS    = 23
	CalcPlusLexerINT   = 24
	CalcPlusLexerVAR   = 25
)
