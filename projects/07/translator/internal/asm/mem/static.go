package mem

import (
	"fmt"
	"translator/internal/lexer"
)

type Static struct {
	cmd      lexer.CMD
	filename string
	val      int
}

func NewStatic(cmd lexer.CMD, filename string, val int) *Static {
	return &Static{
		cmd:      cmd,
		filename: filename,
		val:      val,
	}
}

func (s *Static) Convert() ([]string, error) {
	switch s.cmd {
	case lexer.POP:
		return s.popStatic(), nil
	case lexer.PUSH:
		return s.pushStatic(), nil
	}
	return []string{}, fmt.Errorf("illegal cmd static passed: %s", s.cmd)
}

func (s *Static) popStatic() []string {
	return []string{
		"@SP",
		"AM = M - 1",
		"D = M",

		fmt.Sprintf("@%s.%d", s.filename, s.val),
		"M = D",
	}
}

func (s *Static) pushStatic() []string {
	return []string{
		fmt.Sprintf("@%s.%d", s.filename, s.val),
		"D = M",

		"@SP",
		"AM = M + 1",
		"A = A - 1",
		"M = D",
	}
}
