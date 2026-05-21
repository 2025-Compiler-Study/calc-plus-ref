package executor

import (
	"calcPlus/internal/ast"
	"fmt"
)

func (e *Engine) Evaluate(expr ast.Expression) (int, error) {
	switch expr := expr.(type) {
	case *ast.IntExpression:
		return e.evaluateIntExpression(expr)
	case *ast.VarExpression:
		return e.evaluateVarExpression(expr)
	case *ast.BinaryOperator:
		return e.evaluateBinaryOperator(expr)
	case *ast.BuiltinReadCall:
		return e.evaluateBuiltinRead(expr)
	case *ast.BuiltinWriteCall:
		return e.evaluateBuiltinWrite(expr)
	default:
		return 0, fmt.Errorf("unknown expression type: %T", e)
	}
}

func (e *Engine) evaluateIntExpression(i *ast.IntExpression) (int, error) {
	return i.Value, nil
}

func (e *Engine) evaluateVarExpression(v *ast.VarExpression) (int, error) {
	return e.Variables.GetSymbol(v.Name)
}

func (e *Engine) evaluateBinaryOperator(b *ast.BinaryOperator) (int, error) {
	left, lErr := e.Evaluate(b.Left)
	if lErr != nil {
		return 0, lErr
	}
	right, rErr := e.Evaluate(b.Right)
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
