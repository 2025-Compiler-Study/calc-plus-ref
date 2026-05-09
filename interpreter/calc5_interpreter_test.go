package interpreter

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunCalc5Interpreter(t *testing.T) {
	nestedScopeAndConditionCode := strings.Join([]string{
		"int a, b, result;",
		"a = read();",
		"b = read();",
		"result = 0;",
		"if (a > b) {",
		"    int diff;",
		"    diff = a - b;",
		"    if (diff >= 10) {",
		"        result = diff * 2;",
		"    } else {",
		"        result = diff + 100;",
		"    }",
		"} else {",
		"    int sum;",
		"    sum = a + b;",
		"    result = sum / 2;",
		"}",
		"write(result);",
	}, "\n")

	testCases := map[string]struct {
		code         string
		userInput    string
		expectOutput string
	}{
		"read write block": {
			code: strings.Join([]string{
				"int a, b;",
				"a = 1;",
				"b = read();",
				"{",
				"    int c;",
				"    c = a + b;",
				"    write(c);",
				"}",
				"b = 0;",
			}, "\n"),
			userInput:    "2",
			expectOutput: "3",
		},
		"if without else false branch does nothing": {
			code: strings.Join([]string{
				"int a;",
				"a = 1;",
				"if (a == 2) {",
				"    write(10);",
				"}",
				"write(a);",
			}, "\n"),
			userInput:    "",
			expectOutput: "1",
		},
		"if else": {
			code: strings.Join([]string{
				"int a;",
				"a = 3;",
				"if (a > 5) {",
				"    write(10);",
				"} else {",
				"    write(a + 1);",
				"}",
			}, "\n"),
			userInput:    "",
			expectOutput: "4",
		},
		"inner shadow does not leak": {
			code: strings.Join([]string{
				"int a;",
				"a = 10;",
				"{",
				"    int a;",
				"    a = 20;",
				"    write(a);",
				"}",
				"write(a);",
			}, "\n"),
			userInput:    "",
			expectOutput: "20\n10",
		},
		"nested scope and condition": {
			code:         nestedScopeAndConditionCode,
			userInput:    "20\n7",
			expectOutput: "26",
		},
		"nested scope and condition else path": {
			code:         nestedScopeAndConditionCode,
			userInput:    "4\n12",
			expectOutput: "8",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			stdin := bytes.NewBufferString(tc.userInput)
			stdout := bytes.NewBuffer(nil)

			RunCalc5Interpreter(tc.code, stdin, stdout)

			actual := strings.TrimSuffix(stdout.String(), "\n")
			assert.Equal(t, tc.expectOutput, actual)
		})
	}
}

func TestRunCalc5InterpreterPanic(t *testing.T) {
	testCases := map[string]string{
		"undefined variable read":               "write(a);",
		"undefined variable write":              "a = 1;",
		"re-declare variable in separate lines": "int a; int a;",
		"re-declare in single line":             "int a, a;",
		"inner variable does not escape": strings.Join([]string{
			"{",
			"    int a;",
			"    a = 10;",
			"}",
			"write(a);",
		}, "\n"),
	}

	for name, code := range testCases {
		t.Run(name, func(t *testing.T) {
			assert.Panics(t, func() {
				RunCalc5Interpreter(code, bytes.NewBuffer(nil), bytes.NewBuffer(nil))
			})
		})
	}
}
