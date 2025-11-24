package calc0

import (
	"calcPlus/internal/parser"
	"github.com/antlr4-go/antlr/v4"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculatorL(t *testing.T) {
	type testCase struct {
		input string
		want  int
	}
	testCases := map[string]testCase{
		"just number": {
			input: "1",
			want:  1,
		},
		"operator precedence": {
			input: "1+2*3",
			want:  7,
		},
		"parenthesis": {
			input: "(1+2)*3",
			want:  9,
		},
		"complex expression": {
			input: "10 + 2 * (5 - 9 / 3)",
			want:  14,
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			c := &CalculatorL{}

			input := antlr.NewInputStream(tc.input)
			lexer := parser.NewCalcPlusLexer(input)
			stream := antlr.NewCommonTokenStream(lexer, 0)
			p := parser.NewCalcPlusParser(stream)

			antlr.ParseTreeWalkerDefault.Walk(c, p.Prog())

			assert.Equal(t, tc.want, c.Result())
		})
	}
}
