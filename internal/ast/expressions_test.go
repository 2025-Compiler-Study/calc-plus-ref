package ast

import (
	"fmt"
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
