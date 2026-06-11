package interpreter

import (
	"calcPlus/internal/ast"
	"calcPlus/internal/executor"
	"io"
)

type Interpreter struct {
	Engine *executor.Engine
}

func NewInterpreter(reader io.Reader, writer io.Writer) *Interpreter {
	return &Interpreter{
		Engine: executor.NewEngine(reader, writer),
	}
}

func (i *Interpreter) Run(program []ast.Statement) error {
	for _, stmt := range program {
		if err := i.Engine.Execute(stmt); err != nil {
			return err
		}
	}
	return nil
}
