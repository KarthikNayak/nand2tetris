package math

type Subtract struct{}

func (s *Subtract) Convert() ([]string, error) {
	return []string{
		"@SP",
		"AM = M - 1",
		"D = M",
		"@SP",
		"A = M - 1",
		"M = M - D",
	}, nil
}

func NewSubtract() *Subtract {
	return &Subtract{}
}
