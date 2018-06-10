package usersvc

import (
	"context"
	"time"

	"github.com/alextanhongpin/go-github-scraper/internal/pkg/client/github"
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/constant"
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/schema"
)

// Service represents the model of the user
type Service interface {
	FindLastCreated(ctx context.Context) (string, bool)
	BulkUpsert(ctx context.Context, users []github.User) error
	FindLastFetched(ctx context.Context, limit int) ([]User, error)
	UpdateOne(ctx context.Context, login string) error
	Count(ctx context.Context) (int, error)
	BulkUpdate(ctx context.Context, users []User) error
	WithRepos(ctx context.Context, count int) ([]User, error)
	DistinctCompany(ctx context.Context) ([]string, error)
	FindByCompany(ctx context.Context, company string) ([]schema.User, error)
	AggregateCompany(ctx context.Context, min, max int) ([]schema.Company, error)
	FindOne(ctx context.Context, login string) (*User, error)
}

type service struct {
	model Model
}

// NewService returns a new service
func NewService(m Model) Service {
	return &service{m}
}

func (s *service) FindLastCreated(ctx context.Context) (string, bool) {
	user, err := s.model.FindLastCreated()
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

func (s *service) BulkUpsert(ctx context.Context, users []github.User) error {
	return s.model.BulkUpsert(users)
}

func (s *service) FindLastFetched(ctx context.Context, limit int) ([]User, error) {
	return s.model.FindLastFetched(limit)
}

func (s *service) UpdateOne(ctx context.Context, login string) error {
	return s.model.UpdateOne(login)
}

func (s *service) Count(ctx context.Context) (int, error) {
	return s.model.Count()
}

func (s *service) BulkUpdate(ctx context.Context, users []User) error {
	return s.model.BulkUpdate(users)
}

func (s *service) WithRepos(ctx context.Context, count int) ([]User, error) {
	return s.model.WithRepos(count)
}

func (s *service) DistinctCompany(ctx context.Context) ([]string, error) {
	return s.model.DistinctCompany()
}

func (s *service) FindByCompany(ctx context.Context, company string) ([]schema.User, error) {
	return s.model.FindByCompany(company)
}

func (s *service) AggregateCompany(ctx context.Context, min, max int) ([]schema.Company, error) {
	return s.model.AggregateCompany(min, max)
}

func (s *service) FindOne(ctx context.Context, login string) (*User, error) {
	return s.model.FindOne(login)
}
