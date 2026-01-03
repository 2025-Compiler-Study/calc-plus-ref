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
		variable := VarExpression{Name: "a"}
		fmt.Println(variable.String())
	})

	t.Run("IntExpr", func(t *testing.T) {
		intValue := IntExpression{Value: 1}
		fmt.Println(intValue.String())
	})

	t.Run("BinaryOperator", func(t *testing.T) {
		binOp := BinaryOperator{
			Left:     &IntExpression{1},
			Operator: "+",
			Right:    &IntExpression{2},
		}
		fmt.Println(binOp.String())
	})

	t.Run("Complex BinaryOperator", func(t *testing.T) {
		binOp := BinaryOperator{
			Left:     &IntExpression{1},
			Operator: "+",
			Right: &BinaryOperator{
				Left:     &IntExpression{2},
				Operator: "*",
				Right:    &IntExpression{3},
			},
		}
		fmt.Println(binOp.String())
	})
}

func TestExpression_Evaluate(t *testing.T) {
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

	t.Run("VarExpr (exist)", func(t *testing.T) {
		variable := VarExpression{Name: "a"}
		val, err := variable.Evaluate(symTable)
		assert.NoError(t, err)
		assert.Equal(t, 10, val)
	})

	t.Run("VarExpr (not exist)", func(t *testing.T) {
		variable := VarExpression{Name: "c"}
		_, err := variable.Evaluate(symTable)
		assert.Error(t, err)
	})

	t.Run("IntExpr", func(t *testing.T) {
		intValue := IntExpression{Value: 1}
		val, err := intValue.Evaluate(symTable)
		assert.NoError(t, err)
		assert.Equal(t, 1, val)
	})

	t.Run("BinaryOperator (constants)", func(t *testing.T) {
		binOp := BinaryOperator{
			Left:     &IntExpression{1},
			Operator: "+",
			Right:    &IntExpression{2},
		}
		val, err := binOp.Evaluate(symTable)
		assert.NoError(t, err)
		assert.Equal(t, 3, val)
	})

	t.Run("BinaryOperator (variable)", func(t *testing.T) {
		binOp := BinaryOperator{
			Left:     &VarExpression{Name: "a"},
			Operator: "+",
			Right:    &IntExpression{2},
		}
		val, err := binOp.Evaluate(symTable)
		assert.NoError(t, err)
		assert.Equal(t, 12, val)
	})

	t.Run("BinaryOperator (not exist)", func(t *testing.T) {
		binOp := BinaryOperator{
			Left:     &VarExpression{Name: "c"},
			Operator: "+",
			Right:    &IntExpression{2},
		}
		_, err := binOp.Evaluate(symTable)
		assert.Error(t, err)
	})

	t.Run("Complex BinaryOperator", func(t *testing.T) {
		binOp := BinaryOperator{
			Left:     &IntExpression{1},
			Operator: "+",
			Right: &BinaryOperator{
				Left:     &IntExpression{2},
				Operator: "*",
				Right:    &VarExpression{Name: "b"},
			},
		}
		val, err := binOp.Evaluate(symTable)
		assert.NoError(t, err)
		assert.Equal(t, 41, val)
	})
}
