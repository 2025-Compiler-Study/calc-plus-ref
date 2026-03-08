package ast

import (
	"fmt"
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
