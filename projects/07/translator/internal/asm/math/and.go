package math

type And struct{}

func (a *And) Convert() ([]string, error) {
	return []string{
		"@SP",
		"AM = M - 1",
		"D = M",
		"@SP",
		"A = M - 1",
		"M = D & M",
	}, nil
}

func NewAnd() *And {
	return &And{}
}
