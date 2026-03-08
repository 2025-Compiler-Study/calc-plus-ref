package ast

import (
	"fmt"
	"strings"
)

type Expression interface {
	expression()
	Node
}

type VarExpression struct {
	Name string
}

func NewVarExpression(name string) *VarExpression {
	return &VarExpression{Name: name}
}
func (v *VarExpression) expression()    {}
func (v *VarExpression) String() string { return v.StringDepth(0) }
func (v *VarExpression) StringDepth(d int) string {
	return fmt.Sprintf("%sExprVar: %s", indent(d), v.Name)
}

type IntExpression struct {
	Value int
}

func NewIntExpression(value int) *IntExpression {
	return &IntExpression{Value: value}
}
func (i *IntExpression) expression()    {}
func (i *IntExpression) String() string { return i.StringDepth(0) }
func (i *IntExpression) StringDepth(d int) string {
	return fmt.Sprintf("%sExprInt: %d", indent(d), i.Value)
}

type BinaryOperator struct {
	Left     Expression
	Operator string
	Right    Expression
}

func NewBinaryOperator(left Expression, operator string, right Expression) *BinaryOperator {
	return &BinaryOperator{Left: left, Operator: operator, Right: right}
}
func (b *BinaryOperator) expression()    {}
func (b *BinaryOperator) String() string { return b.StringDepth(0) }
func (b *BinaryOperator) StringDepth(d int) string {
	leftValue := b.Left.StringDepth(d + 1)
	rightValue := b.Right.StringDepth(d + 1)
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("%sBinaryOperator: %s\n", indent(d), b.Operator))
	sb.WriteString(fmt.Sprintf("%s.Left:\n%s\n", indent(d), leftValue))
	sb.WriteString(fmt.Sprintf("%s.Right:\n%s\n", indent(d), rightValue))

	return sb.String()
}
