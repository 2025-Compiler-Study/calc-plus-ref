package executor

import (
	"bytes"
	"calcPlus/internal/ast"
	"calcPlus/internal/testHelper"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEngine_Execute(t *testing.T) {
	testCase := map[string]testHelper.IOTestCase{
		"1": testHelper.IOTestCase{
			UserInput:    "1",
			ExpectOutput: "2",
		},
		"2": testHelper.IOTestCase{
			UserInput:    "2",
			ExpectOutput: "3",
		},
	}

	/* Intended code
	int a, b;
	a = 1;
	b = read();
	{
	  int c;
	  c = a + b;
	  write(c);
	}
	b = 0;
	*/
	program := []ast.Statement{
		ast.NewDeclaration("a"),
		ast.NewDeclaration("b"),
		ast.NewAssignment("a", ast.NewIntExpression(1)),
		ast.NewAssignment("b", ast.NewBuiltinReadCall()),
		ast.NewBlockStatements(
			ast.NewDeclaration("c"),
			ast.NewAssignment(
				"c",
				ast.NewBinaryOperator(
					ast.NewVarExpression("a"),
					"+",
					ast.NewVarExpression("b"))),
			ast.NewAssignment(
				"_",
				ast.NewBuiltinWriteCall(ast.NewVarExpression("c")))),
		ast.NewAssignment("b", ast.NewIntExpression(0)),
	}

	for name, tc := range testCase {
		t.Run(name, func(t *testing.T) {
			stdin, stdout := bytes.NewBuffer([]byte(tc.UserInput)), bytes.NewBuffer([]byte{})
			engine := NewEngine(stdin, stdout)

			for _, statement := range program {
				err := engine.Execute(statement)
				assert.NoError(t, err)
			}
			actual := strings.TrimSuffix(stdout.String(), "\n")
			assert.Equal(t, tc.ExpectOutput, actual)
		})
	}

}
