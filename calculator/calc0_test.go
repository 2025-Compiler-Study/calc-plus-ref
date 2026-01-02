package calculator

import (
	"calcPlus/internal/parser"
	"calcPlus/internal/testHelper"
	"github.com/antlr4-go/antlr/v4"
	"github.com/stretchr/testify/assert"
	"testing"
)

var calc0TestCases = map[string]testHelper.ExprTestCase{
	"just number": {
		Code:   "1",
		Result: 1,
	},
	"operator precedence": {
		Code:   "1+2*3",
		Result: 7,
	},
	"parenthesis": {
		Code:   "(1+2)*3",
		Result: 9,
	},
	"complex expression": {
		Code:   "10 + 2 * (5 - 9 / 3)",
		Result: 14,
	},
}

func TestCalc0_Visitor(t *testing.T) {
	for name, tc := range calc0TestCases {
		t.Run(name, func(t *testing.T) {
			input := antlr.NewInputStream(tc.Code)
			lexer := parser.NewCalcPlusLexer(input)
			stream := antlr.NewCommonTokenStream(lexer, 0)
			p := parser.NewCalcPlusParser(stream)

			tree := p.Expr()
			calculator := &Calc0V{}
			result := calculator.Visit(tree).(int)

			assert.Equal(t, tc.Result, result)
		})
	}
}

func TestCalc0_Listener(t *testing.T) {
	for name, tc := range calc0TestCases {
		t.Run(name, func(t *testing.T) {
			input := antlr.NewInputStream(tc.Code)
			lexer := parser.NewCalcPlusLexer(input)
			stream := antlr.NewCommonTokenStream(lexer, 0)
			p := parser.NewCalcPlusParser(stream)

			calculator := &Calc0L{}
			antlr.ParseTreeWalkerDefault.Walk(calculator, p.Expr())

			assert.Equal(t, tc.Result, calculator.Result())
		})
	}
}
