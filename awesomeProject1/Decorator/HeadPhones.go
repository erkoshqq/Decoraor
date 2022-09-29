package Decorator

import "fmt"

type HeadPhones struct {
	hardware Hardware
}

func (S *HeadPhones) GetDescription() string {
	return fmt.Sprintf("%v, %v", S.hardware.GetDescription(), "HyperX")
}

func (S *HeadPhones) Cost() float64 {
	return S.hardware.Cost() + 100
}

func NewHeadPhones(hardware Hardware) *HeadPhones {
	M := new(HeadPhones)
	M.hardware = hardware
	return M
}
