package interpreter

import (
	"bytes"
	"calcPlus/internal/testHelper"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestVariableDeclarePanic(t *testing.T) {
	testCases := map[string]testHelper.IOTestCase{
		"undefined variable read": {
			Code:         "write(a);",
			UserInput:    "",
			ExpectOutput: "",
		},
		"undefined variable write": {
			Code:         "a = 1;",
			UserInput:    "",
			ExpectOutput: "",
		},
		"re-declare variable in separate lines": {
			Code:         "int a; int a;",
			UserInput:    "",
			ExpectOutput: "",
		},
		"re-declare in single line": {
			Code:         "int a, a;",
			UserInput:    "",
			ExpectOutput: "",
		},
		"declared inner scope, use outer scope": {
			Code:         "{ int a; a = 10; write(a); } write(a);",
			UserInput:    "",
			ExpectOutput: "",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			stdin, stdout := bytes.NewBuffer([]byte(tc.UserInput)), bytes.NewBuffer([]byte{})

			assert.Panics(t, func() {
				RunCalc4Interpreter(tc.Code, stdin, stdout)

				expect := strings.TrimSuffix(tc.ExpectOutput, "\n")
				actual := strings.TrimSuffix(stdout.String(), "\n")
				assert.Equal(t, expect, actual)
			})
		})
	}
}

func TestVariableScope(t *testing.T) {
	type testCase struct {
		Code         string
		UserInput    string
		ExpectOutput string
	}

	testCases := map[string]testCase{
		"define outer scope, use inner scope": {
			Code:         "int a; a = 10; { write(a); } write(a);",
			UserInput:    "",
			ExpectOutput: "10\n10",
		},
		"define outer scope, change inner scope": {
			Code:         "int a; { a = 20; } write(a);",
			UserInput:    "",
			ExpectOutput: "20",
		},
		"shadow in inner scope does not affect outer scope": {
			Code:         "int a; a = 10; { int a; a = 20; } write(a);",
			UserInput:    "",
			ExpectOutput: "10",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			stdin, stdout := bytes.NewBuffer([]byte(tc.UserInput)), bytes.NewBuffer([]byte{})

			RunCalc4Interpreter(tc.Code, stdin, stdout)

			expect := strings.TrimSuffix(tc.ExpectOutput, "\n")
			actual := strings.TrimSuffix(stdout.String(), "\n")
			assert.Equal(t, expect, actual)
		})
	}
}
