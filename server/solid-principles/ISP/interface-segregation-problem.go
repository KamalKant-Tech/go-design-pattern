package main

import (
	"fmt"
	"time"
)

// ISP States:
// A client should never be forced to implement an interface that it doesn't use, or clients shouldn't be forced to depend on methods they do not use.

// This is the case violation of ISP principle.
// In thi below example we can see that we have interface TripRepo and it has three function implementation. But all three we do not use in this code repo.
// Which breaks the second part of the rule "clients shouldn't be forced to depend on methods they do not use."
// And when we write the unit test case and try to mock TripRepo type then we should have to implement the unncessory func of SaveTrip and GetTrip which
// is again the break the rule `A Client should never ne forced to implement an interface that it doed not use.`

type Trip struct {
	From string
	To   string
	Date time.Time
	Bus  string
}

type TripRepo interface {
	SaveTrip() (Trip, error)
	GetTrip() (Trip, error)
	FindTrips() (Trip, error)
}

type usecase struct {
	tripRepo TripRepo
}

func CreateUseCaseService(trip TripRepo) *usecase {
	return &usecase{tripRepo: trip}
}

type TRepo struct{}

// GetTrip implements TripRepo.
func (*TRepo) GetTrip() (Trip, error) {
	panic("unimplemented")
}

// SaveTrip implements TripRepo.
func (*TRepo) SaveTrip() (Trip, error) {
	panic("unimplemented")
}

func (t *TRepo) FindTrips() (Trip, error) {
	return Trip{}, nil
}

func main() {
	tRepo := &TRepo{}
	uCase := CreateUseCaseService(tRepo)
	fmt.Println(uCase)
}
