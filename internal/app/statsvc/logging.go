package statsvc

import (
	"context"
	"time"

	"github.com/alextanhongpin/go-github-scraper/internal/pkg/logger"
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/schema"
	"go.uber.org/zap"
)

func Logging(l *logger.Logger) Middleware {
	return func(s Service) Service {
		return &loggingMiddleware{
			service: s,
			logger:  l,
		}
	}
}

type loggingMiddleware struct {
	logger  *logger.Logger
	service Service
}

func (m *loggingMiddleware) GetUserCount(ctx context.Context) (res *UserCount, err error) {
	defer func(start time.Time) {
		L := logger.Wrap(ctx, m.logger,
			logger.Method("GetUserCount"),
			logger.Duration(start))

		logger.Maybe(L, "get user count", err)
	}(time.Now())

	return m.service.GetUserCount(ctx)
}

func (m *loggingMiddleware) PostUserCount(ctx context.Context, count int) (err error) {
	defer func(start time.Time) {
		L := logger.Wrap(ctx, m.logger,
			logger.Method("PostUserCount"),
			logger.Duration(start),
			zap.Int("count", count))

		logger.Maybe(L, "post user count", err)
	}(time.Now())

	return m.service.PostUserCount(ctx, count)
}

func (m *loggingMiddleware) GetRepoCount(ctx context.Context) (res *RepoCount, err error) {
	defer func(start time.Time) {
		L := logger.Wrap(ctx, m.logger,
			logger.Method("GetRepoCount"),
			logger.Duration(start))

		logger.Maybe(L, "get repo count", err)
	}(time.Now())

	return m.service.GetRepoCount(ctx)
}

func (m *loggingMiddleware) PostRepoCount(ctx context.Context, count int) (err error) {
	defer func(start time.Time) {
		L := logger.Wrap(ctx, m.logger,
			logger.Method("PostRepoCount"),
			logger.Duration(start))

		logger.Maybe(L, "post repo count", err)
	}(time.Now())

	return m.service.PostRepoCount(ctx, count)
}

func (m *loggingMiddleware) GetReposMostRecent(ctx context.Context) (res *ReposMostRecent, err error) {
	defer func(start time.Time) {
		L := logger.Wrap(ctx, m.logger,
			logger.Method("GetReposMostRecent"),
			logger.Duration(start))

		logger.Maybe(L, "get repos most recent", err)
	}(time.Now())

	return m.service.GetReposMostRecent(ctx)
}

func (m *loggingMiddleware) PostReposMostRecent(ctx context.Context, repos []schema.Repo) (err error) {
	defer func(start time.Time) {
		L := logger.Wrap(ctx, m.logger,
			logger.Method("PostReposMostRecent"),
			logger.Duration(start))

		logger.Maybe(L, "post repos most recent", err)
	}(time.Now())

	return m.service.PostReposMostRecent(ctx, repos)
}

func (m *loggingMiddleware) GetRepoCountByUser(ctx context.Context) (res *RepoCountByUser, err error) {
	defer func(start time.Time) {
		L := logger.Wrap(ctx, m.logger,
			logger.Method("GetRepoCountByUser"),
			logger.Duration(start))

		logger.Maybe(L, "get repo count by user", err)
	}(time.Now())

	return m.service.GetRepoCountByUser(ctx)
}

func (m *loggingMiddleware) PostRepoCountByUser(ctx context.Context, users []schema.UserCount) (err error) {
	defer func(start time.Time) {
		L := logger.Wrap(ctx, m.logger,
			logger.Method("PostRepoCountByUser"),
			logger.Duration(start))

		logger.Maybe(L, "post repo count by user", err)
	}(time.Now())

	return m.service.PostRepoCountByUser(ctx, users)
}

func (m *loggingMiddleware) GetReposMostStars(ctx context.Context) (res *ReposMostStars, err error) {
	defer func(start time.Time) {
		L := logger.Wrap(ctx, m.logger,
			logger.Method("GetReposMostStars"),
			logger.Duration(start))

		logger.Maybe(L, "get repos most stars", err)
	}(time.Now())

	return m.service.GetReposMostStars(ctx)
}

func (m *loggingMiddleware) PostReposMostStars(ctx context.Context, repos []schema.Repo) (err error) {
	defer func(start time.Time) {
		L := logger.Wrap(ctx, m.logger,
			logger.Method("PostReposMostStars"),
			logger.Duration(start))

		logger.Maybe(L, "post repos most stars", err)
	}(time.Now())

	return m.service.PostReposMostStars(ctx, repos)
}

func (m *loggingMiddleware) GetReposMostForks(ctx context.Context) (res *ReposMostForks, err error) {
	defer func(start time.Time) {
		L := logger.Wrap(ctx, m.logger,
			logger.Method("GetReposMostForks"),
			logger.Duration(start))

		logger.Maybe(L, "get repos most forks", err)
	}(time.Now())

	return m.service.GetReposMostForks(ctx)
}

func (m *loggingMiddleware) PostReposMostForks(ctx context.Context, repos []schema.Repo) (err error) {
	defer func(start time.Time) {
		L := logger.Wrap(ctx, m.logger,
			logger.Method("PostReposMostForks"),
			logger.Duration(start))

		logger.Maybe(L, "post repos most forks", err)
	}(time.Now())

	return m.service.PostReposMostForks(ctx, repos)
}

func (m *loggingMiddleware) GetMostPopularLanguage(ctx context.Context) (res *MostPopularLanguage, err error) {
	defer func(start time.Time) {
		L := logger.Wrap(ctx, m.logger,
			logger.Method("GetMostPopularLanguage"),
			logger.Duration(start))

		logger.Maybe(L, "get most popular language", err)
	}(time.Now())

	return m.service.GetMostPopularLanguage(ctx)
}

func (m *loggingMiddleware) PostMostPopularLanguage(ctx context.Context, languages []schema.LanguageCount) (err error) {
	defer func(start time.Time) {
		L := logger.Wrap(ctx, m.logger,
			logger.Method("PostMostPopularLanguage"),
			logger.Duration(start))

		logger.Maybe(L, "post most popular language", err)
	}(time.Now())

	return m.service.PostMostPopularLanguage(ctx, languages)
}

func (m *loggingMiddleware) GetLanguageCountByUser(ctx context.Context) (res *LanguageCountByUser, err error) {
	defer func(start time.Time) {
		L := logger.Wrap(ctx, m.logger,
			logger.Method("GetLanguageCountByUser"),
			logger.Duration(start))

		logger.Maybe(L, "get language count by user", err)
	}(time.Now())

	return m.service.GetLanguageCountByUser(ctx)
}

func (m *loggingMiddleware) PostLanguageCountByUser(ctx context.Context, languages []schema.LanguageCount) (err error) {
	defer func(start time.Time) {
		L := logger.Wrap(ctx, m.logger,
			logger.Method("PostLanguageCountByUser"),
			logger.Duration(start))

		logger.Maybe(L, "post language count by user", err)
	}(time.Now())

	return m.service.PostLanguageCountByUser(ctx, languages)
}

func (m *loggingMiddleware) GetMostRecentReposByLanguage(ctx context.Context) (res *MostRecentReposByLanguage, err error) {
	defer func(start time.Time) {
		L := logger.Wrap(ctx, m.logger,
			logger.Method("GetMostRecentReposByLanguage"),
			logger.Duration(start))

		logger.Maybe(L, "get most recent repos by language", err)
	}(time.Now())

	return m.service.GetMostRecentReposByLanguage(ctx)
}

func (m *loggingMiddleware) PostMostRecentReposByLanguage(ctx context.Context, repos []schema.RepoLanguage) (err error) {
	defer func(start time.Time) {
		L := logger.Wrap(ctx, m.logger,
			logger.Method("PostMostRecentReposByLanguage"),
			logger.Duration(start))

		logger.Maybe(L, "post most recent repos by language", err)
	}(time.Now())

	return m.service.PostMostRecentReposByLanguage(ctx, repos)
}

func (m *loggingMiddleware) GetReposByLanguage(ctx context.Context) (res *ReposByLanguage, err error) {
	defer func(start time.Time) {
		L := logger.Wrap(ctx, m.logger,
			logger.Method("GetReposByLanguage"),
			logger.Duration(start))

		logger.Maybe(L, "get repos by language", err)
	}(time.Now())

	return m.service.GetReposByLanguage(ctx)
}

func (m *loggingMiddleware) PostReposByLanguage(ctx context.Context, users []schema.UserCountByLanguage) (err error) {
	defer func(start time.Time) {
		L := logger.Wrap(ctx, m.logger,
			logger.Method("PostReposByLanguage"),
			logger.Duration(start))

		logger.Maybe(L, "post repos by language", err)
	}(time.Now())

	return m.service.PostReposByLanguage(ctx, users)
}

func (m *loggingMiddleware) GetCompanyCount(ctx context.Context) (res *CompanyCount, err error) {
	defer func(start time.Time) {
		L := logger.Wrap(ctx, m.logger,
			logger.Method("GetCompanyCount"),
			logger.Duration(start))

		logger.Maybe(L, "get company count", err)
	}(time.Now())

	return m.service.GetCompanyCount(ctx)
}

func (m *loggingMiddleware) PostCompanyCount(ctx context.Context, count int) (err error) {
	defer func(start time.Time) {
		L := logger.Wrap(ctx, m.logger,
			logger.Method("PostCompanyCount"),
			logger.Duration(start))

		logger.Maybe(L, "post company count", err)
	}(time.Now())

	return m.service.PostCompanyCount(ctx, count)
}

func (m *loggingMiddleware) GetUsersByCompany(ctx context.Context) (res *UsersByCompany, err error) {
	defer func(start time.Time) {
		L := logger.Wrap(ctx, m.logger,
			logger.Method("GetUsersByCompany"),
			logger.Duration(start))

		logger.Maybe(L, "get users by company", err)
	}(time.Now())

	return m.service.GetUsersByCompany(ctx)
}

func (m *loggingMiddleware) PostUsersByCompany(ctx context.Context, users []schema.Company) (err error) {
	defer func(start time.Time) {
		L := logger.Wrap(ctx, m.logger,
			logger.Method("PostUsersByCompany"),
			logger.Duration(start))

		logger.Maybe(L, "post users by company", err)
	}(time.Now())

	return m.service.PostUsersByCompany(ctx, users)
}
