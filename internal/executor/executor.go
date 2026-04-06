package executor

import (
	"calcPlus/internal/symbols"
	"io"
)

type Engine struct {
	Input     io.Reader
	Output    io.Writer
	Variables symbols.ScopedTable[int]
}

func NewEngine(input io.Reader, output io.Writer) *Engine {
	return &Engine{
		Input:     input,
		Output:    output,
		Variables: *symbols.NewScopedTable[int](),
	}
}
