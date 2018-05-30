package github

import (
	"context"
	"time"

	"github.com/alextanhongpin/go-github-scraper/internal/pkg/logger"

	"go.uber.org/zap"
)

// Model represents the api interface for the Github's GraphQL
type model struct {
	store Store
	zlog  *zap.Logger
}

// NewModel returns a new model
func NewModel(store Store, zlog *zap.Logger) API {
	return &model{
		store: store,
		zlog:  zlog,
	}
}

func (m *model) FetchUsersCursor(ctx context.Context, location, start, end string, limit int) (users []User, err error) {
	defer func(s time.Time) {
		zlog := logger.RequestIDFromContext(ctx).
			With(zap.String("method", "FetchUsersCursor"),
				zap.Duration("took", time.Since(s)),
				zap.String("location", location),
				zap.String("start", start),
				zap.String("end", end),
				zap.Int("limit", limit))
		if err != nil {
			zlog.Warn("error fetching users cursor", zap.Error(err))
		} else {
			zlog.Info("fetch users cursor", zap.Int("count", len(users)))
		}
	}(time.Now())

	var res *FetchUsersResponse
	cursor := ""
	hasNextPage := true
	for hasNextPage {
		res, err = m.store.FetchUsers(FetchUsersRequest{
			Location: location,
			Start:    start,
			End:      end,
			Cursor:   cursor,
			Limit:    limit,
		})
		if err != nil {
			break
		}
		hasNextPage = res.Data.Search.PageInfo.HasNextPage
		cursor = res.Data.Search.PageInfo.EndCursor
		for _, edge := range res.Data.Search.Edges {
			users = append(users, edge.Node)
		}
	}
	return
}

func (m *model) FetchReposCursor(ctx context.Context, login, start, end string, limit int) (repos []Repo, err error) {
	defer func(s time.Time) {
		zlog := logger.RequestIDFromContext(ctx).
			With(zap.String("method", "FetchReposCursor"),
				zap.Duration("took", time.Since(s)),
				zap.String("login", login),
				zap.String("start", start),
				zap.String("end", end),
				zap.Int("limit", limit))
		if err != nil {
			zlog.Warn("error fetching repos cursor", zap.Error(err))
		} else {
			zlog.Info("fetch repos cursor", zap.Int("count", len(repos)))
		}
	}(time.Now())

	var res *FetchReposResponse
	cursor := ""
	hasNextPage := true
	for hasNextPage {
		res, err = m.store.FetchRepos(FetchReposRequest{
			Login:  login,
			Start:  start,
			End:    end,
			Cursor: cursor,
			Limit:  limit,
		})
		if err != nil {
			break
		}
		hasNextPage = res.Data.Search.PageInfo.HasNextPage
		cursor = res.Data.Search.PageInfo.EndCursor
		for _, edge := range res.Data.Search.Edges {
			repos = append(repos, edge.Node)
		}
	}
	return
}
