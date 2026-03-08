package ast

import (
	"fmt"
	"testing"
)

func TestBuiltinFunctions_String(t *testing.T) {
	t.Run("read", func(t *testing.T) {
		readCall := NewBuiltinReadCall()
		fmt.Println(readCall.String())
	})

	t.Run("write", func(t *testing.T) {
		writeCall := NewBuiltinWriteCall(NewIntExpression(42))
		fmt.Println(writeCall.String())
	})
}
