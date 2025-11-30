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
		"", "'*'", "'/'", "'+'", "'-'", "'('", "')'", "'='", "';'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "WS", "INT", "VAR",
	}
	staticData.RuleNames = []string{
		"calc0", "expr", "calc1", "stmt",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 11, 44, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 1, 0, 1,
		0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 19, 8, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 27, 8, 1, 10, 1, 12, 1, 30, 9, 1, 1,
		2, 4, 2, 33, 8, 2, 11, 2, 12, 2, 34, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 0, 1, 2, 4, 0, 2, 4, 6, 0, 2, 1, 0, 1, 2, 1, 0, 3, 4, 44, 0,
		8, 1, 0, 0, 0, 2, 18, 1, 0, 0, 0, 4, 32, 1, 0, 0, 0, 6, 38, 1, 0, 0, 0,
		8, 9, 3, 2, 1, 0, 9, 10, 5, 0, 0, 1, 10, 1, 1, 0, 0, 0, 11, 12, 6, 1, -1,
		0, 12, 19, 5, 10, 0, 0, 13, 19, 5, 11, 0, 0, 14, 15, 5, 5, 0, 0, 15, 16,
		3, 2, 1, 0, 16, 17, 5, 6, 0, 0, 17, 19, 1, 0, 0, 0, 18, 11, 1, 0, 0, 0,
		18, 13, 1, 0, 0, 0, 18, 14, 1, 0, 0, 0, 19, 28, 1, 0, 0, 0, 20, 21, 10,
		5, 0, 0, 21, 22, 7, 0, 0, 0, 22, 27, 3, 2, 1, 6, 23, 24, 10, 4, 0, 0, 24,
		25, 7, 1, 0, 0, 25, 27, 3, 2, 1, 5, 26, 20, 1, 0, 0, 0, 26, 23, 1, 0, 0,
		0, 27, 30, 1, 0, 0, 0, 28, 26, 1, 0, 0, 0, 28, 29, 1, 0, 0, 0, 29, 3, 1,
		0, 0, 0, 30, 28, 1, 0, 0, 0, 31, 33, 3, 6, 3, 0, 32, 31, 1, 0, 0, 0, 33,
		34, 1, 0, 0, 0, 34, 32, 1, 0, 0, 0, 34, 35, 1, 0, 0, 0, 35, 36, 1, 0, 0,
		0, 36, 37, 5, 0, 0, 1, 37, 5, 1, 0, 0, 0, 38, 39, 5, 11, 0, 0, 39, 40,
		5, 7, 0, 0, 40, 41, 3, 2, 1, 0, 41, 42, 5, 8, 0, 0, 42, 7, 1, 0, 0, 0,
		4, 18, 26, 28, 34,
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
	CalcPlusParserEOF  = antlr.TokenEOF
	CalcPlusParserT__0 = 1
	CalcPlusParserT__1 = 2
	CalcPlusParserT__2 = 3
	CalcPlusParserT__3 = 4
	CalcPlusParserT__4 = 5
	CalcPlusParserT__5 = 6
	CalcPlusParserT__6 = 7
	CalcPlusParserT__7 = 8
	CalcPlusParserWS   = 9
	CalcPlusParserINT  = 10
	CalcPlusParserVAR  = 11
)

// CalcPlusParser rules.
const (
	CalcPlusParserRULE_calc0 = 0
	CalcPlusParserRULE_expr  = 1
	CalcPlusParserRULE_calc1 = 2
	CalcPlusParserRULE_stmt  = 3
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
		p.SetState(8)
		p.expr(0)
	}
	{
		p.SetState(9)
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
	p.SetState(18)
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
			p.SetState(12)
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
			p.SetState(13)
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
			p.SetState(14)
			p.Match(CalcPlusParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(15)
			p.expr(0)
		}
		{
			p.SetState(16)
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
	p.SetState(28)
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
			p.SetState(26)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMulDivContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CalcPlusParserRULE_expr)
				p.SetState(20)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(21)
					_la = p.GetTokenStream().LA(1)

					if !(_la == CalcPlusParserT__0 || _la == CalcPlusParserT__1) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(22)
					p.expr(6)
				}

			case 2:
				localctx = NewAddSubContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CalcPlusParserRULE_expr)
				p.SetState(23)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(24)
					_la = p.GetTokenStream().LA(1)

					if !(_la == CalcPlusParserT__2 || _la == CalcPlusParserT__3) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(25)
					p.expr(5)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(30)
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
	p.SetState(32)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == CalcPlusParserVAR {
		{
			p.SetState(31)
			p.Stmt()
		}

		p.SetState(34)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(36)
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
	localctx = NewExprAssignContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(38)
		p.Match(CalcPlusParserVAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(39)
		p.Match(CalcPlusParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(40)
		p.expr(0)
	}
	{
		p.SetState(41)
		p.Match(CalcPlusParserT__7)
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
