package linter

import (
	"calcPlus/internal/parser"
	"github.com/antlr4-go/antlr/v4"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestLinter(t *testing.T) {
	type testCase struct {
		Code   string
		errors []TokenPosition
	}
	testCases := map[string]testCase{
		"no errors": {
			Code:   "a = 1;",
			errors: nil,
		},
		"not assigned variable": {
			Code: "a = b + 3;",
			errors: []TokenPosition{
				{Line: 1, Col: 4, Value: "b"},
			},
		},
		"just assigned": {
			Code: "a = a + 1;",
			errors: []TokenPosition{
				{Line: 1, Col: 4, Value: "a"},
			},
		},
		"multiline errors": {
			Code: strings.Join([]string{
				"a = 4 + b;",
				"c = c + 1;",
			}, "\n"),
			errors: []TokenPosition{
				{Line: 1, Col: 8, Value: "b"},
				{Line: 2, Col: 4, Value: "c"},
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			is := antlr.NewInputStream(tc.Code)
			lexer := parser.NewCalcPlusLexer(is)
			stream := antlr.NewCommonTokenStream(lexer, 0)
			p := parser.NewCalcPlusParser(stream)
			l := NewLinter()

			antlr.ParseTreeWalkerDefault.Walk(l, p.Program())

			assert.Equal(t, len(tc.errors), len(l.Errors))
			for i, e := range l.Errors {
				assert.Equal(t, tc.errors[i], e)
			}
		})
	}
}
