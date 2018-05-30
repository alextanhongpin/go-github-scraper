package profilesvc

import (
	"net/http"

	"github.com/alextanhongpin/go-github-scraper/internal/pkg/encoder"
	"github.com/julienschmidt/httprouter"
)

type (
	// Endpoint represents the type alias for the httprouter.Handle
	Endpoint = httprouter.Handle

	// Endpoints represents the services exposed as http routes
	Endpoints interface {
		GetProfile() Endpoint
	}

	endpoints struct {
		svc Service
	}
)

// MakeEndpoints creates a new endpoint based on the service provided and router
func MakeEndpoints(svc Service, r *httprouter.Router) {
	e := &endpoints{svc: svc}

	r.GET("/profiles/:login", e.GetProfile())
}

func (e *endpoints) GetProfile() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		login := ps.ByName("login")
		res, err := e.svc.GetProfile(r.Context(), login)
		encoder.JSON(w, err, res)
	}
}
