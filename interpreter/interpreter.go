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

func (i *Interpreter) Run(program *ast.Program) error {
	for _, function := range program.Functions {
		if function.Name != "main" {
			continue
		}
		for _, stmt := range *function.Body {
			if err := i.Engine.Execute(stmt); err != nil {
				return err
			}
		}
		return nil
	}
	return nil
}
