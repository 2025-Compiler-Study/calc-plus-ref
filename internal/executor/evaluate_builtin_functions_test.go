package executor

import (
	"bytes"
	"calcPlus/internal/ast"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBuiltinReadCall_Evaluate(t *testing.T) {
	type testCase struct {
		input string
		want  int
		isErr bool
	}
	testCases := map[string]testCase{
		"0":                 {input: "0", want: 0, isErr: false},
		"123":               {input: "123", want: 123, isErr: false},
		"empty (err)":       {input: "", want: 0, isErr: true},
		"invalid (non-int)": {input: "abc", want: 0, isErr: true},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			stdin := bytes.NewBuffer([]byte(tc.input))
			engine := NewEngine(stdin, nil)
			ConfigurePreset(engine)

			readCall := ast.NewBuiltinReadCall()
			result, err := engine.Evaluate(readCall)
			if tc.isErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.want, result)
			}
		})
	}
}

func TestBuiltinWriteCall_Execute(t *testing.T) {
	type testCase struct {
		input ast.Expression
		want  string
	}
	testCases := map[string]testCase{
		"IntExpression": {
			input: ast.NewIntExpression(100),
			want:  "100",
		},
		"VarExpression": {
			input: ast.NewVarExpression("a"),
			want:  "10",
		},
		"BinaryOperator": {
			input: ast.NewBinaryOperator(
				ast.NewVarExpression("a"),
				"+",
				ast.NewVarExpression("b"),
			),
			want: "30",
		},
		"ComplexExpression": {
			input: ast.NewBinaryOperator(
				ast.NewIntExpression(1),
				"+",
				ast.NewBinaryOperator(
					ast.NewIntExpression(2),
					"*",
					ast.NewVarExpression("b"),
				),
			),
			want: "41",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			stdout := bytes.NewBuffer([]byte{})
			engine := NewEngine(nil, stdout)
			ConfigurePreset(engine)

			writeCall := ast.NewBuiltinWriteCall(tc.input)
			_, err := engine.Evaluate(writeCall)
			actual := strings.TrimSuffix(stdout.String(), "\n")
			assert.NoError(t, err)
			assert.Equal(t, tc.want, actual)
		})
	}
}
