package transport

import (
	"net/http"

	"github.com/alextanhongpin/go-github-scraper/internal/app/reposvc"
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/encoder"

	"github.com/julienschmidt/httprouter"
)

type repoEndpoints struct {
	service reposvc.Service
}

// NewRepoEndpoints creates a new endpoint based on the service provided and router
func NewRepoEndpoints(s reposvc.Service) Endpoints {
	return &repoEndpoints{s}
}

func (e *repoEndpoints) GetRepoCount() Endpoint {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		ctx := r.Context()
		count, err := e.service.Count(ctx)

		encoder.JSON(w, err, Data{
			"data": count,
		})
	}
}

func (e *repoEndpoints) Wrap(r *httprouter.Router) {
	r.GET("/repos/count", e.GetRepoCount())
}
