package Decorator

type Lenovo struct {
	description string
	cost        float64
}

func (S *Lenovo) GetDescription() string {
	return S.description
}

func (S *Lenovo) Cost() float64 {
	return S.cost
}

func NewLenovo() *Lenovo {
	M := new(Lenovo)
	M.cost = 1500
	M.description = "Lenovo Legion 5 Pro"
	return M
}
