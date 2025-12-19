// Code generated from CalcPlus.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // CalcPlus
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

type CalcPlusParser struct {
	*antlr.BaseParser
}

var CalcPlusParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func calcplusParserInit() {
	staticData := &CalcPlusParserStaticData
	staticData.LiteralNames = []string{
		"", "'*'", "'/'", "'+'", "'-'", "'('", "')'", "'='", "';'", "'read'",
		"'if'", "'else'", "'write'", "'=='", "'!='", "'>'", "'>='", "'<'", "'<='",
		"'{'", "'}'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "WS", "INT", "VAR",
	}
	staticData.RuleNames = []string{
		"calc0", "expr", "calc1", "stmt", "calc2", "cond", "block", "calc3",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 23, 102, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 27, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 5, 1, 35, 8, 1, 10, 1, 12, 1, 38, 9, 1, 1, 2, 4, 2, 41, 8, 2, 11,
		2, 12, 2, 42, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 65,
		8, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 73, 8, 3, 1, 4, 4, 4, 76,
		8, 4, 11, 4, 12, 4, 77, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6,
		5, 6, 88, 8, 6, 10, 6, 12, 6, 91, 9, 6, 1, 6, 1, 6, 1, 7, 4, 7, 96, 8,
		7, 11, 7, 12, 7, 97, 1, 7, 1, 7, 1, 7, 0, 1, 2, 8, 0, 2, 4, 6, 8, 10, 12,
		14, 0, 3, 1, 0, 1, 2, 1, 0, 3, 4, 1, 0, 13, 18, 105, 0, 16, 1, 0, 0, 0,
		2, 26, 1, 0, 0, 0, 4, 40, 1, 0, 0, 0, 6, 72, 1, 0, 0, 0, 8, 75, 1, 0, 0,
		0, 10, 81, 1, 0, 0, 0, 12, 85, 1, 0, 0, 0, 14, 95, 1, 0, 0, 0, 16, 17,
		3, 2, 1, 0, 17, 18, 5, 0, 0, 1, 18, 1, 1, 0, 0, 0, 19, 20, 6, 1, -1, 0,
		20, 27, 5, 22, 0, 0, 21, 27, 5, 23, 0, 0, 22, 23, 5, 5, 0, 0, 23, 24, 3,
		2, 1, 0, 24, 25, 5, 6, 0, 0, 25, 27, 1, 0, 0, 0, 26, 19, 1, 0, 0, 0, 26,
		21, 1, 0, 0, 0, 26, 22, 1, 0, 0, 0, 27, 36, 1, 0, 0, 0, 28, 29, 10, 5,
		0, 0, 29, 30, 7, 0, 0, 0, 30, 35, 3, 2, 1, 6, 31, 32, 10, 4, 0, 0, 32,
		33, 7, 1, 0, 0, 33, 35, 3, 2, 1, 5, 34, 28, 1, 0, 0, 0, 34, 31, 1, 0, 0,
		0, 35, 38, 1, 0, 0, 0, 36, 34, 1, 0, 0, 0, 36, 37, 1, 0, 0, 0, 37, 3, 1,
		0, 0, 0, 38, 36, 1, 0, 0, 0, 39, 41, 3, 6, 3, 0, 40, 39, 1, 0, 0, 0, 41,
		42, 1, 0, 0, 0, 42, 40, 1, 0, 0, 0, 42, 43, 1, 0, 0, 0, 43, 44, 1, 0, 0,
		0, 44, 45, 5, 0, 0, 1, 45, 5, 1, 0, 0, 0, 46, 47, 5, 23, 0, 0, 47, 48,
		5, 7, 0, 0, 48, 49, 3, 2, 1, 0, 49, 50, 5, 8, 0, 0, 50, 73, 1, 0, 0, 0,
		51, 52, 5, 23, 0, 0, 52, 53, 5, 7, 0, 0, 53, 54, 5, 9, 0, 0, 54, 55, 5,
		5, 0, 0, 55, 56, 5, 6, 0, 0, 56, 73, 5, 8, 0, 0, 57, 58, 5, 10, 0, 0, 58,
		59, 5, 5, 0, 0, 59, 60, 3, 10, 5, 0, 60, 61, 5, 6, 0, 0, 61, 64, 3, 12,
		6, 0, 62, 63, 5, 11, 0, 0, 63, 65, 3, 12, 6, 0, 64, 62, 1, 0, 0, 0, 64,
		65, 1, 0, 0, 0, 65, 73, 1, 0, 0, 0, 66, 67, 5, 12, 0, 0, 67, 68, 5, 5,
		0, 0, 68, 69, 3, 2, 1, 0, 69, 70, 5, 6, 0, 0, 70, 71, 5, 8, 0, 0, 71, 73,
		1, 0, 0, 0, 72, 46, 1, 0, 0, 0, 72, 51, 1, 0, 0, 0, 72, 57, 1, 0, 0, 0,
		72, 66, 1, 0, 0, 0, 73, 7, 1, 0, 0, 0, 74, 76, 3, 6, 3, 0, 75, 74, 1, 0,
		0, 0, 76, 77, 1, 0, 0, 0, 77, 75, 1, 0, 0, 0, 77, 78, 1, 0, 0, 0, 78, 79,
		1, 0, 0, 0, 79, 80, 5, 0, 0, 1, 80, 9, 1, 0, 0, 0, 81, 82, 3, 2, 1, 0,
		82, 83, 7, 2, 0, 0, 83, 84, 3, 2, 1, 0, 84, 11, 1, 0, 0, 0, 85, 89, 5,
		19, 0, 0, 86, 88, 3, 6, 3, 0, 87, 86, 1, 0, 0, 0, 88, 91, 1, 0, 0, 0, 89,
		87, 1, 0, 0, 0, 89, 90, 1, 0, 0, 0, 90, 92, 1, 0, 0, 0, 91, 89, 1, 0, 0,
		0, 92, 93, 5, 20, 0, 0, 93, 13, 1, 0, 0, 0, 94, 96, 3, 6, 3, 0, 95, 94,
		1, 0, 0, 0, 96, 97, 1, 0, 0, 0, 97, 95, 1, 0, 0, 0, 97, 98, 1, 0, 0, 0,
		98, 99, 1, 0, 0, 0, 99, 100, 5, 0, 0, 1, 100, 15, 1, 0, 0, 0, 9, 26, 34,
		36, 42, 64, 72, 77, 89, 97,
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

// CalcPlusParserInit initializes any static state used to implement CalcPlusParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewCalcPlusParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func CalcPlusParserInit() {
	staticData := &CalcPlusParserStaticData
	staticData.once.Do(calcplusParserInit)
}

// NewCalcPlusParser produces a new parser instance for the optional input antlr.TokenStream.
func NewCalcPlusParser(input antlr.TokenStream) *CalcPlusParser {
	CalcPlusParserInit()
	this := new(CalcPlusParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &CalcPlusParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "CalcPlus.g4"

	return this
}

// CalcPlusParser tokens.
const (
	CalcPlusParserEOF   = antlr.TokenEOF
	CalcPlusParserT__0  = 1
	CalcPlusParserT__1  = 2
	CalcPlusParserT__2  = 3
	CalcPlusParserT__3  = 4
	CalcPlusParserT__4  = 5
	CalcPlusParserT__5  = 6
	CalcPlusParserT__6  = 7
	CalcPlusParserT__7  = 8
	CalcPlusParserT__8  = 9
	CalcPlusParserT__9  = 10
	CalcPlusParserT__10 = 11
	CalcPlusParserT__11 = 12
	CalcPlusParserT__12 = 13
	CalcPlusParserT__13 = 14
	CalcPlusParserT__14 = 15
	CalcPlusParserT__15 = 16
	CalcPlusParserT__16 = 17
	CalcPlusParserT__17 = 18
	CalcPlusParserT__18 = 19
	CalcPlusParserT__19 = 20
	CalcPlusParserWS    = 21
	CalcPlusParserINT   = 22
	CalcPlusParserVAR   = 23
)

// CalcPlusParser rules.
const (
	CalcPlusParserRULE_calc0 = 0
	CalcPlusParserRULE_expr  = 1
	CalcPlusParserRULE_calc1 = 2
	CalcPlusParserRULE_stmt  = 3
	CalcPlusParserRULE_calc2 = 4
	CalcPlusParserRULE_cond  = 5
	CalcPlusParserRULE_block = 6
	CalcPlusParserRULE_calc3 = 7
)

// ICalc0Context is an interface to support dynamic dispatch.
type ICalc0Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext
	EOF() antlr.TerminalNode

	// IsCalc0Context differentiates from other interfaces.
	IsCalc0Context()
}

type Calc0Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCalc0Context() *Calc0Context {
	var p = new(Calc0Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcPlusParserRULE_calc0
	return p
}

func InitEmptyCalc0Context(p *Calc0Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcPlusParserRULE_calc0
}

func (*Calc0Context) IsCalc0Context() {}

func NewCalc0Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Calc0Context {
	var p = new(Calc0Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcPlusParserRULE_calc0

	return p
}

func (s *Calc0Context) GetParser() antlr.Parser { return s.parser }

func (s *Calc0Context) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Calc0Context) EOF() antlr.TerminalNode {
	return s.GetToken(CalcPlusParserEOF, 0)
}

func (s *Calc0Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Calc0Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Calc0Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcPlusListener); ok {
		listenerT.EnterCalc0(s)
	}
}

func (s *Calc0Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcPlusListener); ok {
		listenerT.ExitCalc0(s)
	}
}

func (s *Calc0Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CalcPlusVisitor:
		return t.VisitCalc0(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CalcPlusParser) Calc0() (localctx ICalc0Context) {
	localctx = NewCalc0Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CalcPlusParserRULE_calc0)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(16)
		p.expr(0)
	}
	{
		p.SetState(17)
		p.Match(CalcPlusParserEOF)
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

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcPlusParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcPlusParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcPlusParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyAll(ctx *ExprContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type MulDivContext struct {
	ExprContext
}

func NewMulDivContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MulDivContext {
	var p = new(MulDivContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *MulDivContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulDivContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *MulDivContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *MulDivContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcPlusListener); ok {
		listenerT.EnterMulDiv(s)
	}
}

func (s *MulDivContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcPlusListener); ok {
		listenerT.ExitMulDiv(s)
	}
}

func (s *MulDivContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CalcPlusVisitor:
		return t.VisitMulDiv(s)

	default:
		return t.VisitChildren(s)
	}
}

type AddSubContext struct {
	ExprContext
}

func NewAddSubContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddSubContext {
	var p = new(AddSubContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *AddSubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddSubContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *AddSubContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *AddSubContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcPlusListener); ok {
		listenerT.EnterAddSub(s)
	}
}

func (s *AddSubContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcPlusListener); ok {
		listenerT.ExitAddSub(s)
	}
}

func (s *AddSubContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CalcPlusVisitor:
		return t.VisitAddSub(s)

	default:
		return t.VisitChildren(s)
	}
}

type VarContext struct {
	ExprContext
}

func NewVarContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VarContext {
	var p = new(VarContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *VarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarContext) VAR() antlr.TerminalNode {
	return s.GetToken(CalcPlusParserVAR, 0)
}

func (s *VarContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcPlusListener); ok {
		listenerT.EnterVar(s)
	}
}

func (s *VarContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcPlusListener); ok {
		listenerT.ExitVar(s)
	}
}

func (s *VarContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CalcPlusVisitor:
		return t.VisitVar(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParensContext struct {
	ExprContext
}

func NewParensContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParensContext {
	var p = new(ParensContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ParensContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParensContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ParensContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcPlusListener); ok {
		listenerT.EnterParens(s)
	}
}

func (s *ParensContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcPlusListener); ok {
		listenerT.ExitParens(s)
	}
}

func (s *ParensContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CalcPlusVisitor:
		return t.VisitParens(s)

	default:
		return t.VisitChildren(s)
	}
}

type IntContext struct {
	ExprContext
}

func NewIntContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntContext {
	var p = new(IntContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *IntContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntContext) INT() antlr.TerminalNode {
	return s.GetToken(CalcPlusParserINT, 0)
}

func (s *IntContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcPlusListener); ok {
		listenerT.EnterInt(s)
	}
}

func (s *IntContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcPlusListener); ok {
		listenerT.ExitInt(s)
	}
}

func (s *IntContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CalcPlusVisitor:
		return t.VisitInt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CalcPlusParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *CalcPlusParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, CalcPlusParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(26)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CalcPlusParserINT:
		localctx = NewIntContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(20)
			p.Match(CalcPlusParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case CalcPlusParserVAR:
		localctx = NewVarContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(21)
			p.Match(CalcPlusParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case CalcPlusParserT__4:
		localctx = NewParensContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(22)
			p.Match(CalcPlusParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(23)
			p.expr(0)
		}
		{
			p.SetState(24)
			p.Match(CalcPlusParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(36)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(34)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMulDivContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CalcPlusParserRULE_expr)
				p.SetState(28)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(29)
					_la = p.GetTokenStream().LA(1)

					if !(_la == CalcPlusParserT__0 || _la == CalcPlusParserT__1) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(30)
					p.expr(6)
				}

			case 2:
				localctx = NewAddSubContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CalcPlusParserRULE_expr)
				p.SetState(31)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(32)
					_la = p.GetTokenStream().LA(1)

					if !(_la == CalcPlusParserT__2 || _la == CalcPlusParserT__3) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(33)
					p.expr(5)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(38)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext())
		if p.HasError() {
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
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICalc1Context is an interface to support dynamic dispatch.
type ICalc1Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllStmt() []IStmtContext
	Stmt(i int) IStmtContext

	// IsCalc1Context differentiates from other interfaces.
	IsCalc1Context()
}

type Calc1Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCalc1Context() *Calc1Context {
	var p = new(Calc1Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcPlusParserRULE_calc1
	return p
}

func InitEmptyCalc1Context(p *Calc1Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcPlusParserRULE_calc1
}

func (*Calc1Context) IsCalc1Context() {}

func NewCalc1Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Calc1Context {
	var p = new(Calc1Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcPlusParserRULE_calc1

	return p
}

func (s *Calc1Context) GetParser() antlr.Parser { return s.parser }

func (s *Calc1Context) EOF() antlr.TerminalNode {
	return s.GetToken(CalcPlusParserEOF, 0)
}

func (s *Calc1Context) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *Calc1Context) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
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

	return t.(IStmtContext)
}

func (s *Calc1Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Calc1Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Calc1Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcPlusListener); ok {
		listenerT.EnterCalc1(s)
	}
}

func (s *Calc1Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcPlusListener); ok {
		listenerT.ExitCalc1(s)
	}
}

func (s *Calc1Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CalcPlusVisitor:
		return t.VisitCalc1(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CalcPlusParser) Calc1() (localctx ICalc1Context) {
	localctx = NewCalc1Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, CalcPlusParserRULE_calc1)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(40)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8393728) != 0) {
		{
			p.SetState(39)
			p.Stmt()
		}

		p.SetState(42)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(44)
		p.Match(CalcPlusParserEOF)
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

// IStmtContext is an interface to support dynamic dispatch.
type IStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsStmtContext differentiates from other interfaces.
	IsStmtContext()
}

type StmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStmtContext() *StmtContext {
	var p = new(StmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcPlusParserRULE_stmt
	return p
}

func InitEmptyStmtContext(p *StmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcPlusParserRULE_stmt
}

func (*StmtContext) IsStmtContext() {}

func NewStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StmtContext {
	var p = new(StmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcPlusParserRULE_stmt

	return p
}

func (s *StmtContext) GetParser() antlr.Parser { return s.parser }

func (s *StmtContext) CopyAll(ctx *StmtContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *StmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type WriteContext struct {
	StmtContext
}

func NewWriteContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *WriteContext {
	var p = new(WriteContext)

	InitEmptyStmtContext(&p.StmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*StmtContext))

	return p
}

func (s *WriteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WriteContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *WriteContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcPlusListener); ok {
		listenerT.EnterWrite(s)
	}
}

func (s *WriteContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcPlusListener); ok {
		listenerT.ExitWrite(s)
	}
}

func (s *WriteContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CalcPlusVisitor:
		return t.VisitWrite(s)

	default:
		return t.VisitChildren(s)
	}
}

type IfElseContext struct {
	StmtContext
	thenBlock IBlockContext
	elseBlock IBlockContext
}

func NewIfElseContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfElseContext {
	var p = new(IfElseContext)

	InitEmptyStmtContext(&p.StmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*StmtContext))

	return p
}

func (s *IfElseContext) GetThenBlock() IBlockContext { return s.thenBlock }

func (s *IfElseContext) GetElseBlock() IBlockContext { return s.elseBlock }

func (s *IfElseContext) SetThenBlock(v IBlockContext) { s.thenBlock = v }

func (s *IfElseContext) SetElseBlock(v IBlockContext) { s.elseBlock = v }

func (s *IfElseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfElseContext) Cond() ICondContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICondContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICondContext)
}

func (s *IfElseContext) AllBlock() []IBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockContext); ok {
			len++
		}
	}

	tst := make([]IBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockContext); ok {
			tst[i] = t.(IBlockContext)
			i++
		}
	}

	return tst
}

func (s *IfElseContext) Block(i int) IBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
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

	return t.(IBlockContext)
}

func (s *IfElseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcPlusListener); ok {
		listenerT.EnterIfElse(s)
	}
}

func (s *IfElseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcPlusListener); ok {
		listenerT.ExitIfElse(s)
	}
}

func (s *IfElseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CalcPlusVisitor:
		return t.VisitIfElse(s)

	default:
		return t.VisitChildren(s)
	}
}

type ReadAssignContext struct {
	StmtContext
}

func NewReadAssignContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ReadAssignContext {
	var p = new(ReadAssignContext)

	InitEmptyStmtContext(&p.StmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*StmtContext))

	return p
}

func (s *ReadAssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReadAssignContext) VAR() antlr.TerminalNode {
	return s.GetToken(CalcPlusParserVAR, 0)
}

func (s *ReadAssignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcPlusListener); ok {
		listenerT.EnterReadAssign(s)
	}
}

func (s *ReadAssignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcPlusListener); ok {
		listenerT.ExitReadAssign(s)
	}
}

func (s *ReadAssignContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CalcPlusVisitor:
		return t.VisitReadAssign(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprAssignContext struct {
	StmtContext
}

func NewExprAssignContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprAssignContext {
	var p = new(ExprAssignContext)

	InitEmptyStmtContext(&p.StmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*StmtContext))

	return p
}

func (s *ExprAssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprAssignContext) VAR() antlr.TerminalNode {
	return s.GetToken(CalcPlusParserVAR, 0)
}

func (s *ExprAssignContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprAssignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcPlusListener); ok {
		listenerT.EnterExprAssign(s)
	}
}

func (s *ExprAssignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcPlusListener); ok {
		listenerT.ExitExprAssign(s)
	}
}

func (s *ExprAssignContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CalcPlusVisitor:
		return t.VisitExprAssign(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CalcPlusParser) Stmt() (localctx IStmtContext) {
	localctx = NewStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, CalcPlusParserRULE_stmt)
	var _la int

	p.SetState(72)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		localctx = NewExprAssignContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(46)
			p.Match(CalcPlusParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(47)
			p.Match(CalcPlusParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(48)
			p.expr(0)
		}
		{
			p.SetState(49)
			p.Match(CalcPlusParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewReadAssignContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(51)
			p.Match(CalcPlusParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(52)
			p.Match(CalcPlusParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(53)
			p.Match(CalcPlusParserT__8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(54)
			p.Match(CalcPlusParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(55)
			p.Match(CalcPlusParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(56)
			p.Match(CalcPlusParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewIfElseContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(57)
			p.Match(CalcPlusParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(58)
			p.Match(CalcPlusParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(59)
			p.Cond()
		}
		{
			p.SetState(60)
			p.Match(CalcPlusParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(61)

			var _x = p.Block()

			localctx.(*IfElseContext).thenBlock = _x
		}
		p.SetState(64)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == CalcPlusParserT__10 {
			{
				p.SetState(62)
				p.Match(CalcPlusParserT__10)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(63)

				var _x = p.Block()

				localctx.(*IfElseContext).elseBlock = _x
			}

		}

	case 4:
		localctx = NewWriteContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(66)
			p.Match(CalcPlusParserT__11)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(67)
			p.Match(CalcPlusParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(68)
			p.expr(0)
		}
		{
			p.SetState(69)
			p.Match(CalcPlusParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(70)
			p.Match(CalcPlusParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
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

// ICalc2Context is an interface to support dynamic dispatch.
type ICalc2Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllStmt() []IStmtContext
	Stmt(i int) IStmtContext

	// IsCalc2Context differentiates from other interfaces.
	IsCalc2Context()
}

type Calc2Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCalc2Context() *Calc2Context {
	var p = new(Calc2Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcPlusParserRULE_calc2
	return p
}

func InitEmptyCalc2Context(p *Calc2Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcPlusParserRULE_calc2
}

func (*Calc2Context) IsCalc2Context() {}

func NewCalc2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Calc2Context {
	var p = new(Calc2Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcPlusParserRULE_calc2

	return p
}

func (s *Calc2Context) GetParser() antlr.Parser { return s.parser }

func (s *Calc2Context) EOF() antlr.TerminalNode {
	return s.GetToken(CalcPlusParserEOF, 0)
}

func (s *Calc2Context) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *Calc2Context) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
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

	return t.(IStmtContext)
}

func (s *Calc2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Calc2Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Calc2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcPlusListener); ok {
		listenerT.EnterCalc2(s)
	}
}

func (s *Calc2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcPlusListener); ok {
		listenerT.ExitCalc2(s)
	}
}

func (s *Calc2Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CalcPlusVisitor:
		return t.VisitCalc2(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CalcPlusParser) Calc2() (localctx ICalc2Context) {
	localctx = NewCalc2Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, CalcPlusParserRULE_calc2)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(75)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8393728) != 0) {
		{
			p.SetState(74)
			p.Stmt()
		}

		p.SetState(77)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(79)
		p.Match(CalcPlusParserEOF)
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

// ICondContext is an interface to support dynamic dispatch.
type ICondContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpr() []IExprContext
	Expr(i int) IExprContext

	// IsCondContext differentiates from other interfaces.
	IsCondContext()
}

type CondContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCondContext() *CondContext {
	var p = new(CondContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcPlusParserRULE_cond
	return p
}

func InitEmptyCondContext(p *CondContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcPlusParserRULE_cond
}

func (*CondContext) IsCondContext() {}

func NewCondContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CondContext {
	var p = new(CondContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcPlusParserRULE_cond

	return p
}

func (s *CondContext) GetParser() antlr.Parser { return s.parser }

func (s *CondContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *CondContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *CondContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CondContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CondContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcPlusListener); ok {
		listenerT.EnterCond(s)
	}
}

func (s *CondContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcPlusListener); ok {
		listenerT.ExitCond(s)
	}
}

func (s *CondContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CalcPlusVisitor:
		return t.VisitCond(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CalcPlusParser) Cond() (localctx ICondContext) {
	localctx = NewCondContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, CalcPlusParserRULE_cond)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(81)
		p.expr(0)
	}
	{
		p.SetState(82)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&516096) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(83)
		p.expr(0)
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

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStmt() []IStmtContext
	Stmt(i int) IStmtContext

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcPlusParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcPlusParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcPlusParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *BlockContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
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

	return t.(IStmtContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcPlusListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcPlusListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CalcPlusVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CalcPlusParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, CalcPlusParserRULE_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(85)
		p.Match(CalcPlusParserT__18)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(89)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8393728) != 0 {
		{
			p.SetState(86)
			p.Stmt()
		}

		p.SetState(91)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(92)
		p.Match(CalcPlusParserT__19)
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

// ICalc3Context is an interface to support dynamic dispatch.
type ICalc3Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllStmt() []IStmtContext
	Stmt(i int) IStmtContext

	// IsCalc3Context differentiates from other interfaces.
	IsCalc3Context()
}

type Calc3Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCalc3Context() *Calc3Context {
	var p = new(Calc3Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcPlusParserRULE_calc3
	return p
}

func InitEmptyCalc3Context(p *Calc3Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcPlusParserRULE_calc3
}

func (*Calc3Context) IsCalc3Context() {}

func NewCalc3Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Calc3Context {
	var p = new(Calc3Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcPlusParserRULE_calc3

	return p
}

func (s *Calc3Context) GetParser() antlr.Parser { return s.parser }

func (s *Calc3Context) EOF() antlr.TerminalNode {
	return s.GetToken(CalcPlusParserEOF, 0)
}

func (s *Calc3Context) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *Calc3Context) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
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

	return t.(IStmtContext)
}

func (s *Calc3Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Calc3Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Calc3Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcPlusListener); ok {
		listenerT.EnterCalc3(s)
	}
}

func (s *Calc3Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcPlusListener); ok {
		listenerT.ExitCalc3(s)
	}
}

func (s *Calc3Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CalcPlusVisitor:
		return t.VisitCalc3(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CalcPlusParser) Calc3() (localctx ICalc3Context) {
	localctx = NewCalc3Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, CalcPlusParserRULE_calc3)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(95)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8393728) != 0) {
		{
			p.SetState(94)
			p.Stmt()
		}

		p.SetState(97)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(99)
		p.Match(CalcPlusParserEOF)
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

func (p *CalcPlusParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *CalcPlusParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
