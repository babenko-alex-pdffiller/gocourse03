package main

import (
	"fmt"
	"math/rand/v2"
)

type RodentType string

const (
	Mice       RodentType = "Mice"
	Rat        RodentType = "Rats"
	Squirrel   RodentType = "Squirrels"
	Beaver     RodentType = "Beaver"
	Vole       RodentType = "Vole"
	GuineaPigs RodentType = "GuineaPigs"
	Jerboas    RodentType = "Jerboas"
	Porcupine  RodentType = "Porcupine"
)

type Location struct {
	SectorID       SectorID
	IsSensorPassed bool
	Times          DayTime
}

type Rodent struct {
	ID        int
	Type      RodentType
	Morning   *Location
	Afternoon *Location
	Evening   *Location
}

func (r *Rodent) isMovingOnTime(dayTime DayTime) bool {
	switch dayTime {
	case Morning:
		return r.Morning != nil
	case Afternoon:
		return r.Afternoon != nil
	case Evening:
		return r.Evening != nil
	}

	return false
}

func (r *Rodent) moveToSector(sector Sector, dayTime DayTime) {
	switch dayTime {
	case Morning:
		r.Morning = &Location{sector.SectorID, sector.HasSensor, dayTime}
	case Afternoon:
		r.Afternoon = &Location{sector.SectorID, sector.HasSensor, dayTime}
	case Evening:
		r.Evening = &Location{sector.SectorID, sector.HasSensor, dayTime}
	}
}

func (r *Rodent) hasPassedThroughTheSensor() bool {
	return (r.Morning != nil && r.Morning.IsSensorPassed) ||
		(r.Afternoon != nil && r.Afternoon.IsSensorPassed) ||
		(r.Evening != nil && r.Evening.IsSensorPassed)
}

type DayTime string

const (
	Morning   DayTime = "Morning"
	Afternoon DayTime = "Afternoon"
	Evening   DayTime = "Evening"
)

type SectorID string

const (
	TopLeft     SectorID = "TopLeft"
	TopRight    SectorID = "TopRight"
	BottomLeft  SectorID = "BottomLeft"
	BottomRight SectorID = "BottomRight"
	Center      SectorID = "Center"
)

type SectorState struct {
	Times   DayTime
	Rodents []*Rodent
}

func (s *SectorState) addRodent(rodent *Rodent) {
	s.Rodents = append(s.Rodents, rodent)
}

type Sector struct {
	SectorID  SectorID
	HasSensor bool
	Morning   SectorState
	Afternoon SectorState
	Evening   SectorState
}

func (s *Sector) addState(state SectorState) {
	switch state.Times {
	case Morning:
		s.Morning = state
	case Afternoon:
		s.Afternoon = state
	case Evening:
		s.Evening = state
	}
}

func main() {
	rodents := []Rodent{
		{ID: 1, Type: Mice},
		{ID: 2, Type: Rat},
		{ID: 3, Type: Squirrel},
		{ID: 4, Type: Beaver},
		{ID: 5, Type: Vole},
		{ID: 6, Type: GuineaPigs},
		{ID: 7, Type: Jerboas},
		{ID: 8, Type: Porcupine},
	}

	sectors := []Sector{
		{SectorID: TopLeft, HasSensor: false, Morning: SectorState{Times: Morning}, Afternoon: SectorState{Times: Afternoon}, Evening: SectorState{Times: Evening}},
		{SectorID: TopRight, HasSensor: false, Morning: SectorState{Times: Morning}, Afternoon: SectorState{Times: Afternoon}, Evening: SectorState{Times: Evening}},
		{SectorID: BottomLeft, HasSensor: false, Morning: SectorState{Times: Morning}, Afternoon: SectorState{Times: Afternoon}, Evening: SectorState{Times: Evening}},
		{SectorID: BottomRight, HasSensor: false, Morning: SectorState{Times: Morning}, Afternoon: SectorState{Times: Afternoon}, Evening: SectorState{Times: Evening}},
		{SectorID: Center, HasSensor: true, Morning: SectorState{Times: Morning}, Afternoon: SectorState{Times: Afternoon}, Evening: SectorState{Times: Evening}},
	}

	dayTimes := [3]DayTime{
		Morning,
		Afternoon,
		Evening,
	}

	for _, dayTime := range dayTimes {
		fmt.Printf("Now is %s\n", dayTime)

		for is, _ := range sectors {
			sectorState := SectorState{Times: dayTime}
			for ir, _ := range rodents {
				if !rodents[ir].isMovingOnTime(dayTime) {
					isMovingToCurrentSector := rand.IntN(2) == 0
					if isMovingToCurrentSector {
						rodents[ir].moveToSector(sectors[is], dayTime)
						sectorState.addRodent(&rodents[ir])
						fmt.Printf("%s is moving to sector %s in %s \n", rodents[ir].Type, sectors[is].SectorID, dayTime)
					}
				}
			}

			sectors[is].addState(sectorState)
		}
	}
	fmt.Println("Daily report:")

	rodentId := rand.IntN(len(rodents)) + 1

	for i, rodent := range rodents {
		if rodent.ID == rodentId {
			printRodentLocation(rodent)
			rodents = append(rodents[:i], rodents[i+1:]...)
		}
	}
}
func printRodentLocation(rodent Rodent) {
	fmt.Printf("%s with ID %d has been in the morning in %s", rodent.Type, rodent.ID, getSectorID(rodent.Morning))
	fmt.Printf(", in the afternoon in %s", getSectorID(rodent.Afternoon))
	fmt.Printf(", in the evening in %s\n", getSectorID(rodent.Evening))

	if rodent.hasPassedThroughTheSensor() {
		fmt.Printf("passed through the sensor\n")
	}
}

func getSectorID(location *Location) SectorID {
	if location == nil {
		return "nil"
	}
	return location.SectorID
}
