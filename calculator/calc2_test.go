package calculator

import (
	"calcPlus/internal/parser"
	"calcPlus/internal/testHelper"
	"github.com/antlr4-go/antlr/v4"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestCalc2(t *testing.T) {
	testCases := map[string]testHelper.VariableTestCase{
		"if condition": {
			Code: strings.Join([]string{
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
			Variables: map[string]int{"a": 2, "b": 3, "c": 9, "d": 10},
		},
		"if without else": {
			Code:      "if (a == 0) { b = 2; }",
			Variables: map[string]int{"b": 2},
		},
		"nested if": {
			Code:      "if (a <= 0) { if (b < 0) { } else { c = 3; } }",
			Variables: map[string]int{"c": 3},
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			input := antlr.NewInputStream(tc.Code)
			lexer := parser.NewCalcPlusLexer(input)
			stream := antlr.NewCommonTokenStream(lexer, 0)
			p := parser.NewCalcPlusParser(stream)

			tree := p.Program()
			calculator := NewCalculator()
			calculator.Visit(tree)

			for n, v := range tc.Variables {
				assert.Equal(t, v, calculator.Variables[n])
			}
		})
	}
}
