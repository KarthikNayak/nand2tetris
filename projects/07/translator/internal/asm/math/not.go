package math

type Not struct{}

func (n *Not) Convert() ([]string, error) {
	return []string{
		"@SP",
		"A = M - 1",
		"M = !M",
	}, nil
}

func NewNot() *Not {
	return &Not{}
}
