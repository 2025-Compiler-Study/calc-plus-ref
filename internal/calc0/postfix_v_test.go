package calc0

import (
	"calcPlus/internal/parser"
	"github.com/antlr4-go/antlr/v4"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPostfixV(t *testing.T) {
	type testCase struct {
		input string
		want  string
	}
	testCases := map[string]testCase{
		"just number": {
			input: "1",
			want:  "1 ",
		},
		"operator precedence": {
			input: "1+2*3",
			want:  "1 2 3 * + ",
		},
		"parenthesis": {
			input: "(1+2)*3",
			want:  "1 2 + 3 * ",
		},
		"complex expression": {
			input: "10 + 2 * (5 - 9 / 3)",
			want:  "10 2 5 9 3 / - * + ",
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			postfixPrinter := &PostfixV{}

			input := antlr.NewInputStream(tc.input)
			lexer := parser.NewCalcPlusLexer(input)
			stream := antlr.NewCommonTokenStream(lexer, 0)
			p := parser.NewCalcPlusParser(stream)

			tree := p.Prog()
			result := postfixPrinter.Visit(tree).(string)

			assert.Equal(t, tc.want, result)
		})
	}
}
