package main

import "fmt"

// Mediator interface
type Mediator interface {
	CanArrive(train Train) bool
	NotifyAboutDeparture()
}

// Concrete Mediator
type StationManager struct {
	trainQueue []Train
}

func NewStationManager() *StationManager {
	return &StationManager{}
}

func (s *StationManager) CanArrive(train Train) bool {
	if len(s.trainQueue) == 0 {
		s.trainQueue = append(s.trainQueue, train)
		return true
	} else {
		fmt.Printf("%T is waiting for the platform.\n", train)
		return false
	}
}

func (s *StationManager) NotifyAboutDeparture() {
	if len(s.trainQueue) > 0 {
		departedTrain := s.trainQueue[0]
		s.trainQueue = s.trainQueue[1:]
		fmt.Printf("%T has departed.\n", departedTrain)
	}
	if len(s.trainQueue) > 0 {
		nextTrain := s.trainQueue[0]
		fmt.Printf("%T can now arrive.\n", nextTrain)
		nextTrain.Arrive()
	}
}

// Train interface
type Train interface {
	Arrive()
	Depart()
}

// PassengerTrain struct
type PassengerTrain struct {
	mediator Mediator
}

func NewPassengerTrain(mediator Mediator) *PassengerTrain {
	return &PassengerTrain{mediator: mediator}
}

func (p *PassengerTrain) Arrive() {
	if p.mediator.CanArrive(p) {
		fmt.Println("Passenger train arrived.")
	} else {
		fmt.Println("Passenger train waiting for arrival.")
	}
}

func (p *PassengerTrain) Depart() {
	fmt.Println("Passenger train leaving...")
	p.mediator.NotifyAboutDeparture()
}

// FreightTrain struct
type FreightTrain struct {
	mediator Mediator
}

func NewFreightTrain(mediator Mediator) *FreightTrain {
	return &FreightTrain{mediator: mediator}
}

func (f *FreightTrain) Arrive() {
	if f.mediator.CanArrive(f) {
		fmt.Println("Freight train arrived.")
	} else {
		fmt.Println("Freight train waiting for arrival.")
	}
}

func (f *FreightTrain) Depart() {
	fmt.Println("Freight train leaving...")
	f.mediator.NotifyAboutDeparture()
}
