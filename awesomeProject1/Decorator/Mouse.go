package Decorator

import "fmt"

type Mouse struct {
	hardware Hardware
}

func (S *Mouse) GetDescription() string {
	return fmt.Sprintf("%v, %v", S.hardware.GetDescription(), "Logitech G305")
}

func (S *Mouse) Cost() float64 {
	return S.hardware.Cost() + 50
}

func NewMouse(hardware Hardware) *Mouse {
	M := new(Mouse)
	M.hardware = hardware
	return M
}
