package ast

import (
	"bytes"
	"calcPlus/internal/symbols"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"strings"
	"testing"
)

func TestBuiltinFunctions_String(t *testing.T) {
	t.Run("read", func(t *testing.T) {
		readCall := BuiltinReadCall{Reader: os.Stdin}
		fmt.Println(readCall.String())
	})

	t.Run("write", func(t *testing.T) {
		writeCall := BuiltinWriteCall{
			Writer:   os.Stdout,
			Argument: &IntExpression{Value: 42},
		}
		fmt.Println(writeCall.String())
	})
}

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

	symTable := symbols.NewSimpleTable[int]()
	presets := map[string]int{
		"a": 10,
		"b": 20,
	}
	for k, v := range presets {
		err := symTable.Register(k)
		require.NoError(t, err)
		err = symTable.SetSymbol(k, v)
		require.NoError(t, err)
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			stdin := bytes.NewBuffer([]byte(tc.input))

			readCall := BuiltinReadCall{Reader: stdin}
			result, err := readCall.Evaluate(symTable)
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
	symTable := symbols.NewSimpleTable[int]()
	presets := map[string]int{
		"a": 10,
		"b": 20,
	}
	for k, v := range presets {
		err := symTable.Register(k)
		require.NoError(t, err)
		err = symTable.SetSymbol(k, v)
		require.NoError(t, err)
	}

	type testCase struct {
		input Expression
		want  string
	}
	testCases := map[string]testCase{
		"IntExpression": {
			input: &IntExpression{Value: 100},
			want:  "100",
		},
		"VarExpression": {
			input: &VarExpression{Name: "a"},
			want:  "10",
		},
		"BinaryOperator": {
			input: &BinaryOperator{
				Left:     &VarExpression{Name: "a"},
				Operator: "+",
				Right:    &VarExpression{Name: "b"},
			},
			want: "30",
		},
		"ComplexExpression": {
			input: &BinaryOperator{
				Left:     &IntExpression{1},
				Operator: "+",
				Right: &BinaryOperator{
					Left:     &IntExpression{2},
					Operator: "*",
					Right:    &VarExpression{Name: "b"},
				},
			},
			want: "41",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			stdout := bytes.NewBuffer([]byte{})
			writeCall := BuiltinWriteCall{
				Writer:   stdout,
				Argument: tc.input,
			}
			_, err := writeCall.Evaluate(symTable)
			actual := strings.TrimSuffix(stdout.String(), "\n")
			assert.NoError(t, err)
			assert.Equal(t, tc.want, actual)
		})
	}
}
