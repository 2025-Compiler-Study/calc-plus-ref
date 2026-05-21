package ast

import (
	"fmt"
	"strings"
)

type Statement interface {
	statement()
	Node
}

type Declaration struct {
	Name string
}

func NewDeclaration(name string) *Declaration {
	return &Declaration{Name: name}
}
func (d *Declaration) statement()     {}
func (d *Declaration) String() string { return d.StringDepth(0) }
func (d *Declaration) StringDepth(depth int) string {
	return fmt.Sprintf("%sDecl: %s\n", indent(depth), d.Name)
}

type Assignment struct {
	Name  string
	Value Expression
}

func NewAssignment(name string, value Expression) *Assignment {
	return &Assignment{Name: name, Value: value}
}
func (a *Assignment) statement()     {}
func (a *Assignment) String() string { return a.StringDepth(0) }
func (a *Assignment) StringDepth(d int) string {
	expr := a.Value.StringDepth(d + 1)
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("%sAssignVar: %s\n", indent(d), a.Name))
	sb.WriteString(fmt.Sprintf("%s.AssignExpr: \n%s\n", indent(d), expr))

	return sb.String()
}

type BlockStatements []Statement

func NewBlockStatements(statements ...Statement) *BlockStatements {
	var value BlockStatements
	value = append(value, statements...)
	return &value
}
func (b *BlockStatements) statement()     {}
func (b *BlockStatements) String() string { return b.StringDepth(0) }
func (b *BlockStatements) StringDepth(d int) string {
	sb := strings.Builder{}
	for _, stmt := range *b {
		sb.WriteString(fmt.Sprintf("%s", stmt.StringDepth(d)))
	}

	return sb.String()
}

type IfElse struct {
	Condition Expression
	ThenBlock *BlockStatements
	ElseBlock *BlockStatements
}

func NewIfElse(condition Expression, thenBlock, elseBlock *BlockStatements) *IfElse {
	return &IfElse{Condition: condition, ThenBlock: thenBlock, ElseBlock: elseBlock}
}
func (i *IfElse) statement()     {}
func (i *IfElse) String() string { return i.StringDepth(0) }
func (i *IfElse) StringDepth(d int) string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("%sIfCondition:\n%s", indent(d), i.Condition.StringDepth(d+1)))
	sb.WriteString(fmt.Sprintf("%s.ThenBlock:\n%s", indent(d), i.ThenBlock.StringDepth(d+1)))
	if i.ElseBlock == nil || len(*i.ElseBlock) > 0 {
		sb.WriteString(fmt.Sprintf("%s.ElseBlock:\n%s", indent(d), i.ElseBlock.StringDepth(d+1)))
	}

	return sb.String()
}
