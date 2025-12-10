package calc2

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
		"if condition": {
			code: strings.Join([]string{
				"a = 1;",
				"if (a >= 1) {",
				"    b = a + 2;",
				"    c = b * 3;",
				"} else {",
				"    b = a;",
				"    c = b;",
				"}",
				"a = a + 1;",
				"d = (5 - e) * 2;",
			}, "\n"),
			vars: map[string]int{"a": 2, "b": 3, "c": 9, "d": 10},
		},
		"if without else": {
			code: "if (a == 0) { b = 2; }",
			vars: map[string]int{"b": 2},
		},
		"nested if": {
			code: "if (a <= 0) { if (b < 0) { } else { c = 3; } }",
			vars: map[string]int{"c": 3},
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			input := antlr.NewInputStream(tc.code)
			lexer := parser.NewCalcPlusLexer(input)
			stream := antlr.NewCommonTokenStream(lexer, 0)
			p := parser.NewCalcPlusParser(stream)

			tree := p.Calc2()
			calculator := NewCalculator()
			calculator.Visit(tree)

			for n, v := range tc.vars {
				assert.Equal(t, v, calculator.Variables[n])
			}
		})
	}
}
