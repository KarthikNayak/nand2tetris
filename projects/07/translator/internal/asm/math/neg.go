package math

type Negate struct{}

func (n *Negate) Convert() ([]string, error) {
	return []string{
		"@SP",
		"A = M - 1",
		"M = -M",
	}, nil
}

func NewNegate() *Negate {
	return &Negate{}
}
