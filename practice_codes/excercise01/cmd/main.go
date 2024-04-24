package main

import (
	"excercise01/acura"
	"excercise01/ford"
	"excercise01/honda"
	"excercise01/tesla"
	"excercise01/vehicle"
	"fmt"
)

func main() {
	myGarage := make([]vehicle.Vehicle, 4)
	myGarage[0], _ = ford.NewExplorer(1908)
	myGarage[1], _ = honda.NewPilot(2010)
	myGarage[2], _ = acura.NewTSX(2006)
	myGarage[3], _ = tesla.NewModel3(2018)
	for _, vehicle := range myGarage {
		fmt.Println(vehicle)
	}
}
