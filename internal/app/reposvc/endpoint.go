package reposvc

import (
	"net/http"

	"github.com/alextanhongpin/go-github-scraper/internal/pkg/encoder"
	"github.com/julienschmidt/httprouter"
)

// Endpoints represents the services exposed as http routes
type (
	Endpoints interface {
		GetRepoCount() httprouter.Handle
	}

	endpoints struct {
		svc Service
	}
)

// MakeEndpoints creates a new endpoint based on the service provided and router
func MakeEndpoints(svc Service, r *httprouter.Router) {
	e := &endpoints{svc: svc}

	r.GET("/repos/count", e.GetRepoCount())
}

func (e *endpoints) GetRepoCount() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		count, err := e.svc.Count(r.Context())
		encoder.JSON(w, err, GetRepoCountResponse{count})
	}
}
