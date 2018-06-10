package usersvc

import (
	"errors"
	"log"

	"github.com/alextanhongpin/go-github-scraper/internal/pkg/client/github"
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/schema"
)

type (
	// Model contains business logic for user-related operations
	Model interface {
		AggregateCompany(min, max int) ([]schema.Company, error)
		BulkUpsert(users []github.User) error
		BulkUpdate(users []User) error
		Count() (int, error)
		Drop() error
		FindByCompany(company string) ([]schema.User, error)
		FindOne(login string) (*User, error)
		FindLastCreated() (*User, error)
		FindLastFetched(limit int) ([]User, error)
		MostRecent(limit int) ([]User, error)
		Init() error
		UpdateOne(login string) error
		PickLogin() ([]string, error)
		WithRepos(count int) ([]User, error)
		DistinctCompany() ([]string, error)
	}

	model struct {
		store Store
	}
)

var (
	ErrInvalidLogin = errors.New("login provided is invalid")
)

// NewModel returns a new model with the store
func NewModel(store Store) Model {
	m := model{store: store}
	if err := m.Init(); err != nil {
		log.Fatal(err)
	}
	return &m
}

func (m *model) Init() error {
	return m.store.Init()
}

func (m *model) MostRecent(limit int) ([]User, error) {
	limit = setLimit(limit)
	return m.store.FindAll(limit, []string{"-createdAt"})
}

func (m *model) BulkUpsert(users []github.User) error {
	if len(users) == 0 {
		return nil
	}
	return m.store.BulkUpsert(users)
}

func (m *model) BulkUpdate(users []User) error {
	if len(users) == 0 {
		return nil
	}
	return m.store.BulkUpdate(users)
}

func (m *model) Drop() error {
	return m.store.Drop()
}

// FindLastCreated returns the last created date in the format YYYY-MM-DD, and a boolean to indicate
// if the value returned exists or is default
func (m *model) FindLastCreated() (*User, error) {
	return m.store.FindLastCreated()
}

func (m *model) FindByCompany(company string) ([]schema.User, error) {
	return m.store.FindByCompany(company)
}

func (m *model) FindLastFetched(limit int) ([]User, error) {
	limit = setLimit(limit)
	return m.store.FindAll(limit, []string{"fetchedAt"})
}

func (m *model) Count() (int, error) {
	return m.store.Count()
}

func (m *model) UpdateOne(login string) (err error) {
	if login == "" {
		return ErrInvalidLogin
	}
	return m.store.UpdateOne(login)
}

func (m *model) FindOne(login string) (*User, error) {
	if login == "" {
		return nil, ErrInvalidLogin
	}
	return m.store.FindOne(login)
}

func (m *model) PickLogin() ([]string, error) {
	return m.store.PickLogin()
}

func (m *model) WithRepos(count int) ([]User, error) {
	return m.store.WithRepos(count)
}

func (m *model) AggregateCompany(min, max int) ([]schema.Company, error) {
	return m.store.AggregateCompany(min, max)
}

func (m *model) DistinctCompany() ([]string, error) {
	return m.store.DistinctCompany()
}

func setLimit(limit int) int {
	if limit < 0 {
		return 10
	}
	if limit > 100 {
		return 100
	}
	return limit
}
