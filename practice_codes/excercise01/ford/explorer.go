package ford

import (
	"errors"
	"excercise01/vehicle"
	"fmt"
	"time"
)

type Explorer struct {
	vehicleData
}

type vehicleData struct {
	year  uint16
	seats uint8
}

func NewExplorer(year uint16) (e *Explorer, err error) {
	//validate year
	currentYear := uint16(time.Now().Year())
	if year < 1910 || year > currentYear {
		return nil, errors.New("invalid year for Ford Explorer")
	}
	e = &Explorer{}
	e.year = year
	e.seats = 5
	return e, nil
}

func (e *Explorer) Model() string {
	return "Explorer"
}

func (e *Explorer) Make() string {
	return "ford"
}

func (e *Explorer) Year() uint16 {
	if e == nil {
		return 0
	}
	return e.Year()
}

func (e *Explorer) SeatingCap() uint8 {
	if e == nil {
		return 5
	}
	return e.seats
}

func (e *Explorer) Type() int {
	return vehicle.FOUR_DOOR_PICKUP
}

func (e *Explorer) PowerTrain() string {
	return "gas"
}

func (e *Explorer) String() string {
	var s = fmt.Sprintf(`Model: %v
	Make: %v
	Year: %v
	Type: %v
	Seats: %v,
	Power Train: %v`,
		e.Model(), e.Make(), e.year, vehicle.Type(e.Type()), e.seats, e.PowerTrain(),
	)

	return s
}
