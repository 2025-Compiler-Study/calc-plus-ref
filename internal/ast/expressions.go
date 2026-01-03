package ast

import (
	"calcPlus/internal/symbols"
	"fmt"
	"strings"
)

type Expression interface {
	Evaluate(symbols.Table[int]) (int, error)
	Node
}

type VarExpression struct {
	Name string
}

func (v *VarExpression) Evaluate(t symbols.Table[int]) (int, error) {
	return t.GetSymbol(v.Name)
}
func (v *VarExpression) String() string { return v.StringDepth(0) }
func (v *VarExpression) StringDepth(d int) string {
	return fmt.Sprintf("%sExprVar: %s", indent(d), v.Name)
}

type IntExpression struct {
	Value int
}

func (i *IntExpression) Evaluate(_ symbols.Table[int]) (int, error) { return i.Value, nil }
func (i *IntExpression) String() string                             { return i.StringDepth(0) }
func (i *IntExpression) StringDepth(d int) string {
	return fmt.Sprintf("%sExprInt: %d", indent(d), i.Value)
}

type BinaryOperator struct {
	Left     Expression
	Operator string
	Right    Expression
}

func (b *BinaryOperator) Evaluate(t symbols.Table[int]) (int, error) {
	left, lErr := b.Left.Evaluate(t)
	if lErr != nil {
		return 0, lErr
	}
	right, rErr := b.Right.Evaluate(t)
	if rErr != nil {
		return 0, rErr
	}
	switch b.Operator {
	case "+":
		return left + right, nil
	case "-":
		return left - right, nil
	case "*":
		return left * right, nil
	case "/":
		return left / right, nil
	default:
		return 0, fmt.Errorf("unknown operator: %s", b.Operator)
	}
}
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
