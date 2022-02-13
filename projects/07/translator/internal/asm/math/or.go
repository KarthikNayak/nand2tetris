package math

type Or struct{}

func (o *Or) Convert() ([]string, error) {
	return []string{
		"@SP",
		"AM = M - 1",
		"D = M",
		"@SP",
		"A = M - 1",
		"M = D | M",
	}, nil
}

func NewOr() *Or {
	return &Or{}
}
