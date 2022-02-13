package mem

import (
	"fmt"
	"translator/internal/lexer"
)

type Temp struct {
	cmd lexer.CMD
	val int
}

func (t *Temp) Convert() ([]string, error) {
	switch t.cmd {
	case lexer.POP:
		return t.popTemp(), nil
	case lexer.PUSH:
		return t.pushTemp(), nil
	}
	return []string{}, fmt.Errorf("illegal cmd temp passed: %s", t.cmd)
}

func NewTemp(cmd lexer.CMD, val int) *Temp {
	return &Temp{
		cmd: cmd,
		val: val + 5,
	}
}

func (t *Temp) popTemp() []string {
	return []string{
		"@SP",
		"AM = M - 1",
		"D = M",

		fmt.Sprintf("@%d", t.val),
		"M = D",
	}
}

func (t *Temp) pushTemp() []string {
	return []string{
		fmt.Sprintf("@%d", t.val),
		"D = M",

		"@SP",
		"AM = M + 1",
		"A = A - 1",
		"M = D",
	}
}
