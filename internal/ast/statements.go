package ast

import (
	"calcPlus/internal/symbols"
	"fmt"
	"strings"
)

type Statement interface {
	Execute(*symbols.ScopedTable[int]) error
	Node
}

type Declaration struct {
	Name string
}

func (d *Declaration) String() string { return d.StringDepth(0) }
func (d *Declaration) StringDepth(depth int) string {
	return fmt.Sprintf("%sDecl: %s\n", indent(depth), d.Name)
}
func (d *Declaration) Execute(env *symbols.ScopedTable[int]) error {
	return env.Register(d.Name)
}

type Assignment struct {
	Name  string
	Value Expression
}

func (a *Assignment) String() string { return a.StringDepth(0) }
func (a *Assignment) StringDepth(d int) string {
	expr := a.Value.StringDepth(d + 1)
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("%sAssignVar: %s\n", indent(d), a.Name))
	sb.WriteString(fmt.Sprintf("%s.AssignExpr: \n%s\n", indent(d), expr))

	return sb.String()
}
func (a *Assignment) Execute(env *symbols.ScopedTable[int]) error {
	value, err := a.Value.Evaluate(env)
	if err != nil {
		return err
	}
	return env.SetSymbol(a.Name, value)
}

type BlockStatements []Statement

func (b *BlockStatements) String() string { return b.StringDepth(0) }
func (b *BlockStatements) StringDepth(d int) string {
	sb := strings.Builder{}
	for _, stmt := range *b {
		sb.WriteString(fmt.Sprintf("%s", stmt.StringDepth(d)))
	}

	return sb.String()
}
func (b *BlockStatements) Execute(env *symbols.ScopedTable[int]) error {
	env.PushScope()
	for _, stmt := range *b {
		err := stmt.Execute(env)
		if err != nil {
			return err
		}
	}
	env.PopScope()
	return nil
}

type IfElse struct {
	Condition Expression
	ThenBlock BlockStatements
	ElseBlock BlockStatements
}

func (i *IfElse) String() string { return i.StringDepth(0) }
func (i *IfElse) StringDepth(d int) string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("%sIfCondition:\n%s", indent(d), i.Condition.StringDepth(d+1)))
	sb.WriteString(fmt.Sprintf("%s.ThenBlock:\n%s", indent(d), i.ThenBlock.StringDepth(d+1)))
	if len(i.ElseBlock) > 0 {
		sb.WriteString(fmt.Sprintf("%s.ElseBlock:\n%s", indent(d), i.ElseBlock.StringDepth(d+1)))
	}

	return sb.String()
}
func (i *IfElse) Execute(env *symbols.ScopedTable[int]) error {
	condition, err := i.Condition.Evaluate(env)
	if err != nil {
		return err
	}
	if condition == 1 {
		return i.ThenBlock.Execute(env)
	} else {
		return i.ElseBlock.Execute(env)
	}
}
