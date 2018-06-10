package usersvc

import (
	"log"
	"time"

	"github.com/alextanhongpin/go-github-scraper/internal/pkg/client/github"
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/constant"
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
		FindLastCreated() (string, bool)
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
	return m.store.FindAll(limit, []string{"-createdAt"})
}

func (m *model) BulkUpsert(users []github.User) error {
	return m.store.BulkUpsert(users)
}

func (m *model) BulkUpdate(users []User) error {
	return m.store.BulkUpdate(users)
}

func (m *model) Drop() error {
	return m.store.Drop()
}

// FindLastCreated returns the last created date in the format YYYY-MM-DD, and a boolean to indicate
// if the value returned exists or is default
func (m *model) FindLastCreated() (string, bool) {
	user, err := m.store.FindLastCreated()
	if err != nil || user == nil {
		return constant.GithubCreatedAt, false
	}
	t, err := time.Parse(time.RFC3339, user.CreatedAt)
	if err != nil {
		return constant.GithubCreatedAt, false
	}
	t = t.AddDate(0, -1, 0)
	return t.Format("2006-01-02"), true
}

func (m *model) FindByCompany(company string) ([]schema.User, error) {
	return m.store.FindByCompany(company)
}

func (m *model) FindLastFetched(limit int) ([]User, error) {
	return m.store.FindAll(limit, []string{"fetchedAt"})
}

func (m *model) Count() (int, error) {
	return m.store.Count()
}

func (m *model) UpdateOne(login string) (err error) {
	return m.store.UpdateOne(login)
}

func (m *model) FindOne(login string) (*User, error) {
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
