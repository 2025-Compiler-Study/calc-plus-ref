package executor

import (
	"calcPlus/internal/ast"
	"calcPlus/internal/symbols"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestEvaluateIntExpression(t *testing.T) {
	t.Run("value", func(t *testing.T) {
		intValue := ast.NewIntExpression(1)
		val, err := evaluateIntExpression(intValue)
		assert.NoError(t, err)
		assert.Equal(t, 1, val)
	})
}

func TestEvaluateVarExpression(t *testing.T) {
	symTable := symbols.NewSimpleTable[int]()
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

	t.Run("exist variable", func(t *testing.T) {
		variable := ast.NewVarExpression("a")
		val, err := evaluateVarExpression(variable, symTable)
		assert.NoError(t, err)
		assert.Equal(t, 10, val)
	})

	t.Run("not existing variable", func(t *testing.T) {
		variable := ast.NewVarExpression("c")
		_, err := evaluateVarExpression(variable, symTable)
		assert.Error(t, err)
	})
}

func TestEvaluateBinaryOperator(t *testing.T) {
	symTable := symbols.NewSimpleTable[int]()
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

	t.Run("add constants", func(t *testing.T) {
		binOp := ast.NewBinaryOperator(
			ast.NewIntExpression(1),
			"+",
			ast.NewIntExpression(2),
		)
		val, err := evaluateBinaryOperator(binOp, symTable)
		assert.NoError(t, err)
		assert.Equal(t, 3, val)
	})

	t.Run("add variable and constant", func(t *testing.T) {
		binOp := ast.NewBinaryOperator(
			ast.NewVarExpression("a"),
			"+",
			ast.NewIntExpression(2),
		)
		val, err := evaluateBinaryOperator(binOp, symTable)
		assert.NoError(t, err)
		assert.Equal(t, 12, val)
	})

	t.Run("use not existing variable", func(t *testing.T) {
		binOp := ast.NewBinaryOperator(
			ast.NewVarExpression("c"),
			"+",
			ast.NewIntExpression(2),
		)
		_, err := evaluateBinaryOperator(binOp, symTable)
		assert.Error(t, err)
	})

	t.Run("compare true", func(t *testing.T) {
		binOp := ast.NewBinaryOperator(
			ast.NewIntExpression(1),
			">",
			ast.NewIntExpression(0),
		)
		val, err := evaluateBinaryOperator(binOp, symTable)
		assert.NoError(t, err)
		assert.Equal(t, 1, val)
	})

	t.Run("compare false", func(t *testing.T) {
		binOp := ast.NewBinaryOperator(
			ast.NewIntExpression(1),
			"==",
			ast.NewIntExpression(2),
		)
		val, err := evaluateBinaryOperator(binOp, symTable)
		assert.NoError(t, err)
		assert.Equal(t, 0, val)
	})

	t.Run("not supported operator", func(t *testing.T) {
		binOp := ast.NewBinaryOperator(
			ast.NewIntExpression(1),
			"**",
			ast.NewIntExpression(2),
		)
		_, err := evaluateBinaryOperator(binOp, symTable)
		assert.Error(t, err)
	})

	t.Run("complex expression", func(t *testing.T) {
		binOp := ast.NewBinaryOperator(
			ast.NewIntExpression(1),
			"+",
			ast.NewBinaryOperator(
				ast.NewIntExpression(2),
				"*",
				ast.NewVarExpression("b"),
			),
		)
		val, err := evaluateBinaryOperator(binOp, symTable)
		assert.NoError(t, err)
		assert.Equal(t, 41, val)
	})
}
