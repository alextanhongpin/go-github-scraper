package transport

import (
	"net/http"

	"github.com/alextanhongpin/go-github-scraper/internal/app/usersvc"
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/encoder"

	"github.com/julienschmidt/httprouter"
)

// Endpoints represents the services exposed as http routes
type userEndpoints struct {
	service usersvc.Service
}

// NewUserEndpoints creates a new endpoint
func NewUserEndpoints(s usersvc.Service) Endpoints {
	return &userEndpoints{s}
}

func (e *userEndpoints) GetUserCount() Endpoint {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		ctx := r.Context()
		count, err := e.service.Count(ctx)
		encoder.JSON(w, err, Data{
			"count": count,
		})
	}
}

func (e *userEndpoints) GetUser() Endpoint {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		ctx := r.Context()
		login := ps.ByName("login")
		user, err := e.service.FindOne(ctx, login)
		encoder.JSON(w, err, user)
	}
}

func (e *userEndpoints) Wrap(r *httprouter.Router) {
	r.GET("/users/:login", e.GetUser())
	r.GET("/users", e.GetUserCount())
}
