package ast

import (
	"fmt"
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
