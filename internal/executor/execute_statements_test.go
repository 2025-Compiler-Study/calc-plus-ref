package executor

import (
	"calcPlus/internal/ast"
	"calcPlus/internal/symbols"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestExecuteDeclaration(t *testing.T) {
	symTable := symbols.NewScopedTable[int]()
	presets := map[string]int{
		"a": 10,
		"b": 20,
	}
	for k, v := range presets {
		err := symTable.Register(k)
		require.NoError(t, err)
		err = symTable.SetSymbol(k, v)
		require.NoError(t, err)
	}

	t.Run("declaring undeclared name", func(t *testing.T) {
		declaration := ast.NewDeclaration("new")
		err := Execute(declaration, symTable)
		assert.NoError(t, err)
	})

	t.Run("declaring already declared", func(t *testing.T) {
		declaration := ast.NewDeclaration("a")
		err := Execute(declaration, symTable)
		assert.Error(t, err)
	})
}

func TestExecuteAssignment(t *testing.T) {
	symTable := symbols.NewScopedTable[int]()
	presets := map[string]int{
		"a": 10,
		"b": 20,
	}
	for k, v := range presets {
		err := symTable.Register(k)
		require.NoError(t, err)
		err = symTable.SetSymbol(k, v)
		require.NoError(t, err)
	}

	t.Run("normal case", func(t *testing.T) {
		assignment := ast.NewAssignment(
			"a",
			ast.NewIntExpression(1),
		)
		err := Execute(assignment, symTable)
		assert.NoError(t, err)
	})

	t.Run("assign to undeclared variable", func(t *testing.T) {
		assignment := ast.NewAssignment(
			"undeclared",
			ast.NewIntExpression(1),
		)
		err := Execute(assignment, symTable)
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
		err := Execute(assignment, symTable)
		assert.Error(t, err)
	})
}

func TestExecuteBlockStatement(t *testing.T) {
	symTable := symbols.NewScopedTable[int]()
	presets := map[string]int{
		"a": 10,
		"b": 20,
	}
	for k, v := range presets {
		err := symTable.Register(k)
		require.NoError(t, err)
		err = symTable.SetSymbol(k, v)
		require.NoError(t, err)
	}

	t.Run("multiple statements", func(t *testing.T) {
		block := ast.NewBlockStatements(
			ast.NewDeclaration("value"),
			ast.NewAssignment("value", ast.NewIntExpression(1)),
		)
		err := Execute(block, symTable)
		assert.NoError(t, err)
	})

	t.Run("empty block", func(t *testing.T) {
		block := ast.NewBlockStatements()
		err := Execute(block, symTable)
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
		err := Execute(block, symTable)
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
		err := Execute(block, symTable)
		assert.NoError(t, err)
	})
}

func TestExecuteIfElse(t *testing.T) {
	symTable := symbols.NewScopedTable[int]()
	presets := map[string]int{
		"a": 10,
		"b": 20,
		"c": 30,
	}
	for k, v := range presets {
		err := symTable.Register(k)
		require.NoError(t, err)
		err = symTable.SetSymbol(k, v)
	}

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
		err := Execute(ifElse, symTable)
		assert.NoError(t, err)
		val, err := symTable.GetSymbol("a")
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
		err := Execute(ifElse, symTable)
		assert.NoError(t, err)
		val, err := symTable.GetSymbol("b")
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
		err := Execute(ifElse, symTable)
		assert.NoError(t, err)
		val, err := symTable.GetSymbol("c")
		assert.NoError(t, err)
		assert.NotEqual(t, 100, val)
	})
}
