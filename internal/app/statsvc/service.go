package statsvc

import (
	"context"

	"github.com/alextanhongpin/go-github-scraper/internal/pkg/database"
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/logger"
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/schema"
)

// Service represents the analytic service
type Service interface {
	Init(ctx context.Context) error
	GetUserCount(ctx context.Context) (*UserCount, error)
	PostUserCount(ctx context.Context, count int) error
	GetRepoCount(ctx context.Context) (*RepoCount, error)
	PostRepoCount(ctx context.Context, count int) error
	GetReposMostRecent(ctx context.Context) (*ReposMostRecent, error)
	PostReposMostRecent(ctx context.Context, data []schema.Repo) error
	GetRepoCountByUser(ctx context.Context) (*RepoCountByUser, error)
	PostRepoCountByUser(ctx context.Context, users []schema.UserCount) error
	GetReposMostStars(ctx context.Context) (*ReposMostStars, error)
	PostReposMostStars(ctx context.Context, repos []schema.Repo) error
	GetMostPopularLanguage(ctx context.Context) (*MostPopularLanguage, error)
	PostMostPopularLanguage(ctx context.Context, languages []schema.LanguageCount) error
	GetLanguageCountByUser(ctx context.Context) (*LanguageCountByUser, error)
	PostLanguageCountByUser(ctx context.Context, languages []schema.LanguageCount) error
	GetMostRecentReposByLanguage(ctx context.Context) (*MostRecentReposByLanguage, error)
	PostMostRecentReposByLanguage(ctx context.Context, repos []schema.RepoLanguage) error
	GetReposByLanguage(ctx context.Context) (*ReposByLanguage, error)
	PostReposByLanguage(ctx context.Context, users []schema.UserCountByLanguage) error
}

// New returns a new analytic service model
func New(db *database.DB, l *logger.Logger) Service {
	return NewModel(NewStore(db, database.Stats), l)
}
