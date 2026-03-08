package ast

import (
	"fmt"
	"io"
	"strings"
)

type BuiltinReadCall struct {
	Reader io.Reader
}

func NewBuiltinReadCall() *BuiltinReadCall {
	return &BuiltinReadCall{}
}
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
func (f *BuiltinWriteCall) String() string { return f.StringDepth(0) }
func (f *BuiltinWriteCall) StringDepth(d int) string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("%sFunctionCall: (built-in)write\n", indent(d)))
	sb.WriteString(fmt.Sprintf("%s.Arg.0:\n%s\n", indent(d), f.Argument.StringDepth(d+1)))

	return sb.String()
}
