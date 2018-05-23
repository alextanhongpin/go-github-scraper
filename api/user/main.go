package user

import "errors"

// Service provides the interface for the Service struct
type Service interface {
	GetOne() error
	InsertMany() error
}

// service is a struct that holds service configuration
type service struct{}

// New returns a new service
func New() Service {
	return &service{}
}

func (s *service) GetOne() error {
	return errors.New("Not implemented")
}

func (s *service) InsertMany() error {
	return errors.New("Not implemented")
}
