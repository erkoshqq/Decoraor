package main

import (
	"awesomeProject1/Decorator"
	"fmt"
)

func main() {
	var Lenovo Decorator.Hardware = Decorator.NewLenovo()
	Lenovo = Decorator.NewMouse(Lenovo)
	Lenovo = Decorator.NewMechanicalKeyboard(Lenovo)
	Lenovo = Decorator.NewHeadPhones(Lenovo)
	fmt.Printf("%v = $ %v ", Lenovo.GetDescription(), Lenovo.Cost())
}
