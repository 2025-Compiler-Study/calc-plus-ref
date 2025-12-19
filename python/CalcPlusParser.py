# Generated from CalcPlus.g4 by ANTLR 4.13.2
# encoding: utf-8
from antlr4 import *
from io import StringIO
import sys
if sys.version_info[1] > 5:
	from typing import TextIO
else:
	from typing.io import TextIO

def serializedATN():
    return [
        4,1,23,102,2,0,7,0,2,1,7,1,2,2,7,2,2,3,7,3,2,4,7,4,2,5,7,5,2,6,7,
        6,2,7,7,7,1,0,1,0,1,0,1,1,1,1,1,1,1,1,1,1,1,1,1,1,3,1,27,8,1,1,1,
        1,1,1,1,1,1,1,1,1,1,5,1,35,8,1,10,1,12,1,38,9,1,1,2,4,2,41,8,2,11,
        2,12,2,42,1,2,1,2,1,3,1,3,1,3,1,3,1,3,1,3,1,3,1,3,1,3,1,3,1,3,1,
        3,1,3,1,3,1,3,1,3,1,3,1,3,3,3,65,8,3,1,3,1,3,1,3,1,3,1,3,1,3,3,3,
        73,8,3,1,4,4,4,76,8,4,11,4,12,4,77,1,4,1,4,1,5,1,5,1,5,1,5,1,6,1,
        6,5,6,88,8,6,10,6,12,6,91,9,6,1,6,1,6,1,7,4,7,96,8,7,11,7,12,7,97,
        1,7,1,7,1,7,0,1,2,8,0,2,4,6,8,10,12,14,0,3,1,0,1,2,1,0,3,4,1,0,13,
        18,105,0,16,1,0,0,0,2,26,1,0,0,0,4,40,1,0,0,0,6,72,1,0,0,0,8,75,
        1,0,0,0,10,81,1,0,0,0,12,85,1,0,0,0,14,95,1,0,0,0,16,17,3,2,1,0,
        17,18,5,0,0,1,18,1,1,0,0,0,19,20,6,1,-1,0,20,27,5,22,0,0,21,27,5,
        23,0,0,22,23,5,5,0,0,23,24,3,2,1,0,24,25,5,6,0,0,25,27,1,0,0,0,26,
        19,1,0,0,0,26,21,1,0,0,0,26,22,1,0,0,0,27,36,1,0,0,0,28,29,10,5,
        0,0,29,30,7,0,0,0,30,35,3,2,1,6,31,32,10,4,0,0,32,33,7,1,0,0,33,
        35,3,2,1,5,34,28,1,0,0,0,34,31,1,0,0,0,35,38,1,0,0,0,36,34,1,0,0,
        0,36,37,1,0,0,0,37,3,1,0,0,0,38,36,1,0,0,0,39,41,3,6,3,0,40,39,1,
        0,0,0,41,42,1,0,0,0,42,40,1,0,0,0,42,43,1,0,0,0,43,44,1,0,0,0,44,
        45,5,0,0,1,45,5,1,0,0,0,46,47,5,23,0,0,47,48,5,7,0,0,48,49,3,2,1,
        0,49,50,5,8,0,0,50,73,1,0,0,0,51,52,5,23,0,0,52,53,5,7,0,0,53,54,
        5,9,0,0,54,55,5,5,0,0,55,56,5,6,0,0,56,73,5,8,0,0,57,58,5,10,0,0,
        58,59,5,5,0,0,59,60,3,10,5,0,60,61,5,6,0,0,61,64,3,12,6,0,62,63,
        5,11,0,0,63,65,3,12,6,0,64,62,1,0,0,0,64,65,1,0,0,0,65,73,1,0,0,
        0,66,67,5,12,0,0,67,68,5,5,0,0,68,69,3,2,1,0,69,70,5,6,0,0,70,71,
        5,8,0,0,71,73,1,0,0,0,72,46,1,0,0,0,72,51,1,0,0,0,72,57,1,0,0,0,
        72,66,1,0,0,0,73,7,1,0,0,0,74,76,3,6,3,0,75,74,1,0,0,0,76,77,1,0,
        0,0,77,75,1,0,0,0,77,78,1,0,0,0,78,79,1,0,0,0,79,80,5,0,0,1,80,9,
        1,0,0,0,81,82,3,2,1,0,82,83,7,2,0,0,83,84,3,2,1,0,84,11,1,0,0,0,
        85,89,5,19,0,0,86,88,3,6,3,0,87,86,1,0,0,0,88,91,1,0,0,0,89,87,1,
        0,0,0,89,90,1,0,0,0,90,92,1,0,0,0,91,89,1,0,0,0,92,93,5,20,0,0,93,
        13,1,0,0,0,94,96,3,6,3,0,95,94,1,0,0,0,96,97,1,0,0,0,97,95,1,0,0,
        0,97,98,1,0,0,0,98,99,1,0,0,0,99,100,5,0,0,1,100,15,1,0,0,0,9,26,
        34,36,42,64,72,77,89,97
    ]

class CalcPlusParser ( Parser ):

    grammarFileName = "CalcPlus.g4"

    atn = ATNDeserializer().deserialize(serializedATN())

    decisionsToDFA = [ DFA(ds, i) for i, ds in enumerate(atn.decisionToState) ]

    sharedContextCache = PredictionContextCache()

    literalNames = [ "<INVALID>", "'*'", "'/'", "'+'", "'-'", "'('", "')'", 
                     "'='", "';'", "'read'", "'if'", "'else'", "'write'", 
                     "'=='", "'!='", "'>'", "'>='", "'<'", "'<='", "'{'", 
                     "'}'" ]

    symbolicNames = [ "<INVALID>", "<INVALID>", "<INVALID>", "<INVALID>", 
                      "<INVALID>", "<INVALID>", "<INVALID>", "<INVALID>", 
                      "<INVALID>", "<INVALID>", "<INVALID>", "<INVALID>", 
                      "<INVALID>", "<INVALID>", "<INVALID>", "<INVALID>", 
                      "<INVALID>", "<INVALID>", "<INVALID>", "<INVALID>", 
                      "<INVALID>", "WS", "INT", "VAR" ]

    RULE_calc0 = 0
    RULE_expr = 1
    RULE_calc1 = 2
    RULE_stmt = 3
    RULE_calc2 = 4
    RULE_cond = 5
    RULE_block = 6
    RULE_calc3 = 7

    ruleNames =  [ "calc0", "expr", "calc1", "stmt", "calc2", "cond", "block", 
                   "calc3" ]

    EOF = Token.EOF
    T__0=1
    T__1=2
    T__2=3
    T__3=4
    T__4=5
    T__5=6
    T__6=7
    T__7=8
    T__8=9
    T__9=10
    T__10=11
    T__11=12
    T__12=13
    T__13=14
    T__14=15
    T__15=16
    T__16=17
    T__17=18
    T__18=19
    T__19=20
    WS=21
    INT=22
    VAR=23

    def __init__(self, input:TokenStream, output:TextIO = sys.stdout):
        super().__init__(input, output)
        self.checkVersion("4.13.2")
        self._interp = ParserATNSimulator(self, self.atn, self.decisionsToDFA, self.sharedContextCache)
        self._predicates = None




    class Calc0Context(ParserRuleContext):
        __slots__ = 'parser'

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser

        def expr(self):
            return self.getTypedRuleContext(CalcPlusParser.ExprContext,0)


        def EOF(self):
            return self.getToken(CalcPlusParser.EOF, 0)

        def getRuleIndex(self):
            return CalcPlusParser.RULE_calc0

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterCalc0" ):
                listener.enterCalc0(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitCalc0" ):
                listener.exitCalc0(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitCalc0" ):
                return visitor.visitCalc0(self)
            else:
                return visitor.visitChildren(self)




    def calc0(self):

        localctx = CalcPlusParser.Calc0Context(self, self._ctx, self.state)
        self.enterRule(localctx, 0, self.RULE_calc0)
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 16
            self.expr(0)
            self.state = 17
            self.match(CalcPlusParser.EOF)
        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx


    class ExprContext(ParserRuleContext):
        __slots__ = 'parser'

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser


        def getRuleIndex(self):
            return CalcPlusParser.RULE_expr

     
        def copyFrom(self, ctx:ParserRuleContext):
            super().copyFrom(ctx)


    class MulDivContext(ExprContext):

        def __init__(self, parser, ctx:ParserRuleContext): # actually a CalcPlusParser.ExprContext
            super().__init__(parser)
            self.copyFrom(ctx)

        def expr(self, i:int=None):
            if i is None:
                return self.getTypedRuleContexts(CalcPlusParser.ExprContext)
            else:
                return self.getTypedRuleContext(CalcPlusParser.ExprContext,i)


        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterMulDiv" ):
                listener.enterMulDiv(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitMulDiv" ):
                listener.exitMulDiv(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitMulDiv" ):
                return visitor.visitMulDiv(self)
            else:
                return visitor.visitChildren(self)


    class AddSubContext(ExprContext):

        def __init__(self, parser, ctx:ParserRuleContext): # actually a CalcPlusParser.ExprContext
            super().__init__(parser)
            self.copyFrom(ctx)

        def expr(self, i:int=None):
            if i is None:
                return self.getTypedRuleContexts(CalcPlusParser.ExprContext)
            else:
                return self.getTypedRuleContext(CalcPlusParser.ExprContext,i)


        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterAddSub" ):
                listener.enterAddSub(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitAddSub" ):
                listener.exitAddSub(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitAddSub" ):
                return visitor.visitAddSub(self)
            else:
                return visitor.visitChildren(self)


    class VarContext(ExprContext):

        def __init__(self, parser, ctx:ParserRuleContext): # actually a CalcPlusParser.ExprContext
            super().__init__(parser)
            self.copyFrom(ctx)

        def VAR(self):
            return self.getToken(CalcPlusParser.VAR, 0)

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterVar" ):
                listener.enterVar(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitVar" ):
                listener.exitVar(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitVar" ):
                return visitor.visitVar(self)
            else:
                return visitor.visitChildren(self)


    class ParensContext(ExprContext):

        def __init__(self, parser, ctx:ParserRuleContext): # actually a CalcPlusParser.ExprContext
            super().__init__(parser)
            self.copyFrom(ctx)

        def expr(self):
            return self.getTypedRuleContext(CalcPlusParser.ExprContext,0)


        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterParens" ):
                listener.enterParens(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitParens" ):
                listener.exitParens(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitParens" ):
                return visitor.visitParens(self)
            else:
                return visitor.visitChildren(self)


    class IntContext(ExprContext):

        def __init__(self, parser, ctx:ParserRuleContext): # actually a CalcPlusParser.ExprContext
            super().__init__(parser)
            self.copyFrom(ctx)

        def INT(self):
            return self.getToken(CalcPlusParser.INT, 0)

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterInt" ):
                listener.enterInt(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitInt" ):
                listener.exitInt(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitInt" ):
                return visitor.visitInt(self)
            else:
                return visitor.visitChildren(self)



    def expr(self, _p:int=0):
        _parentctx = self._ctx
        _parentState = self.state
        localctx = CalcPlusParser.ExprContext(self, self._ctx, _parentState)
        _prevctx = localctx
        _startState = 2
        self.enterRecursionRule(localctx, 2, self.RULE_expr, _p)
        self._la = 0 # Token type
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 26
            self._errHandler.sync(self)
            token = self._input.LA(1)
            if token in [22]:
                localctx = CalcPlusParser.IntContext(self, localctx)
                self._ctx = localctx
                _prevctx = localctx

                self.state = 20
                self.match(CalcPlusParser.INT)
                pass
            elif token in [23]:
                localctx = CalcPlusParser.VarContext(self, localctx)
                self._ctx = localctx
                _prevctx = localctx
                self.state = 21
                self.match(CalcPlusParser.VAR)
                pass
            elif token in [5]:
                localctx = CalcPlusParser.ParensContext(self, localctx)
                self._ctx = localctx
                _prevctx = localctx
                self.state = 22
                self.match(CalcPlusParser.T__4)
                self.state = 23
                self.expr(0)
                self.state = 24
                self.match(CalcPlusParser.T__5)
                pass
            else:
                raise NoViableAltException(self)

            self._ctx.stop = self._input.LT(-1)
            self.state = 36
            self._errHandler.sync(self)
            _alt = self._interp.adaptivePredict(self._input,2,self._ctx)
            while _alt!=2 and _alt!=ATN.INVALID_ALT_NUMBER:
                if _alt==1:
                    if self._parseListeners is not None:
                        self.triggerExitRuleEvent()
                    _prevctx = localctx
                    self.state = 34
                    self._errHandler.sync(self)
                    la_ = self._interp.adaptivePredict(self._input,1,self._ctx)
                    if la_ == 1:
                        localctx = CalcPlusParser.MulDivContext(self, CalcPlusParser.ExprContext(self, _parentctx, _parentState))
                        self.pushNewRecursionContext(localctx, _startState, self.RULE_expr)
                        self.state = 28
                        if not self.precpred(self._ctx, 5):
                            from antlr4.error.Errors import FailedPredicateException
                            raise FailedPredicateException(self, "self.precpred(self._ctx, 5)")
                        self.state = 29
                        _la = self._input.LA(1)
                        if not(_la==1 or _la==2):
                            self._errHandler.recoverInline(self)
                        else:
                            self._errHandler.reportMatch(self)
                            self.consume()
                        self.state = 30
                        self.expr(6)
                        pass

                    elif la_ == 2:
                        localctx = CalcPlusParser.AddSubContext(self, CalcPlusParser.ExprContext(self, _parentctx, _parentState))
                        self.pushNewRecursionContext(localctx, _startState, self.RULE_expr)
                        self.state = 31
                        if not self.precpred(self._ctx, 4):
                            from antlr4.error.Errors import FailedPredicateException
                            raise FailedPredicateException(self, "self.precpred(self._ctx, 4)")
                        self.state = 32
                        _la = self._input.LA(1)
                        if not(_la==3 or _la==4):
                            self._errHandler.recoverInline(self)
                        else:
                            self._errHandler.reportMatch(self)
                            self.consume()
                        self.state = 33
                        self.expr(5)
                        pass

             
                self.state = 38
                self._errHandler.sync(self)
                _alt = self._interp.adaptivePredict(self._input,2,self._ctx)

        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.unrollRecursionContexts(_parentctx)
        return localctx


    class Calc1Context(ParserRuleContext):
        __slots__ = 'parser'

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser

        def EOF(self):
            return self.getToken(CalcPlusParser.EOF, 0)

        def stmt(self, i:int=None):
            if i is None:
                return self.getTypedRuleContexts(CalcPlusParser.StmtContext)
            else:
                return self.getTypedRuleContext(CalcPlusParser.StmtContext,i)


        def getRuleIndex(self):
            return CalcPlusParser.RULE_calc1

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterCalc1" ):
                listener.enterCalc1(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitCalc1" ):
                listener.exitCalc1(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitCalc1" ):
                return visitor.visitCalc1(self)
            else:
                return visitor.visitChildren(self)




    def calc1(self):

        localctx = CalcPlusParser.Calc1Context(self, self._ctx, self.state)
        self.enterRule(localctx, 4, self.RULE_calc1)
        self._la = 0 # Token type
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 40 
            self._errHandler.sync(self)
            _la = self._input.LA(1)
            while True:
                self.state = 39
                self.stmt()
                self.state = 42 
                self._errHandler.sync(self)
                _la = self._input.LA(1)
                if not ((((_la) & ~0x3f) == 0 and ((1 << _la) & 8393728) != 0)):
                    break

            self.state = 44
            self.match(CalcPlusParser.EOF)
        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx


    class StmtContext(ParserRuleContext):
        __slots__ = 'parser'

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser


        def getRuleIndex(self):
            return CalcPlusParser.RULE_stmt

     
        def copyFrom(self, ctx:ParserRuleContext):
            super().copyFrom(ctx)



    class WriteContext(StmtContext):

        def __init__(self, parser, ctx:ParserRuleContext): # actually a CalcPlusParser.StmtContext
            super().__init__(parser)
            self.copyFrom(ctx)

        def expr(self):
            return self.getTypedRuleContext(CalcPlusParser.ExprContext,0)


        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterWrite" ):
                listener.enterWrite(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitWrite" ):
                listener.exitWrite(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitWrite" ):
                return visitor.visitWrite(self)
            else:
                return visitor.visitChildren(self)


    class IfElseContext(StmtContext):

        def __init__(self, parser, ctx:ParserRuleContext): # actually a CalcPlusParser.StmtContext
            super().__init__(parser)
            self.thenBlock = None # BlockContext
            self.elseBlock = None # BlockContext
            self.copyFrom(ctx)

        def cond(self):
            return self.getTypedRuleContext(CalcPlusParser.CondContext,0)

        def block(self, i:int=None):
            if i is None:
                return self.getTypedRuleContexts(CalcPlusParser.BlockContext)
            else:
                return self.getTypedRuleContext(CalcPlusParser.BlockContext,i)


        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterIfElse" ):
                listener.enterIfElse(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitIfElse" ):
                listener.exitIfElse(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitIfElse" ):
                return visitor.visitIfElse(self)
            else:
                return visitor.visitChildren(self)


    class ReadAssignContext(StmtContext):

        def __init__(self, parser, ctx:ParserRuleContext): # actually a CalcPlusParser.StmtContext
            super().__init__(parser)
            self.copyFrom(ctx)

        def VAR(self):
            return self.getToken(CalcPlusParser.VAR, 0)

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterReadAssign" ):
                listener.enterReadAssign(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitReadAssign" ):
                listener.exitReadAssign(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitReadAssign" ):
                return visitor.visitReadAssign(self)
            else:
                return visitor.visitChildren(self)


    class ExprAssignContext(StmtContext):

        def __init__(self, parser, ctx:ParserRuleContext): # actually a CalcPlusParser.StmtContext
            super().__init__(parser)
            self.copyFrom(ctx)

        def VAR(self):
            return self.getToken(CalcPlusParser.VAR, 0)
        def expr(self):
            return self.getTypedRuleContext(CalcPlusParser.ExprContext,0)


        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterExprAssign" ):
                listener.enterExprAssign(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitExprAssign" ):
                listener.exitExprAssign(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitExprAssign" ):
                return visitor.visitExprAssign(self)
            else:
                return visitor.visitChildren(self)



    def stmt(self):

        localctx = CalcPlusParser.StmtContext(self, self._ctx, self.state)
        self.enterRule(localctx, 6, self.RULE_stmt)
        self._la = 0 # Token type
        try:
            self.state = 72
            self._errHandler.sync(self)
            la_ = self._interp.adaptivePredict(self._input,5,self._ctx)
            if la_ == 1:
                localctx = CalcPlusParser.ExprAssignContext(self, localctx)
                self.enterOuterAlt(localctx, 1)
                self.state = 46
                self.match(CalcPlusParser.VAR)
                self.state = 47
                self.match(CalcPlusParser.T__6)
                self.state = 48
                self.expr(0)
                self.state = 49
                self.match(CalcPlusParser.T__7)
                pass

            elif la_ == 2:
                localctx = CalcPlusParser.ReadAssignContext(self, localctx)
                self.enterOuterAlt(localctx, 2)
                self.state = 51
                self.match(CalcPlusParser.VAR)
                self.state = 52
                self.match(CalcPlusParser.T__6)
                self.state = 53
                self.match(CalcPlusParser.T__8)
                self.state = 54
                self.match(CalcPlusParser.T__4)
                self.state = 55
                self.match(CalcPlusParser.T__5)
                self.state = 56
                self.match(CalcPlusParser.T__7)
                pass

            elif la_ == 3:
                localctx = CalcPlusParser.IfElseContext(self, localctx)
                self.enterOuterAlt(localctx, 3)
                self.state = 57
                self.match(CalcPlusParser.T__9)
                self.state = 58
                self.match(CalcPlusParser.T__4)
                self.state = 59
                self.cond()
                self.state = 60
                self.match(CalcPlusParser.T__5)
                self.state = 61
                localctx.thenBlock = self.block()
                self.state = 64
                self._errHandler.sync(self)
                _la = self._input.LA(1)
                if _la==11:
                    self.state = 62
                    self.match(CalcPlusParser.T__10)
                    self.state = 63
                    localctx.elseBlock = self.block()


                pass

            elif la_ == 4:
                localctx = CalcPlusParser.WriteContext(self, localctx)
                self.enterOuterAlt(localctx, 4)
                self.state = 66
                self.match(CalcPlusParser.T__11)
                self.state = 67
                self.match(CalcPlusParser.T__4)
                self.state = 68
                self.expr(0)
                self.state = 69
                self.match(CalcPlusParser.T__5)
                self.state = 70
                self.match(CalcPlusParser.T__7)
                pass


        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx


    class Calc2Context(ParserRuleContext):
        __slots__ = 'parser'

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser

        def EOF(self):
            return self.getToken(CalcPlusParser.EOF, 0)

        def stmt(self, i:int=None):
            if i is None:
                return self.getTypedRuleContexts(CalcPlusParser.StmtContext)
            else:
                return self.getTypedRuleContext(CalcPlusParser.StmtContext,i)


        def getRuleIndex(self):
            return CalcPlusParser.RULE_calc2

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterCalc2" ):
                listener.enterCalc2(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitCalc2" ):
                listener.exitCalc2(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitCalc2" ):
                return visitor.visitCalc2(self)
            else:
                return visitor.visitChildren(self)




    def calc2(self):

        localctx = CalcPlusParser.Calc2Context(self, self._ctx, self.state)
        self.enterRule(localctx, 8, self.RULE_calc2)
        self._la = 0 # Token type
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 75 
            self._errHandler.sync(self)
            _la = self._input.LA(1)
            while True:
                self.state = 74
                self.stmt()
                self.state = 77 
                self._errHandler.sync(self)
                _la = self._input.LA(1)
                if not ((((_la) & ~0x3f) == 0 and ((1 << _la) & 8393728) != 0)):
                    break

            self.state = 79
            self.match(CalcPlusParser.EOF)
        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx


    class CondContext(ParserRuleContext):
        __slots__ = 'parser'

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser

        def expr(self, i:int=None):
            if i is None:
                return self.getTypedRuleContexts(CalcPlusParser.ExprContext)
            else:
                return self.getTypedRuleContext(CalcPlusParser.ExprContext,i)


        def getRuleIndex(self):
            return CalcPlusParser.RULE_cond

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterCond" ):
                listener.enterCond(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitCond" ):
                listener.exitCond(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitCond" ):
                return visitor.visitCond(self)
            else:
                return visitor.visitChildren(self)




    def cond(self):

        localctx = CalcPlusParser.CondContext(self, self._ctx, self.state)
        self.enterRule(localctx, 10, self.RULE_cond)
        self._la = 0 # Token type
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 81
            self.expr(0)
            self.state = 82
            _la = self._input.LA(1)
            if not((((_la) & ~0x3f) == 0 and ((1 << _la) & 516096) != 0)):
                self._errHandler.recoverInline(self)
            else:
                self._errHandler.reportMatch(self)
                self.consume()
            self.state = 83
            self.expr(0)
        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx


    class BlockContext(ParserRuleContext):
        __slots__ = 'parser'

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser

        def stmt(self, i:int=None):
            if i is None:
                return self.getTypedRuleContexts(CalcPlusParser.StmtContext)
            else:
                return self.getTypedRuleContext(CalcPlusParser.StmtContext,i)


        def getRuleIndex(self):
            return CalcPlusParser.RULE_block

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterBlock" ):
                listener.enterBlock(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitBlock" ):
                listener.exitBlock(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitBlock" ):
                return visitor.visitBlock(self)
            else:
                return visitor.visitChildren(self)




    def block(self):

        localctx = CalcPlusParser.BlockContext(self, self._ctx, self.state)
        self.enterRule(localctx, 12, self.RULE_block)
        self._la = 0 # Token type
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 85
            self.match(CalcPlusParser.T__18)
            self.state = 89
            self._errHandler.sync(self)
            _la = self._input.LA(1)
            while (((_la) & ~0x3f) == 0 and ((1 << _la) & 8393728) != 0):
                self.state = 86
                self.stmt()
                self.state = 91
                self._errHandler.sync(self)
                _la = self._input.LA(1)

            self.state = 92
            self.match(CalcPlusParser.T__19)
        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx


    class Calc3Context(ParserRuleContext):
        __slots__ = 'parser'

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser

        def EOF(self):
            return self.getToken(CalcPlusParser.EOF, 0)

        def stmt(self, i:int=None):
            if i is None:
                return self.getTypedRuleContexts(CalcPlusParser.StmtContext)
            else:
                return self.getTypedRuleContext(CalcPlusParser.StmtContext,i)


        def getRuleIndex(self):
            return CalcPlusParser.RULE_calc3

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterCalc3" ):
                listener.enterCalc3(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitCalc3" ):
                listener.exitCalc3(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitCalc3" ):
                return visitor.visitCalc3(self)
            else:
                return visitor.visitChildren(self)




    def calc3(self):

        localctx = CalcPlusParser.Calc3Context(self, self._ctx, self.state)
        self.enterRule(localctx, 14, self.RULE_calc3)
        self._la = 0 # Token type
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 95 
            self._errHandler.sync(self)
            _la = self._input.LA(1)
            while True:
                self.state = 94
                self.stmt()
                self.state = 97 
                self._errHandler.sync(self)
                _la = self._input.LA(1)
                if not ((((_la) & ~0x3f) == 0 and ((1 << _la) & 8393728) != 0)):
                    break

            self.state = 99
            self.match(CalcPlusParser.EOF)
        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx



    def sempred(self, localctx:RuleContext, ruleIndex:int, predIndex:int):
        if self._predicates == None:
            self._predicates = dict()
        self._predicates[1] = self.expr_sempred
        pred = self._predicates.get(ruleIndex, None)
        if pred is None:
            raise Exception("No predicate with index:" + str(ruleIndex))
        else:
            return pred(localctx, predIndex)

    def expr_sempred(self, localctx:ExprContext, predIndex:int):
            if predIndex == 0:
                return self.precpred(self._ctx, 5)
         

            if predIndex == 1:
                return self.precpred(self._ctx, 4)
         




