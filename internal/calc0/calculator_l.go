package calc0

import (
	"calcPlus/internal/parser"
	"github.com/antlr4-go/antlr/v4"
	"strconv"
)

type CalculatorL struct {
	parser.BaseCalcPlusListener
	Stack []int
}

func (c *CalculatorL) push(value int) {
	c.Stack = append(c.Stack, value)
}

func (c *CalculatorL) pop() int {
	last := len(c.Stack) - 1
	value := c.Stack[last]
	c.Stack = c.Stack[:last]
	return value
}

func (c *CalculatorL) Result() int {
	return c.Stack[0]
}

func (c *CalculatorL) ExitInt(ctx *parser.IntContext) {
	value, _ := strconv.Atoi(ctx.GetText())
	c.push(value)
}

func (c *CalculatorL) ExitAddSub(ctx *parser.AddSubContext) {
	right := c.pop()
	left := c.pop()

	op := ctx.GetChild(1).(antlr.TerminalNode).GetText()

	if op == "+" {
		c.push(left + right)
	} else {
		c.push(left - right)
	}
}

func (c *CalculatorL) ExitMulDiv(ctx *parser.MulDivContext) {
	right := c.pop()
	left := c.pop()

	op := ctx.GetChild(1).(antlr.TerminalNode).GetText()

	if op == "*" {
		c.push(left * right)
	} else {
		c.push(left / right)
	}
}
