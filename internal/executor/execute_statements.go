package executor

import (
	"calcPlus/internal/ast"
	"calcPlus/internal/symbols"
	"fmt"
)

func Execute(s ast.Statement, symTable *symbols.ScopedTable[int]) error {
	switch s := s.(type) {
	case *ast.Declaration:
		return executeDeclaration(s, symTable)
	case *ast.Assignment:
		return executeAssignment(s, symTable)
	case *ast.BlockStatements:
		return executeBlockStatements(s, symTable)
	case *ast.IfElse:
		return executeIfElse(s, symTable)
	default:
		return fmt.Errorf("unsupported statement type: %T", s)
	}
}

func executeDeclaration(d *ast.Declaration, symTable *symbols.ScopedTable[int]) error {
	return symTable.Register(d.Name)
}

func executeAssignment(a *ast.Assignment, symTable *symbols.ScopedTable[int]) error {
	value, err := Evaluate(a.Value, symTable)
	if err != nil {
		return err
	}
	return symTable.SetSymbol(a.Name, value)
}

func executeBlockStatements(b *ast.BlockStatements, symTable *symbols.ScopedTable[int]) error {
	if b == nil || len(*b) == 0 {
		return nil
	}
	symTable.PushScope()
	for _, statement := range *b {
		if err := Execute(statement, symTable); err != nil {
			return err
		}
	}
	symTable.PopScope()
	return nil
}

func executeIfElse(i *ast.IfElse, symTable *symbols.ScopedTable[int]) error {
	condition, err := Evaluate(i.Condition, symTable)
	if err != nil {
		return err
	}
	if condition == 1 {
		return executeBlockStatements(i.ThenBlock, symTable)
	}
	return executeBlockStatements(i.ElseBlock, symTable)
}
