package mem

import (
	"fmt"
	"translator/internal/lexer"
)

type Location struct {
	cmd lexer.CMD
	loc string
	val int
}

func (l *Location) Convert() ([]string, error) {
	switch l.cmd {
	case lexer.POP:
		return l.popLocation(), nil
	case lexer.PUSH:
		return l.pushLocation(), nil
	}
	return []string{}, fmt.Errorf("illegal cmd location passed: %s", l.cmd)
}

func NewLocation(cmd lexer.CMD, loc string, val int) *Location {
	return &Location{
		cmd: cmd,
		val: val,
		loc: loc,
	}
}

func (l *Location) popLocation() []string {
	return []string{
		fmt.Sprintf("@%d", l.val),
		"D = A",

		"@" + l.loc,
		"D = D + M",

		"@R13",
		"M = D",

		"@SP",
		"AM = M - 1",
		"D = M",

		"@R13",
		"A = M",
		"M = D",
	}
}

func (l *Location) pushLocation() []string {
	return []string{
		fmt.Sprintf("@%d", l.val),
		"D = A",

		"@" + l.loc,
		"A = D + M",
		"D = M",

		"@SP",
		"AM = M + 1",
		"A = A - 1",
		"M = D",
	}
}
