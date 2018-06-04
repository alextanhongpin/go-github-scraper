package statsvc

import (
	"net/http"

	"github.com/alextanhongpin/go-github-scraper/internal/pkg/encoder"
	"github.com/julienschmidt/httprouter"
)

// Endpoints represents the services exposed as http routes
type (
	Endpoints interface {
		GetStats() httprouter.Handle
	}

	endpoints struct {
		svc Service
	}
)

// MakeEndpoints creates a new set of endpoints based on the service provided and router
func MakeEndpoints(svc Service, r *httprouter.Router) {
	e := &endpoints{svc: svc}

	r.GET("/stats", e.GetStats())
}

func (e *endpoints) GetStats() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		ctx := r.Context()
		var res interface{}
		var err error
		switch r.URL.Query().Get("type") {
		case EnumUserCount:
			res, err = e.svc.GetUserCount(ctx)
		case EnumRepoCount:
			res, err = e.svc.GetRepoCount(ctx)
		case EnumReposMostRecent:
			res, err = e.svc.GetReposMostRecent(ctx)
		case EnumRepoCountByUser:
			res, err = e.svc.GetRepoCountByUser(ctx)
		case EnumReposMostStars:
			res, err = e.svc.GetReposMostStars(ctx)
		case EnumReposMostForks:
			res, err = e.svc.GetReposMostForks(ctx)
		case EnumMostPopularLanguage:
			res, err = e.svc.GetMostPopularLanguage(ctx)
		case EnumMostRecentReposByLanguage:
			res, err = e.svc.GetMostRecentReposByLanguage(ctx)
		case EnumReposByLanguage:
			res, err = e.svc.GetReposByLanguage(ctx)
		case EnumCompanyCount:
			res, err = e.svc.GetCompanyCount(ctx)
		case EnumUsersByCompany:
			res, err = e.svc.GetUsersByCompany(ctx)
		default:
			res = IndexResponse{
				Paths: []string{
					"/stats?type=user_count",
					"/stats?type=repo_count",
					"/stats?type=company_count",
					"/stats?type=repos_most_recent",
					"/stats?type=repo_count_by_user",
					"/stats?type=repos_most_stars",
					"/stats?type=repos_most_forks",
					"/stats?type=users_by_company",
					"/stats?type=languages_most_popular",
					"/stats?type=repos_most_recent_by_language",
					"/stats?type=repos_by_language",
				},
			}
		}
		encoder.JSON(w, err, res)
	}
}
