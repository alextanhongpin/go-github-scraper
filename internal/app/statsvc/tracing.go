package statsvc

import (
	"context"

	"github.com/alextanhongpin/go-github-scraper/internal/pkg/schema"
	"go.opencensus.io/trace"
)

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

func (m *tracingMiddleware) GetUserCount(ctx context.Context) (*UserCount, error) {
	ctx, span := trace.StartSpan(ctx, "GetUserCount")
	defer span.End()

	return m.service.GetUserCount(ctx)
}

func (m *tracingMiddleware) PostUserCount(ctx context.Context, count int) error {
	ctx, span := trace.StartSpan(ctx, "PostUserCount")
	defer span.End()

	return m.service.PostUserCount(ctx, count)
}

func (m *tracingMiddleware) GetRepoCount(ctx context.Context) (*RepoCount, error) {
	ctx, span := trace.StartSpan(ctx, "GetRepoCount")
	defer span.End()

	return m.service.GetRepoCount(ctx)
}

func (m *tracingMiddleware) PostRepoCount(ctx context.Context, count int) error {
	ctx, span := trace.StartSpan(ctx, "PostRepoCount")
	defer span.End()

	return m.service.PostRepoCount(ctx, count)
}

func (m *tracingMiddleware) GetReposMostRecent(ctx context.Context) (*ReposMostRecent, error) {
	ctx, span := trace.StartSpan(ctx, "GetReposMostRecent")
	defer span.End()

	return m.service.GetReposMostRecent(ctx)
}

func (m *tracingMiddleware) PostReposMostRecent(ctx context.Context, data []schema.Repo) error {
	ctx, span := trace.StartSpan(ctx, "PostReposMostRecent")
	defer span.End()

	return m.service.PostReposMostRecent(ctx, data)
}

func (m *tracingMiddleware) GetRepoCountByUser(ctx context.Context) (*RepoCountByUser, error) {
	ctx, span := trace.StartSpan(ctx, "GetRepoCountByUser")
	defer span.End()

	return m.service.GetRepoCountByUser(ctx)
}

func (m *tracingMiddleware) PostRepoCountByUser(ctx context.Context, users []schema.UserCount) error {
	ctx, span := trace.StartSpan(ctx, "PostRepoCountByUser")
	defer span.End()

	return m.service.PostRepoCountByUser(ctx, users)
}

func (m *tracingMiddleware) GetReposMostStars(ctx context.Context) (*ReposMostStars, error) {
	ctx, span := trace.StartSpan(ctx, "GetReposMostStars")
	defer span.End()

	return m.service.GetReposMostStars(ctx)
}

func (m *tracingMiddleware) PostReposMostStars(ctx context.Context, repos []schema.Repo) error {
	ctx, span := trace.StartSpan(ctx, "PostReposMostStars")
	defer span.End()

	return m.service.PostReposMostStars(ctx, repos)
}

func (m *tracingMiddleware) GetReposMostForks(ctx context.Context) (*ReposMostForks, error) {
	ctx, span := trace.StartSpan(ctx, "GetReposMostForks")
	defer span.End()

	return m.service.GetReposMostForks(ctx)
}

func (m *tracingMiddleware) PostReposMostForks(ctx context.Context, repos []schema.Repo) error {
	ctx, span := trace.StartSpan(ctx, "PostReposMostForks")
	defer span.End()

	return m.service.PostReposMostForks(ctx, repos)
}

func (m *tracingMiddleware) GetMostPopularLanguage(ctx context.Context) (*MostPopularLanguage, error) {
	ctx, span := trace.StartSpan(ctx, "GetMostPopularLanguage")
	defer span.End()

	return m.service.GetMostPopularLanguage(ctx)
}

func (m *tracingMiddleware) PostMostPopularLanguage(ctx context.Context, languages []schema.LanguageCount) error {
	ctx, span := trace.StartSpan(ctx, "PostMostPopularLanguage")
	defer span.End()

	return m.service.PostMostPopularLanguage(ctx, languages)
}

func (m *tracingMiddleware) GetLanguageCountByUser(ctx context.Context) (*LanguageCountByUser, error) {
	ctx, span := trace.StartSpan(ctx, "GetLanguageCountByUser")
	defer span.End()

	return m.service.GetLanguageCountByUser(ctx)
}

func (m *tracingMiddleware) PostLanguageCountByUser(ctx context.Context, languages []schema.LanguageCount) error {
	ctx, span := trace.StartSpan(ctx, "PostLanguageCountByUser")
	defer span.End()

	return m.service.PostLanguageCountByUser(ctx, languages)
}

func (m *tracingMiddleware) GetMostRecentReposByLanguage(ctx context.Context) (*MostRecentReposByLanguage, error) {
	ctx, span := trace.StartSpan(ctx, "GetMostRecentReposByLanguage")
	defer span.End()

	return m.service.GetMostRecentReposByLanguage(ctx)
}

func (m *tracingMiddleware) PostMostRecentReposByLanguage(ctx context.Context, repos []schema.RepoLanguage) error {
	ctx, span := trace.StartSpan(ctx, "PostMostRecentReposByLanguage")
	defer span.End()

	return m.service.PostMostRecentReposByLanguage(ctx, repos)
}

func (m *tracingMiddleware) GetReposByLanguage(ctx context.Context) (*ReposByLanguage, error) {
	ctx, span := trace.StartSpan(ctx, "GetReposByLanguage")
	defer span.End()

	return m.service.GetReposByLanguage(ctx)
}

func (m *tracingMiddleware) PostReposByLanguage(ctx context.Context, users []schema.UserCountByLanguage) error {
	ctx, span := trace.StartSpan(ctx, "PostReposByLanguage")
	defer span.End()

	return m.service.PostReposByLanguage(ctx, users)
}

func (m *tracingMiddleware) GetCompanyCount(ctx context.Context) (*CompanyCount, error) {
	ctx, span := trace.StartSpan(ctx, "GetCompanyCount")
	defer span.End()

	return m.service.GetCompanyCount(ctx)
}

func (m *tracingMiddleware) PostCompanyCount(ctx context.Context, count int) error {
	ctx, span := trace.StartSpan(ctx, "PostCompanyCount")
	defer span.End()

	return m.service.PostCompanyCount(ctx, count)
}

func (m *tracingMiddleware) GetUsersByCompany(ctx context.Context) (*UsersByCompany, error) {
	ctx, span := trace.StartSpan(ctx, "GetUsersByCompany")
	defer span.End()

	return m.service.GetUsersByCompany(ctx)
}

func (m *tracingMiddleware) PostUsersByCompany(ctx context.Context, users []schema.Company) error {
	ctx, span := trace.StartSpan(ctx, "PostUsersByCompany")
	defer span.End()

	return m.service.PostUsersByCompany(ctx, users)
}
