package calculator

import (
	"calcPlus/internal/parser"
	"github.com/antlr4-go/antlr/v4"
	"strconv"
)

type Calc0L struct {
	parser.BaseCalcPlusListener
	Stack []int
}

func (c *Calc0L) push(value int) {
	c.Stack = append(c.Stack, value)
}

func (c *Calc0L) pop() int {
	last := len(c.Stack) - 1
	value := c.Stack[last]
	c.Stack = c.Stack[:last]
	return value
}

func (c *Calc0L) Result() int {
	return c.Stack[0]
}

func (c *Calc0L) ExitInt(ctx *parser.IntContext) {
	value, _ := strconv.Atoi(ctx.GetText())
	c.push(value)
}

func (c *Calc0L) ExitAddSub(ctx *parser.AddSubContext) {
	right := c.pop()
	left := c.pop()

	op := ctx.GetChild(1).(antlr.TerminalNode).GetText()

	if op == "+" {
		c.push(left + right)
	} else {
		c.push(left - right)
	}
}

func (c *Calc0L) ExitMulDiv(ctx *parser.MulDivContext) {
	right := c.pop()
	left := c.pop()

	op := ctx.GetChild(1).(antlr.TerminalNode).GetText()

	if op == "*" {
		c.push(left * right)
	} else {
		c.push(left / right)
	}
}
