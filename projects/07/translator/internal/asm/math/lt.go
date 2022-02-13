package math

import "fmt"

type Lt struct {
	index int
}

func (l *Lt) Convert() ([]string, error) {
	return []string{
		"@SP",
		"AM = M - 1",
		"D = M",

		"@SP",
		"A = M - 1",
		"D = M - D",

		fmt.Sprintf("@TRUE%d", l.index),
		"D; JLT",

		"@SP",
		"A = M - 1",
		"M = 0",
		fmt.Sprintf("@END%d", l.index),
		"D; JMP",

		fmt.Sprintf("(TRUE%d)", l.index),
		"@SP",
		"A = M - 1",
		"M = -1",

		fmt.Sprintf("(END%d)", l.index),
	}, nil
}

func NewLt(index int) *Lt {
	return &Lt{index: index}
}
