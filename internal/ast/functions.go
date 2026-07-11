package ast

import (
	"fmt"
	"strings"
)

// Program is the AST root for Calc-6. A Calc-6 source file consists only of
// function definitions.
type Program struct {
	Functions []*FunctionDefinition
}

func NewProgram(functions ...*FunctionDefinition) *Program {
	return &Program{Functions: functions}
}

func (p *Program) String() string { return p.StringDepth(0) }
func (p *Program) StringDepth(depth int) string {
	var sb strings.Builder
	for _, function := range p.Functions {
		sb.WriteString(function.StringDepth(depth))
	}
	return sb.String()
}

type Parameter struct {
	Name string
}

type FunctionDefinition struct {
	Name       string
	ReturnsInt bool
	Parameters []Parameter
	Body       *BlockStatements
}

func NewFunctionDefinition(name string, returnsInt bool, parameters []Parameter, body *BlockStatements) *FunctionDefinition {
	return &FunctionDefinition{Name: name, ReturnsInt: returnsInt, Parameters: parameters, Body: body}
}

func (f *FunctionDefinition) String() string { return f.StringDepth(0) }
func (f *FunctionDefinition) StringDepth(depth int) string {
	returnType := "void"
	if f.ReturnsInt {
		returnType = "int"
	}

	parameters := make([]string, 0, len(f.Parameters))
	for _, parameter := range f.Parameters {
		parameters = append(parameters, "int "+parameter.Name)
	}

	return fmt.Sprintf("%sFunction: %s %s(%s)\n%s", indent(depth), returnType, f.Name,
		strings.Join(parameters, ", "), f.Body.StringDepth(depth+1))
}

type FunctionCall struct {
	Name      string
	Arguments []Expression
}

func NewFunctionCall(name string, arguments ...Expression) *FunctionCall {
	return &FunctionCall{Name: name, Arguments: arguments}
}

func (f *FunctionCall) expression()    {}
func (f *FunctionCall) String() string { return f.StringDepth(0) }
func (f *FunctionCall) StringDepth(depth int) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("%sFunctionCall: %s\n", indent(depth), f.Name))
	for index, argument := range f.Arguments {
		sb.WriteString(fmt.Sprintf("%s.Arg.%d:\n%s\n", indent(depth), index, argument.StringDepth(depth+1)))
	}
	return sb.String()
}
