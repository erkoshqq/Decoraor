package Decorator

import "fmt"

type MechanicalKeyboard struct {
	hardware Hardware
}

func (S *MechanicalKeyboard) GetDescription() string {
	return fmt.Sprintf("%v, %v", S.hardware.GetDescription(), "Mechanical Keyboard Logitech G413")
}

func (S *MechanicalKeyboard) Cost() float64 {
	return S.hardware.Cost() + 100
}

func NewMechanicalKeyboard(hardware Hardware) *MechanicalKeyboard {
	M := new(MechanicalKeyboard)
	M.hardware = hardware
	return M
}
