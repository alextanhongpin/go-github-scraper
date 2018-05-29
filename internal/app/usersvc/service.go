package usersvc

import (
	"context"

	"github.com/alextanhongpin/go-github-scraper/internal/pkg/client/github"
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/database"
)

// Service represents the model of the user
type Service interface {
	BulkUpsert(ctx context.Context, users []github.User) error
	Count(ctx context.Context) (int, error)
	Drop(ctx context.Context) error
	Init(ctx context.Context) error
	FindOne(ctx context.Context, login string) (*User, error)
	FindLastCreated(ctx context.Context) (string, bool)
	FindLastFetched(ctx context.Context, limit int) ([]User, error)
	MostRecent(ctx context.Context, limit int) ([]User, error)
	PickLogin(ctx context.Context) ([]string, error)
	UpdateOne(ctx context.Context, login string) error
}

// New returns a new user service
func New(db *database.DB) Service {
	return NewModel(NewStore(db, database.Users))
}
