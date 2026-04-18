package executor

import (
	"calcPlus/internal/ast"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecuteDeclaration(t *testing.T) {
	engine := NewEngine(nil, nil)
	ConfigurePreset(engine)

	t.Run("declaring undeclared name", func(t *testing.T) {
		declaration := ast.NewDeclaration("new")
		err := engine.Execute(declaration)
		assert.NoError(t, err)
	})

	t.Run("declaring already declared", func(t *testing.T) {
		declaration := ast.NewDeclaration("a")
		err := engine.Execute(declaration)
		assert.Error(t, err)
	})
}

func TestExecuteAssignment(t *testing.T) {
	engine := NewEngine(nil, nil)
	ConfigurePreset(engine)

	t.Run("normal case", func(t *testing.T) {
		assignment := ast.NewAssignment(
			"a",
			ast.NewIntExpression(1),
		)
		err := engine.Execute(assignment)
		assert.NoError(t, err)
	})

	t.Run("assign to undeclared variable", func(t *testing.T) {
		assignment := ast.NewAssignment(
			"undeclared",
			ast.NewIntExpression(1),
		)
		err := engine.Execute(assignment)
		assert.Error(t, err)
	})

	t.Run("expression uses undeclared variable", func(t *testing.T) {
		assignment := ast.NewAssignment(
			"a",
			ast.NewBinaryOperator(
				ast.NewVarExpression("undeclared"),
				"+",
				ast.NewIntExpression(1),
			),
		)
		err := engine.Execute(assignment)
		assert.Error(t, err)
	})
}

func TestExecuteBlockStatement(t *testing.T) {
	engine := NewEngine(nil, nil)
	ConfigurePreset(engine)

	t.Run("multiple statements", func(t *testing.T) {
		block := ast.NewBlockStatements(
			ast.NewDeclaration("value"),
			ast.NewAssignment("value", ast.NewIntExpression(1)),
		)
		err := engine.Execute(block)
		assert.NoError(t, err)
	})

	t.Run("empty block", func(t *testing.T) {
		block := ast.NewBlockStatements()
		err := engine.Execute(block)
		assert.NoError(t, err)
	})

	t.Run("use inner scope variable from outer scope error", func(t *testing.T) {
		block := ast.NewBlockStatements(
			ast.NewBlockStatements(
				ast.NewDeclaration("inner"),
				ast.NewAssignment("inner", ast.NewVarExpression("a")),
			),
			ast.NewAssignment("inner", ast.NewIntExpression(1)),
		)
		err := engine.Execute(block)
		assert.Error(t, err)
	})

	t.Run("shadowing outer scope variable", func(t *testing.T) {
		block := ast.NewBlockStatements(
			ast.NewDeclaration("value"),
			ast.NewBlockStatements(
				ast.NewDeclaration("value"),
				ast.NewAssignment("value", ast.NewVarExpression("a")),
			),
			ast.NewAssignment("value", ast.NewIntExpression(1)),
		)
		err := engine.Execute(block)
		assert.NoError(t, err)
	})
}

func TestExecuteIfElse(t *testing.T) {
	engine := NewEngine(nil, nil)
	ConfigurePreset(engine)

	t.Run("condition true", func(t *testing.T) {
		ifElse := ast.NewIfElse(
			ast.NewBinaryOperator(
				ast.NewIntExpression(1),
				"==",
				ast.NewIntExpression(1),
			),
			ast.NewBlockStatements(
				ast.NewAssignment("a", ast.NewIntExpression(30)),
			),
			nil,
		)
		err := engine.Execute(ifElse)
		assert.NoError(t, err)
		val, err := engine.Variables.GetSymbol("a")
		assert.NoError(t, err)
		assert.Equal(t, 30, val)
	})

	t.Run("condition false", func(t *testing.T) {
		ifElse := ast.NewIfElse(
			ast.NewBinaryOperator(
				ast.NewIntExpression(1),
				"==",
				ast.NewIntExpression(2),
			),
			ast.NewBlockStatements(
				ast.NewAssignment("b", ast.NewIntExpression(30)),
			),
			ast.NewBlockStatements(
				ast.NewAssignment("b", ast.NewIntExpression(40)),
			),
		)
		err := engine.Execute(ifElse)
		assert.NoError(t, err)
		val, err := engine.Variables.GetSymbol("b")
		assert.NoError(t, err)
		assert.Equal(t, 40, val)
	})

	t.Run("condition false, no else block", func(t *testing.T) {
		ifElse := ast.NewIfElse(
			ast.NewBinaryOperator(
				ast.NewIntExpression(1),
				"==",
				ast.NewIntExpression(2),
			),
			ast.NewBlockStatements(
				ast.NewAssignment("c", ast.NewIntExpression(100)),
			),
			nil,
		)
		err := engine.Execute(ifElse)
		assert.NoError(t, err)
		_, err = engine.Variables.GetSymbol("c")
		assert.Error(t, err)
	})
}
