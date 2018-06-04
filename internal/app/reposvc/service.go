// Package reposvc defines the service for the repo microservice
// The interface defines the method that exists for the service
// and should have ctx as the first arguments to support cancellation/deadlines etc
// Logging, tracing etc should be done at the service level, not store
package reposvc

import (
	"context"

	"github.com/alextanhongpin/go-github-scraper/internal/app/usersvc"
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/client/github"
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/database"
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/logger"
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/schema"
)

// Service represents the interface for the repo service
type Service interface {
	BulkUpsert(ctx context.Context, repos []github.Repo) error
	Count(ctx context.Context) (int, error)
	Drop(ctx context.Context) error
	Init(ctx context.Context) error
	FindLastCreatedByUser(ctx context.Context, login string) (string, bool)
	LanguageCountByUser(ctx context.Context, login string, limit int) ([]schema.LanguageCount, error)
	MostPopularLanguage(ctx context.Context, limit int) ([]schema.LanguageCount, error)
	MostRecent(ctx context.Context, limit int) ([]schema.Repo, error)
	MostRecentReposByLanguage(ctx context.Context, language string, limit int) ([]schema.Repo, error)
	MostStars(ctx context.Context, limit int) ([]schema.Repo, error)
	MostForks(ctx context.Context, limit int) ([]schema.Repo, error)
	RepoCountByUser(ctx context.Context, limit int) ([]schema.UserCount, error)
	ReposByLanguage(ctx context.Context, language string, limit int) ([]schema.UserCount, error)
	DistinctLogin(ctx context.Context) ([]string, error)
	GetProfile(ctx context.Context, login string) usersvc.User
}

// New returns a new service with store
func New(db *database.DB, l *logger.Logger) Service {
	return NewModel(NewStore(db, database.Repos), l)
}
