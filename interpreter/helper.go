package interpreter

import (
	"calcPlus/internal/parser"
	"github.com/antlr4-go/antlr/v4"
	"io"
)

func RunCalc3Interpreter(code string, reader io.Reader, writer io.Writer) {
	is := antlr.NewInputStream(code)
	lexer := parser.NewCalcPlusLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewCalcPlusParser(stream)

	tree := p.Program()
	calculator := NewCalc3Interpreter(reader, writer)
	calculator.Visit(tree)
}

func RunCalc4Interpreter(code string, reader io.Reader, writer io.Writer) {
	is := antlr.NewInputStream(code)
	lexer := parser.NewCalcPlusLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewCalcPlusParser(stream)

	tree := p.Program()
	calculator := NewCalc4Interpreter(reader, writer)
	calculator.Visit(tree)
}
