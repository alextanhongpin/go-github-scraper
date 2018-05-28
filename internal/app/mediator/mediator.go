// Package mediator handles multiple service orchestrations
// and is categorized as domain model
package mediator

import (
	"context"
	"sync"
	"time"

	"github.com/alextanhongpin/go-github-scraper/api/github"
	"github.com/alextanhongpin/go-github-scraper/internal/schema"
	"github.com/alextanhongpin/go-github-scraper/internal/util"
	"github.com/alextanhongpin/go-github-scraper/service/analyticsvc"
	"github.com/alextanhongpin/go-github-scraper/service/profilesvc"
	"github.com/alextanhongpin/go-github-scraper/service/reposvc"
	"github.com/alextanhongpin/go-github-scraper/service/usersvc"

	"go.uber.org/zap"
)

type (
	// Service represents the methods the mediator service must implement
	Service interface {
		FetchUsers(location string, months int, perPage int) error
		FetchRepos(userPerPage, repoPerPage int) error
		UpdateUserCount() error
		UpdateRepoCount() error
		UpdateReposMostRecent(perPage int) error
		UpdateRepoCountByUser(perPage int) error
		UpdateReposMostStars(perPage int) error
		UpdateLanguagesMostPopular(perPage int) error
		UpdateMostRecentReposByLanguage(perPage int) error
		UpdateReposByLanguage(perPage int) error
		UpdateProfile(numWorkers int) error
	}

	// Mediator holds the services in used
	Mediator struct {
		Analytic analyticsvc.Service
		Github   github.API
		Profile  profilesvc.Service
		Repo     reposvc.Service
		User     usersvc.Service
	}

	model struct {
		Mediator
		zlog *zap.Logger
	}
)

// New returns a new mediator service
func New(m Mediator, zlog *zap.Logger) Service {
	return &model{
		m,
		zlog,
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

func (m *model) FetchUsers(location string, months int, perPage int) error {
	zlog := util.LoggerWithRequestID(m.zlog)

	start, ok := m.User.FindLastCreated()
	end := makeEndDate(start, months)

	zlog.Info("fetch users since",
		zap.String("start", start),
		zap.String("end", end),
		zap.Bool("default", ok))

	users, err := m.Github.FetchUsersCursor(location, start, end, perPage)
	if err != nil {
		zlog.Warn("error fetching users", zap.Error(err))
		return err
	}

	// Note that bulk upsert can only hold 1000 users max
	if err := m.User.BulkUpsert(users); err != nil {
		zlog.Warn("error upserting users", zap.Error(err))
		return err
	}

	zlog.Info("upserted users",
		zap.String("start", start),
		zap.String("end", end),
		zap.Int("count", len(users)))
	return nil
}

func (m *model) FetchRepos(userPerPage, repoPerPage int) error {
	zlog := util.LoggerWithRequestID(m.zlog)

	users, err := m.User.FindLastFetched(userPerPage)
	if err != nil {
		zlog.Error("unable to find last fetched users", zap.Error(err))
		return err
	}

	for _, user := range users {
		login := user.Login
		if login == "" {
			continue
		}

		start, ok := m.Repo.FindLastCreatedByUser(login)
		end := util.NewCurrentFormattedDate()

		zlog.Info("fetch repos since",
			zap.String("start", start),
			zap.String("end", end),
			zap.String("for", login),
			zap.Bool("default", ok))

		repos, err := m.Github.FetchReposCursor(login, start, end, repoPerPage)
		if err != nil {
			zlog.Warn("error fetching repos", zap.Error(err))
			return err
		}

		if err := m.Repo.BulkUpsert(repos); err != nil {
			zlog.Warn("error upserting repos", zap.Error(err))
			return err
		}

		if err = m.User.UpdateOne(login); err != nil {
			zlog.Warn("error updating users", zap.Error(err))
			return err
		}

		zlog.Info("upserted repos",
			zap.Int("count", len(repos)),
			zap.String("start", start),
			zap.String("end", end),
			zap.String("for", login))
	}
	return nil
}

// UpdateUserCount updates the analytic type `user_count`
func (m *model) UpdateUserCount() error {
	zlog := util.LoggerWithRequestID(m.zlog)

	count, err := m.User.Count()
	if err != nil {
		zlog.Warn("error getting user count", zap.Error(err))
		return err
	}
	if err := m.Analytic.PostUserCount(count); err != nil {
		zlog.Warn("error updating user count", zap.Error(err))
		return err
	}
	zlog.Info("updated user count", zap.Int("count", count))
	return nil
}

// UpdateRepoCount updates the analytic type `repo_count`
func (m *model) UpdateRepoCount() error {
	zlog := util.LoggerWithRequestID(m.zlog)

	count, err := m.Repo.Count()
	if err != nil {
		zlog.Warn("error getting repo count", zap.Error(err))
		return err
	}
	if err := m.Analytic.PostRepoCount(count); err != nil {
		zlog.Warn("error updating repo count", zap.Error(err))
		return err
	}
	zlog.Info("updated repo count", zap.Int("count", count))
	return nil
}

// UpdateReposMostRecent updates the analytic type `repos_most_recent`
func (m *model) UpdateReposMostRecent(perPage int) error {
	zlog := util.LoggerWithRequestID(m.zlog)

	repos, err := m.Repo.MostRecent(perPage)
	if err != nil {
		zlog.Warn("error fetching most recent repos", zap.Error(err))
		return err
	}

	if err := m.Analytic.PostReposMostRecent(repos); err != nil {
		zlog.Warn("error updating most recent repos", zap.Error(err))
		return err
	}
	zlog.Info("updated most recent repos", zap.Int("count", len(repos)))
	return nil
}

// UpdateRepoCountByUser updates the analytic type `repo_count_by_user`
func (m *model) UpdateRepoCountByUser(perPage int) error {
	zlog := util.LoggerWithRequestID(m.zlog)

	users, err := m.Repo.RepoCountByUser(perPage)
	if err != nil {
		zlog.Warn("error fetching repo count by users", zap.Error(err))
		return err
	}

	if err := m.Analytic.PostRepoCountByUser(users); err != nil {
		zlog.Warn("error updating repo count by users", zap.Error(err))
		return err
	}
	zlog.Info("updated repo count by users", zap.Int("count", len(users)))
	return nil
}

// UpdateReposMostStars updates the analytic type `repos_most_stars`
func (m *model) UpdateReposMostStars(perPage int) error {
	zlog := util.LoggerWithRequestID(m.zlog)

	repos, err := m.Repo.MostStars(perPage)
	if err != nil {
		zlog.Warn("error fetching most stars repos", zap.Error(err))
		return err
	}

	if err := m.Analytic.PostReposMostStars(repos); err != nil {
		zlog.Warn("error updating repo count by users", zap.Error(err))
		return err
	}
	zlog.Info("updated repos most stars", zap.Int("count", len(repos)))
	return nil
}

// UpdateLanguagesMostPopular updates the analytic type `languages_most_popular`
func (m *model) UpdateLanguagesMostPopular(perPage int) error {
	zlog := util.LoggerWithRequestID(m.zlog)

	languages, err := m.Repo.MostPopularLanguage(perPage)
	if err != nil {
		zlog.Warn("error fetching language most popular", zap.Error(err))
		return err
	}

	if err := m.Analytic.PostMostPopularLanguage(languages); err != nil {
		zlog.Warn("error updating language most popular", zap.Error(err))
		return err
	}
	zlog.Info("updated popular languages", zap.Int("count", len(languages)))
	return nil
}

// UpdateMostRecentReposByLanguage updates the analytic type `repos_most_recent_by_language`
func (m *model) UpdateMostRecentReposByLanguage(perPage int) error {
	zlog := util.LoggerWithRequestID(m.zlog)

	languages, err := m.Repo.MostPopularLanguage(perPage)
	if err != nil {
		zlog.Warn("error fetching language most popular", zap.Error(err))
		return err
	}

	var reposByLanguages []schema.RepoLanguage
	for _, lang := range languages {
		r, err := m.Repo.MostRecentReposByLanguage(lang.Name, perPage)
		if err != nil {
			zlog.Warn("error fetching most recent repos by language", zap.Error(err))
			return err
		}
		reposByLanguages = append(reposByLanguages, schema.RepoLanguage{
			Language: lang.Name,
			Repos:    r,
		})
	}

	if err := m.Analytic.PostMostRecentReposByLanguage(reposByLanguages); err != nil {
		zlog.Warn("error updating most recent repos by language", zap.Error(err))
		return err
	}
	zlog.Info("updated most recent repos by language", zap.Int("count", len(reposByLanguages)))
	return nil
}

// UpdateReposByLanguage updates the analytic type `repos_by_language`
func (m *model) UpdateReposByLanguage(perPage int) error {
	zlog := util.LoggerWithRequestID(m.zlog)

	languages, err := m.Repo.MostPopularLanguage(perPage)
	if err != nil {
		zlog.Warn("error fetching language most popular", zap.Error(err))
		return err
	}

	var userCountByLanguage []schema.UserCountByLanguage
	for _, lang := range languages {
		users, err := m.Repo.ReposByLanguage(lang.Name, perPage)
		if err != nil {
			zlog.Warn("error fetching user repo count by language", zap.Error(err))
			return err
		}

		userCountByLanguage = append(userCountByLanguage, schema.UserCountByLanguage{
			Language: lang.Name,
			Users:    users,
		})
	}

	if err := m.Analytic.PostReposByLanguage(userCountByLanguage); err != nil {
		zlog.Warn("error updating user repo count by language", zap.Error(err))
		return err
	}

	zlog.Info("updated repos by languages", zap.Int("count", len(userCountByLanguage)))
	return nil
}

func (m *model) UpdateProfile(numWorkers int) error {
	zlog := util.LoggerWithRequestID(m.zlog)

	logins, err := m.Repo.DistinctLogin()
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

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	var profiles []schema.Profile
	for p := range fanIn(ctx, numWorkers, toStream(ctx, logins...)) {
		profiles = append(profiles, p)
	}

	if err = m.Profile.BulkUpsert(profiles); err != nil {
		zlog.Warn("error upserting", zap.Error(err))
		return err
	}

	zlog.Info("upserted new profiles", zap.Int("count", len(profiles)))
	return nil
}
