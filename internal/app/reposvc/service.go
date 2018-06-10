// Package reposvc defines the service for the repo microservice
// The interface defines the method that exists for the service
// and should have ctx as the first arguments to support cancellation/deadlines etc
// Logging, tracing etc should be done at the service level, not store
// Orchestration of models should also be done at service level (or hiding complex implentation
// of the models through facade) - the models are only responsible for validating the inputs
// that enters the store and also business logic
package reposvc

import (
	"context"
	"sort"
	"time"

	"github.com/alextanhongpin/go-github-scraper/internal/app/usersvc"
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/bow"
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/client/github"
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/constant"
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/schema"
)

// Service represents the interface for the repo service
type (
	Service interface {
		BulkUpsert(ctx context.Context, repos []github.Repo) error
		Count(ctx context.Context) (int, error)
		LastCreatedBy(ctx context.Context, login string) (string, bool)
		MostPopularLanguage(ctx context.Context, limit int) ([]schema.LanguageCount, error)
		MostRecent(ctx context.Context, limit int) ([]schema.Repo, error)
		MostRecentReposByLanguage(ctx context.Context, language string, limit int) ([]schema.Repo, error)
		MostStars(ctx context.Context, limit int) ([]schema.Repo, error)
		MostForks(ctx context.Context, limit int) ([]schema.Repo, error)
		RepoCountByUser(ctx context.Context, limit int) ([]schema.UserCount, error)
		ReposByLanguage(ctx context.Context, language string, limit int) ([]schema.UserCount, error)
		Distinct(ctx context.Context, login string) ([]string, error)
		GetProfile(ctx context.Context, login string) (*usersvc.User, error)
	}

	service struct {
		model Model
	}
)

// NewService returns a new service
func NewService(model Model) Service {
	return &service{
		model: model,
	}
}

func (s *service) BulkUpsert(ctx context.Context, repos []github.Repo) error {
	return s.model.BulkUpsert(repos)
}

func (s *service) Count(ctx context.Context) (int, error) {
	return s.model.Count()
}

func (s *service) LastCreatedBy(ctx context.Context, login string) (string, bool) {
	repo, err := s.model.LastCreatedBy(login)
	if err != nil || repo == nil {
		return constant.GithubCreatedAt, false
	}
	t, err := time.Parse(time.RFC3339, repo.CreatedAt)
	if err != nil {
		return constant.GithubCreatedAt, false
	}
	// Deduct a day
	t = t.AddDate(0, -1, 0)
	return t.Format("2006-01-02"), true
}

func (s *service) MostPopularLanguage(ctx context.Context, limit int) ([]schema.LanguageCount, error) {
	return s.model.MostPopularLanguage(limit)
}

func (s *service) MostRecent(ctx context.Context, limit int) ([]schema.Repo, error) {
	return s.model.MostRecent(limit)
}

func (s *service) MostRecentReposByLanguage(ctx context.Context, language string, limit int) ([]schema.Repo, error) {
	return s.model.MostRecentReposByLanguage(language, limit)
}

func (s *service) MostStars(ctx context.Context, limit int) ([]schema.Repo, error) {
	return s.model.MostStars(limit)
}

func (s *service) MostForks(ctx context.Context, limit int) ([]schema.Repo, error) {
	return s.model.MostForks(limit)
}

func (s *service) RepoCountByUser(ctx context.Context, limit int) ([]schema.UserCount, error) {
	return s.model.RepoCountByUser(limit)
}

func (s *service) ReposByLanguage(ctx context.Context, language string, limit int) ([]schema.UserCount, error) {
	return s.model.ReposByLanguage(language, limit)
}

func (s *service) Distinct(ctx context.Context, field string) ([]string, error) {
	return s.model.Distinct(field)
}

func (s *service) GetProfile(ctx context.Context, login string) (*usersvc.User, error) {
	var watchers, stargazers, forks int64
	var languageList []string
	var descriptions []string

	var keywords []schema.Keyword
	var languages []schema.LanguageCount

	repos, err := s.model.ReposBy(login)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(repos); i++ {
		repo := repos[i]
		stargazers += repo.Stargazers
		watchers += repo.Watchers
		forks += repo.Forks
		languageList = append(languageList, repo.Languages...)
		descriptions = append(descriptions, repo.Description)
	}

	topKeywords := bow.Top(bow.Parse(descriptions...), 20)
	for _, k := range topKeywords {
		keywords = append(keywords, schema.Keyword{ID: k.Key, Value: k.Value})
	}
	sort.SliceStable(keywords, func(i, j int) bool {
		return keywords[i].Value > keywords[j].Value
	})

	topLanguages := bow.Top(languageList, 20)
	for _, k := range topLanguages {
		languages = append(languages, schema.LanguageCount{Name: k.Key, Count: k.Value})
	}
	sort.SliceStable(languages, func(i, j int) bool {
		return languages[i].Count > languages[j].Count
	})

	return &usersvc.User{
		Login: login,
		Profile: schema.Profile{
			Watchers:   watchers,
			Stargazers: stargazers,
			Forks:      forks,
			Keywords:   keywords,
			Languages:  languages,
		},
	}, nil
}
