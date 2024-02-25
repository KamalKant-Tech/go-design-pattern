package main

import (
	"fmt"
	"time"
)

// ISP States:
// A client should never be forced to implement an interface that it doesn't use, or clients shouldn't be forced to depend on methods they do not use.

// Define an interface with only functions which are being utilised in the code and do not have unnecessory func implementation.

type Trip struct {
	From string
	To   string
	Date time.Time
	Bus  string
}

type TRepo struct{}

type TripRepo interface {
	FindTrips() (Trip, error)
}

type usecase struct {
	tripRepo TripRepo
}

func CreateUseCaseService(trip TripRepo) *usecase {
	return &usecase{tripRepo: trip}
}

func (t *TRepo) FindTrips() (Trip, error) {
	return Trip{}, nil
}

func main() {
	tRepo := &TRepo{}
	uCase := CreateUseCaseService(tRepo)
	fmt.Println(uCase)
}
