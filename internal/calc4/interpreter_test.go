package calc4

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestVariableDeclarePanic(t *testing.T) {
	type testCase struct {
		code         string
		userInput    string
		expectOutput string
	}

	testCases := map[string]testCase{
		"undefined variable read": {
			code:         "write(a);",
			userInput:    "",
			expectOutput: "",
		},
		"undefined variable write": {
			code:         "a = 1;",
			userInput:    "",
			expectOutput: "",
		},
		"re-declare variable in separate lines": {
			code:         "int a; int a;",
			userInput:    "",
			expectOutput: "",
		},
		"re-declare in single line": {
			code:         "int a, a;",
			userInput:    "",
			expectOutput: "",
		},
		"declared inner scope, use outer scope": {
			code:         "{ int a; a = 10; write(a); } write(a);",
			userInput:    "",
			expectOutput: "",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			stdin, stdout := bytes.NewBuffer([]byte(tc.userInput)), bytes.NewBuffer([]byte{})

			assert.Panics(t, func() {
				RunInterpreter(tc.code, stdin, stdout)

				expect := strings.TrimSuffix(tc.expectOutput, "\n")
				actual := strings.TrimSuffix(stdout.String(), "\n")
				assert.Equal(t, expect, actual)
			})
		})
	}
}

func TestVariableScope(t *testing.T) {
	type testCase struct {
		code         string
		userInput    string
		expectOutput string
	}

	testCases := map[string]testCase{
		"define outer scope, use inner scope": {
			code:         "int a; a = 10; { write(a); } write(a);",
			userInput:    "",
			expectOutput: "10\n10",
		},
		"define outer scope, change inner scope": {
			code:         "int a; { a = 20; } write(a);",
			userInput:    "",
			expectOutput: "20",
		},
		"shadow in inner scope does not affect outer scope": {
			code:         "int a; a = 10; { int a; a = 20; } write(a);",
			userInput:    "",
			expectOutput: "10",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			stdin, stdout := bytes.NewBuffer([]byte(tc.userInput)), bytes.NewBuffer([]byte{})

			RunInterpreter(tc.code, stdin, stdout)

			expect := strings.TrimSuffix(tc.expectOutput, "\n")
			actual := strings.TrimSuffix(stdout.String(), "\n")
			assert.Equal(t, expect, actual)
		})
	}
}
