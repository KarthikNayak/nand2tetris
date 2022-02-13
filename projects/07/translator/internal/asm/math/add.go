package math

type Add struct{}

func (a *Add) Convert() ([]string, error) {
	return []string{
		"@SP",
		"AM = M - 1",
		"D = M",
		"@SP",
		"A = M - 1",
		"M = D + M",
	}, nil
}

func NewAdd() *Add {
	return &Add{}
}
