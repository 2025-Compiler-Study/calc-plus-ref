package ast

import (
	"fmt"
	"strings"
)

type BuiltinReadCall struct {
}

func NewBuiltinReadCall() *BuiltinReadCall {
	return &BuiltinReadCall{}
}
func (r *BuiltinReadCall) expression()    {}
func (r *BuiltinReadCall) String() string { return r.StringDepth(0) }
func (r *BuiltinReadCall) StringDepth(d int) string {
	return fmt.Sprintf("%sFunctionCall: (built-in)read", indent(d))
}

type BuiltinWriteCall struct {
	Argument Expression
}

func NewBuiltinWriteCall(argument Expression) *BuiltinWriteCall {
	return &BuiltinWriteCall{Argument: argument}
}
func (w *BuiltinWriteCall) expression()    {}
func (w *BuiltinWriteCall) String() string { return w.StringDepth(0) }
func (w *BuiltinWriteCall) StringDepth(d int) string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("%sFunctionCall: (built-in)write\n", indent(d)))
	sb.WriteString(fmt.Sprintf("%s.Arg.0:\n%s\n", indent(d), w.Argument.StringDepth(d+1)))

	return sb.String()
}
