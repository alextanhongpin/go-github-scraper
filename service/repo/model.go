package repo

import (
	"log"
	"time"

	"github.com/alextanhongpin/go-github-scraper/api/github"
)

// Model represents the interface for the repo business logic
type (
	Model interface {
		Init() error
		MostRecent(limit int) ([]Repo, error)
		MostStars(limit int) ([]Repo, error)
		Count() (int, error)
		MostPopularLanguage(limit int) ([]LanguageCount, error)
		RepoCountByUser(limit int) ([]UserCount, error)
		LanguageCountByUser(login string, limit int) ([]LanguageCount, error)
		MostRecentReposByLanguage(language string, limit int) ([]Repo, error)
		ReposByLanguage(language string, limit int) ([]Repo, error)
		BulkUpsert(repos []github.Repo) error
		FindLastCreatedByUser(login string) (string, bool)
	}

	model struct {
		store Store
	}
)

// NewModel returns a pointer to the Model
func NewModel(store Store) Model {
	return &model{
		store: store,
	}
}

func (m *model) Init() error {
	return m.store.Init()
}

func (m *model) MostRecent(limit int) ([]Repo, error) {
	return m.store.FindAll(limit, []string{"-updatedAt"})
}

func (m *model) MostStars(limit int) ([]Repo, error) {
	return m.store.FindAll(limit, []string{"-stargazers"})
}

func (m *model) Count() (int, error) {
	return m.store.Count()
}

func (m *model) MostPopularLanguage(limit int) ([]LanguageCount, error) {
	return m.store.AggregateLanguages(limit)
}

func (m *model) RepoCountByUser(limit int) ([]UserCount, error) {
	return m.store.AggregateReposByUser(limit)
}

func (m *model) LanguageCountByUser(login string, limit int) ([]LanguageCount, error) {
	return m.store.AggregateLanguageByUser(login, limit)
}

func (m *model) MostRecentReposByLanguage(language string, limit int) ([]Repo, error) {
	return m.store.AggregateMostRecentReposByLanguage(language, limit)
}

func (m *model) ReposByLanguage(language string, limit int) ([]Repo, error) {
	return m.store.AggregateReposByLanguage(language, limit)
}

func (m *model) BulkUpsert(repos []github.Repo) error {
	log.Printf("repo.model.BulkUpsert repoCount:%d\n", len(repos))
	return m.store.BulkUpsert(repos)
}

// FindLastCreatedByUser returns the date of the last created repo in the format 2006-01-02
func (m *model) FindLastCreatedByUser(login string) (string, bool) {
	log.Printf("repo.model.FindLastCreatedByUser login:%s\n", login)
	repo, err := m.store.FindLastCreatedByUser(login)
	if err != nil || repo == nil {
		// Github's creation date
		return "2008-04-01", false
	}
	t, err := time.Parse(time.RFC3339, repo.CreatedAt)
	if err != nil {
		return "2008-04-01", false
	}
	return t.Format("2006-01-02"), true
}
