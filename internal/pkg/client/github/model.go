package github

import (
	"context"

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

func (m *model) FetchUsersCursor(ctx context.Context, location, start, end string, limit int) ([]User, error) {
	zlog := logger.RequestIDFromContext(ctx)

	zlog.Info("fetching users",
		zap.String("location", location),
		zap.String("start", start),
		zap.String("end", end),
		zap.Int("limit", limit))

	cursor := ""
	hasNextPage := true
	var users []User
	for hasNextPage {
		res, err := m.store.FetchUsers(FetchUsersRequest{
			Location: location,
			Start:    start,
			End:      end,
			Cursor:   cursor,
			Limit:    limit,
		})
		if err != nil {
			zlog.Warn("error fetching users", zap.Error(err))
			break
		}
		hasNextPage = res.Data.Search.PageInfo.HasNextPage
		cursor = res.Data.Search.PageInfo.EndCursor
		for _, edge := range res.Data.Search.Edges {
			users = append(users, edge.Node)
		}
	}
	zlog.Info("fetched users",
		zap.Int("total", len(users)))
	return users, nil
}

func (m *model) FetchReposCursor(ctx context.Context, login, start, end string, limit int) ([]Repo, error) {
	zlog := logger.RequestIDFromContext(ctx)
	zlog.Info("fetching repos",
		zap.String("login", login),
		zap.String("start", start),
		zap.String("end", end),
		zap.Int("limit", limit))

	cursor := ""
	hasNextPage := true
	var repos []Repo
	for hasNextPage {
		res, err := m.store.FetchRepos(FetchReposRequest{
			Login:  login,
			Start:  start,
			End:    end,
			Cursor: cursor,
			Limit:  limit,
		})
		if err != nil {
			zlog.Warn("error fetching repos", zap.Error(err))
			break
		}
		hasNextPage = res.Data.Search.PageInfo.HasNextPage
		cursor = res.Data.Search.PageInfo.EndCursor
		for _, edge := range res.Data.Search.Edges {
			repos = append(repos, edge.Node)
		}
	}
	zlog.Info("fetched repos",
		zap.Int("total", len(repos)))
	return repos, nil
}
