package math

import "fmt"

type Gt struct {
	index int
}

func (g *Gt) Convert() ([]string, error) {
	return []string{
		"@SP",
		"AM = M - 1",
		"D = M",

		"@SP",
		"A = M - 1",
		"D = M - D",

		fmt.Sprintf("@TRUE%d", g.index),
		"D; JGT",

		"@SP",
		"A = M - 1",
		"M = 0",
		fmt.Sprintf("@END%d", g.index),
		"D; JMP",

		fmt.Sprintf("(TRUE%d)", g.index),
		"@SP",
		"A = M - 1",
		"M = -1",

		fmt.Sprintf("(END%d)", g.index),
	}, nil
}

func NewGt(index int) *Gt {
	return &Gt{index: index}
}
