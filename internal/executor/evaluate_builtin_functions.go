package executor

import (
	"calcPlus/internal/ast"
	"fmt"
	"strconv"
)

func (e *Engine) evaluateBuiltinRead(r *ast.BuiltinReadCall) (int, error) {
	buffer := ""
	if _, err := fmt.Fscanln(e.Input, &buffer); err != nil {
		return 0, err
	}
	return strconv.Atoi(buffer)
}

func (e *Engine) evaluateBuiltinWrite(r *ast.BuiltinWriteCall) (int, error) {
	value, err := e.Evaluate(r.Argument)
	if err != nil {
		return 0, err
	}

	return fmt.Fprintln(e.Output, value)
}
