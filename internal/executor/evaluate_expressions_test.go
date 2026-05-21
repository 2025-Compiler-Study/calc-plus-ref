package executor

import (
	"calcPlus/internal/ast"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvaluateIntExpression(t *testing.T) {
	engine := NewEngine(nil, nil)
	t.Run("value", func(t *testing.T) {
		intValue := ast.NewIntExpression(1)
		val, err := engine.Evaluate(intValue)
		assert.NoError(t, err)
		assert.Equal(t, 1, val)
	})
}

func TestEvaluateVarExpression(t *testing.T) {
	engine := NewEngine(nil, nil)
	ConfigurePreset(engine)

	t.Run("exist variable", func(t *testing.T) {
		variable := ast.NewVarExpression("a")
		val, err := engine.Evaluate(variable)
		assert.NoError(t, err)
		assert.Equal(t, 10, val)
	})

	t.Run("not existing variable", func(t *testing.T) {
		variable := ast.NewVarExpression("c")
		_, err := engine.Evaluate(variable)
		assert.Error(t, err)
	})
}

func TestEvaluateBinaryOperator(t *testing.T) {
	engine := NewEngine(nil, nil)
	ConfigurePreset(engine)
	t.Run("add constants", func(t *testing.T) {
		binOp := ast.NewBinaryOperator(
			ast.NewIntExpression(1),
			"+",
			ast.NewIntExpression(2),
		)
		val, err := engine.Evaluate(binOp)
		assert.NoError(t, err)
		assert.Equal(t, 3, val)
	})

	t.Run("add variable and constant", func(t *testing.T) {
		binOp := ast.NewBinaryOperator(
			ast.NewVarExpression("a"),
			"+",
			ast.NewIntExpression(2),
		)
		val, err := engine.Evaluate(binOp)
		assert.NoError(t, err)
		assert.Equal(t, 12, val)
	})

	t.Run("use not existing variable", func(t *testing.T) {
		binOp := ast.NewBinaryOperator(
			ast.NewVarExpression("c"),
			"+",
			ast.NewIntExpression(2),
		)
		_, err := engine.Evaluate(binOp)
		assert.Error(t, err)
	})

	t.Run("compare true", func(t *testing.T) {
		binOp := ast.NewBinaryOperator(
			ast.NewIntExpression(1),
			">",
			ast.NewIntExpression(0),
		)
		val, err := engine.Evaluate(binOp)
		assert.NoError(t, err)
		assert.Equal(t, 1, val)
	})

	t.Run("compare false", func(t *testing.T) {
		binOp := ast.NewBinaryOperator(
			ast.NewIntExpression(1),
			"==",
			ast.NewIntExpression(2),
		)
		val, err := engine.Evaluate(binOp)
		assert.NoError(t, err)
		assert.Equal(t, 0, val)
	})

	t.Run("not supported operator", func(t *testing.T) {
		binOp := ast.NewBinaryOperator(
			ast.NewIntExpression(1),
			"**",
			ast.NewIntExpression(2),
		)
		_, err := engine.Evaluate(binOp)
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
		val, err := engine.Evaluate(binOp)
		assert.NoError(t, err)
		assert.Equal(t, 41, val)
	})
}
