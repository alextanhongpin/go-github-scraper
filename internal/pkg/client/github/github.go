package github

import (
	"context"
	"net/http"

	"github.com/alextanhongpin/go-github-scraper/internal/pkg/logger"
)

// API represents the interface for Github's API
type API interface {
	FetchUsersCursor(ctx context.Context, location, start, end string, limit int) ([]User, error)
	FetchReposCursor(ctx context.Context, login, start, end string, limit int) ([]Repo, error)
}

// New returns a new github api
func New(client *http.Client, token, endpoint string, l *logger.Logger) API {
	return NewModel(NewStore(client, token, endpoint), l)
}
