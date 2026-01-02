package ast

import (
	"fmt"
	"strings"
)

type Expression interface {
	Node
}

type VarExpression struct {
	Name string
}

func (v *VarExpression) String() string { return v.StringDepth(0) }
func (v *VarExpression) StringDepth(d int) string {
	indent := strings.Repeat("..", d)
	return fmt.Sprintf("%sExprVar: %s", indent, v.Name)
}

type IntExpression struct {
	Value int
}

func (i *IntExpression) String() string { return i.StringDepth(0) }
func (i *IntExpression) StringDepth(d int) string {
	indent := strings.Repeat("..", d)
	return fmt.Sprintf("%sExprInt: %d", indent, i.Value)
}

type BinaryOperator struct {
	Left     Expression
	Operator string
	Right    Expression
}

func (b *BinaryOperator) String() string { return b.StringDepth(0) }
func (b *BinaryOperator) StringDepth(d int) string {
	indent := strings.Repeat("..", d)
	leftValue := b.Left.StringDepth(d + 1)
	rightValue := b.Right.StringDepth(d + 1)
	result := strings.Builder{}
	result.WriteString(fmt.Sprintf("%sLeft:\n%s\n", indent, leftValue))
	result.WriteString(fmt.Sprintf("%sOperator: %s\n", indent, b.Operator))
	result.WriteString(fmt.Sprintf("%sRight:\n%s\n", indent, rightValue))

	return result.String()
}
