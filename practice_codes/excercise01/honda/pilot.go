package honda

import (
	"errors"
	"excercise01/vehicle"
	"fmt"
	"time"
)

type (
	// Pilot is one of Honda's vehicle that implements vehicle.Vehicle
	Pilot struct {
		vehicleData
	}
	// vehicleData is private to package
	vehicleData struct {
		year  uint16
		seats uint8
	}
)

// NewPilot returns an initialized honda.Pilot value
func NewPilot(year uint16) (e *Pilot, err error) {
	// verify validity of year parameter
	currentYear := uint16(time.Now().Year())
	if year < 2006 || year > currentYear {
		return nil, errors.New("invalid year for Honda Pilot")
	}

	e = &Pilot{}
	e.year = year
	e.seats = 8
	return
}

func (e *Pilot) Model() string {
	return "Pilot"
}
func (e *Pilot) Make() string {
	return "Honda"
}
func (e *Pilot) Year() uint16 {
	if e == nil {
		return 0
	}
	return e.year
}
func (e *Pilot) SeatingCap() uint8 {
	if e == nil {
		return 8
	}
	return e.seats
}
func (e *Pilot) Type() int {
	return vehicle.FOUR_DOOR_SUV
}

func (e *Pilot) PowerTrain() string {
	return "gas"
}

func (e *Pilot) String() string {
	var s = fmt.Sprintf(`Model: %v
	Make: %v
	Year: %v
	Type: %v
	Seats: %v,
	Power Train: %v`,
		e.Model(), e.Make(), e.Year(), vehicle.Type(e.Type()), e.SeatingCap(), e.PowerTrain())
	return s
}
