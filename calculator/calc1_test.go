package calculator

import (
	"calcPlus/internal/parser"
	"calcPlus/internal/testHelper"
	"github.com/antlr4-go/antlr/v4"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestCalc1(t *testing.T) {
	testCases := map[string]testHelper.VariableTestCase{
		"simple assignment": {
			Code:      "a = 1;",
			Variables: map[string]int{"a": 1},
		},
		"complex expression": {
			Code:      "a = 1 + 2 * 3;",
			Variables: map[string]int{"a": 7},
		},
		"multiple assignments": {
			Code:      "a = 1; b = 2;",
			Variables: map[string]int{"a": 1, "b": 2},
		},
		"variable use": {
			Code:      "a = 1; b = a + 2;",
			Variables: map[string]int{"a": 1, "b": 3},
		},
		"multiple variable use": {
			Code:      "a = 1; b = 2; c = a + b;",
			Variables: map[string]int{"a": 1, "b": 2, "c": 3},
		},
		"undefined variable": {
			Code:      "a=b+1;",
			Variables: map[string]int{"a": 1},
		},
		"variable with some names": {
			Code:      "aa = 1; BB = aa + 2",
			Variables: map[string]int{"aa": 1, "BB": 3},
		},
		"multiline": {
			Code: strings.Join([]string{
				"a = 1;",
				"b = a + 2;",
				"c = b * 3;",
				"a = a + 1;",
				"d = (5 - e) * 2;",
			}, "\n"),
			Variables: map[string]int{"a": 2, "b": 3, "c": 9, "d": 10},
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
