package statsvc

import (
	"log"

	"github.com/alextanhongpin/go-github-scraper/internal/pkg/schema"
)

// Model represents the interface for the analytic business logic
type (
	Model interface {
		Init() error
		GetUserCount() (*UserCount, error)
		PostUserCount(count int) error
		GetRepoCount() (*RepoCount, error)
		PostRepoCount(count int) error
		GetReposMostRecent() (*ReposMostRecent, error)
		PostReposMostRecent(data []schema.Repo) error
		GetRepoCountByUser() (*RepoCountByUser, error)
		PostRepoCountByUser(users []schema.UserCount) error
		GetReposMostStars() (*ReposMostStars, error)
		PostReposMostStars(repos []schema.Repo) error
		GetReposMostForks() (*ReposMostForks, error)
		PostReposMostForks(repos []schema.Repo) error
		GetMostPopularLanguage() (*MostPopularLanguage, error)
		PostMostPopularLanguage(languages []schema.LanguageCount) error
		GetLanguageCountByUser() (*LanguageCountByUser, error)
		PostLanguageCountByUser(languages []schema.LanguageCount) error
		GetMostRecentReposByLanguage() (*MostRecentReposByLanguage, error)
		PostMostRecentReposByLanguage(repos []schema.RepoLanguage) error
		GetReposByLanguage() (*ReposByLanguage, error)
		PostReposByLanguage(users []schema.UserCountByLanguage) error
		GetCompanyCount() (*CompanyCount, error)
		PostCompanyCount(count int) error
		GetUsersByCompany() (*UsersByCompany, error)
		PostUsersByCompany(users []schema.Company) error
	}

	model struct {
		store Store
	}
)

// NewModel returns a new analytic model
func NewModel(s Store) Model {
	m := model{
		store: s,
	}
	if err := m.Init(); err != nil {
		log.Fatal(err)
	}
	return &m
}

func (m *model) Init() error {
	return m.store.Init()
}

func (m *model) GetUserCount() (*UserCount, error) {
	return m.store.GetUserCount()
}

func (m *model) PostUserCount(count int) error {
	return m.store.PostUserCount(count)
}

func (m *model) GetRepoCount() (*RepoCount, error) {
	return m.store.GetRepoCount()
}

func (m *model) PostRepoCount(count int) error {
	return m.store.PostRepoCount(count)
}

func (m *model) GetReposMostRecent() (*ReposMostRecent, error) {
	return m.store.GetReposMostRecent()
}

func (m *model) PostReposMostRecent(repos []schema.Repo) error {
	if len(repos) == 0 {
		return nil
	}
	return m.store.PostReposMostRecent(repos)
}

func (m *model) GetRepoCountByUser() (*RepoCountByUser, error) {
	return m.store.GetRepoCountByUser()
}

func (m *model) PostRepoCountByUser(users []schema.UserCount) error {
	if len(users) == 0 {
		return nil
	}
	return m.store.PostRepoCountByUser(users)
}

func (m *model) GetReposMostStars() (*ReposMostStars, error) {
	return m.store.GetReposMostStars()
}

func (m *model) PostReposMostStars(repos []schema.Repo) error {
	if len(repos) == 0 {
		return nil
	}
	return m.store.PostReposMostStars(repos)
}

func (m *model) GetReposMostForks() (*ReposMostForks, error) {
	return m.store.GetReposMostForks()
}

func (m *model) PostReposMostForks(repos []schema.Repo) error {
	if len(repos) == 0 {
		return nil
	}
	return m.store.PostReposMostForks(repos)
}

func (m *model) GetMostPopularLanguage() (*MostPopularLanguage, error) {
	return m.store.GetMostPopularLanguage()
}

func (m *model) PostMostPopularLanguage(languages []schema.LanguageCount) error {
	if len(languages) == 0 {
		return nil
	}
	return m.store.PostMostPopularLanguage(languages)
}

func (m *model) GetLanguageCountByUser() (*LanguageCountByUser, error) {
	return m.store.GetLanguageCountByUser()
}

func (m *model) PostLanguageCountByUser(languages []schema.LanguageCount) error {
	if len(languages) == 0 {
		return nil
	}
	return m.store.PostLanguageCountByUser(languages)
}

func (m *model) GetMostRecentReposByLanguage() (*MostRecentReposByLanguage, error) {
	return m.store.GetMostRecentReposByLanguage()
}

func (m *model) PostMostRecentReposByLanguage(repos []schema.RepoLanguage) error {
	if len(repos) == 0 {
		return nil
	}
	return m.store.PostMostRecentReposByLanguage(repos)
}

func (m *model) GetReposByLanguage() (*ReposByLanguage, error) {
	return m.store.GetReposByLanguage()
}

func (m *model) PostReposByLanguage(users []schema.UserCountByLanguage) error {
	if len(users) == 0 {
		return nil
	}
	return m.store.PostReposByLanguage(users)
}

func (m *model) GetCompanyCount() (*CompanyCount, error) {
	return m.store.GetCompanyCount()
}

func (m *model) PostCompanyCount(count int) error {
	return m.store.PostCompanyCount(count)
}

func (m *model) GetUsersByCompany() (*UsersByCompany, error) {
	return m.store.GetUsersByCompany()
}

func (m *model) PostUsersByCompany(users []schema.Company) error {
	if len(users) == 0 {
		return nil
	}
	return m.store.PostUsersByCompany(users)
}
