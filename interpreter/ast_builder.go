package interpreter

import (
	"calcPlus/internal/ast"
	"calcPlus/internal/parser"
	"calcPlus/internal/symbols"
	"fmt"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
)

type ASTBuilder struct {
	parser.BaseCalcPlusVisitor

	program   []ast.Statement
	variables *symbols.ScopedTable[int]
}

func NewASTBuilder() *ASTBuilder {
	return &ASTBuilder{
		variables: symbols.NewScopedTable[int](),
	}
}

func (b *ASTBuilder) Build(tree antlr.ParseTree) []ast.Statement {
	b.program = nil
	b.variables = symbols.NewScopedTable[int]()
	b.Visit(tree)
	return b.program
}

func (b *ASTBuilder) Visit(tree antlr.ParseTree) any {
	return tree.Accept(b)
}

func (b *ASTBuilder) VisitProgram(ctx *parser.ProgramContext) any {
	for _, funcDef := range ctx.AllFuncDef() {
		if funcDef.IDENT().GetText() != "main" {
			continue
		}

		b.program = append(b.program, b.Visit(funcDef.Block()).([]ast.Statement)...)
		return nil
	}

	panic("main function not found")
}

func (b *ASTBuilder) VisitDeclare(ctx *parser.DeclareContext) any {
	var declarations []ast.Statement
	for _, ident := range ctx.AllIDENT() {
		name := ident.GetText()
		if err := b.variables.Register(name); err != nil {
			panic(err)
		}

		declarations = append(declarations, ast.NewDeclaration(name))
	}
	return declarations
}

func (b *ASTBuilder) VisitExprAssign(ctx *parser.ExprAssignContext) any {
	varName := ctx.IDENT().GetText()
	if _, err := b.variables.GetSymbol(varName); err != nil {
		panic(err)
	}

	expr := b.Visit(ctx.Expr()).(ast.Expression)
	return []ast.Statement{ast.NewAssignment(varName, expr)}
}

func (b *ASTBuilder) VisitExprStmt(ctx *parser.ExprStmtContext) any {
	expr := b.Visit(ctx.Expr()).(ast.Expression)
	return []ast.Statement{ast.NewAssignment("_", expr)}
}

func (b *ASTBuilder) VisitIfElse(ctx *parser.IfElseContext) any {
	cond := b.Visit(ctx.Cond()).(ast.Expression)
	thenBlock := b.Visit(ctx.GetThenBlock()).([]ast.Statement)
	var elseBlock []ast.Statement
	if ctx.GetElseBlock() != nil {
		elseBlock = b.Visit(ctx.GetElseBlock()).([]ast.Statement)
	}

	return []ast.Statement{ast.NewIfElse(
		cond,
		ast.NewBlockStatements(thenBlock...),
		ast.NewBlockStatements(elseBlock...),
	)}
}

func (b *ASTBuilder) VisitStmtBlock(ctx *parser.StmtBlockContext) any {
	statements := b.Visit(ctx.Block()).([]ast.Statement)
	return []ast.Statement{ast.NewBlockStatements(statements...)}
}

func (b *ASTBuilder) VisitReturn(ctx *parser.ReturnContext) any {
	panic("return is not supported by the AST builder")
}

func (b *ASTBuilder) VisitBlock(ctx *parser.BlockContext) any {
	b.variables.PushScope()
	var astStatements []ast.Statement
	for _, stmt := range ctx.AllStmt() {
		blockStatements := b.Visit(stmt).([]ast.Statement)
		astStatements = append(astStatements, blockStatements...)
	}
	b.variables.PopScope()

	return astStatements
}

func (b *ASTBuilder) VisitMulDiv(ctx *parser.MulDivContext) any {
	left := b.Visit(ctx.Expr(0)).(ast.Expression)
	right := b.Visit(ctx.Expr(1)).(ast.Expression)
	op := ctx.GetChild(1).(antlr.TerminalNode).GetText()

	return ast.NewBinaryOperator(left, op, right)
}

func (b *ASTBuilder) VisitAddSub(ctx *parser.AddSubContext) any {
	left := b.Visit(ctx.Expr(0)).(ast.Expression)
	right := b.Visit(ctx.Expr(1)).(ast.Expression)
	op := ctx.GetChild(1).(antlr.TerminalNode).GetText()

	return ast.NewBinaryOperator(left, op, right)
}

func (b *ASTBuilder) VisitFuncCall(ctx *parser.FuncCallContext) any {
	name := ctx.IDENT().GetText()
	switch name {
	case "read":
		if ctx.ArgList() != nil {
			panic("read does not accept arguments")
		}
		return ast.NewBuiltinReadCall()
	case "write":
		if ctx.ArgList() == nil || len(ctx.ArgList().AllExpr()) != 1 {
			panic("write requires exactly one argument")
		}
		return ast.NewBuiltinWriteCall(b.Visit(ctx.ArgList().Expr(0)).(ast.Expression))
	default:
		panic(fmt.Sprintf("function call is not supported by the AST builder: %s", name))
	}
}

func (b *ASTBuilder) VisitVar(ctx *parser.VarContext) any {
	if _, err := b.variables.GetSymbol(ctx.GetText()); err != nil {
		panic(err)
	}

	return ast.NewVarExpression(ctx.GetText())
}

func (b *ASTBuilder) VisitParens(ctx *parser.ParensContext) any {
	return b.Visit(ctx.Expr())
}

func (b *ASTBuilder) VisitInt(ctx *parser.IntContext) any {
	value, err := strconv.Atoi(ctx.GetText())
	if err != nil {
		panic(err)
	}
	return ast.NewIntExpression(value)
}

func (b *ASTBuilder) VisitCond(ctx *parser.CondContext) any {
	left := b.Visit(ctx.Expr(0)).(ast.Expression)
	right := b.Visit(ctx.Expr(1)).(ast.Expression)
	op := ctx.GetChild(1).(antlr.TerminalNode).GetText()

	return ast.NewBinaryOperator(left, op, right)
}
