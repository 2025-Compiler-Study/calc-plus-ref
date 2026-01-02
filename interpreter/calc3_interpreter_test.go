package interpreter

import (
	"bytes"
	"calcPlus/internal/testHelper"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestCalc3_IO_Interaction(t *testing.T) {
	testCases := map[string]testHelper.IOTestCase{
		"bulk_rw": {
			Code: strings.Join([]string{
				"a = read();",
				"b = read();",
				"c = read();",
				"write(a);",
				"write(b);",
				"write(c);",
			}, "\n"),
			UserInput:    "20\n40\n6",
			ExpectOutput: "20\n40\n6",
		},
		"echo": {
			Code:         "a = read(); write(a);",
			UserInput:    "10",
			ExpectOutput: "10",
		},
		"helloworld": {
			Code:         "write(1);",
			UserInput:    "",
			ExpectOutput: "1",
		},
		"multi_echo": {
			Code:         "a = read(); write(a); b = read(); write(b);",
			UserInput:    "20\n40",
			ExpectOutput: "20\n40",
		},
		"multi_write": {
			Code:         "a = 2; write(a); write(3);",
			UserInput:    "",
			ExpectOutput: "2\n3",
		},
		"var_write": {
			Code:         "a = 2; write(a);",
			UserInput:    "",
			ExpectOutput: "2",
		},
		"wrong input": {
			Code:         "a = read(); write(a);",
			UserInput:    "wrong",
			ExpectOutput: "0",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			stdin, stdout := bytes.NewBuffer([]byte(tc.UserInput)), bytes.NewBuffer([]byte{})

			RunCalc3Interpreter(tc.Code, stdin, stdout)

			expect := strings.TrimSuffix(tc.ExpectOutput, "\n")
			actual := strings.TrimSuffix(stdout.String(), "\n")
			assert.Equal(t, expect, actual)
		})
	}
}

func TestCalc3_ComplexCode(t *testing.T) {
	minMaxCode := strings.Join([]string{
		"a = read();",
		"b = read();",
		"if (a > b) {",
		"    write(a);",
		"} else {",
		"    write(b);",
		"}",
	}, "\n")
	oddEvenCode := strings.Join([]string{
		"a = read();",
		"if (a / 2 * 2 == a) {",
		"    write(0);",
		"} else {",
		"    write(1);",
		"}",
	}, "\n")

	testCases := map[string]testHelper.IOTestCase{
		"max (right case)": {
			Code:         minMaxCode,
			UserInput:    "10\n20",
			ExpectOutput: "20",
		},
		"max (left case)": {
			Code:         minMaxCode,
			UserInput:    "20\n10",
			ExpectOutput: "20",
		},
		"odd_even (even case)": {
			Code:         oddEvenCode,
			UserInput:    "10",
			ExpectOutput: "0",
		},
		"odd_even (odd case)": {
			Code:         oddEvenCode,
			UserInput:    "9",
			ExpectOutput: "1",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			stdin, stdout := bytes.NewBuffer([]byte(tc.UserInput)), bytes.NewBuffer([]byte{})

			RunCalc3Interpreter(tc.Code, stdin, stdout)

			expect := strings.TrimSuffix(tc.ExpectOutput, "\n")
			actual := strings.TrimSuffix(stdout.String(), "\n")
			assert.Equal(t, expect, actual)
		})
	}
}
