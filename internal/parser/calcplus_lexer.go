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
		"", "'*'", "'/'", "'+'", "'-'", "'('", "')'", "'='", "';'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "WS", "INT", "VAR",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "WS",
		"INT", "VAR",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 11, 56, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4,
		1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 4, 8, 41, 8, 8, 11, 8, 12, 8,
		42, 1, 8, 1, 8, 1, 9, 4, 9, 48, 8, 9, 11, 9, 12, 9, 49, 1, 10, 4, 10, 53,
		8, 10, 11, 10, 12, 10, 54, 0, 0, 11, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11,
		6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 1, 0, 3, 3, 0, 9, 10, 13, 13, 32,
		32, 1, 0, 48, 57, 2, 0, 65, 90, 97, 122, 58, 0, 1, 1, 0, 0, 0, 0, 3, 1,
		0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1,
		0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19,
		1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 1, 23, 1, 0, 0, 0, 3, 25, 1, 0, 0, 0, 5,
		27, 1, 0, 0, 0, 7, 29, 1, 0, 0, 0, 9, 31, 1, 0, 0, 0, 11, 33, 1, 0, 0,
		0, 13, 35, 1, 0, 0, 0, 15, 37, 1, 0, 0, 0, 17, 40, 1, 0, 0, 0, 19, 47,
		1, 0, 0, 0, 21, 52, 1, 0, 0, 0, 23, 24, 5, 42, 0, 0, 24, 2, 1, 0, 0, 0,
		25, 26, 5, 47, 0, 0, 26, 4, 1, 0, 0, 0, 27, 28, 5, 43, 0, 0, 28, 6, 1,
		0, 0, 0, 29, 30, 5, 45, 0, 0, 30, 8, 1, 0, 0, 0, 31, 32, 5, 40, 0, 0, 32,
		10, 1, 0, 0, 0, 33, 34, 5, 41, 0, 0, 34, 12, 1, 0, 0, 0, 35, 36, 5, 61,
		0, 0, 36, 14, 1, 0, 0, 0, 37, 38, 5, 59, 0, 0, 38, 16, 1, 0, 0, 0, 39,
		41, 7, 0, 0, 0, 40, 39, 1, 0, 0, 0, 41, 42, 1, 0, 0, 0, 42, 40, 1, 0, 0,
		0, 42, 43, 1, 0, 0, 0, 43, 44, 1, 0, 0, 0, 44, 45, 6, 8, 0, 0, 45, 18,
		1, 0, 0, 0, 46, 48, 7, 1, 0, 0, 47, 46, 1, 0, 0, 0, 48, 49, 1, 0, 0, 0,
		49, 47, 1, 0, 0, 0, 49, 50, 1, 0, 0, 0, 50, 20, 1, 0, 0, 0, 51, 53, 7,
		2, 0, 0, 52, 51, 1, 0, 0, 0, 53, 54, 1, 0, 0, 0, 54, 52, 1, 0, 0, 0, 54,
		55, 1, 0, 0, 0, 55, 22, 1, 0, 0, 0, 4, 0, 42, 49, 54, 1, 6, 0, 0,
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
	CalcPlusLexerT__0 = 1
	CalcPlusLexerT__1 = 2
	CalcPlusLexerT__2 = 3
	CalcPlusLexerT__3 = 4
	CalcPlusLexerT__4 = 5
	CalcPlusLexerT__5 = 6
	CalcPlusLexerT__6 = 7
	CalcPlusLexerT__7 = 8
	CalcPlusLexerWS   = 9
	CalcPlusLexerINT  = 10
	CalcPlusLexerVAR  = 11
)
