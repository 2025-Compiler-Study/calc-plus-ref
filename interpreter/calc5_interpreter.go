package interpreter

import (
	"calcPlus/internal/ast"
	"calcPlus/internal/executor"
	"calcPlus/internal/parser"
	"calcPlus/internal/symbols"
	"io"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
)

type Calc5Interpreter struct {
	parser.BaseCalcPlusVisitor
	program []ast.Statement  // parse tree -> ast 결과물
	Engine  *executor.Engine // ast 결과물을 실행하기 위한 환경

	errors    []error
	Variables *symbols.ScopedTable[int]
	// parse tree 순회중 변수 사용 조건 위배 확인용 (의미검사용)
}

func NewCalc5Interpreter(reader io.Reader, writer io.Writer) *Calc5Interpreter {
	return &Calc5Interpreter{
		Engine:    executor.NewEngine(reader, writer),
		Variables: symbols.NewScopedTable[int](),
	}
}

func (c *Calc5Interpreter) Visit(tree antlr.ParseTree) any {
	result := tree.Accept(c)

	if _, ok := tree.(*parser.ProgramContext); !ok {
		return result
	}

	if len(c.errors) > 0 {
		panic(c.errors)
	}

	for _, stmt := range c.program {
		err := c.Engine.Execute(stmt)

		if err != nil {
			panic(err)
		}
	}
	return result
}

func (c *Calc5Interpreter) VisitProgram(ctx *parser.ProgramContext) any {
	for _, stmt := range ctx.AllStmt() {
		var astStatements []ast.Statement
		astStatements = c.Visit(stmt).([]ast.Statement)
		c.program = append(c.program, astStatements...)
	}
	return nil
}

func (c *Calc5Interpreter) VisitDeclare(ctx *parser.DeclareContext) any {
	var declarations []ast.Statement
	for _, v := range ctx.AllVAR() {
		// Check variable declaration
		err := c.Variables.Register(v.GetText())
		if err != nil {
			panic(err)
		}

		declarations = append(declarations, ast.NewDeclaration(v.GetText()))
	}
	return declarations
}

func (c *Calc5Interpreter) VisitExprAssign(ctx *parser.ExprAssignContext) any {
	// Check write target variable exists
	varName := ctx.VAR().GetText()
	_, err := c.Variables.GetSymbol(varName)
	if err != nil {
		panic(err)
	}

	// Evaluate expression
	expr := c.Visit(ctx.Expr()).(ast.Expression)
	assignment := ast.NewAssignment(varName, expr)
	return []ast.Statement{assignment}
}

func (c *Calc5Interpreter) VisitReadAssign(ctx *parser.ReadAssignContext) any {
	varName := ctx.VAR().GetText()
	_, err := c.Variables.GetSymbol(varName)
	if err != nil {
		panic(err)
	}

	// Builtin read call is expression
	return []ast.Statement{ast.NewAssignment(varName, ast.NewBuiltinReadCall())}
}

func (c *Calc5Interpreter) VisitWrite(ctx *parser.WriteContext) any {
	// Evaluate write arguments
	expr := c.Visit(ctx.Expr()).(ast.Expression)

	// Builtin write call is expression
	writeCall := ast.NewBuiltinWriteCall(expr)

	// Wrap write call as a sink assignment
	assignment := ast.NewAssignment("_", writeCall)

	return []ast.Statement{assignment}
}

func (c *Calc5Interpreter) VisitIfElse(ctx *parser.IfElseContext) any {
	// Evaluate condition
	cond := c.Visit(ctx.Cond()).(ast.Expression)
	// Evaluate then and else blocks
	thenBlock := c.Visit(ctx.GetThenBlock()).([]ast.Statement)
	var elseBlock []ast.Statement
	if ctx.GetElseBlock() != nil {
		elseBlock = c.Visit(ctx.GetElseBlock()).([]ast.Statement)
	}

	ifElse := ast.NewIfElse(
		cond,
		ast.NewBlockStatements(thenBlock...),
		ast.NewBlockStatements(elseBlock...))
	return []ast.Statement{ifElse}
}

func (c *Calc5Interpreter) VisitStmtBlock(ctx *parser.StmtBlockContext) any {
	statements := c.Visit(ctx.Block()).([]ast.Statement)
	return []ast.Statement{ast.NewBlockStatements(statements...)}
}

func (c *Calc5Interpreter) VisitBlock(ctx *parser.BlockContext) any {
	c.Variables.PushScope()
	var astStatements []ast.Statement
	for _, stmt := range ctx.AllStmt() {
		blockStatements := c.Visit(stmt).([]ast.Statement)
		astStatements = append(astStatements, blockStatements...)
	}
	c.Variables.PopScope()

	return astStatements
}

func (c *Calc5Interpreter) VisitMulDiv(ctx *parser.MulDivContext) any {
	left := c.Visit(ctx.Expr(0)).(ast.Expression)
	right := c.Visit(ctx.Expr(1)).(ast.Expression)

	op := ctx.GetChild(1).(antlr.TerminalNode).GetText()

	return ast.NewBinaryOperator(left, op, right)
}

func (c *Calc5Interpreter) VisitAddSub(ctx *parser.AddSubContext) any {
	left := c.Visit(ctx.Expr(0)).(ast.Expression)
	right := c.Visit(ctx.Expr(1)).(ast.Expression)

	op := ctx.GetChild(1).(antlr.TerminalNode).GetText()

	return ast.NewBinaryOperator(left, op, right)
}

func (c *Calc5Interpreter) VisitVar(ctx *parser.VarContext) any {
	// Check variable exists
	_, err := c.Variables.GetSymbol(ctx.GetText())
	if err != nil {
		panic(err)
	}

	return ast.NewVarExpression(ctx.GetText())
}

func (c *Calc5Interpreter) VisitParens(ctx *parser.ParensContext) any {
	return c.Visit(ctx.Expr())
}

func (c *Calc5Interpreter) VisitInt(ctx *parser.IntContext) any {
	value, err := strconv.Atoi(ctx.GetText())
	if err != nil {
		panic(err)
	}
	return ast.NewIntExpression(value)
}

func (c *Calc5Interpreter) VisitCond(ctx *parser.CondContext) any {
	left := c.Visit(ctx.Expr(0)).(ast.Expression)
	right := c.Visit(ctx.Expr(1)).(ast.Expression)

	op := ctx.GetChild(1).(antlr.TerminalNode).GetText()

	return ast.NewBinaryOperator(left, op, right)
}
