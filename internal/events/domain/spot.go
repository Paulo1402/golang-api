package domain

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrInvalidSpotNumber   = errors.New(("invalid spot number"))
	ErrSpotNotFound        = errors.New(("spot not found"))
	ErrSpotAlreadyReserved = errors.New(("spot already reserved"))

	ErrSpotNameRequired = errors.New("spot name is required")
	ErrSpotNameLength   = errors.New("spot name must have two characters")
	ErrSpotNameStart    = errors.New("spot name must start with a letter")
	ErrSpotNameEnd      = errors.New("spot name must end with a number")
)

type SpotStatus string

const (
	SpotStatusAvailable SpotStatus = "available"
	SpotStatusSold      SpotStatus = "sold"
)

type Spot struct {
	ID       string
	EventID  string
	Name     string
	Status   SpotStatus
	TicketID string
}

func NewSpot(event *Event, name string) (*Spot, error) {
	spot := &Spot{
		ID:      uuid.New().String(),
		EventID: event.ID,
		Name:    name,
		Status:  SpotStatusAvailable,
	}

	if err := spot.Validate(); err != nil {
		return nil, err
	}

	return spot, nil
}

func (s Spot) Validate() error {
	if s.Name == "" {
		return ErrSpotNameRequired
	}

	if len(s.Name) != 2 {
		return ErrSpotNameLength
	}

	if s.Name[0] < 'A' || s.Name[0] > 'Z' {
		return ErrSpotNameStart
	}

	if s.Name[1] < '0' || s.Name[1] > '9' {
		return ErrSpotNameEnd
	}

	return nil
}

func (s *Spot) Reserve(ticketID string) error {
	if s.Status == SpotStatusSold {
		return ErrSpotAlreadyReserved
	}

	s.Status = SpotStatusSold
	s.TicketID = ticketID

	return nil
}