package statsvc

import (
	"context"

	"github.com/alextanhongpin/go-github-scraper/internal/pkg/schema"
)

type (
	// Service represents the analytic service
	Service interface {
		GetUserCount(ctx context.Context) (*UserCount, error)
		PostUserCount(ctx context.Context, count int) error
		GetRepoCount(ctx context.Context) (*RepoCount, error)
		PostRepoCount(ctx context.Context, count int) error
		GetReposMostRecent(ctx context.Context) (*ReposMostRecent, error)
		PostReposMostRecent(ctx context.Context, data []schema.Repo) error
		GetRepoCountByUser(ctx context.Context) (*RepoCountByUser, error)
		PostRepoCountByUser(ctx context.Context, users []schema.UserCount) error
		GetReposMostStars(ctx context.Context) (*ReposMostStars, error)
		PostReposMostStars(ctx context.Context, repos []schema.Repo) error
		GetReposMostForks(ctx context.Context) (*ReposMostForks, error)
		PostReposMostForks(ctx context.Context, repos []schema.Repo) error
		GetMostPopularLanguage(ctx context.Context) (*MostPopularLanguage, error)
		PostMostPopularLanguage(ctx context.Context, languages []schema.LanguageCount) error
		GetLanguageCountByUser(ctx context.Context) (*LanguageCountByUser, error)
		PostLanguageCountByUser(ctx context.Context, languages []schema.LanguageCount) error
		GetMostRecentReposByLanguage(ctx context.Context) (*MostRecentReposByLanguage, error)
		PostMostRecentReposByLanguage(ctx context.Context, repos []schema.RepoLanguage) error
		GetReposByLanguage(ctx context.Context) (*ReposByLanguage, error)
		PostReposByLanguage(ctx context.Context, users []schema.UserCountByLanguage) error
		GetCompanyCount(ctx context.Context) (*CompanyCount, error)
		PostCompanyCount(ctx context.Context, count int) error
		GetUsersByCompany(ctx context.Context) (*UsersByCompany, error)
		PostUsersByCompany(ctx context.Context, users []schema.Company) error
	}

	service struct {
		model Model
	}
)

func NewService(m Model) Service {
	return &service{m}
}

func (s *service) GetUserCount(ctx context.Context) (*UserCount, error) {
	return s.model.GetUserCount()
}

func (s *service) PostUserCount(ctx context.Context, count int) error {
	return s.model.PostUserCount(count)
}

func (s *service) GetRepoCount(ctx context.Context) (*RepoCount, error) {
	return s.model.GetRepoCount()
}

func (s *service) PostRepoCount(ctx context.Context, count int) error {
	return s.model.PostRepoCount(count)
}

func (s *service) GetReposMostRecent(ctx context.Context) (*ReposMostRecent, error) {
	return s.model.GetReposMostRecent()
}

func (s *service) PostReposMostRecent(ctx context.Context, data []schema.Repo) error {
	return s.model.PostReposMostRecent(data)
}

func (s *service) GetRepoCountByUser(ctx context.Context) (*RepoCountByUser, error) {
	return s.model.GetRepoCountByUser()
}

func (s *service) PostRepoCountByUser(ctx context.Context, users []schema.UserCount) error {
	return s.model.PostRepoCountByUser(users)
}

func (s *service) GetReposMostStars(ctx context.Context) (*ReposMostStars, error) {
	return s.model.GetReposMostStars()
}

func (s *service) PostReposMostStars(ctx context.Context, repos []schema.Repo) error {
	return s.model.PostReposMostStars(repos)
}

func (s *service) GetReposMostForks(ctx context.Context) (*ReposMostForks, error) {
	return s.model.GetReposMostForks()
}

func (s *service) PostReposMostForks(ctx context.Context, repos []schema.Repo) error {
	return s.model.PostReposMostForks(repos)
}

func (s *service) GetMostPopularLanguage(ctx context.Context) (*MostPopularLanguage, error) {
	return s.model.GetMostPopularLanguage()
}

func (s *service) PostMostPopularLanguage(ctx context.Context, languages []schema.LanguageCount) error {
	return s.model.PostMostPopularLanguage(languages)
}

func (s *service) GetLanguageCountByUser(ctx context.Context) (*LanguageCountByUser, error) {
	return s.model.GetLanguageCountByUser()
}

func (s *service) PostLanguageCountByUser(ctx context.Context, languages []schema.LanguageCount) error {
	return s.model.PostLanguageCountByUser(languages)
}

func (s *service) GetMostRecentReposByLanguage(ctx context.Context) (*MostRecentReposByLanguage, error) {
	return s.model.GetMostRecentReposByLanguage()
}

func (s *service) PostMostRecentReposByLanguage(ctx context.Context, repos []schema.RepoLanguage) error {
	return s.model.PostMostRecentReposByLanguage(repos)
}

func (s *service) GetReposByLanguage(ctx context.Context) (*ReposByLanguage, error) {
	return s.model.GetReposByLanguage()
}

func (s *service) PostReposByLanguage(ctx context.Context, users []schema.UserCountByLanguage) error {
	return s.model.PostReposByLanguage(users)
}

func (s *service) GetCompanyCount(ctx context.Context) (*CompanyCount, error) {
	return s.model.GetCompanyCount()
}

func (s *service) PostCompanyCount(ctx context.Context, count int) error {
	return s.model.PostCompanyCount(count)
}

func (s *service) GetUsersByCompany(ctx context.Context) (*UsersByCompany, error) {
	return s.model.GetUsersByCompany()
}

func (s *service) PostUsersByCompany(ctx context.Context, users []schema.Company) error {
	return s.model.PostUsersByCompany(users)
}
