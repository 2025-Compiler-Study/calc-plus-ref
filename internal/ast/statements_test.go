package ast

import (
	"calcPlus/internal/symbols"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestStatement_String(t *testing.T) {
	t.Run("Declaration", func(t *testing.T) {
		declaration := NewDeclaration("a")
		fmt.Println(declaration.String())
	})

	t.Run("Assignment", func(t *testing.T) {
		assignment := NewAssignment("a", NewIntExpression(1))
		fmt.Println(assignment.String())
	})

	t.Run("BlockStatements", func(t *testing.T) {
		block := NewBlockStatements(
			NewDeclaration("a"),
			NewAssignment("a", NewIntExpression(1)),
		)
		fmt.Println(block.String())
	})

	t.Run("IfElse (all block)", func(t *testing.T) {
		ifElse := NewIfElse(
			NewBinaryOperator(
				NewIntExpression(1),
				"==",
				NewIntExpression(1),
			),
			NewBlockStatements(
				NewDeclaration("a"),
				NewAssignment("a", NewIntExpression(1)),
			),
			NewBlockStatements(
				NewDeclaration("a"),
				NewAssignment("a", NewIntExpression(2)),
			),
		)
		fmt.Println(ifElse.String())
	})

	t.Run("IfElse (skip else)", func(t *testing.T) {
		ifElse := NewIfElse(
			NewBinaryOperator(
				NewIntExpression(1),
				"==",
				NewIntExpression(1),
			),
			NewBlockStatements(
				NewDeclaration("a"),
				NewAssignment("a", NewIntExpression(1)),
			),
			NewBlockStatements(),
		)
		fmt.Println(ifElse.String())
	})
}

func TestDeclaration_Execute(t *testing.T) {
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
		declaration := NewDeclaration("new")
		err := declaration.Execute(symTable)
		assert.NoError(t, err)
	})

	t.Run("declaring already declared", func(t *testing.T) {
		declaration := NewDeclaration("a")
		err := declaration.Execute(symTable)
		assert.Error(t, err)
	})
}

func TestAssignment_Execute(t *testing.T) {
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
		assignment := NewAssignment(
			"a",
			NewIntExpression(1),
		)
		err := assignment.Execute(symTable)
		assert.NoError(t, err)
	})

	t.Run("assign to undeclared variable", func(t *testing.T) {
		assignment := NewAssignment(
			"undeclared",
			NewIntExpression(1),
		)
		err := assignment.Execute(symTable)
		assert.Error(t, err)
	})

	t.Run("expression uses undeclared variable", func(t *testing.T) {
		assignment := NewAssignment(
			"a",
			NewBinaryOperator(
				NewVarExpression("undeclared"),
				"+",
				NewIntExpression(1),
			),
		)
		err := assignment.Execute(symTable)
		assert.Error(t, err)
	})
}

func TestBlockStatements_Execute(t *testing.T) {
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
		block := NewBlockStatements(
			NewDeclaration("value"),
			NewAssignment("value", NewIntExpression(1)),
		)
		err := block.Execute(symTable)
		assert.NoError(t, err)
	})

	t.Run("empty block", func(t *testing.T) {
		block := NewBlockStatements()
		err := block.Execute(symTable)
		assert.NoError(t, err)
	})

	t.Run("use inner scope variable from outer scope error", func(t *testing.T) {
		block := NewBlockStatements(
			NewBlockStatements(
				NewDeclaration("inner"),
				NewAssignment("inner", NewVarExpression("a")),
			),
			NewAssignment("inner", NewIntExpression(1)),
		)
		err := block.Execute(symTable)
		assert.Error(t, err)
	})

	t.Run("shadowing outer scope variable", func(t *testing.T) {
		block := NewBlockStatements(
			NewDeclaration("value"),
			NewBlockStatements(
				NewDeclaration("value"),
				NewAssignment("value", NewVarExpression("a")),
			),
			NewAssignment("value", NewIntExpression(1)),
		)
		err := block.Execute(symTable)
		assert.NoError(t, err)
	})
}

func TestIfElse_Execute(t *testing.T) {
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
		ifElse := NewIfElse(
			NewBinaryOperator(
				NewIntExpression(1),
				"==",
				NewIntExpression(1),
			),
			NewBlockStatements(
				NewAssignment("a", NewIntExpression(30)),
			),
			nil,
		)
		err := ifElse.Execute(symTable)
		assert.NoError(t, err)
		val, err := symTable.GetSymbol("a")
		assert.NoError(t, err)
		assert.Equal(t, 30, val)
	})

	t.Run("condition false", func(t *testing.T) {
		ifElse := NewIfElse(
			NewBinaryOperator(
				NewIntExpression(1),
				"==",
				NewIntExpression(2),
			),
			NewBlockStatements(
				NewAssignment("b", NewIntExpression(30)),
			),
			NewBlockStatements(
				NewAssignment("b", NewIntExpression(40)),
			),
		)
		err := ifElse.Execute(symTable)
		assert.NoError(t, err)
		val, err := symTable.GetSymbol("b")
		assert.NoError(t, err)
		assert.Equal(t, 40, val)
	})

	t.Run("condition false, no else block", func(t *testing.T) {
		ifElse := NewIfElse(
			NewBinaryOperator(
				NewIntExpression(1),
				"==",
				NewIntExpression(2),
			),
			NewBlockStatements(
				NewAssignment("c", NewIntExpression(100)),
			),
			nil,
		)
		err := ifElse.Execute(symTable)
		assert.NoError(t, err)
		val, err := symTable.GetSymbol("c")
		assert.NoError(t, err)
		assert.NotEqual(t, 100, val)
	})
}
