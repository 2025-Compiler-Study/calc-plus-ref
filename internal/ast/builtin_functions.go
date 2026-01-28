package ast

import (
	"calcPlus/internal/symbols"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type BuiltinReadCall struct {
	Reader io.Reader
}

func (r *BuiltinReadCall) String() string { return r.StringDepth(0) }
func (r *BuiltinReadCall) StringDepth(d int) string {
	return fmt.Sprintf("%sFunctionCall: (built-in)read", indent(d))
}
func (r *BuiltinReadCall) Evaluate(_ symbols.Table[int]) (int, error) {
	buffer := ""
	if _, err := fmt.Fscanln(r.Reader, &buffer); err != nil {
		return 0, err
	}
	return strconv.Atoi(buffer)
}

type BuiltinWriteCall struct {
	Writer   io.Writer
	Argument Expression
}

func (f *BuiltinWriteCall) String() string { return f.StringDepth(0) }
func (f *BuiltinWriteCall) StringDepth(d int) string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("%sFunctionCall: (built-in)write\n", indent(d)))
	sb.WriteString(fmt.Sprintf("%s.Arg.0:\n%s\n", indent(d), f.Argument.StringDepth(d+1)))

	return sb.String()
}
func (f *BuiltinWriteCall) Evaluate(env symbols.Table[int]) (int, error) {
	value, err := f.Argument.Evaluate(env)
	if err != nil {
		return 0, err
	}

	return fmt.Fprintln(f.Writer, value)
}
