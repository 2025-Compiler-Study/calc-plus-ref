import { AbstractParseTreeVisitor } from 'antlr4ts/tree/AbstractParseTreeVisitor';
import { CalcPlusVisitor } from './CalcPlusVisitor';
import {
  Calc3Context, ExprAssignContext, ReadAssignContext, WriteContext,
  IfElseContext, MulDivContext, AddSubContext, IntContext, VarContext,
  ParensContext, CondContext
} from './CalcPlusParser';

export class Calculator extends AbstractParseTreeVisitor<number> implements CalcPlusVisitor<number> {

    private memory: Map<string, number> = new Map();

    constructor(
        private inputProvider: () => string,
        private outputProvider: (msg: string) => void
    ) {
        super();
    }

    protected defaultResult(): number {
        return 0;
    }

    visitCalc3(ctx: Calc3Context): number {
        // 모든 stmt 순차 실행
        ctx.stmt().forEach(stmt => this.visit(stmt));
        return 0;
    }

    visitExprAssign(ctx: ExprAssignContext): number {
        const name = ctx.VAR().text;
        const value = this.visit(ctx.expr());
        this.memory.set(name, value);
        return value;
    }

    visitReadAssign(ctx: ReadAssignContext): number {
        const name = ctx.VAR().text;
        try {
            const inputStr = this.inputProvider();
            let value = parseInt(inputStr, 10);

            if (Number.isNaN(value)) {
                value = 0;
            }
            this.memory.set(name, value);
            return value;
        } catch (e) {
            this.memory.set(name, 0);
            return 0;
        }
    }

    visitWrite(ctx: WriteContext): number {
        const value = this.visit(ctx.expr());
        this.outputProvider(value.toString());
        return value;
    }

    visitIfElse(ctx: IfElseContext): number {
        const cond = this.visit(ctx.cond()) !== 0;
        if (cond) {
            this.visit(ctx._thenBlock);
        } else if (ctx._elseBlock) {
            this.visit(ctx._elseBlock);
        }
        return 0;
    }

    visitMulDiv(ctx: MulDivContext): number {
        const left = this.visit(ctx.expr(0));
        const right = this.visit(ctx.expr(1));
        const op = ctx.getChild(1).text;
        return op === '*' ? left * right : Math.floor(left / right);
    }

    visitAddSub(ctx: AddSubContext): number {
        const left = this.visit(ctx.expr(0));
        const right = this.visit(ctx.expr(1));
        const op = ctx.getChild(1).text;
        return op === '+' ? left + right : left - right;
    }

    visitInt(ctx: IntContext): number {
        return parseInt(ctx.text, 10);
    }

    visitVar(ctx: VarContext): number {
        const name = ctx.VAR().text;
        return this.memory.get(name) || 0;
    }

    visitParens(ctx: ParensContext): number {
        return this.visit(ctx.expr());
    }

    visitCond(ctx: CondContext): number {
        const left = this.visit(ctx.getChild(0));
        const right = this.visit(ctx.getChild(2));
        const op = ctx.getChild(1).text;

        let result = false;
        switch(op) {
            case '==': result = left === right; break;
            case '!=': result = left !== right; break;
            case '<':  result = left < right; break;
            case '<=': result = left <= right; break;
            case '>':  result = left > right; break;
            case '>=': result = left >= right; break;
        }
        return result ? 1 : 0;
    }
}