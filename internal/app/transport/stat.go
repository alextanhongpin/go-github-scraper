package transport

import (
	"net/http"

	"github.com/alextanhongpin/go-github-scraper/internal/app/statsvc"
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/encoder"

	"github.com/julienschmidt/httprouter"
)

// Endpoints represents the services exposed as http routes
type statEndpoints struct {
	service statsvc.Service
}

// NewStatEndpoints creates a new set of endpoints based on the service provided and router
func NewStatEndpoints(s statsvc.Service) Endpoints {
	return &statEndpoints{s}
}

func (e *statEndpoints) GetStats() Endpoint {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		ctx := r.Context()
		var res interface{}
		var err error
		switch r.URL.Query().Get("type") {
		case statsvc.EnumUserCount:
			res, err = e.service.GetUserCount(ctx)
		case statsvc.EnumRepoCount:
			res, err = e.service.GetRepoCount(ctx)
		case statsvc.EnumReposMostRecent:
			res, err = e.service.GetReposMostRecent(ctx)
		case statsvc.EnumRepoCountByUser:
			res, err = e.service.GetRepoCountByUser(ctx)
		case statsvc.EnumReposMostStars:
			res, err = e.service.GetReposMostStars(ctx)
		case statsvc.EnumReposMostForks:
			res, err = e.service.GetReposMostForks(ctx)
		case statsvc.EnumMostPopularLanguage:
			res, err = e.service.GetMostPopularLanguage(ctx)
		case statsvc.EnumMostRecentReposByLanguage:
			res, err = e.service.GetMostRecentReposByLanguage(ctx)
		case statsvc.EnumReposByLanguage:
			res, err = e.service.GetReposByLanguage(ctx)
		case statsvc.EnumCompanyCount:
			res, err = e.service.GetCompanyCount(ctx)
		case statsvc.EnumUsersByCompany:
			res, err = e.service.GetUsersByCompany(ctx)
		default:
			res = Data{
				"paths": []string{
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

func (e *statEndpoints) Wrap(r *httprouter.Router) {
	r.GET("/stats", e.GetStats())
}
