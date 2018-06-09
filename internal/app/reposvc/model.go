package reposvc

// Models should contain validations, and performs business logic.
// Custom errors should also be thrown here.
// If the model does not fulfil the business requirements, it should not call the store.
// For orchestration (facade pattern), perform it at the service level

import (
	"errors"
	"log"
	"sort"
	"time"

	"github.com/alextanhongpin/go-github-scraper/internal/app/usersvc"
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/bow"
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/client/github"
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/constant"
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/schema"
)

var (
	ErrInvalidLanguage = errors.New("language provided is not valid")
)

type (
	// Model represents the interface for the repo service
	Model interface {
		BulkUpsert(repos []github.Repo) error
		Count() (int, error)
		Drop() error
		Init() error
		LastCreatedBy(login string) (string, bool)
		LanguageCountByUser(login string, limit int) ([]schema.LanguageCount, error)
		MostPopularLanguage(limit int) ([]schema.LanguageCount, error)
		MostRecent(limit int) ([]schema.Repo, error)
		MostRecentReposByLanguage(language string, limit int) ([]schema.Repo, error)
		MostStars(limit int) ([]schema.Repo, error)
		MostForks(limit int) ([]schema.Repo, error)
		RepoCountByUser(limit int) ([]schema.UserCount, error)
		ReposByLanguage(language string, limit int) ([]schema.UserCount, error)
		Distinct(field string) ([]string, error)
		GetProfile(login string) (*usersvc.User, error)
	}

	model struct {
		store Store
	}
)

// NewModel returns a pointer to the Model
func NewModel(store Store) Model {
	m := model{store: store}
	if err := m.Init(); err != nil {
		log.Fatal(err)
	}
	return &m
}

// BulkUpsert inserts a list of docs if they do not exists, or updates them if they exist and values differs
func (m *model) BulkUpsert(repos []github.Repo) error {
	if len(repos) == 0 {
		return nil
	}
	return m.store.BulkUpsert(repos)
}

// Count returns the total count of the repos
func (m *model) Count() (int, error) {
	return m.store.Count()
}

// Drop drops the collection
func (m *model) Drop() error {
	return m.store.Drop()
}

// Perform initialization of the service, such as setting up
// tables for the storage or indexes
func (m *model) Init() error {
	return m.store.Init()
}

// LastCreatedBy returns the last created datetime in the format YYYY-MM-DD for a particular user
func (m *model) LastCreatedBy(login string) (string, bool) {
	repo, err := m.store.LastCreatedBy(login)
	if err != nil || repo == nil {
		// Github's creation date
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

// LanguageCountByUser returns the top languages for a particular user
func (m *model) LanguageCountByUser(login string, limit int) ([]schema.LanguageCount, error) {
	return m.store.LanguagesBy(login, limit)
}

// MostPopularLanguage returns the most frequent language based on repo count in descending order
func (m *model) MostPopularLanguage(limit int) ([]schema.LanguageCount, error) {
	if limit < 1 {
		limit = 10
	}
	return m.store.Languages(limit)
}

// MostRecent returns a limited results of repo that are recently updated
func (m *model) MostRecent(limit int) ([]schema.Repo, error) {
	if limit < 1 {
		limit = 10
	}
	return m.store.FindAll(limit, []string{"-updatedAt"})
}

// MostRecentReposByLanguage returns the most recent repos that are updated for a given language
func (m *model) MostRecentReposByLanguage(language string, limit int) ([]schema.Repo, error) {
	if limit < 1 {
		limit = 10
	}
	if language == "" {
		return nil, ErrInvalidLanguage
	}
	return m.store.GroupByLanguageSortByMostRecent(language, limit)
}

// MostStars returns a limited results of repos with the most stars
func (m *model) MostStars(limit int) ([]schema.Repo, error) {
	if limit < 1 {
		limit = 10
	}
	return m.store.FindAll(limit, []string{"-stargazers"})
}

// MostForks returns a limited results of repos with the most forks
func (m *model) MostForks(limit int) ([]schema.Repo, error) {
	if limit < 1 {
		limit = 10
	}
	return m.store.FindAll(limit, []string{"-forks"})
}

// RepoCountByUser returns the users with most repos sorted in descending order
func (m *model) RepoCountByUser(limit int) ([]schema.UserCount, error) {
	if limit < 1 {
		limit = 10
	}
	return m.store.GroupByUser(limit)
}

// ReposByLanguage returns the users with the most repo in the particular language
func (m *model) ReposByLanguage(language string, limit int) ([]schema.UserCount, error) {
	if limit < 1 {
		limit = 10
	}

	if language == "" {
		return nil, ErrInvalidLanguage
	}
	return m.store.GroupByLanguage(language, limit)
}

// Distinct should return the distinct results for a given field
func (m *model) Distinct(field string) ([]string, error) {
	return m.store.Distinct(field)
}

// GetProfile should return a profile for a given login
func (m *model) GetProfile(login string) (*usersvc.User, error) {
	var watchers, stargazers, forks int64
	var languageList []string
	var descriptions []string

	var keywords []schema.Keyword
	var languages []schema.LanguageCount

	// TODO: Fix error here
	repos, err := m.store.ReposBy(login)
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
