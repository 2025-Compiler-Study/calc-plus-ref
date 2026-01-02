package calc4

import (
	"calcPlus/internal/parser"
	"github.com/antlr4-go/antlr/v4"
	"io"
)

func RunInterpreter(code string, reader io.Reader, writer io.Writer) {
	is := antlr.NewInputStream(code)
	lexer := parser.NewCalcPlusLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewCalcPlusParser(stream)

	tree := p.Calc4()
	calculator := NewInterpreter(reader, writer)
	calculator.Visit(tree)
}
