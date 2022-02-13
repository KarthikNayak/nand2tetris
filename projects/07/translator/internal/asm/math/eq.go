package math

import "fmt"

type Eq struct {
	index int
}

func (e *Eq) Convert() ([]string, error) {
	return []string{
		"@SP",
		"AM = M - 1",
		"D = M",

		"@SP",
		"A = M - 1",
		"D = M - D",

		fmt.Sprintf("@TRUE%d", e.index),
		"D; JEQ",

		"@SP",
		"A = M - 1",
		"M = 0",
		fmt.Sprintf("@END%d", e.index),
		"D; JMP",

		fmt.Sprintf("(TRUE%d)", e.index),
		"@SP",
		"A = M - 1",
		"M = -1",

		fmt.Sprintf("(END%d)", e.index),
	}, nil
}

func NewEq(index int) *Eq {
	return &Eq{index: index}
}
