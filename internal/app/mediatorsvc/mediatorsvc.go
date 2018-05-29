// Package mediatorsvc handles multiple service orchestrations
// and is categorized as domain model
package mediatorsvc

import (
	"context"
	"sync"
	"time"

	"github.com/alextanhongpin/go-github-scraper/internal/app/profilesvc"
	"github.com/alextanhongpin/go-github-scraper/internal/app/reposvc"
	"github.com/alextanhongpin/go-github-scraper/internal/app/statsvc"
	"github.com/alextanhongpin/go-github-scraper/internal/app/usersvc"
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/client/github"
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/logger"
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/moment"
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/schema"

	"go.uber.org/zap"
)

type (
	// Service represents the methods the mediator service must implement
	Service interface {
		FetchUsers(ctx context.Context, location string, months int, perPage int) error
		FetchRepos(ctx context.Context, userPerPage, repoPerPage int) error
		UpdateUserCount(ctx context.Context) error
		UpdateRepoCount(ctx context.Context) error
		UpdateReposMostRecent(ctx context.Context, perPage int) error
		UpdateRepoCountByUser(ctx context.Context, perPage int) error
		UpdateReposMostStars(ctx context.Context, perPage int) error
		UpdateLanguagesMostPopular(ctx context.Context, perPage int) error
		UpdateMostRecentReposByLanguage(ctx context.Context, perPage int) error
		UpdateReposByLanguage(ctx context.Context, perPage int) error
		UpdateProfile(ctx context.Context, numWorkers int) error
	}

	// Mediator holds the services in used
	Mediator struct {
		Analytic statsvc.Service
		Github   github.API
		Profile  profilesvc.Service
		Repo     reposvc.Service
		User     usersvc.Service
	}

	model struct {
		Mediator
	}
)

// New returns a new mediator service
func New(m Mediator) Service {
	return &model{
		m,
	}
}

// makeEndDate sets the end date n months away to max present day
func makeEndDate(start string, months int) string {
	t1, err := time.Parse("2006-01-02", start)
	if err != nil {
		return time.Now().Format("2006-01-02")
	}
	t2 := t1.Add(time.Duration(months) * 30 * 24 * time.Hour)
	if t2.Unix() > time.Now().Unix() {
		return time.Now().Format("2006-01-02")
	}
	return t2.Format("2006-01-02")
}

func (m *model) FetchUsers(ctx context.Context, location string, months int, perPage int) error {
	zlog := logger.RequestIDFromContext(ctx)

	start, _ := m.User.FindLastCreated(ctx)
	end := makeEndDate(start, months)

	users, err := m.Github.FetchUsersCursor(ctx, location, start, end, perPage)
	if err != nil {
		zlog.Warn("error fetching users", zap.Error(err))
		return err
	}

	// Note that bulk upsert can only hold 1000 users max
	if err := m.User.BulkUpsert(ctx, users); err != nil {
		zlog.Warn("error upserting users", zap.Error(err))
		return err
	}
	return nil
}

func (m *model) FetchRepos(ctx context.Context, userPerPage, repoPerPage int) error {
	zlog := logger.RequestIDFromContext(ctx)

	users, err := m.User.FindLastFetched(ctx, userPerPage)
	if err != nil {
		zlog.Error("unable to find last fetched users", zap.Error(err))
		return err
	}

	for _, user := range users {
		login := user.Login
		if login == "" {
			continue
		}

		start, _ := m.Repo.FindLastCreatedByUser(ctx, login)
		end := moment.NewCurrentFormattedDate()

		repos, err := m.Github.FetchReposCursor(ctx, login, start, end, repoPerPage)
		if err != nil {
			zlog.Warn("error fetching repos", zap.Error(err))
			return err
		}

		if err := m.Repo.BulkUpsert(ctx, repos); err != nil {
			zlog.Warn("error upserting repos", zap.Error(err))
			return err
		}

		if err = m.User.UpdateOne(ctx, login); err != nil {
			zlog.Warn("UpdateOne",
				zap.Bool("error", true),
				zap.String("login", login),
				zap.Error(err))
			return err
		}
	}
	return nil
}

// UpdateUserCount updates the analytic type `user_count`
func (m *model) UpdateUserCount(ctx context.Context) error {
	zlog := logger.RequestIDFromContext(ctx)

	count, err := m.User.Count(ctx)
	if err != nil {
		zlog.Warn("error getting user count", zap.Error(err))
		return err
	}
	if err := m.Analytic.PostUserCount(ctx, count); err != nil {
		zlog.Warn("error updating user count", zap.Error(err))
		return err
	}
	zlog.Info("updated user count", zap.Int("count", count))
	return nil
}

// UpdateRepoCount updates the analytic type `repo_count`
func (m *model) UpdateRepoCount(ctx context.Context) error {
	zlog := logger.RequestIDFromContext(ctx)

	count, err := m.Repo.Count(ctx)
	if err != nil {
		zlog.Warn("error getting repo count", zap.Error(err))
		return err
	}
	if err := m.Analytic.PostRepoCount(ctx, count); err != nil {
		zlog.Warn("error updating repo count", zap.Error(err))
		return err
	}
	zlog.Info("updated repo count", zap.Int("count", count))
	return nil
}

// UpdateReposMostRecent updates the analytic type `repos_most_recent`
func (m *model) UpdateReposMostRecent(ctx context.Context, perPage int) error {
	zlog := logger.RequestIDFromContext(ctx)

	repos, err := m.Repo.MostRecent(ctx, perPage)
	if err != nil {
		zlog.Warn("error fetching most recent repos", zap.Error(err))
		return err
	}

	if err := m.Analytic.PostReposMostRecent(ctx, repos); err != nil {
		zlog.Warn("error updating most recent repos", zap.Error(err))
		return err
	}
	zlog.Info("updated most recent repos", zap.Int("count", len(repos)))
	return nil
}

// UpdateRepoCountByUser updates the analytic type `repo_count_by_user`
func (m *model) UpdateRepoCountByUser(ctx context.Context, perPage int) error {
	zlog := logger.RequestIDFromContext(ctx)

	users, err := m.Repo.RepoCountByUser(ctx, perPage)
	if err != nil {
		zlog.Warn("error fetching repo count by users", zap.Error(err))
		return err
	}

	if err := m.Analytic.PostRepoCountByUser(ctx, users); err != nil {
		zlog.Warn("error updating repo count by users", zap.Error(err))
		return err
	}
	zlog.Info("updated repo count by users", zap.Int("count", len(users)))
	return nil
}

// UpdateReposMostStars updates the analytic type `repos_most_stars`
func (m *model) UpdateReposMostStars(ctx context.Context, perPage int) error {
	zlog := logger.RequestIDFromContext(ctx)

	repos, err := m.Repo.MostStars(ctx, perPage)
	if err != nil {
		zlog.Warn("error fetching most stars repos", zap.Error(err))
		return err
	}

	if err := m.Analytic.PostReposMostStars(ctx, repos); err != nil {
		zlog.Warn("error updating repo count by users", zap.Error(err))
		return err
	}
	zlog.Info("updated repos most stars", zap.Int("count", len(repos)))
	return nil
}

// UpdateLanguagesMostPopular updates the analytic type `languages_most_popular`
func (m *model) UpdateLanguagesMostPopular(ctx context.Context, perPage int) error {
	zlog := logger.RequestIDFromContext(ctx)

	languages, err := m.Repo.MostPopularLanguage(ctx, perPage)
	if err != nil {
		zlog.Warn("error fetching language most popular", zap.Error(err))
		return err
	}

	if err := m.Analytic.PostMostPopularLanguage(ctx, languages); err != nil {
		zlog.Warn("error updating language most popular", zap.Error(err))
		return err
	}
	zlog.Info("updated popular languages", zap.Int("count", len(languages)))
	return nil
}

// UpdateMostRecentReposByLanguage updates the analytic type `repos_most_recent_by_language`
func (m *model) UpdateMostRecentReposByLanguage(ctx context.Context, perPage int) error {
	zlog := logger.RequestIDFromContext(ctx)

	languages, err := m.Repo.MostPopularLanguage(ctx, perPage)
	if err != nil {
		zlog.Warn("error fetching language most popular", zap.Error(err))
		return err
	}

	var reposByLanguages []schema.RepoLanguage
	for _, lang := range languages {
		r, err := m.Repo.MostRecentReposByLanguage(ctx, lang.Name, perPage)
		if err != nil {
			zlog.Warn("error fetching most recent repos by language", zap.Error(err))
			return err
		}
		reposByLanguages = append(reposByLanguages, schema.RepoLanguage{
			Language: lang.Name,
			Repos:    r,
		})
	}

	if err := m.Analytic.PostMostRecentReposByLanguage(ctx, reposByLanguages); err != nil {
		zlog.Warn("error updating most recent repos by language", zap.Error(err))
		return err
	}
	zlog.Info("updated most recent repos by language", zap.Int("count", len(reposByLanguages)))
	return nil
}

// UpdateReposByLanguage updates the analytic type `repos_by_language`
func (m *model) UpdateReposByLanguage(ctx context.Context, perPage int) error {
	zlog := logger.RequestIDFromContext(ctx)

	languages, err := m.Repo.MostPopularLanguage(ctx, perPage)
	if err != nil {
		zlog.Warn("error fetching language most popular", zap.Error(err))
		return err
	}

	var userCountByLanguage []schema.UserCountByLanguage
	for _, lang := range languages {
		users, err := m.Repo.ReposByLanguage(ctx, lang.Name, perPage)
		if err != nil {
			zlog.Warn("error fetching user repo count by language", zap.Error(err))
			return err
		}

		userCountByLanguage = append(userCountByLanguage, schema.UserCountByLanguage{
			Language: lang.Name,
			Users:    users,
		})
	}

	if err := m.Analytic.PostReposByLanguage(ctx, userCountByLanguage); err != nil {
		zlog.Warn("error updating user repo count by language", zap.Error(err))
		return err
	}

	zlog.Info("updated repos by languages", zap.Int("count", len(userCountByLanguage)))
	return nil
}

func (m *model) UpdateProfile(ctx context.Context, numWorkers int) error {
	zlog := logger.RequestIDFromContext(ctx)

	logins, err := m.Repo.DistinctLogin(ctx)
	if err != nil {
		zlog.Warn("error getting distinct login", zap.Error(err))
		return err
	}

	zlog.Info("got distinct logins",
		zap.Int("count", len(logins)))

	toStream := func(ctx context.Context, args ...string) <-chan string {
		c := make(chan string)
		go func() {
			defer close(c)
			for _, i := range args {
				select {
				case <-ctx.Done():
					return
				case c <- i:
				}
			}
		}()
		return c
	}

	fanIn := func(ctx context.Context, numWorkers int, in <-chan string) <-chan schema.Profile {
		c := make(chan schema.Profile)

		var wg sync.WaitGroup
		wg.Add(numWorkers)

		multiplex := func(in <-chan string) {
			defer wg.Done()
			for i := range in {
				select {
				case <-ctx.Done():
					return
				case c <- m.Repo.GetProfile(ctx, i):
				}
			}
		}

		for i := 0; i < numWorkers; i++ {
			go multiplex(in)
		}

		go func() {
			defer close(c)
			wg.Wait()
		}()

		return c
	}

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	var profiles []schema.Profile
	for p := range fanIn(ctx, numWorkers, toStream(ctx, logins...)) {
		profiles = append(profiles, p)
	}

	if err = m.Profile.BulkUpsert(ctx, profiles); err != nil {
		zlog.Warn("error upserting", zap.Error(err))
		return err
	}

	return nil
}
