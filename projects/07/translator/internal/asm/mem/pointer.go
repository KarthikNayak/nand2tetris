package mem

import (
	"fmt"
	"translator/internal/lexer"
)

type Pointer struct {
	cmd lexer.CMD
	val int
}

func (p *Pointer) Convert() ([]string, error) {
	switch p.cmd {
	case lexer.POP:
		return p.popPointer(), nil
	case lexer.PUSH:
		return p.pushPointer(), nil
	}
	return []string{}, fmt.Errorf("illegal cmd pointer passed: %s", p.cmd)
}

func NewPointer(cmd lexer.CMD, filename string, val int) *Pointer {
	return &Pointer{
		cmd: cmd,
		val: val,
	}
}

func (p *Pointer) popPointer() []string {
	v := "THIS"
	if p.val == 1 {
		v = "THAT"
	}

	return []string{
		"@SP",
		"AM = M - 1",
		"D = M",

		"@" + v,
		"M = D",
	}
}

func (p *Pointer) pushPointer() []string {
	v := "THIS"
	if p.val == 1 {
		v = "THAT"
	}

	return []string{
		"@" + v,
		"D = A",

		"@SP",
		"AM = M + 1",
		"A = A - 1",
		"M = D",
	}
}
