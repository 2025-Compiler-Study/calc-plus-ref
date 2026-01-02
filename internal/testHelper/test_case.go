package testHelper

type ExprTestCase struct {
	Code   string
	Result int
}

type PostfixTestCase struct {
	Code    string
	Postfix string
}

type ErrorTestCase struct {
	Code string
}

type VariableTestCase struct {
	Code      string
	Variables map[string]int
}

type IOTestCase struct {
	Code         string
	UserInput    string
	ExpectOutput string
}
