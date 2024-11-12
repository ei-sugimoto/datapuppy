package domain

import (
	"github.com/ei-sugimoto/datapuppy/internal"
)

type Host struct {
	ID   int
	Name string
	Port int
}

func NewHost(name string, port int) (*Host, error) {
	if err := ValidateHost(name, port); err != nil {
		return nil, err
	}
	return &Host{
		Name: name,
		Port: port,
	}, nil
}

func ValidateHost(name string, port int) error {
	if name == "" {
		return internal.ErrNameIsRequired
	}
	if port <= 0 {
		return internal.ErrPortMustBeGreaterThanZero
	}
	return nil
}
