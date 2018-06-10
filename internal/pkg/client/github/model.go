package github

import (
	"errors"
)

var (
	ErrStartFieldRequired = errors.New(`field "start" is required`)
	ErrEndFieldRequired   = errors.New(`field "end" is required`)
)

// Model represents the api interface for the Github's GraphQL
type (
	Model interface {
		FetchUsers(req FetchUsersRequest) (*FetchUsersResponse, error)
		FetchRepos(req FetchReposRequest) (*FetchReposResponse, error)
	}

	model struct {
		store Store
	}
)

// NewModel returns a new model
func NewModel(store Store) Model {
	return &model{
		store: store,
	}
}

func (m *model) FetchUsers(req FetchUsersRequest) (*FetchUsersResponse, error) {
	if req.Start == "" {
		return nil, ErrStartFieldRequired
	}
	if req.End == "" {
		return nil, ErrEndFieldRequired
	}
	return m.store.FetchUsers(req)
}

func (m *model) FetchRepos(req FetchReposRequest) (*FetchReposResponse, error) {
	if req.Start == "" {
		return nil, ErrStartFieldRequired
	}
	if req.End == "" {
		return nil, ErrEndFieldRequired
	}
	return m.store.FetchRepos(req)
}
