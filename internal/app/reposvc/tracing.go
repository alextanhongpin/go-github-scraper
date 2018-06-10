package reposvc

import (
	"context"

	"github.com/alextanhongpin/go-github-scraper/internal/app/usersvc"
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/client/github"
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/schema"

	"go.opencensus.io/trace"
)

// Tracing is a middleware that adds tracing capabilities to the service
func Tracing() Middleware {
	return func(s Service) Service {
		return &tracingMiddleware{
			service: s,
		}
	}
}

type tracingMiddleware struct {
	service Service
}

func (m *tracingMiddleware) BulkUpsert(ctx context.Context, repos []github.Repo) error {
	ctx, span := trace.StartSpan(ctx, "BulkUpsert")
	defer span.End()

	span.AddAttributes(trace.Int64Attribute("perPage", int64(len(repos))))

	return m.service.BulkUpsert(ctx, repos)
}

func (m *tracingMiddleware) Count(ctx context.Context) (int, error) {
	ctx, span := trace.StartSpan(ctx, "Count")
	defer span.End()

	return m.service.Count(ctx)
}

func (m *tracingMiddleware) LastCreatedBy(ctx context.Context, login string) (string, bool) {
	ctx, span := trace.StartSpan(ctx, "LastCreatedBy")
	defer span.End()

	span.AddAttributes(trace.StringAttribute("login", login))

	return m.service.LastCreatedBy(ctx, login)
}

func (m *tracingMiddleware) MostPopularLanguage(ctx context.Context, limit int) ([]schema.LanguageCount, error) {
	ctx, span := trace.StartSpan(ctx, "MostPopularLanguage")
	defer span.End()

	span.AddAttributes(trace.Int64Attribute("limit", int64(limit)))

	return m.service.MostPopularLanguage(ctx, limit)
}

func (m *tracingMiddleware) MostRecent(ctx context.Context, limit int) ([]schema.Repo, error) {
	ctx, span := trace.StartSpan(ctx, "MostRecent")
	defer span.End()

	span.AddAttributes(trace.Int64Attribute("limit", int64(limit)))

	return m.service.MostRecent(ctx, limit)
}

func (m *tracingMiddleware) MostRecentReposByLanguage(ctx context.Context, language string, limit int) ([]schema.Repo, error) {
	ctx, span := trace.StartSpan(ctx, "MostRecentReposByLanguage")
	defer span.End()

	span.AddAttributes(
		trace.StringAttribute("language", language),
		trace.Int64Attribute("limit", int64(limit)))

	return m.service.MostRecentReposByLanguage(ctx, language, limit)
}

func (m *tracingMiddleware) MostStars(ctx context.Context, limit int) ([]schema.Repo, error) {
	ctx, span := trace.StartSpan(ctx, "MostStars")
	defer span.End()

	span.AddAttributes(trace.Int64Attribute("limit", int64(limit)))

	return m.service.MostStars(ctx, limit)
}

func (m *tracingMiddleware) MostForks(ctx context.Context, limit int) ([]schema.Repo, error) {
	ctx, span := trace.StartSpan(ctx, "MostForks")
	defer span.End()

	span.AddAttributes(trace.Int64Attribute("limit", int64(limit)))

	return m.service.MostForks(ctx, limit)
}

func (m *tracingMiddleware) RepoCountByUser(ctx context.Context, limit int) ([]schema.UserCount, error) {
	ctx, span := trace.StartSpan(ctx, "RepoCountByUser")
	defer span.End()

	span.AddAttributes(trace.Int64Attribute("limit", int64(limit)))

	return m.service.RepoCountByUser(ctx, limit)
}

func (m *tracingMiddleware) ReposByLanguage(ctx context.Context, language string, limit int) ([]schema.UserCount, error) {
	ctx, span := trace.StartSpan(ctx, "ReposByLanguage")
	defer span.End()

	span.AddAttributes(
		trace.StringAttribute("language", language),
		trace.Int64Attribute("limit", int64(limit)))

	return m.service.ReposByLanguage(ctx, language, limit)
}

func (m *tracingMiddleware) Distinct(ctx context.Context, login string) ([]string, error) {
	ctx, span := trace.StartSpan(ctx, "Distinct")
	defer span.End()

	span.AddAttributes(trace.StringAttribute("login", login))

	return m.service.Distinct(ctx, login)
}

func (m *tracingMiddleware) GetProfile(ctx context.Context, login string) (*usersvc.User, error) {
	ctx, span := trace.StartSpan(ctx, "GetProfile")
	defer span.End()

	span.AddAttributes(trace.StringAttribute("login", login))

	return m.service.GetProfile(ctx, login)
}
