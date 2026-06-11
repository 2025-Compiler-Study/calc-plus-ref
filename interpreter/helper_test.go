package interpreter

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
