import sys
from CalcPlusVisitor import CalcPlusVisitor
from CalcPlusParser import CalcPlusParser


class Calculator(CalcPlusVisitor):
    def __init__(self):
        self.memory = {}

    def visitCalc3(self, ctx: CalcPlusParser.Calc3Context):
        for stmt in ctx.stmt():
            self.visit(stmt)
        return None

    def visitExprAssign(self, ctx: CalcPlusParser.ExprAssignContext):
        var_name = ctx.VAR().getText()
        value = self.visit(ctx.expr())
        self.memory[var_name] = value
        return value

    def visitReadAssign(self, ctx: CalcPlusParser.ReadAssignContext):
        var_name = ctx.VAR().getText()
        try:
            user_input = input()
            value = int(user_input)
        except (ValueError, EOFError):
            value = 0

        self.memory[var_name] = value
        return value

    def visitWrite(self, ctx: CalcPlusParser.WriteContext):
        value = self.visit(ctx.expr())
        print(value)
        return None

    def visitIfElse(self, ctx: CalcPlusParser.IfElseContext):
        cond = self.visit(ctx.cond())
        if cond:
            self.visit(ctx.thenBlock)
        elif ctx.elseBlock:
            self.visit(ctx.elseBlock)
        return None

    def visitMulDiv(self, ctx: CalcPlusParser.MulDivContext):
        left = self.visit(ctx.expr(0))
        right = self.visit(ctx.expr(1))
        op = ctx.getChild(1).getText()
        return left * right if op == '*' else int(left / right)

    def visitAddSub(self, ctx: CalcPlusParser.AddSubContext):
        left = self.visit(ctx.expr(0))
        right = self.visit(ctx.expr(1))
        op = ctx.getChild(1).getText()
        return left + right if op == '+' else left - right

    def visitInt(self, ctx: CalcPlusParser.IntContext):
        return int(ctx.INT().getText())

    def visitVar(self, ctx: CalcPlusParser.VarContext):
        var_name = ctx.VAR().getText()
        return self.memory.get(var_name, 0)

    def visitParens(self, ctx: CalcPlusParser.ParensContext):
        return self.visit(ctx.expr())

    def visitCond(self, ctx: CalcPlusParser.CondContext):
        left = self.visit(ctx.getChild(0))
        right = self.visit(ctx.getChild(2))
        op = ctx.getChild(1).getText()

        if op == '==':
            return left == right
        elif op == '!=':
            return left != right
        elif op == '<':
            return left < right
        elif op == '<=':
            return left <= right
        elif op == '>':
            return left > right
        elif op == '>=':
            return left >= right
        return False