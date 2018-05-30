// Package reposvc defines the service for the repo microservice
// The interface defines the method that exists for the service
// and should have ctx as the first arguments to support cancellation/deadlines etc
// Logging, tracing etc should be done at the service level, not store
package reposvc

import (
	"context"

	"github.com/alextanhongpin/go-github-scraper/internal/pkg/client/github"
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/database"
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
	RepoCountByUser(ctx context.Context, limit int) ([]schema.UserCount, error)
	ReposByLanguage(ctx context.Context, language string, limit int) ([]schema.UserCount, error)
	WatchersFor(ctx context.Context, login string) (int64, error)
	StargazersFor(ctx context.Context, login string) (int64, error)
	ForksFor(ctx context.Context, login string) (int64, error)
	KeywordsFor(ctx context.Context, login string, limit int) ([]schema.Keyword, error)
	DistinctLogin(ctx context.Context) ([]string, error)
	GetProfile(ctx context.Context, login string) schema.Profile
}

// New returns a new service with store
func New(db *database.DB) Service {
	return NewModel(NewStore(db, database.Repos))
}