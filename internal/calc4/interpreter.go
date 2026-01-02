package calc4

import (
	"calcPlus/internal/parser"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"io"
	"strconv"
)

type Interpreter struct {
	parser.BaseCalcPlusVisitor
	SymbolTable ScopeSymbolTable
	Reader      io.Reader
	Writer      io.Writer
}

func NewInterpreter(reader io.Reader, writer io.Writer) *Interpreter {
	return &Interpreter{
		SymbolTable: *NewScopeSymbolTable(),
		Reader:      reader,
		Writer:      writer,
	}
}

func (c *Interpreter) Visit(tree antlr.ParseTree) any {
	return tree.Accept(c)
}

func (c *Interpreter) VisitCalc4(ctx *parser.Calc4Context) any {
	for _, stmt := range ctx.AllStmt() {
		c.Visit(stmt)
	}
	return nil
}

func (c *Interpreter) VisitExprAssign(ctx *parser.ExprAssignContext) any {
	result := c.Visit(ctx.Expr()).(int)
	err := c.SymbolTable.SetVariable(ctx.VAR().GetText(), result)
	if err != nil {
		panic(err)
	}

	return nil
}

func (c *Interpreter) VisitReadAssign(ctx *parser.ReadAssignContext) any {
	buffer := ""
	_, _ = fmt.Fscanln(c.Reader, &buffer)

	varName := ctx.VAR().GetText()
	value, _ := strconv.Atoi(buffer)

	err := c.SymbolTable.SetVariable(varName, value)
	if err != nil {
		panic(err)
	}

	return nil
}

func (c *Interpreter) VisitIfElse(ctx *parser.IfElseContext) any {
	if c.Visit(ctx.Cond()).(bool) {
		c.Visit(ctx.GetThenBlock())
	} else if ctx.GetElseBlock() != nil {
		c.Visit(ctx.GetElseBlock())
	}

	return nil
}

func (c *Interpreter) VisitWrite(ctx *parser.WriteContext) any {
	_, _ = fmt.Fprintln(c.Writer, c.Visit(ctx.Expr()).(int))
	return nil
}

func (c *Interpreter) VisitMulDiv(ctx *parser.MulDivContext) any {
	left := c.Visit(ctx.Expr(0)).(int)
	right := c.Visit(ctx.Expr(1)).(int)

	op := ctx.GetChild(1).(antlr.TerminalNode).GetText()
	if op == "*" {
		return left * right
	} else {
		return left / right
	}
}

func (c *Interpreter) VisitAddSub(ctx *parser.AddSubContext) any {
	left := c.Visit(ctx.Expr(0)).(int)
	right := c.Visit(ctx.Expr(1)).(int)

	op := ctx.GetChild(1).(antlr.TerminalNode).GetText()

	if op == "+" {
		return left + right
	} else {
		return left - right
	}
}

func (c *Interpreter) VisitInt(ctx *parser.IntContext) any {
	val, _ := strconv.Atoi(ctx.GetText())
	return val
}

func (c *Interpreter) VisitVar(ctx *parser.VarContext) any {
	value, err := c.SymbolTable.GetVariable(ctx.GetText())
	if err != nil {
		panic(err)
	}
	return value
}

func (c *Interpreter) VisitParens(ctx *parser.ParensContext) any {
	return c.Visit(ctx.Expr()).(int)
}

func (c *Interpreter) VisitCond(ctx *parser.CondContext) any {
	left := c.Visit(ctx.Expr(0)).(int)
	right := c.Visit(ctx.Expr(1)).(int)

	op := ctx.GetChild(1).(antlr.TerminalNode).GetText()
	switch op {
	case "==":
		return left == right
	case "!=":
		return left != right
	case ">":
		return left > right
	case ">=":
		return left >= right
	case "<":
		return left < right
	case "<=":
		return left <= right
	}
	return false
}

func (c *Interpreter) VisitBlock(ctx *parser.BlockContext) any {
	c.SymbolTable.PushScope()
	for _, stmt := range ctx.AllStmt() {
		c.Visit(stmt)
	}
	c.SymbolTable.PopScope()
	return nil
}

func (c *Interpreter) VisitDeclare(ctx *parser.DeclareContext) any {
	for _, v := range ctx.AllVAR() {
		err := c.SymbolTable.Register(v.GetText())
		if err != nil {
			panic(err)
		}
	}
	return nil
}

func (c *Interpreter) VisitStmtBlock(ctx *parser.StmtBlockContext) any {
	return c.Visit(ctx.Block())
}
