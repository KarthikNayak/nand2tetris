package mem

import (
	"fmt"
	"translator/internal/lexer"
)

type Constant struct {
	cmd lexer.CMD
	val int
}

func (c *Constant) Convert() ([]string, error) {
	switch c.cmd {
	case lexer.POP:
		return c.popConstant()
	case lexer.PUSH:
		return c.pushConstant(), nil
	}
	return []string{}, fmt.Errorf("illegal cmd constant passed: %s", c.cmd)
}

func NewConstant(cmd lexer.CMD, val int) *Constant {
	return &Constant{
		cmd: cmd,
		val: val,
	}
}

func (c *Constant) pushConstant() []string {
	return []string{
		fmt.Sprintf("@%d", c.val),
		"D = A",

		"@SP",
		"AM = M + 1",
		"A = A - 1",
		"M = D",
	}
}

func (c *Constant) popConstant() ([]string, error) {
	return []string{}, fmt.Errorf("constant stack doesn't support pop")
}
