package ast

import (
	"calcPlus/internal/symbols"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestExpression_String(t *testing.T) {
	t.Run("VarExpr", func(t *testing.T) {
		variable := NewVarExpression("a")
		fmt.Println(variable.String())
	})

	t.Run("IntExpr", func(t *testing.T) {
		intValue := NewIntExpression(1)
		fmt.Println(intValue.String())
	})

	t.Run("BinaryOperator(Arithmetic)", func(t *testing.T) {
		binOp := NewBinaryOperator(
			NewIntExpression(1),
			"+",
			NewIntExpression(2),
		)
		fmt.Println(binOp.String())
	})

	t.Run("Complex BinaryOperator", func(t *testing.T) {
		binOp := NewBinaryOperator(
			NewIntExpression(1),
			"+",
			NewBinaryOperator(
				NewIntExpression(2),
				"*",
				NewIntExpression(3),
			),
		)
		fmt.Println(binOp.String())
	})
}

func TestVarExpr_Evaluate(t *testing.T) {
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
		variable := NewVarExpression("a")
		val, err := variable.Evaluate(symTable)
		assert.NoError(t, err)
		assert.Equal(t, 10, val)
	})

	t.Run("not existing variable", func(t *testing.T) {
		variable := NewVarExpression("c")
		_, err := variable.Evaluate(symTable)
		assert.Error(t, err)
	})
}

func TestIntExpr_Evaluate(t *testing.T) {
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

	t.Run("integer", func(t *testing.T) {
		intValue := NewIntExpression(1)
		val, err := intValue.Evaluate(symTable)
		assert.NoError(t, err)
		assert.Equal(t, 1, val)
	})
}

func TestBinaryOperator_Evaluate(t *testing.T) {
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
		binOp := NewBinaryOperator(
			NewIntExpression(1),
			"+",
			NewIntExpression(2),
		)
		val, err := binOp.Evaluate(symTable)
		assert.NoError(t, err)
		assert.Equal(t, 3, val)
	})

	t.Run("add variable and constant", func(t *testing.T) {
		binOp := NewBinaryOperator(
			NewVarExpression("a"),
			"+",
			NewIntExpression(2),
		)
		val, err := binOp.Evaluate(symTable)
		assert.NoError(t, err)
		assert.Equal(t, 12, val)
	})

	t.Run("use not existing variable", func(t *testing.T) {
		binOp := NewBinaryOperator(
			NewVarExpression("c"),
			"+",
			NewIntExpression(2),
		)
		_, err := binOp.Evaluate(symTable)
		assert.Error(t, err)
	})

	t.Run("compare true", func(t *testing.T) {
		binOp := NewBinaryOperator(
			NewIntExpression(1),
			">",
			NewIntExpression(0),
		)
		val, err := binOp.Evaluate(symTable)
		assert.NoError(t, err)
		assert.Equal(t, 1, val)
	})

	t.Run("compare false", func(t *testing.T) {
		binOp := NewBinaryOperator(
			NewIntExpression(1),
			"==",
			NewIntExpression(2),
		)
		val, err := binOp.Evaluate(symTable)
		assert.NoError(t, err)
		assert.Equal(t, 0, val)
	})

	t.Run("not supported operator", func(t *testing.T) {
		binOp := NewBinaryOperator(
			NewIntExpression(1),
			"**",
			NewIntExpression(2),
		)
		_, err := binOp.Evaluate(symTable)
		assert.Error(t, err)
	})

	t.Run("complex expression", func(t *testing.T) {
		binOp := NewBinaryOperator(
			NewIntExpression(1),
			"+",
			NewBinaryOperator(
				NewIntExpression(2),
				"*",
				NewVarExpression("b"),
			),
		)
		val, err := binOp.Evaluate(symTable)
		assert.NoError(t, err)
		assert.Equal(t, 41, val)
	})
}
