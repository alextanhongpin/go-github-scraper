package github

import (
	"net/http"
)

// New returns a new github api
func New(client *http.Client, token, endpoint string, middlewares ...Middleware) Service {
	store := NewStore(client, token, endpoint)
	model := NewModel(store)
	service := NewService(model)

	service = Decorate(service, middlewares...)
	return service
}
