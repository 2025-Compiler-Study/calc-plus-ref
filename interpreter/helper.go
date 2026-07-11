package interpreter

import (
	"calcPlus/internal/ast"
	"calcPlus/internal/parser"
	"io"

	"github.com/antlr4-go/antlr/v4"
)

func BuildAST(code string) *ast.Program {
	is := antlr.NewInputStream(code)
	lexer := parser.NewCalcPlusLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewCalcPlusParser(stream)

	tree := p.Program()
	return NewASTBuilder().Build(tree)
}

func RunInterpreter(program *ast.Program, reader io.Reader, writer io.Writer) {
	if err := NewInterpreter(reader, writer).Run(program); err != nil {
		panic(err)
	}
}

func RunCode(code string, reader io.Reader, writer io.Writer) {
	RunInterpreter(BuildAST(code), reader, writer)
}
