package usersvc

import (
	"context"

	"github.com/alextanhongpin/go-github-scraper/internal/pkg/client/github"
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/database"
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/logger"
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/schema"
)

// Service represents the model of the user
type Service interface {
	AggregateCompany(ctx context.Context, min, max int) ([]schema.Company, error)
	BulkUpsert(ctx context.Context, users []github.User) error
	BulkUpdate(ctx context.Context, users []User) error
	Count(ctx context.Context) (int, error)
	Drop(ctx context.Context) error
	FindByCompany(ctx context.Context, company string) ([]schema.User, error)
	FindOne(ctx context.Context, login string) (*User, error)
	FindLastCreated(ctx context.Context) (string, bool)
	FindLastFetched(ctx context.Context, limit int) ([]User, error)
	MostRecent(ctx context.Context, limit int) ([]User, error)
	Init(ctx context.Context) error
	UpdateOne(ctx context.Context, login string) error
	PickLogin(ctx context.Context) ([]string, error)
	WithRepos(ctx context.Context, count int) ([]User, error)
	DistinctCompany(ctx context.Context) ([]string, error)
}

// New returns a new user service
func New(db *database.DB, l *logger.Logger) Service {
	return NewModel(NewStore(db, database.Users), l)
}
