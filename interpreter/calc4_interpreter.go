package interpreter

import (
	"calcPlus/internal/parser"
	"calcPlus/internal/symbols"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"io"
	"strconv"
)

type Calc4Interpreter struct {
	parser.BaseCalcPlusVisitor
	SymbolTable symbols.ScopedTable[int]
	Reader      io.Reader
	Writer      io.Writer
}

func NewCalc4Interpreter(reader io.Reader, writer io.Writer) *Calc4Interpreter {
	return &Calc4Interpreter{
		SymbolTable: *symbols.NewScopedTable[int](),
		Reader:      reader,
		Writer:      writer,
	}
}

func (c *Calc4Interpreter) Visit(tree antlr.ParseTree) any {
	return tree.Accept(c)
}

func (c *Calc4Interpreter) VisitProgram(ctx *parser.ProgramContext) any {
	for _, stmt := range ctx.AllStmt() {
		c.Visit(stmt)
	}
	return nil
}

func (c *Calc4Interpreter) VisitExprAssign(ctx *parser.ExprAssignContext) any {
	result := c.Visit(ctx.Expr()).(int)
	err := c.SymbolTable.SetSymbol(ctx.VAR().GetText(), result)
	if err != nil {
		panic(err)
	}

	return nil
}

func (c *Calc4Interpreter) VisitReadAssign(ctx *parser.ReadAssignContext) any {
	buffer := ""
	_, _ = fmt.Fscanln(c.Reader, &buffer)

	varName := ctx.VAR().GetText()
	value, _ := strconv.Atoi(buffer)

	err := c.SymbolTable.SetSymbol(varName, value)
	if err != nil {
		panic(err)
	}

	return nil
}

func (c *Calc4Interpreter) VisitIfElse(ctx *parser.IfElseContext) any {
	if c.Visit(ctx.Cond()).(bool) {
		c.Visit(ctx.GetThenBlock())
	} else if ctx.GetElseBlock() != nil {
		c.Visit(ctx.GetElseBlock())
	}

	return nil
}

func (c *Calc4Interpreter) VisitWrite(ctx *parser.WriteContext) any {
	_, _ = fmt.Fprintln(c.Writer, c.Visit(ctx.Expr()).(int))
	return nil
}

func (c *Calc4Interpreter) VisitMulDiv(ctx *parser.MulDivContext) any {
	left := c.Visit(ctx.Expr(0)).(int)
	right := c.Visit(ctx.Expr(1)).(int)

	op := ctx.GetChild(1).(antlr.TerminalNode).GetText()
	if op == "*" {
		return left * right
	} else {
		return left / right
	}
}

func (c *Calc4Interpreter) VisitAddSub(ctx *parser.AddSubContext) any {
	left := c.Visit(ctx.Expr(0)).(int)
	right := c.Visit(ctx.Expr(1)).(int)

	op := ctx.GetChild(1).(antlr.TerminalNode).GetText()

	if op == "+" {
		return left + right
	} else {
		return left - right
	}
}

func (c *Calc4Interpreter) VisitInt(ctx *parser.IntContext) any {
	val, _ := strconv.Atoi(ctx.GetText())
	return val
}

func (c *Calc4Interpreter) VisitVar(ctx *parser.VarContext) any {
	value, err := c.SymbolTable.GetSymbol(ctx.GetText())
	if err != nil {
		panic(err)
	}
	return value
}

func (c *Calc4Interpreter) VisitParens(ctx *parser.ParensContext) any {
	return c.Visit(ctx.Expr()).(int)
}

func (c *Calc4Interpreter) VisitCond(ctx *parser.CondContext) any {
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

func (c *Calc4Interpreter) VisitBlock(ctx *parser.BlockContext) any {
	c.SymbolTable.PushScope()
	for _, stmt := range ctx.AllStmt() {
		c.Visit(stmt)
	}
	c.SymbolTable.PopScope()
	return nil
}

func (c *Calc4Interpreter) VisitDeclare(ctx *parser.DeclareContext) any {
	for _, v := range ctx.AllVAR() {
		err := c.SymbolTable.Register(v.GetText())
		if err != nil {
			panic(err)
		}
	}
	return nil
}

func (c *Calc4Interpreter) VisitStmtBlock(ctx *parser.StmtBlockContext) any {
	return c.Visit(ctx.Block())
}
