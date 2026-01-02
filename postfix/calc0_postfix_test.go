package postfix

import (
	"calcPlus/internal/parser"
	"calcPlus/internal/testHelper"
	"github.com/antlr4-go/antlr/v4"
	"github.com/stretchr/testify/assert"
	"testing"
)

var postfixTestCases = map[string]testHelper.PostfixTestCase{
	"just number": {
		Code:    "1",
		Postfix: "1 ",
	},
	"operator precedence": {
		Code:    "1+2*3",
		Postfix: "1 2 3 * + ",
	},
	"parenthesis": {
		Code:    "(1+2)*3",
		Postfix: "1 2 + 3 * ",
	},
	"complex expression": {
		Code:    "10 + 2 * (5 - 9 / 3)",
		Postfix: "10 2 5 9 3 / - * + ",
	},
}

func TestCalc0_Listener(t *testing.T) {
	for name, tc := range postfixTestCases {
		t.Run(name, func(t *testing.T) {
			input := antlr.NewInputStream(tc.Code)
			lexer := parser.NewCalcPlusLexer(input)
			stream := antlr.NewCommonTokenStream(lexer, 0)
			p := parser.NewCalcPlusParser(stream)

			postfixPrinter := &Calc0PostfixL{}
			antlr.ParseTreeWalkerDefault.Walk(postfixPrinter, p.Expr())
			assert.Equal(t, tc.Postfix, postfixPrinter.PostfixExpr)
		})
	}
}

func TestCalc0_Visitor(t *testing.T) {
	for name, tc := range postfixTestCases {
		t.Run(name, func(t *testing.T) {
			input := antlr.NewInputStream(tc.Code)
			lexer := parser.NewCalcPlusLexer(input)
			stream := antlr.NewCommonTokenStream(lexer, 0)
			p := parser.NewCalcPlusParser(stream)

			tree := p.Expr()
			postfixPrinter := &Calc0PostfixV{}
			result := postfixPrinter.Visit(tree).(string)
			assert.Equal(t, tc.Postfix, result)
		})
	}
}
