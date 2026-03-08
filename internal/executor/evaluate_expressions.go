package executor

import (
	"calcPlus/internal/ast"
	"calcPlus/internal/symbols"
	"fmt"
)

func Evaluate(e ast.Expression, symTable symbols.Table[int]) (int, error) {
	switch e := e.(type) {
	case *ast.IntExpression:
		return evaluateIntExpression(e)
	case *ast.VarExpression:
		return evaluateVarExpression(e, symTable)
	case *ast.BinaryOperator:
		return evaluateBinaryOperator(e, symTable)
	default:
		return 0, fmt.Errorf("unknown expression type: %T", e)
	}
}

func evaluateIntExpression(i *ast.IntExpression) (int, error) {
	return i.Value, nil
}

func evaluateVarExpression(v *ast.VarExpression, symTable symbols.Table[int]) (int, error) {
	return symTable.GetSymbol(v.Name)
}

func evaluateBinaryOperator(b *ast.BinaryOperator, symTable symbols.Table[int]) (int, error) {
	left, lErr := Evaluate(b.Left, symTable)
	if lErr != nil {
		return 0, lErr
	}
	right, rErr := Evaluate(b.Right, symTable)
	if rErr != nil {
		return 0, rErr
	}
	logicalResult := false
	switch b.Operator {
	// Arithmetic operators
	case "+":
		return left + right, nil
	case "-":
		return left - right, nil
	case "*":
		return left * right, nil
	case "/":
		return left / right, nil
	// Logical operators
	case "==":
		logicalResult = left == right
	case "!=":
		logicalResult = left != right
	case ">":
		logicalResult = left > right
	case ">=":
		logicalResult = left >= right
	case "<":
		logicalResult = left < right
	case "<=":
		logicalResult = left <= right
	default:
		return 0, fmt.Errorf("unknown operator: %s", b.Operator)
	}

	// Fallback logical operators result
	if logicalResult {
		return 1, nil
	}
	return 0, nil
}
