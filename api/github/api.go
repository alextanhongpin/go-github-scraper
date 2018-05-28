package github

import (
	"net/http"

	"go.uber.org/zap"
)

// API represents the interface for Github's API
type API interface {
	FetchUsersCursor(location, start, end string, limit int) ([]User, error)
	FetchReposCursor(login, start, end string, limit int) ([]Repo, error)
}

// New returns a new github api
func New(client *http.Client, token, endpoint string, zlog *zap.Logger) API {
	return NewModel(NewStore(client, token, endpoint), zlog)
}
