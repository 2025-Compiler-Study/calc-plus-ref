package interpreter

import (
	"bytes"
	"calcPlus/internal/ast"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildASTBuildsEveryFunctionDefinition(t *testing.T) {
	code := `
func int add(int left, int right) {
	return left + right;
}

func main() {
	int result;
	result = add(1, 2);
	return;
}
`

	program := BuildAST(code)
	if assert.Len(t, program.Functions, 2) {
		add := program.Functions[0]
		assert.Equal(t, "add", add.Name)
		assert.True(t, add.ReturnsInt)
		assert.Equal(t, []ast.Parameter{{Name: "left"}, {Name: "right"}}, add.Parameters)
		if assert.Len(t, *add.Body, 1) {
			returnStatement, ok := (*add.Body)[0].(*ast.Return)
			if assert.True(t, ok) {
				_, ok = returnStatement.Value.(*ast.BinaryOperator)
				assert.True(t, ok)
			}
		}

		main := program.Functions[1]
		assert.Equal(t, "main", main.Name)
		assert.False(t, main.ReturnsInt)
		assert.Empty(t, main.Parameters)
		if assert.Len(t, *main.Body, 3) {
			assignment := (*main.Body)[1].(*ast.Assignment)
			call, ok := assignment.Value.(*ast.FunctionCall)
			if assert.True(t, ok) {
				assert.Equal(t, "add", call.Name)
				assert.Len(t, call.Arguments, 2)
			}

			returnStatement := (*main.Body)[2].(*ast.Return)
			assert.Nil(t, returnStatement.Value)
		}
	}
}

func TestBuildASTAndRunInterpreter(t *testing.T) {
	code := `
func int main() {
	int a, b;
	a = 1;
	b = read();
	write(a + b);
}
`
	stdin := bytes.NewBufferString("2\n")
	stdout := bytes.NewBuffer(nil)

	program := BuildAST(code)
	RunInterpreter(program, stdin, stdout)

	assert.Equal(t, "3", strings.TrimSuffix(stdout.String(), "\n"))
}

func TestRunCode(t *testing.T) {
	code := `
func int main() {
	int a;
	a = read();
	write(a + 1);
}
`
	stdin := bytes.NewBufferString("2\n")
	stdout := bytes.NewBuffer(nil)

	RunCode(code, stdin, stdout)

	assert.Equal(t, "3", strings.TrimSuffix(stdout.String(), "\n"))
}
