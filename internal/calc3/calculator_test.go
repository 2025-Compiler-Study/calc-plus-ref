package calc3

import (
	"bytes"
	"calcPlus/internal/parser"
	"github.com/antlr4-go/antlr/v4"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestIO(t *testing.T) {
	type testCase struct {
		code         string
		userInput    string
		expectOutput string
	}

	testCases := map[string]testCase{
		"hello world": {
			code:         "write(1);",
			userInput:    "",
			expectOutput: "1",
		},
		"var write": {
			code:         "a = 2; write(a);",
			userInput:    "",
			expectOutput: "2",
		},
		"multiple writes": {
			code:         "a = 2; write(a); write(3);",
			userInput:    "",
			expectOutput: "2\n3",
		},
		"echo": {
			code:         "a = read(); write(a);",
			userInput:    "10",
			expectOutput: "10",
		},
		"multiple echo": {
			code:         "a = read(); write(a); b = read(); write(b);",
			userInput:    "20\n40",
			expectOutput: "20\n40",
		},
		"wrong input": {
			code:         "a = read(); write(a);",
			userInput:    "wrong",
			expectOutput: "0",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			is := antlr.NewInputStream(tc.code)
			lexer := parser.NewCalcPlusLexer(is)
			stream := antlr.NewCommonTokenStream(lexer, 0)
			p := parser.NewCalcPlusParser(stream)

			tree := p.Calc3()
			stdin, stdout := bytes.NewBuffer([]byte(tc.userInput)), bytes.NewBuffer([]byte{})
			calculator := NewCalculator(stdin, stdout)
			calculator.Visit(tree)

			expect := strings.TrimSuffix(tc.expectOutput, "\n")
			actual := strings.TrimSuffix(stdout.String(), "\n")
			assert.Equal(t, expect, actual)
		})
	}
}
