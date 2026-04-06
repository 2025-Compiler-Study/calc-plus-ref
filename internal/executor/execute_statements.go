package executor

import (
	"calcPlus/internal/ast"
	"fmt"
)

func (e *Engine) Execute(stmt ast.Statement) error {
	switch stmt := stmt.(type) {
	case *ast.Declaration:
		return e.executeDeclaration(stmt)
	case *ast.Assignment:
		return e.executeAssignment(stmt)
	case *ast.BlockStatements:
		return e.executeBlockStatements(stmt)
	case *ast.IfElse:
		return e.executeIfElse(stmt)
	default:
		return fmt.Errorf("unsupported statement type: %T", stmt)
	}
}

func (e *Engine) executeDeclaration(d *ast.Declaration) error {
	return e.Variables.Register(d.Name)
}

func (e *Engine) executeAssignment(a *ast.Assignment) error {
	value, err := e.Evaluate(a.Value)
	if err != nil {
		return err
	}
	return e.Variables.SetSymbol(a.Name, value)
}

func (e *Engine) executeBlockStatements(b *ast.BlockStatements) error {
	if b == nil || len(*b) == 0 {
		return nil
	}
	e.Variables.PushScope()
	for _, statement := range *b {
		if err := e.Execute(statement); err != nil {
			return err
		}
	}
	e.Variables.PopScope()
	return nil
}

func (e *Engine) executeIfElse(i *ast.IfElse) error {
	condition, err := e.Evaluate(i.Condition)
	if err != nil {
		return err
	}
	if condition == 1 {
		return e.executeBlockStatements(i.ThenBlock)
	}
	return e.executeBlockStatements(i.ElseBlock)
}
