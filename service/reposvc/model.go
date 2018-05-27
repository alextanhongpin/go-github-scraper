package reposvc

import (
	"log"
	"time"

	"github.com/alextanhongpin/go-github-scraper/api/github"
	"github.com/alextanhongpin/go-github-scraper/internal/schema"
)

// Model represents the interface for the repo business logic
type model struct {
	store Store
}

// NewModel returns a pointer to the Model
func NewModel(store Store) Service {
	m := model{store: store}
	if err := m.Init(); err != nil {
		log.Fatal(err)
	}
	return &m
}

// BulkUpsert inserts a list of docs if they do not exists, or updates them if they exist and values differs
func (m *model) BulkUpsert(repos []github.Repo) error {
	log.Printf("repo.model.BulkUpsert repoCount:%d\n", len(repos))
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

// FindLastCreatedByUser returns the last created datetime in the format YYYY-MM-DD for a particular user
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

// LanguageCountByUser returns the top languages for a particular user
func (m *model) LanguageCountByUser(login string, limit int) ([]schema.LanguageCount, error) {
	return m.store.AggregateLanguageByUser(login, limit)
}

// MostPopularLanguage returns the most frequent language based on repo count in descending order
func (m *model) MostPopularLanguage(limit int) ([]schema.LanguageCount, error) {
	return m.store.AggregateLanguages(limit)
}

// MostRecent returns a limited results of repo that are recently updated
func (m *model) MostRecent(limit int) ([]schema.Repo, error) {
	return m.store.FindAll(limit, []string{"-updatedAt"})
}

// MostRecentReposByLanguage returns the most recent repos that are updated for a given language
func (m *model) MostRecentReposByLanguage(language string, limit int) ([]schema.Repo, error) {
	return m.store.AggregateMostRecentReposByLanguage(language, limit)
}

// MostStars returns a limited results of repos with the most stars
func (m *model) MostStars(limit int) ([]schema.Repo, error) {
	return m.store.FindAll(limit, []string{"-stargazers"})
}

// RepoCountByUser returns the users with most repos sorted in descending order
func (m *model) RepoCountByUser(limit int) ([]schema.UserCount, error) {
	return m.store.AggregateReposByUser(limit)
}

// ReposByLanguage returns the users with the most repo in the particular language
func (m *model) ReposByLanguage(language string, limit int) ([]schema.UserCount, error) {
	return m.store.AggregateReposByLanguage(language, limit)
}

func (m *model) WatchersFor(login string) (int64, error) {
	return m.store.WatchersFor(login)
}

func (m *model) StargazersFor(login string) (int64, error) {
	return m.store.StargazersFor(login)
}

func (m *model) ForksFor(login string) (int64, error) {
	return m.store.ForksFor(login)
}

func (m *model) WordCount(login string, limit int) ([]WordCount, error) {
	return m.store.WordCount(login, limit)
}
