package calc1

import (
	"calcPlus/internal/parser"
	"github.com/antlr4-go/antlr/v4"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestCalculator(t *testing.T) {
	type testCase struct {
		code string
		vars map[string]int
	}
	testCases := map[string]testCase{
		"simple assignment": {
			code: "a = 1;",
			vars: map[string]int{"a": 1},
		},
		"complex expression": {
			code: "a = 1 + 2 * 3;",
			vars: map[string]int{"a": 7},
		},
		"multiple assignments": {
			code: "a = 1; b = 2;",
			vars: map[string]int{"a": 1, "b": 2},
		},
		"variable use": {
			code: "a = 1; b = a + 2;",
			vars: map[string]int{"a": 1, "b": 3},
		},
		"multiple variable use": {
			code: "a = 1; b = 2; c = a + b;",
			vars: map[string]int{"a": 1, "b": 2, "c": 3},
		},
		"undefined variable": {
			code: "a=b+1;",
			vars: map[string]int{"a": 1},
		},
		"variable with some names": {
			code: "aa = 1; BB = aa + 2",
			vars: map[string]int{"aa": 1, "BB": 3},
		},
		"multiline": {
			code: strings.Join([]string{
				"a = 1;",
				"b = a + 2;",
				"c = b * 3;",
				"a = a + 1;",
				"d = (5 - e) * 2;",
			}, "\n"),
			vars: map[string]int{"a": 2, "b": 3, "c": 9, "d": 10},
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			input := antlr.NewInputStream(tc.code)
			lexer := parser.NewCalcPlusLexer(input)
			stream := antlr.NewCommonTokenStream(lexer, 0)
			p := parser.NewCalcPlusParser(stream)

			tree := p.Calc1()
			calculator := NewCalculator()
			calculator.Visit(tree)

			for n, v := range tc.vars {
				assert.Equal(t, v, calculator.Variables[n])
			}
		})
	}
}
