package usersvc

import (
	"net/http"

	"github.com/alextanhongpin/go-github-scraper/internal/pkg/encoder"

	"github.com/julienschmidt/httprouter"
)

// Endpoints represents the services exposed as http routes
type (
	Endpoints interface {
		GetUser() httprouter.Handle
		GetUserCount() httprouter.Handle
	}

	endpoints struct {
		svc Service
	}
)

// NewEndpoints creates a new endpoint
func NewEndpoints(svc Service) Endpoints {
	return &endpoints{svc: svc}
}

// MakeEndpoints creates a new endpoint based on the service provided and router
func MakeEndpoints(svc Service, r *httprouter.Router) {
	e := NewEndpoints(svc)

	r.GET("/users/:login", e.GetUser())
	r.GET("/users", e.GetUserCount())
}

func (e *endpoints) GetUserCount() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		ctx := r.Context()
		count, err := e.svc.Count(ctx)
		encoder.JSON(w, err, GetUserCountResponse{count})
	}
}

func (e *endpoints) GetUser() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		ctx := r.Context()
		login := ps.ByName("login")
		user, err := e.svc.FindOne(ctx, login)
		encoder.JSON(w, err, user)
	}
}
